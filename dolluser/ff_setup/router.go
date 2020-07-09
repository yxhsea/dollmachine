package ff_setup

import (
	"github.com/gin-gonic/gin"
	"dollmachine/dolluser/ff_controller/ff_test"
	"dollmachine/dolluser/ff_controller/ff_v1address"
	"dollmachine/dolluser/ff_controller/ff_v1award"
	"dollmachine/dolluser/ff_controller/ff_v1flow"
	"dollmachine/dolluser/ff_controller/ff_v1game"
	"dollmachine/dolluser/ff_controller/ff_v1login"
	"dollmachine/dolluser/ff_controller/ff_v1notice"
	"dollmachine/dolluser/ff_controller/ff_v1pay"
	"dollmachine/dolluser/ff_controller/ff_v1room"
	"dollmachine/dolluser/ff_controller/ff_v1suggest"
	"dollmachine/dolluser/ff_controller/ff_v1user"
	"dollmachine/dolluser/ff_middleware"
)

func SetupRouter(router *gin.Engine) {
	prefixV1 := "/dev/usr/v1"

	router.GET(prefixV1+"/test/info", ff_test.GetTest)

	//登录接口 TODO::模拟登录
	router.POST(prefixV1+"/login/auth", ff_v1login.LoginAuth)
	router.POST(prefixV1+"/login/check", nil)

	apiUserV1 := router.Group(prefixV1)
	apiUserV1.Use(ff_middleware.UserAuth())
	{
		//用户信息
		apiUserV1.GET("/user/info", ff_v1user.GetUserInfo)

		//房间
		apiUserV1.GET("/room/info", ff_v1room.GetRoomInfo)
		apiUserV1.GET("/room/list", ff_v1room.GetRoomList)
		apiUserV1.GET("/room/award", ff_v1room.GetRoomAward)

		//公告
		apiUserV1.GET("/notice/info", ff_v1Notice.GetNoticeInfo)
		apiUserV1.GET("/notice/list", ff_v1Notice.GetNoticeList)

		//收货地址
		apiUserV1.GET("/address/info", ff_v1address.GetAddressInfo)
		apiUserV1.POST("/address/add", ff_v1address.AddAddress)

		//意见反馈、投诉、其他
		apiUserV1.GET("/suggest/list", ff_v1suggest.GetSuggestList)
		apiUserV1.POST("/suggest/add", ff_v1suggest.AddSuggest)

		//流水记录
		apiUserV1.GET("/flow/list", ff_v1flow.GetFlowList)

		//游戏记录
		apiUserV1.GET("/game/list", ff_v1game.GetGameList)

		//中奖记录(抓中记录)
		apiUserV1.GET("/award/list", ff_v1award.GetAwardList)

		/******************************** 支付 *****************************/
		//充值
		apiUserV1.GET("/pay/recharge/query", ff_v1pay.QueryRecharge)
		apiUserV1.POST("/pay/recharge/create", ff_v1pay.CreateRecharge)
		apiUserV1.POST("/pay/recharge/notify", ff_v1pay.NotifyRecharge)
		//邮费
		apiUserV1.GET("/pay/express/query", ff_v1pay.QueryExpress)
		apiUserV1.POST("/pay/express/create", ff_v1pay.CreateExpress)
		apiUserV1.POST("/pay/express/notify", ff_v1pay.NotifyExpress)
	}
}
