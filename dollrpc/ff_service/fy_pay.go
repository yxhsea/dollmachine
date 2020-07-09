package ff_service

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollrpc/ff_core/ff_common/ff_json"
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_fuyou/ff_basepay"
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_payconf"
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_payment"
	"net/url"
	"strconv"
	"strings"
)

type RequestSendPay struct {
	RechargeId int64   `json:"recharge_id"` //交易id
	Channel    string  `json:"channel"`     //交易方式
	Subject    string  `json:"subject"`     //主题
	Amount     float64 `json:"amount"`      //金额
	ClientIp   int64   `json:"client_ip"`   //交易者IP
	CreatedAt  int64   `json:"created_at"`  //创建时间
	Currency   string  `json:"currency"`    //货币

	OpenId    string `json:"open_id"`     //富友openId
	SubOpenId string `json:"sub_open_id"` //子商户AppId
	SubAppId  string `json:"sub_app_id"`  //子商户openId

	NotifyUrl string `json:"notify_url"` //回调地址
}

type ResponseSendPay struct {
	Channel    string            `json:"channel"`
	Credential string            `json:"credential"` // 需要返回给前端的内容
	Action     string            `json:"action"`     // 返回类型标识 告知前端需要执行什么操作 例如 QrCode二维码形式 WxCreatePay调起微信支付
	RechargeId int64             `json:"recharge_id"`
	SnapShot   map[string]string `json:"-"`
}

//flowId := "1142" + time.Now().Format("20060102") 流水号生成规则 前缀 + 当前日期

type FyPay int

//发起支付
func (p *FyPay) SendPay(req *RequestSendPay, res *ResponseSendPay) error {
	payGate, err := ff_payment.GetPayment(req.Channel, req.NotifyUrl, "")
	if err != nil {
		logrus.Errorf("GetPayment, Error : ", err.Error())
		return err
	}

	recharge := &ff_basepay.PmtRecharge{}
	recharge.Subject = req.Subject
	recharge.RechargeId = req.RechargeId
	recharge.Currency = req.Currency
	recharge.Amount = req.Amount
	recharge.ClientIp = req.ClientIp
	recharge.CreatedAt = req.CreatedAt
	recharge.Channel = req.Channel

	meta := ff_basepay.RechargeExtMeta{}
	meta.OpenId = req.OpenId
	meta.SubAppId = req.SubAppId
	meta.SubOpenId = req.SubOpenId

	//调用充值方法
	createRechargeResponse, err := payGate.RechargeCreate(recharge, &ff_basepay.PmtRechargeExt{Meta: ff_json.MarshalToStringNoError(meta)})
	if err != nil {
		logrus.Errorf("RechargeCreate, Error : ", err.Error())
		return err
	}

	*res = ResponseSendPay{
		Channel:    createRechargeResponse.Channel,
		Credential: createRechargeResponse.Credential,
		Action:     createRechargeResponse.Action,
		RechargeId: createRechargeResponse.RechargeId,
		SnapShot:   createRechargeResponse.SnapShot,
	}

	return nil
}

//解析回调数据
func (p *FyPay) ParseNotify(req *string, res *ff_basepay.QueryRechargeResponse) error {
	payGate, err := ff_payment.GetPayment("", "", "")
	if err != nil {
		return err
	}

	notifyStr, err := url.PathUnescape(*req)
	if err != nil {
		return err
	}

	//再次解 去除+等信息
	notifyStr, err = url.QueryUnescape(notifyStr)
	if err != nil {
		return err
	}

	notifyStr = strings.Replace(notifyStr, "req=", "", 1)

	res, err = payGate.ParseRechargeNotify(notifyStr)

	return nil
}

type RequestValidNotify struct {
	NotifyStr  string `json:"notify_str"`
	Additional string `json:"additional"`
}

//验证回调数据
func (p *FyPay) ValidNotify(req *RequestValidNotify, res *bool) error {
	payGate, err := ff_payment.GetPayment("", "", "")
	if err != nil {
		return err
	}

	var payConf *ff_payconf.PayConf
	payConf = ff_payconf.NewPayConf()
	payConfMeta, err := payConf.GetPayConfMeta(req.Additional)
	if err != nil {
		return err
	}
	payGate.SetPayConfMeta(payConfMeta)
	//验证数据是否来自官方
	*res, err = payGate.ValidRechargeNotify(req.NotifyStr)
	if err != nil {
		return err
	}
	return nil
}

//查询支付状态
func (p *FyPay) QueryPayState(req *string, res *ff_basepay.QueryRechargeResponse) error {
	payGate, err := ff_payment.GetPayment("", "", "")
	if err != nil {
		return err
	}

	recharge := &ff_basepay.PmtRecharge{}
	recharge.RechargeId, _ = strconv.ParseInt(*req, 10, 64)
	res, err = payGate.RechargeQuery(recharge, &ff_basepay.PmtRechargeExt{})
	if err != nil {
		return err
	}

	return nil
}
