package ff_payment

import (
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_fuyou/ff_basepay"
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_fuyou/ff_jspay"
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_fuyou/ff_letpay"
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_fuyou/ff_wechat_qr"
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_payconf"
)

/**
  payChannel 支付方式
  meta 商户Id
*/
func GetPayment(payChannel string, notifyUrl string, meta string) (ff_basepay.IPayGate, error) {
	var payConfMeta *ff_payconf.PayConfMeta

	var payConf *ff_payconf.PayConf
	payConf = ff_payconf.NewPayConf()

	var err error
	if meta == "" {
		payConfMeta, err = payConf.GetDefaultPayConfMeta()
	} else {
		payConfMeta, err = payConf.GetPayConfMeta(meta)
	}
	if err != nil {
		return nil, err
	}

	//todo 可以加参数由外部指定
	payConfMeta.SetChannel(payChannel)

	var payGate ff_basepay.IPayGate
	switch payConfMeta.Channel {
	case ff_payconf.WechatPay_Qr_FuYou:
		payGate = ff_wechat_qr.NewWeChatQrPay()
	case ff_payconf.WECHATPAY_LETPAY_FUYOU:
		payGate = ff_letpay.NewLetPay()
	case ff_payconf.WECHATPAY_JSPAY_FUYOU:
		payGate = ff_jspay.NewJsPay()
	}
	payGate.SetRechargeNotifyUrl(notifyUrl) //todo 临时url 设置回调地址
	payGate.SetPayConfMeta(payConfMeta)     //设置商户MchId

	return payGate, nil
}
