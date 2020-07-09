package ff_router

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"gopkg.in/chanxuehong/wechat.v2/mp/menu"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/callback/request"
	"dollmachine/dollwechat/ff_wechat/ff_v1"
	"dollmachine/dollwechat/ff_config/ff_vars"
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

var (
	// 下面两个变量不一定非要作为全局变量, 根据自己的场景来选择.
	msgHandler core.Handler
	msgServer  *core.Server
)

func SetWxRouter(){
	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(defaultMsgHandler)
	mux.DefaultEventHandleFunc(defaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, ff_v1.TextMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, ff_v1.MenuClickEventHandler)
	mux.EventHandleFunc(request.EventTypeScan, ff_v1.ScanEventHandler)
	mux.EventHandleFunc(request.EventTypeSubscribe, ff_v1.SubscribeEventHandler)

	msgHandler = mux
	msgServer = core.NewServer(ff_vars.WxOriId, ff_vars.WxMpAppId, ff_vars.WxToken, ff_vars.WxEncodedAesKey, msgHandler, nil)
}

func wxCallbackHandler(c *gin.Context) {
	msgServer.ServeHTTP(c.Writer, c.Request, nil)
}

func defaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func defaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}