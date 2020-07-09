package ff_router

import (
	"dollmachine/dollwechat/ff_controller/ff_v1"
	"github.com/gin-gonic/gin"
)

func SetHttpRouter(router *gin.Engine, mode string){
	prefixUrl := "/" + mode
	/*router.PanicHandler = ff_c50x.ApiServerPanic
    router.MethodNotAllowed = http.HandlerFunc(ff_c50x.MethodNotAllowed)
	router.NotFound = http.HandlerFunc(ff_c50x.NotFound)*/

	prefixV1 := prefixUrl + "/wechat/v1"
	router.GET(prefixV1 + "/test", ff_v1.Test)
	router.GET(prefixV1 + "/menu", ff_v1.CreateMenu)

	prefixV1User := prefixV1 + "/user"
	//公众号授权登录
	router.GET(prefixV1User + "/auth_url", ff_v1.GetAuthUrl)
	router.POST(prefixV1User + "/sign_up", ff_v1.SignUp)
	//二维码登录
	router.GET(prefixV1User + "/qr_code", ff_v1.GetWxQrCode)
	router.POST(prefixV1User + "/login", ff_v1.GetLoginUser)


	//微信服务器回调
	router.GET(prefixV1 + "/callback", wxCallbackHandler)
	router.POST(prefixV1 + "/callback", wxCallbackHandler)
}