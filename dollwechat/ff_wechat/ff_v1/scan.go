package ff_v1

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"strings"
	"dollmachine/dollwechat/ff_config/ff_const"
	"dollmachine/dollwechat/ff_wechat/ff_base"
	log "github.com/sirupsen/logrus"
)

func ScanEventHandler(ctx *core.Context){
	log.Debugf("收到scan事件:\n%s\n", ctx.MsgPlaintext)
	msg := ctx.MixedMsg

	//登录事件
	eKey := strings.Split(msg.EventKey, "-")[0]
	if eKey == ff_const.FFUserLoginQrCodeKey {
		//set redis
		userNickName := ff_base.UserLogin(msg.EventKey, msg.FromUserName)

		//reply message
		ff_base.ReplyContent(msg.FromUserName,userNickName, msg.CreateTime)
	}
}
