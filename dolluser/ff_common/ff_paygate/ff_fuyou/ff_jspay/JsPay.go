package ff_jspay

import (
	"dollmachine/dolluser/ff_common/ff_paygate/ff_fuyou/ff_basepay"
)

//公众号支付

type JsPay struct {
	ff_basepay.WxPreCreatePay
}

func NewJsPay() *JsPay {
	pay := &JsPay{}
	pay.SetService()
	return pay
}

func (p *JsPay) SetRechargeNotifyUrl(notifyUrl string) {
	p.WxPreCreatePay.SetRechargeNotifyUrl(notifyUrl)
}

func (p *JsPay) SetService() {
	p.WxPreCreatePay.SetService("JSAPI")
}

func (p *JsPay) RechargeCreate(pmtRecharge *ff_basepay.PmtRecharge, pmtRechargeExt *ff_basepay.PmtRechargeExt) (*ff_basepay.CreateRechargeResponse, error) {
	return p.WxPreCreatePay.RechargeCreate(pmtRecharge, pmtRechargeExt)
}
