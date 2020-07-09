package ff_v1

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"strings"
	"dollmachine/dollwechat/ff_config/ff_const"
	"dollmachine/dollwechat/ff_wechat/ff_base"
	log "github.com/sirupsen/logrus"
)

func SubscribeEventHandler(ctx *core.Context){
	log.Printf("收到订阅 Subscribe 事件:\n%s\n", ctx.MsgPlaintext)
	msg := ctx.MixedMsg

	//登录事件
	eKey := strings.Split(msg.EventKey, "-")[0]
	if eKey == "qrscene"+"_"+ff_const.FFUserLoginQrCodeKey {
		userLoginKey := strings.Split(msg.EventKey, "_")[1]

		//set redis
		userNickName := ff_base.UserLogin(userLoginKey, msg.FromUserName)

		//reply message
		ff_base.ReplyContent(msg.FromUserName,userNickName, msg.CreateTime)
	}
}