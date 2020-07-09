package ff_letpay

import (
	"dollmachine/dolluser/ff_common/ff_paygate/ff_fuyou/ff_basepay"
)

//小程序支付

type LetPay struct {
	ff_basepay.WxPreCreatePay
}

func NewLetPay() *LetPay {
	pay := &LetPay{}
	pay.SetService()
	return pay
}

func (p *LetPay) SetRechargeNotifyUrl(notifyUrl string) {
	p.WxPreCreatePay.SetRechargeNotifyUrl(notifyUrl)
}

func (p *LetPay) SetService() {
	p.WxPreCreatePay.SetService("LETPAY")
}

func (p *LetPay) RechargeCreate(pmtRecharge *ff_basepay.PmtRecharge, pmtRechargeExt *ff_basepay.PmtRechargeExt) (*ff_basepay.CreateRechargeResponse, error) {
	return p.WxPreCreatePay.RechargeCreate(pmtRecharge, pmtRechargeExt)
}
