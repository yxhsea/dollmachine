package ff_alipay_qr

import (
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_fuyou/ff_basepay"
)

//支付宝扫码支付

type AlipayQrPay struct {
	ff_basepay.BaseQrPay
}

func NewAlipayQrPay() *AlipayQrPay {
	pay := &AlipayQrPay{}
	pay.SetService()
	return pay
}

func (p *AlipayQrPay) SetRechargeNotifyUrl(notifyUrl string) {
	p.BaseQrPay.SetRechargeNotifyUrl(notifyUrl)
}

func (p *AlipayQrPay) SetService() {
	p.BaseQrPay.SetService("ALIPAY")
}

func (p *AlipayQrPay) RechargeCreate(pmtRecharge *ff_basepay.PmtRecharge, pmtRechargeExt *ff_basepay.PmtRechargeExt) (*ff_basepay.CreateRechargeResponse, error) {
	return p.BaseQrPay.RechargeCreate(pmtRecharge, pmtRechargeExt)
}
