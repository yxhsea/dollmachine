package ff_setup

import (
	"github.com/gin-gonic/gin"
	"dollmachine/dollmerchant/ff_controller/ff_v1device"
	"dollmachine/dollmerchant/ff_controller/ff_v1device_type"
	"dollmachine/dollmerchant/ff_controller/ff_v1flow"
	"dollmachine/dollmerchant/ff_controller/ff_v1gift"
	"dollmachine/dollmerchant/ff_controller/ff_v1gift_type"
	"dollmachine/dollmerchant/ff_controller/ff_v1login"
	"dollmachine/dollmerchant/ff_controller/ff_v1mch"
	"dollmachine/dollmerchant/ff_controller/ff_v1notice"
	"dollmachine/dollmerchant/ff_controller/ff_v1place"
	"dollmachine/dollmerchant/ff_controller/ff_v1room"
	"dollmachine/dollmerchant/ff_controller/ff_v1settle"
	"dollmachine/dollmerchant/ff_controller/ff_v1staff"
	"dollmachine/dollmerchant/ff_controller/ff_v1suggest"
	"dollmachine/dollmerchant/ff_controller/ff_v1test"
	"dollmachine/dollmerchant/ff_controller/ff_v1upload"
	"dollmachine/dollmerchant/ff_controller/ff_v1win"
	"dollmachine/dollmerchant/ff_middleware"
)

func SetupRouter(router *gin.Engine) {
	prefixV1 := "/dev/mch/v1"

	router.PUT(prefixV1+"/test/info", ff_v1test.GetTest)

	//登录接口
	router.POST(prefixV1+"/login/sign_in", ff_v1login.SignIn)
	//微信签名
	router.GET(prefixV1+"/login/wx_sign", ff_v1login.WxSign)

	apiMchV1 := router.Group(prefixV1)
	apiMchV1.Use(ff_middleware.MerchantAuth())
	{
		//注销登陆
		apiMchV1.GET("/login/sign_out", ff_v1login.SignOut)

		//商户信息
		apiMchV1.GET("/mch/info", ff_v1mch.GetMchInfo)

		//职工管理
		apiMchV1.GET("/staff/info", ff_v1staff.GetStaffInfo)
		apiMchV1.GET("/staff/list", ff_v1staff.GetStaffList)
		apiMchV1.POST("/staff/add", ff_v1staff.AddStaff)
		apiMchV1.PUT("/staff/upd", ff_v1staff.UpdStaff)
		apiMchV1.PUT("/staff/pwd", ff_v1staff.PwdStaff)
		apiMchV1.DELETE("/staff/upd", ff_v1staff.DelStaff)

		//设备类型管理
		apiMchV1.GET("/dev_type/info", ff_v1device_type.GetDeviceTypeInfo)
		apiMchV1.GET("/dev_type/list", ff_v1device_type.GetDeviceTypeList)

		//礼品类型管理
		apiMchV1.GET("/gift_type/info", ff_v1gift_type.GetGiftTypeInfo)
		apiMchV1.GET("/gift_type/list", ff_v1gift_type.GetGiftTypeList)
		apiMchV1.POST("/gift_type/add", ff_v1gift_type.AddGiftType)
		apiMchV1.PUT("/gift_type/upd", ff_v1gift_type.UpdGiftType)
		apiMchV1.DELETE("/gift_type/del", ff_v1gift_type.DelGiftType)

		//礼品管理
		apiMchV1.GET("/gift/info", ff_v1gift.GetGiftInfo)
		apiMchV1.GET("/gift/list", ff_v1gift.GetGiftList)
		apiMchV1.POST("/gift/add", ff_v1gift.AddGift)
		apiMchV1.PUT("/gift/upd", ff_v1gift.UpdGift)
		apiMchV1.DELETE("/gift/del", ff_v1gift.DelGift)

		//投放地址管理
		apiMchV1.GET("/place/info", ff_v1place.GetPlaceInfo)
		apiMchV1.GET("/place/list", ff_v1place.AddPlace)
		apiMchV1.POST("/place/add", ff_v1place.AddPlace)
		apiMchV1.PUT("/place/upd", ff_v1place.UpdPlace)
		apiMchV1.DELETE("/place/del", ff_v1place.DelPlace)

		//设备管理
		apiMchV1.GET("/device/info", ff_v1device.GetDeviceInfo)
		apiMchV1.GET("/device/list", ff_v1device.GetDeviceList)
		apiMchV1.POST("/device/bind", ff_v1device.BindDevice)
		apiMchV1.POST("/device/unbind", ff_v1device.UnbindDevice)
		apiMchV1.PUT("/device/upd", ff_v1device.UpdDevice)

		//房间管理
		apiMchV1.GET("/room/info", ff_v1room.GetRoomInfo)
		apiMchV1.GET("/room/list", ff_v1room.GetRoomList)
		apiMchV1.POST("/room/add", ff_v1room.AddRoom)
		apiMchV1.PUT("/room/upd", ff_v1room.UpdRoom)
		apiMchV1.DELETE("/room/del", ff_v1room.Delroom)

		//上传图片
		apiMchV1.POST("/upload/img", ff_v1upload.UploadImg)
		//上传文件
		apiMchV1.POST("/upload/file", ff_v1upload.UploadFile)

		//中奖管理
		apiMchV1.GET("/win/online/list", ff_v1win.GetWinOnlineList)
		apiMchV1.GET("/win/offline/list", ff_v1win.GetWinOfflineList)
		apiMchV1.GET("/win/express/list", ff_v1win.GetWinExpressList)
		apiMchV1.POST("/win/cash", ff_v1win.WinCash)
		apiMchV1.POST("/win/send", ff_v1win.WinSend)

		//公告管理
		apiMchV1.GET("/notice/info", ff_v1notice.GetNoticeInfo)
		apiMchV1.GET("/notice/list", ff_v1notice.GetNoticeList)
		apiMchV1.POST("/notice/add", ff_v1notice.AddNotice)
		apiMchV1.PUT("/notice/upd", ff_v1notice.UpdNotice)
		apiMchV1.DELETE("/notice/del", ff_v1notice.DelNotice)

		//意见管理
		apiMchV1.GET("/suggest/info", ff_v1suggest.GetSuggestInfo)
		apiMchV1.GET("/suggest/list", ff_v1suggest.GetSuggestList)
		apiMchV1.POST("/suggest/reply", ff_v1suggest.SuggestReply)

		//流水记录
		apiMchV1.GET("/flow/record/list", ff_v1flow.GetFlowRecordList)

		//结算中心
		apiMchV1.GET("/settle/current/month/detail", ff_v1settle.GetCurrentMonthDetail)
		apiMchV1.GET("/settle/month/detail/list", ff_v1settle.GetMonthDetailList)
		apiMchV1.GET("/settle/user/recharge/list", ff_v1settle.GetUserRechargeList)
		apiMchV1.GET("/settle/user/integral/list", ff_v1settle.GetUserIntegralList)
		apiMchV1.GET("/settle/apply/withdraw/list", ff_v1settle.GetApplyDrawRecordList)
		apiMchV1.POST("/settle/apply/withdraw", ff_v1settle.ApplyForWithDraw)
	}
}
