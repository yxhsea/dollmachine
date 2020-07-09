package ff_v1

import (
	"dollmachine/dollwechat/ff_config/ff_vars"
	"dollmachine/dollwechat/ff_common/ff_res"
	"gopkg.in/chanxuehong/wechat.v2/oauth2"
	mpoauth2 "gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	"dollmachine/dollwechat/ff_service/user_auth"
	"github.com/gin-gonic/gin"
)

const (
	oauth2Scope  = "snsapi_base"
)

func GetAuthUrl(ctx *gin.Context){
	redirectUrl := ctx.Query("redirect_url")
	state := ctx.Query("state")

	if len(redirectUrl) <= 0 {
		ccErr := ff_res.NewCCErr("redirect_url参数不能为空",ff_res.ErrorCode0ServiceError, ff_res.ErrorCode1TypeInvalidRequestErr,ff_res.ErrorCode2TypeInvalidParameter,ff_res.ErrorCode3TypeRequestError,ff_res.ErrorMsgParameterError)
		ff_res.NetHttpNewFailJsonResp(ctx.Request, ctx.Writer, ccErr)
		return
	}
	authUrl := mpoauth2.AuthCodeURL(ff_vars.WxMpAppId, redirectUrl, oauth2Scope, state)

	ff_res.NetHttpNewSuccessJsonResp(ctx.Request, ctx.Writer, map[string]interface{}{"auth_url":authUrl})
	return
}

func SignUp(ctx *gin.Context)  {
	code := ctx.Query("code")
	_ = ctx.Query("state") //备用字段
	fromTag := ctx.Query("from_tag")

	if len(code) <= 0 {
		ccErr := ff_res.NewCCErr(ff_res.ErrorMsgParameterError,ff_res.ErrorCode0ServiceError, ff_res.ErrorCode1TypeInvalidRequestErr,ff_res.ErrorCode2TypeInvalidParameter,ff_res.ErrorCode3TypeRequestError,"code参数不能为空")
		ff_res.NetHttpNewFailJsonResp(ctx.Request, ctx.Writer, ccErr)
		return
	}

	//通过code换取token
	oauth2Endpoint := mpoauth2.NewEndpoint(ff_vars.WxMpAppId, ff_vars.WxMpAppSecret)
	oauth2Client := oauth2.Client{Endpoint: oauth2Endpoint}
	token, err := oauth2Client.ExchangeToken(code)
	if err != nil {
		ccErr := ff_res.NewCCErr(ff_res.ErrorMsgThirdServiceError,ff_res.ErrorCode0ServiceError, ff_res.ErrorCode1TypeInvalidRequestErr,ff_res.ErrorCode2TypeInvalidParameter,ff_res.ErrorCode3TypeRequestError,"ExchangeToken失败")
		ff_res.NetHttpNewFailJsonResp(ctx.Request, ctx.Writer, ccErr)
		return
	}

	//拉取用户信息
	userInfo, err := mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)
	if err != nil {
		ccErr := ff_res.NewCCErr(ff_res.ErrorMsgThirdServiceError,ff_res.ErrorCode0ServiceError, ff_res.ErrorCode1TypeInvalidRequestErr,ff_res.ErrorCode2TypeInvalidParameter,ff_res.ErrorCode3TypeRequestError,"获取用户信息失败")
		ff_res.NetHttpNewFailJsonResp(ctx.Request, ctx.Writer, ccErr)
		return
	}

	userSession := &user_auth.UserInfo{
		OpenId:userInfo.OpenId,
		Nickname:userInfo.Nickname,
		Sex:userInfo.Sex,
		City:userInfo.City,
		Province:userInfo.Province,
		Country:userInfo.Country,
		HeadImageURL:userInfo.HeadImageURL,
		UnionId:userInfo.UnionId,
		FromTag:fromTag,
	}
	userMap := user_auth.AddOrUpdateUserLogin(userSession)

	ff_res.NetHttpNewSuccessJsonResp(ctx.Request, ctx.Writer, userMap)
	return
}
