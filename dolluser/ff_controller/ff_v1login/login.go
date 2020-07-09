package ff_v1login

import (
	"github.com/Unknwon/com"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_common/ff_err"
	"dollmachine/dolluser/ff_service/mch_merchant"
	"dollmachine/dolluser/ff_service/user"
	"net/http"
)

type Login struct {
	UserId     string `form:"user_id" valid:"required~用户ID不能为空,numeric~用户ID必须是数字"`
	MerchantId string `form:"merchant_id" valid:"required~商户ID不能为空,numeric~商户ID必须是数字"`
}

// @Summary 登录授权
// @tags login
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param user_id formData string true "用户ID"
// @Param merchant_id formData string true "商户ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /login/auth [post]
func LoginAuth(ctx *gin.Context) {
	var code int
	var loginInfo Login
	err := ctx.ShouldBind(&loginInfo)
	if err != nil {
		logrus.Errorf("Login auth should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&loginInfo)
	if err != nil {
		logrus.Errorf("Login auth valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	//参数类型转换
	userId, _ := com.StrTo(loginInfo.UserId).Int64()
	merchantId, _ := com.StrTo(loginInfo.MerchantId).Int64()

	//检查系统中是否存在该商户
	mchService := mch_merchant.NewMerchantService()
	if !mchService.CheckIsExitsByMchMerchantId(merchantId) {
		code = ff_err.ERROR_SYSTEM_NOT_EXIST_MERCHANT
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//检查系统中是否存在该用户
	userService := user.NewUserService()
	if !userService.CheckIsExitsByUserId(userId) {
		code = ff_err.ERROR_SYSTEM_NOT_EXIST_USER
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//开始执行登录操作
	userSessionStr, flag := user.NewUserService().UserLogin(userId, merchantId)
	if !flag {
		code = ff_err.ERROR_AUTH_TOKEN
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//返回登录信息
	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": userSessionStr})
	return
}

func LoginCheck(ctx *gin.Context) {

}
