package ff_basepay

import (
	"github.com/pkg/errors"
	"dollmachine/dolluser/ff_common/ff_paygate/ff_payconf"
)

type IPayGate interface {
	SetRechargeNotifyUrl(url string)
	SetPayConfMeta(payConfMeta *ff_payconf.PayConfMeta)
	GetPayConfMeta() *ff_payconf.PayConfMeta
	RechargeCreate(pmtRecharge *PmtRecharge, pmtRechargeExt *PmtRechargeExt) (*CreateRechargeResponse, error)
	RechargeQuery(pmtRecharge *PmtRecharge, pmtRechargeExt *PmtRechargeExt) (*QueryRechargeResponse, error)
	ParseRechargeNotify(notifyStr string) (*QueryRechargeResponse, error)
	ValidRechargeNotify(notifyStr string) (bool, error)
	RechargeRefund(pmtRecharge *PmtRecharge) error
	RechargeCancel(pmtRecharge *PmtRecharge) error
	RechargeClose(pmtRecharge *PmtRecharge) error
	TransferCreate() error
	TransferQuery() error
	Sign(sendMap map[string]string) (string, error)
	Verify(respMap map[string]string) (bool, error)
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
