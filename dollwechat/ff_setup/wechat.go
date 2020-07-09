package ff_setup

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
	"dollmachine/dollwechat/ff_config/ff_vars"
)

func SetupWxMpClient(wxMpAppId string, wxMpAppSecret string, wxOriId string, wxToken string, wxEncodedAesKey string) error {
	ff_vars.WxMpAppId = wxMpAppId
	ff_vars.WxMpAppSecret = wxMpAppSecret
	ff_vars.WxOriId = wxOriId
	ff_vars.WxToken = wxToken
	ff_vars.WxEncodedAesKey = wxEncodedAesKey
	ff_vars.WxMpAccessTokenServer = core.NewDefaultAccessTokenServer(wxMpAppId, wxMpAppSecret, nil)
	ff_vars.WxMpWechatClient = core.NewClient(ff_vars.WxMpAccessTokenServer, nil)
	return nil
}
