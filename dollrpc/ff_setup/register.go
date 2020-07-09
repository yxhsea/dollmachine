package ff_setup

import (
	"dollmachine/dollrpc/ff_test"
	"net/rpc"
	"dollmachine/dollrpc/ff_service"
)

func RegisterService(){
	rpc.Register(new(ff_test.Arith))
	rpc.Register(new(ff_service.FyPay))
}