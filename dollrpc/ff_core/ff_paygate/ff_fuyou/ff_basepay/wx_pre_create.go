package ff_basepay

import (
	"fmt"
	"dollmachine/dollrpc/ff_core/ff_common/ff_convert"
	"dollmachine/dollrpc/ff_core/ff_common/ff_json"
	"dollmachine/dollrpc/ff_core/ff_common/ff_random"
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_payconf"
	"strconv"
	"strings"
	"time"
)

type WxPreCreatePay struct {
	BasePay
}

func (p *WxPreCreatePay) SetRechargeNotifyUrl(notifyUrl string) {
	p.BasePay.SetRechargeNotifyUrl(notifyUrl)
}

func (p *WxPreCreatePay) SetService(service string) {
	p.BasePay.SetService(service)
}

func (p *WxPreCreatePay) RechargeCreate(pmtRecharge *PmtRecharge, pmtRechargeExt *PmtRechargeExt) (*CreateRechargeResponse, error) {
	meta := RechargeExtMeta{}
	ff_json.Unmarshal(pmtRechargeExt.Meta, &meta)

	//TODO::交易备用字段, 待addn_inf加上之后修复
	var rechargeId string
	if pmtRecharge.RechargeId <= 0 {
		rechargeId = pmtRecharge.RechargeBakId
	} else {
		rechargeId = fmt.Sprint(pmtRecharge.RechargeId)
	}

	var sendMap = map[string]string{
		"version":        "1.0",
		"ins_cd":         p.PayConfMeta.MchPlatform, //接入机构在富友的唯一代码
		"mchnt_cd":       p.PayConfMeta.MchId,       //富友分配给二级商户的商户号
		"term_id":        "dollmachine",             //终端号
		"random_str":     ff_random.KrandAll(32),    //ff_random.KrandAll(16),
		"sign":           "",
		"goods_des":      pmtRecharge.Subject, //可选 商品描述, 商品或支付单简要描述
		"goods_detail":   "",                  //可选 商品详情, 商品名称明细
		"goods_tag":      "",                  //可选 商品标记
		"product_id":     "",                  //可选 商品标识
		"addn_inf":       pmtRecharge.AddnInf, //附加参数
		"mchnt_order_no": rechargeId,
		"curr_type":      strings.ToUpper(pmtRecharge.Currency),
		"order_amt":      strconv.FormatInt(int64(pmtRecharge.Amount*100), 10),
		"term_ip":        ff_convert.IpIntToString(pmtRecharge.ClientIp),
		"txn_begin_ts":   time.Unix(pmtRecharge.CreatedAt, 0).Format("20060102150405"),
		"notify_url":     p.RechargeNotifyUrl,
		"limit_pay":      "",             //限制支付,no_credit:不能使用信用卡
		"trade_type":     p.FuyouService, //JSAPI--公众号支付、APP--app支付、FWC--支付宝服务窗、QQ--QQ页面支付、QQJSAPI--QQJSAPI、JDJSAPI--京东JS、LETPAY-小程序
		"openid":         meta.OpenId,    //用户标识(微信公众号服务商模式必填，其他不填)
		"sub_openid":     meta.SubOpenId, //子商户用户标识 支付宝服务窗为用户buyer_id（此场景必填） 微信公众号为用户的openid(非服务商模式必填)
		"sub_appid":      meta.SubAppId,  //子商户公众号id,微信交易为商户的app_id(非服务商模式必填)
	}

	var wxPreCreateResp WxPreCreateResp
	err := p.BasePay.Post(sendMap, ff_payconf.WxPreCreate, &wxPreCreateResp)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	var createRechargeResponse = &CreateRechargeResponse{}
	createRechargeResponse.Channel = pmtRecharge.Channel
	createRechargeResponse.RechargeId = pmtRecharge.RechargeId
	createRechargeResponse.Action = "WxCreatePay"

	sdkMap := map[string]string{
		"appid":     wxPreCreateResp.SdkAppid,
		"timestamp": wxPreCreateResp.SdkTimestamp,
		"nonce_str": wxPreCreateResp.SdkNoncestr,
		"package":   wxPreCreateResp.SdkPackage,
		"sign_type": wxPreCreateResp.SdkSigntype,
		"pay_sign":  wxPreCreateResp.SdkPaysign,
	}
	credential := ff_json.MarshalToStringNoError(sdkMap)

	fmt.Println("sign", credential)

	createRechargeResponse.Credential = credential
	createRechargeResponse.SnapShot = make(map[string]string)
	createRechargeResponse.SnapShot["resp"] = ff_json.MarshalToStringNoError(wxPreCreateResp)

	return createRechargeResponse, nil
}
