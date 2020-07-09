package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"time"
)

type RequestArgs struct {
	RechargeId int64   `json:"recharge_id"` //交易id
	Channel    string  `json:"channel"`     //交易方式
	Subject    string  `json:"subject"`     //主题
	Amount     float64 `json:"amount"`      //金额
	ClientIp   int64   `json:"client_ip"`   //交易者IP
	CreatedAt  int64   `json:"created_at"`  //创建时间
	Currency   string  `json:"currency"`    //货币

	OpenId    string `json:"open_id"`     //富友openId
	SubOpenId string `json:"sub_open_id"` //子商户AppId
	SubAppId  string `json:"sub_app_id"`  //子商户openId

	NotifyUrl string `json:"notify_url"` //回调地址
}

func main() {
	client, err := jsonrpc.Dial("tcp", ":9225")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	nowTime := time.Now().Unix()
	req := &RequestArgs{
		RechargeId: 1142*1000000000000000 + nowTime,
		Channel:    "wechatpay_qr_fy",
		Subject:    "rpc-test",
		Amount:     0.01,
		ClientIp:   3232235545,
		CreatedAt:  nowTime,

		OpenId:    "",
		SubOpenId: "",
		SubAppId:  "",

		NotifyUrl: "https://wawafront.tunnel.aioil.cn/ks/v1/oauth/recharge/notify",
	}

	var reply interface{}
	err = client.Call("FyPay.SendPay", req, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	replyStr, _ := json.Marshal(reply)
	fmt.Println(string(replyStr)) //{"action":"QRCode","channel":"wechatpay_qr_fy","credential":"weixin://wxpay/bizpayurl?pr=i4NigN9","recharge_id":1142000001526537700}
}
