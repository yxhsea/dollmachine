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
		Channel:    "wechatpay_jspay_fy",
		Subject:    "rpc-test",
		Amount:     0.01,
		ClientIp:   3232235545,
		CreatedAt:  nowTime,

		OpenId:    "",
		SubAppId:  "wxc7d98f96c6bb4a79",
		SubOpenId: "ocuGf0igqRAhgDocsuhgxJgXTw2w",

		NotifyUrl: "https://wawafront.tunnel.aioil.cn/ks/v1/oauth/recharge/notify",
	}

	var reply interface{}
	err = client.Call("FyPay.SendPay", req, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	replyStr, _ := json.Marshal(reply)
	fmt.Println(string(replyStr))
	//{"action":"WxCreatePay","channel":"wechatpay_jspay_fy","credential":"{\"appid\":\"wx0781c1dea664cd9a\",\"nonce_str\":\"1526538581282\",\"package\":\"prepay_id=wx17142941236747ad13dc50ff3793574533\",\"pay_sign\":\"4352679D376CE48D9496438D005B6FA8\",\"sign_type\":\"MD5\",\"timestamp\":\"1526538581282\"}","recharge_id":1142000001526538600}
}
