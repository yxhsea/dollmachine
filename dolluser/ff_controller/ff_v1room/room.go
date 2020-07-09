package ff_v1room

import (
	"github.com/Unknwon/com"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_common/ff_err"
	"dollmachine/dolluser/ff_common/ff_header"
	"dollmachine/dolluser/ff_service/room"
	"net/http"
)

type GetRoomInfoReq struct {
	RoomId string `form:"room_id" valid:"required~房间ID不能为空,numeric~房间ID必须是数字"`
}

// @Summary 获取房间信息
// @tags room
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param room_id query string true "房间ID"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /room/info [get]
func GetRoomInfo(ctx *gin.Context) {
	var code int
	var reqRoom GetRoomInfoReq
	err := ctx.ShouldBind(&reqRoom)
	if err != nil {
		logrus.Errorf("GetRoomInfo should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&reqRoom)
	if err != nil {
		logrus.Errorf("GetRoomInfo valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	roomId, _ := com.StrTo(reqRoom.RoomId).Int64()
	//查询房间信息
	userService := room.NewRoomService()
	roomInfo, err := userService.GetRoomInfo(roomId)
	if err != nil {
		logrus.Errorf("Query room information failure. Error : %v", err)
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": roomInfo})
	return
}

type RoomListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取房间列表
// @tags room
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /room/list [get]
func GetRoomList(ctx *gin.Context) {
	var code int
	var roomListReq RoomListReq
	err := ctx.ShouldBind(&roomListReq)
	if err != nil {
		logrus.Errorf("GetRoomList should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&roomListReq)
	if err != nil {
		logrus.Errorf("GetRoomList valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestUserSession().MerchantId
	offset, _ := com.StrTo(roomListReq.Offset).Int()
	pageSize, _ := com.StrTo(roomListReq.PageSize).Int()
	totalSize, _ := com.StrTo(roomListReq.TotalSize).Int()
	roomService := room.NewRoomService()
	roomList, page, err := roomService.GetRoomList(merchantId, offset, pageSize, totalSize)
	if err != nil {
		logrus.Errorf("Query room list information failure. Error : %v", err)
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": roomList, "page": page})
	return
}

type RoomAwardReq struct {
	RoomId string `form:"room_id" valid:"required~房间ID不能为空,numeric~房间ID必须是数字"`
	Limit  string `form:"limit" valid:"required~数据条数不能为空,numeric~房间ID必须是数字"`
}

// @Summary 获取房间中奖列表
// @tags room
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param room_id query string true "房间ID"
// @Param limit query string true "查询多少条数据"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /room/award [get]
func GetRoomAward(ctx *gin.Context) {
	var code int
	var roomAwardReq RoomAwardReq
	err := ctx.ShouldBind(&roomAwardReq)
	if err != nil {
		logrus.Errorf("GetRoomAward should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&roomAwardReq)
	if err != nil {
		logrus.Errorf("GetRoomAward valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	//参数类型转换
	merchantId := ff_header.NewFFHeader(ctx).GetRequestUserSession().MerchantId
	roomId, _ := com.StrTo(roomAwardReq.RoomId).Int64()
	limit, _ := com.StrTo(roomAwardReq.Limit).Int()

	roomService := room.NewRoomService()
	roomList, err := roomService.GetRoomAward(merchantId, roomId, limit)
	if err != nil {
		logrus.Errorf("Query room award list information failure. Error : %v", err)
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": roomList})
	return
}
