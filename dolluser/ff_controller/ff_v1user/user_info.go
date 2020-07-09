package ff_v1user

import (
	"github.com/gin-gonic/gin"
	"dollmachine/dolluser/ff_common/ff_err"
	"dollmachine/dolluser/ff_common/ff_header"
	"dollmachine/dolluser/ff_service/user"
	"net/http"
)

// @Summary 获取用户信息
// @tags user
// @Produce  json
// @Param Token header string true "Token令牌"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /user/info [get]
func GetUserInfo(ctx *gin.Context) {
	var code int
	userId := ff_header.NewFFHeader(ctx).GetRequestUserSession().UserId

	//查询用户信息
	userService := user.NewUserService()
	userInfo := userService.GetUserInfo(userId)

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": userInfo})
	return
}
