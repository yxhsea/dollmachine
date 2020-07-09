package ff_middleware

import (
	"github.com/gin-gonic/gin"
	"dollmachine/dollplatform/ff_cache/platform_session"
	"dollmachine/dollplatform/ff_common/ff_err"
	"net/http"
)

func PlatformAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		var data interface{}

		code = ff_err.SUCCESS
		token := ctx.Request.Header.Get("Token")
		if token == "" {
			code = ff_err.INVALID_PARAMS
		} else {
			//校验Token
			flag := platform_session.NewPlatformSession().CheckIsExitsByToken(token)
			if !flag {
				code = ff_err.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}

		if code != ff_err.SUCCESS {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  ff_err.GetMsg(code),
				"data": data,
			})

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
