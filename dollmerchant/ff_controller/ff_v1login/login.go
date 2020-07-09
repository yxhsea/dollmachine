package ff_v1login

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_common/ff_header"
	"dollmachine/dollmerchant/ff_common/ff_json"
	"dollmachine/dollmerchant/ff_common/ff_random"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"dollmachine/dollmerchant/ff_service/login"
	"gopkg.in/chanxuehong/wechat.v2/mp/jssdk"
	"io/ioutil"
	"net/http"
	"time"
)

type SignInReq struct {
	LoginToken string `form:"login_token" valid:"required~登录账号不能为空"`
	Password   string `form:"password" valid:"required~登录密码不能为空"`
}

// @Summary 登录授权
// @tags Login
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
	mchSessionStr, flag := loginServer.SignIn(req.LoginToken, req.Password)
	if !flag {
		code = ff_err.ERROR_AUTH_TOKEN
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//返回登录信息
	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": mchSessionStr})
	return
}

// @Summary 注销登录
// @tags Login
// @Produce  json
// @Param Token header string true "Token"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /login/sign_out [get]
func SignOut(ctx *gin.Context) {
	var code int
	mchSession := ff_header.NewFFHeader(ctx).GetRequestMerchantSession()
	token := mchSession.Token
	loginServer := login.NewLoginService()
	if !loginServer.SignOut(token) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
	}
	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
	return
}

//ticket
type TicketRes struct {
	ErrCode   int    `json:"err_code"`
	ErrMsg    string `json:"err_msg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

// @Summary 微信签名
// @tags Login
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param url formData string true "签名URL"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /login/wx_sign [post]
func WxSign(ctx *gin.Context) {
	var code int
	WxUrl := string(ctx.Query("url"))
	token, err := ff_vars.MpWechatClient.Token()
	if err != nil {
		logrus.Error("token获取失败")
		code = ff_err.ERROR
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + token + "&type=jsapi"
	res, err := http.Get(url)
	if err != nil {
		logrus.Error("ticket网络请求失败")
		code = ff_err.ERROR
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	var ticketRes *TicketRes
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Error("解析Body失败")
		code = ff_err.ERROR
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	ff_json.Unmarshal(string(body), &ticketRes)
	if ticketRes.ErrCode != 0 {
		logrus.Error("获取ticket失败")
		code = ff_err.ERROR
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	jsApiTicket := ticketRes.Ticket
	nonceStr := ff_random.KrandAll(16)
	timestamp := fmt.Sprint(time.Now().Unix())
	url = WxUrl
	haveSignature := jssdk.WXConfigSign(jsApiTicket, nonceStr, timestamp, url)

	dataMap := map[string]interface{}{
		"appId":     ff_vars.MpAppId,
		"timestamp": timestamp,
		"nonceStr":  nonceStr,
		"signature": haveSignature,
	}
	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": dataMap})
	return
}
