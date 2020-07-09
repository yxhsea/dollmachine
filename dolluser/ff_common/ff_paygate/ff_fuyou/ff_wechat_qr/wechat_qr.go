package ff_wechat_qr

import (
	"dollmachine/dolluser/ff_common/ff_paygate/ff_fuyou/ff_basepay"
)

//微信扫码支付

type WeChatQrPay struct {
	ff_basepay.BaseQrPay
}

func NewWeChatQrPay() *WeChatQrPay {
	pay := &WeChatQrPay{}
	pay.SetService()
	return pay
}

func (p *WeChatQrPay) SetRechargeNotifyUrl(notifyUrl string) {
	p.BaseQrPay.SetRechargeNotifyUrl(notifyUrl)
}

func (p *WeChatQrPay) SetService() {
	p.BaseQrPay.SetService("WECHAT")
}

func (p *WeChatQrPay) RechargeCreate(pmtRecharge *ff_basepay.PmtRecharge, pmtRechargeExt *ff_basepay.PmtRechargeExt) (*ff_basepay.CreateRechargeResponse, error) {
	return p.BaseQrPay.RechargeCreate(pmtRecharge, pmtRechargeExt)
}
