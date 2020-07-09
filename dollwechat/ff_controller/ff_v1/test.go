package ff_v1

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	"dollmachine/dollwechat/ff_config/ff_vars"
	"dollmachine/dollwechat/ff_common/ff_res"
	"github.com/gin-gonic/gin"
)

func Test(ctx *gin.Context){
	redirectUrl := ctx.Query("redirect_url")
	state := ctx.Query("state")
	if len(redirectUrl) <= 0 {
		ccErr := ff_res.NewCCErr("参数不能为空",ff_res.ErrorCode0ServiceError, ff_res.ErrorCode1TypeInvalidRequestErr,ff_res.ErrorCode2TypeInvalidParameter,ff_res.ErrorCode3TypeRequestError,"")
		ff_res.NetHttpNewFailJsonResp(ctx.Request, ctx.Writer, ccErr)
		return
	}
	authUrl := oauth2.AuthCodeURL(ff_vars.WxMpAppId, redirectUrl, "snsapi_userinfo", state)
	ff_res.NetHttpNewSuccessJsonResp(ctx.Request, ctx.Writer, map[string]interface{}{"auth_url":authUrl})
	return
}
