package ff_payment

import (
	"fmt"
	"dollmachine/dollrpc/ff_core/ff_common/ff_json"
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_fuyou/ff_basepay"
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_payconf"
	"testing"
	"time"
)

func TestGetPayment(t *testing.T) {
	payGate, err := GetPayment(ff_payconf.WechatPay_Qr_FuYou, ff_payconf.NotifyUrl, "")
	if err != nil {
		fmt.Println("GetPayment, Error : ", err.Error())
	}

	nowTime := time.Now().Unix()

	recharge := &ff_basepay.PmtRecharge{}
	recharge.Subject = "test"
	recharge.RechargeId = 1142*1000000000000000 + nowTime
	recharge.Currency = "cny"
	recharge.Amount = 0.01
	recharge.ClientIp = 3232235545
	recharge.CreatedAt = nowTime
	recharge.Channel = ff_payconf.WechatPay_Qr_FuYou

	meta := ff_basepay.RechargeExtMeta{}
	meta.OpenId = "ooIeqsza0qteX1OidgsOAIjxS1Ws"
	meta.SubOpenId = ""
	meta.SubAppId = ""

	//调用充值方法
	createRechargeResponse, err := payGate.RechargeCreate(recharge, &ff_basepay.PmtRechargeExt{Meta: ff_json.MarshalToStringNoError(meta)})
	if err != nil {
		fmt.Println("RechargeCreate, Error : ", err.Error())
	}

	fmt.Println(createRechargeResponse)
}

func TestGetPayment2(t *testing.T) {
	payGate, err := GetPayment(ff_payconf.WECHATPAY_JSPAY_FUYOU, ff_payconf.NotifyUrl, "")
	if err != nil {
		fmt.Println("GetPayment, Error : ", err.Error())
	}

	nowTime := time.Now().Unix()

	recharge := &ff_basepay.PmtRecharge{}
	recharge.Subject = "test"
	recharge.RechargeId = 1142*5000000000000000 + nowTime
	recharge.Currency = "cny"
	recharge.Amount = 0.01
	recharge.ClientIp = 3232235545
	recharge.CreatedAt = nowTime
	recharge.Channel = ff_payconf.WECHATPAY_JSPAY_FUYOU

	meta := ff_basepay.RechargeExtMeta{}
	//meta.OpenId = "ooIeqsza0qteX1OidgsOAIjxS1Ws"
	meta.SubOpenId = "ocuGf0igqRAhgDocsuhgxJgXTw2w"
	meta.SubAppId = "wxc7d98f96c6bb4a79"

	//调用充值方法
	createRechargeResponse, err := payGate.RechargeCreate(recharge, &ff_basepay.PmtRechargeExt{Meta: ff_json.MarshalToStringNoError(meta)})
	if err != nil {
		fmt.Println("RechargeCreate, Error : ", err.Error())
	}

	fmt.Println(createRechargeResponse)
}
