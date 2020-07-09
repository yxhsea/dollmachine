package ff_setup

import (
	"github.com/gin-gonic/gin"
	"dollmachine/dollplatform/ff_controller/ff_v1device_type"
	"dollmachine/dollplatform/ff_controller/ff_v1login"
	"dollmachine/dollplatform/ff_controller/ff_v1mch"
	"dollmachine/dollplatform/ff_controller/ff_v1test"
	"dollmachine/dollplatform/ff_middleware"
)

func SetupRouter(router *gin.Engine) {
	prefixV1 := "/dev/plf/v1"

	router.PUT(prefixV1+"/test/info", ff_v1test.GetTest)

	//登录接口
	router.POST(prefixV1+"/login/sign_in", ff_v1login.SignIn)

	apiPlfV1 := router.Group(prefixV1)
	apiPlfV1.Use(ff_middleware.PlatformAuth())
	{
		//注销登陆
		apiPlfV1.GET("/login/sign_out", ff_v1login.SignOut)

		//商户管理
		apiPlfV1.GET("/mch/info", ff_v1mch.GetMchInfo)
		apiPlfV1.GET("/mch/list", ff_v1mch.GetMchList)
		apiPlfV1.POST("/mch/add", ff_v1mch.AddMch)
		apiPlfV1.PUT("/mch/upd", ff_v1mch.UpdMch)
		apiPlfV1.PUT("/mch/pwd", ff_v1mch.UpdMchPwd)
		apiPlfV1.PUT("/mch/state", ff_v1mch.UpdMchState)
		apiPlfV1.DELETE("/mch/del", nil)

		//设备类型管理
		apiPlfV1.GET("/dev_type/info", ff_v1device_type.GetDeviceTypeInfo)
		apiPlfV1.GET("/dev_type/list", ff_v1device_type.GetDeviceTypeList)
		apiPlfV1.POST("/dev_type/add", ff_v1device_type.AddDeviceType)
		apiPlfV1.PUT("/dev_type/upd", ff_v1device_type.UpdDeviceType)
	}
}
