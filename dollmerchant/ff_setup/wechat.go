package ff_setup

import (
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

func SetupWxMiniAppClient(appId string, appSecret string, mpAppId string, mpAppSecret string, pcAppId string, pcAppSecret string) error {
	ff_vars.MpAppId = mpAppId
	ff_vars.MpAppSecret = mpAppSecret
	ff_vars.MpAccessTokenServer = core.NewDefaultAccessTokenServer(mpAppId, mpAppSecret, nil)
	ff_vars.MpWechatClient = core.NewClient(ff_vars.MpAccessTokenServer, nil)

	ff_vars.PcAppId = pcAppId
	ff_vars.PcAppSecret = pcAppSecret
	ff_vars.PcAccessTokenServer = core.NewDefaultAccessTokenServer(pcAppId, pcAppSecret, nil)
	ff_vars.PcWechatClient = core.NewClient(ff_vars.MpAccessTokenServer, nil)

	ff_vars.MiniUserAppId = appId
	ff_vars.MiniUserAppSecret = appSecret
	ff_vars.MiniUserAccessTokenServer = core.NewDefaultAccessTokenServer(appId, appSecret, nil)
	ff_vars.MiniUserWechatClient = core.NewClient(ff_vars.MiniUserAccessTokenServer, nil)
	return nil
}
