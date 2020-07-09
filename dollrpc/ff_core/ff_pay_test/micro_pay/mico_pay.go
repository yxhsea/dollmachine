package micro_pay

import (
	"encoding/json"
	"fmt"
	"dollmachine/dollrpc/ff_core/ff_pay_test/base_pay"
)

type MicroPay struct {
	base_pay.BasePay
}

func (p *MicroPay) CreatePay(req *MicroPayRequestMsg) error {
	var sendMap map[string]string
	byteData, err := json.Marshal(sendMap)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteData, &sendMap)
	if err != nil {
		return err
	}

	var respData MicroPayResponseMsg
	err = p.Post(sendMap, "", &respData)
	if err != nil {
		return err
	}
	fmt.Println(respData)
	return nil
}
