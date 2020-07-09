package ff_basepay

import (
	"github.com/pkg/errors"
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_payconf"
)

type IPayGate interface {
	SetRechargeNotifyUrl(url string)                                                                          //设置支付回调地址
	SetPayConfMeta(payConfMeta *ff_payconf.PayConfMeta)                                                       //设置支付默认配置
	GetPayConfMeta() *ff_payconf.PayConfMeta                                                                  //获取支付默认配置
	RechargeCreate(pmtRecharge *PmtRecharge, pmtRechargeExt *PmtRechargeExt) (*CreateRechargeResponse, error) //支付发起
	RechargeQuery(pmtRecharge *PmtRecharge, pmtRechargeExt *PmtRechargeExt) (*QueryRechargeResponse, error)   //支付查询
	ParseRechargeNotify(notifyStr string) (*QueryRechargeResponse, error)                                     //解析支付回调
	ValidRechargeNotify(notifyStr string) (bool, error)                                                       //验证支付回调
	RechargeRefund(pmtRecharge *PmtRecharge) error                                                            //支付拒绝
	RechargeCancel(pmtRecharge *PmtRecharge) error                                                            //支付取消
	RechargeClose(pmtRecharge *PmtRecharge) error                                                             //支付关闭
	TransferCreate() error
	TransferQuery() error
	Sign(sendMap map[string]string) (string, error) //签名
	Verify(respMap map[string]string) (bool, error) //校验
}

type PayGate struct {
	PayConfMeta         *ff_payconf.PayConfMeta
	RechargeCallbackUrl string
	RechargeNotifyUrl   string
	TransferCallbackUrl string
	TransferNotifyUrl   string
}

func (p *PayGate) SetRechargeNotifyUrl(url string) {
	p.RechargeNotifyUrl = url
}

func (p *PayGate) SetPayConfMeta(payConfMeta *ff_payconf.PayConfMeta) {
	p.PayConfMeta = payConfMeta
}

func (p *PayGate) GetPayConfMeta() *ff_payconf.PayConfMeta {
	return p.PayConfMeta
}

var paymentNotImpErr = errors.New("支付接口不存在")

func (p *PayGate) RechargeCreate(pmtRecharge *PmtRecharge, pmtRechargeExt *PmtRechargeExt) (*CreateRechargeResponse, error) {
	return nil, paymentNotImpErr
}
func (p *PayGate) RechargeQuery(pmtRecharge *PmtRecharge, pmtRechargeExt *PmtRechargeExt) (*QueryRechargeResponse, error) {
	return nil, paymentNotImpErr
}

func (p *PayGate) ParseRechargeNotify(notifyStr string) (*QueryRechargeResponse, error) {
	return nil, paymentNotImpErr
}
func (p *PayGate) ValidRechargeNotify(notifyStr string) (bool, error) {
	return false, paymentNotImpErr
}
func (p *PayGate) RechargeRefund(pmtRecharge *PmtRecharge) error {
	return paymentNotImpErr
}
func (p *PayGate) RechargeCancel(pmtRecharge *PmtRecharge) error {
	return paymentNotImpErr
}
func (p *PayGate) RechargeClose(pmtRecharge *PmtRecharge) error {
	return paymentNotImpErr
}
func (p *PayGate) TransferCreate() error { return paymentNotImpErr }
func (p *PayGate) TransferQuery() error  { return paymentNotImpErr }
func (p *PayGate) Sign(sendMap map[string]string) (string, string, error) {
	return "", "", paymentNotImpErr
}
func (p *PayGate) Verify(respMap map[string]string) (bool, error) {
	return false, paymentNotImpErr
}
