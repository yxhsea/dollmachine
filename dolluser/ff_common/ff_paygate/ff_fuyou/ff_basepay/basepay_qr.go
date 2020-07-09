package ff_basepay

import (
	"dollmachine/dolluser/ff_common/ff_convert"
	"dollmachine/dolluser/ff_common/ff_json"
	"dollmachine/dolluser/ff_common/ff_paygate/ff_payconf"
	"dollmachine/dolluser/ff_common/ff_random"
	"strconv"
	"strings"
	"time"
)

type BaseQrPay struct {
	BasePay
}

func (p *BaseQrPay) SetRechargeNotifyUrl(notifyUrl string) {
	p.BasePay.SetRechargeNotifyUrl(notifyUrl)
}

func (p *BaseQrPay) SetService(service string) {
	p.BasePay.SetService(service)
}

func (p *BaseQrPay) RechargeCreate(pmtRecharge *PmtRecharge, pmtRechargeExt *PmtRechargeExt) (*CreateRechargeResponse, error) {

	var sendMap = map[string]string{
		"version":                "1",
		"ins_cd":                 p.PayConfMeta.MchPlatform, //接入机构在富友的唯一代码
		"mchnt_cd":               p.PayConfMeta.MchId,       //富友分配给二级商户的商户号
		"term_id":                "dollmachine",             //终端号
		"random_str":             ff_random.KrandAll(16),    //ff_random.KrandAll(16),
		"sign":                   "",
		"order_type":             p.FuyouService, //订单类型:ALIPAY , WECHAT, JD(京东钱包),QQ(QQ钱包),UNIONPAY
		"goods_des":              pmtRecharge.Subject,
		"goods_detail":           "",
		"goods_tag":              "",
		"addn_inf":               pmtRecharge.AddnInf, //附加参数
		"mchnt_order_no":         strconv.FormatInt(pmtRecharge.RechargeId, 10),
		"curr_type":              strings.ToUpper(pmtRecharge.Currency),
		"order_amt":              strconv.FormatInt(int64(pmtRecharge.Amount*100), 10),
		"term_ip":                ff_convert.IpIntToString(pmtRecharge.ClientIp),
		"txn_begin_ts":           time.Unix(pmtRecharge.CreatedAt, 0).Format("20060102150405"),
		"notify_url":             p.RechargeNotifyUrl,
		"reserved_sub_appid":     "",
		"reserved_limit_pay":     "",
		"reserved_expire_minute": "0",
	}

	var preCreateResp PreCreateResp
	err := p.BasePay.Post(sendMap, ff_payconf.PreCreate, &preCreateResp)
	if err != nil {
		return nil, err
	}

	var createRechargeResponse = &CreateRechargeResponse{}
	createRechargeResponse.Channel = pmtRecharge.Channel
	createRechargeResponse.RechargeId = pmtRecharge.RechargeId
	createRechargeResponse.Action = "QRCode"
	createRechargeResponse.Credential = preCreateResp.QrCode
	createRechargeResponse.SnapShot = make(map[string]string)
	createRechargeResponse.SnapShot["resp"] = ff_json.MarshalToStringNoError(preCreateResp)

	return createRechargeResponse, nil
}

func (p *BaseQrPay) RechargeQuery(pmtRecharge *PmtRecharge, pmtRechargeExt *PmtRechargeExt) (*QueryRechargeResponse, error) {
	return p.BasePay.RechargeQuery(pmtRecharge, pmtRechargeExt)
}
