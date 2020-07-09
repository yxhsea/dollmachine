package ff_basepay

import (
	"fmt"
	"dollmachine/dolluser/ff_common/ff_convert"
	"dollmachine/dolluser/ff_common/ff_paygate/ff_fuyou/ff_alipay_qr"
	"dollmachine/dolluser/ff_common/ff_paygate/ff_payconf"
	"dollmachine/dolluser/ff_config/ff_vars"
	"testing"
	"time"
)

func TestBasePay_RechargeCreate(t *testing.T) {

	ff_vars.PayConf = ff_payconf.NewPayConf()

	nowTime := time.Now().Unix()
	recharge := &PmtRecharge{}
	recharge.RechargeId = nowTime + 123456789
	recharge.Subject = "卡盟测试"
	recharge.Currency = "cny"
	recharge.Amount = 0.01
	recharge.ClientIp = ff_convert.StringIpToInt("127.0.0.1")
	recharge.CreatedAt = nowTime

	rechargeExt := &PmtRechargeExt{}

	pay := &ff_alipay_qr.AlipayQrPay{}
	payConfMeta, _ := ff_vars.PayConf.GetDefaultPayConfMeta()
	pay.SetPayConfMeta(payConfMeta)
	pay.SetService()
	pay.SetRechargeNotifyUrl("http://www.wrx.cn")
	resp, err := pay.RechargeCreate(recharge, rechargeExt)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(recharge.RechargeId, resp)
}

func TestBasePay_RechargeQuery(t *testing.T) {
	ff_vars.PayConf = ff_payconf.NewPayConf()

	recharge := &PmtRecharge{}
	recharge.RechargeId = 1

	rechargeExt := &PmtRechargeExt{}

	pay := &ff_alipay_qr.AlipayQrPay{}
	payConfMeta, _ := ff_vars.PayConf.GetDefaultPayConfMeta()
	pay.SetPayConfMeta(payConfMeta)
	pay.SetService()
	pay.SetRechargeNotifyUrl("http://www.wrx.cn")
	resp, err := pay.RechargeQuery(recharge, rechargeExt)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
}
