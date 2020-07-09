package ff_v1login

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dollplatform/ff_common/ff_err"
	"dollmachine/dollplatform/ff_common/ff_header"
	"dollmachine/dollplatform/ff_service/login"
	"net/http"
)

type SignInReq struct {
	LoginToken string `form:"login_token" valid:"required~登录账号不能为空"`
	Password   string `form:"password" valid:"required~登录密码不能为空"`
}

// @Summary 登录授权
// @tags login
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param login_token formData string true "登录账号"
// @Param password formData string true "登录密码"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /login/sign_in [post]
func SignIn(ctx *gin.Context) {
	var code int
	var req SignInReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("sign in should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("sign in valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	loginServer := login.NewLoginService()
	//检查系统中是否存在该商户
	if !loginServer.CheckIsExitsByLoginToken(req.LoginToken) {
		code = ff_err.ERROR_SYSTEM_NOT_EXIST_MERCHANT
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//检查密码是否正确
	if !loginServer.CheckPassword(req.LoginToken, req.Password) {
		code = ff_err.ERROR_PASSWORD_INVALID
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//开始执行登录操作
	plfSessionStr, flag := loginServer.SignIn(req.LoginToken, req.Password)
	if !flag {
		code = ff_err.ERROR_AUTH_TOKEN
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//返回登录信息
	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": plfSessionStr})
	return
}

// @Summary 注销登录
// @tags login
// @Produce  json
// @Param Token header string true "Token"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /login/sign_out [get]
func SignOut(ctx *gin.Context) {
	var code int
	plfSession := ff_header.NewFFHeader(ctx).GetRequestPlatformSession()
	token := plfSession.Token
	loginServer := login.NewLoginService()
	if !loginServer.SignOut(token) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
	}
	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
	return
}
