package ff_v1room

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_common/ff_header"
	"dollmachine/dollmerchant/ff_service/device"
	"dollmachine/dollmerchant/ff_service/room"
	"net/http"
)

type GetRoomInfoReq struct {
	RoomId string `form:"room_id" valid:"required~房间ID不能为空,numeric~房间ID应为数字"`
}

// @Summary 获取房间
// @tags Room
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param room_id query string true "房间ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /room/info [get]
func GetRoomInfo(ctx *gin.Context) {
	var code int
	var req GetRoomInfoReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get room info should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get room info valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	roomService := room.NewRoomService()
	if !roomService.CheckIsExitsRoomId(req.RoomId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "房间不存在", "data": ""})
		return
	}

	roomInfo, err := roomService.GetRoomInfo(req.RoomId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取房间失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取房间成功", "data": roomInfo})
	return
}

type GetRoomListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取房间列表
// @tags Room
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /room/list [get]
func GetRoomList(ctx *gin.Context) {
	var code int
	var req GetRoomListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get room list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get room list valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	roomService := room.NewRoomService()
	roomList, page, err := roomService.GetRoomList(req.Offset, req.PageSize, req.TotalSize)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取房间列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取房间列表成功", "data": roomList, "page": page})
	return
}

type AddRoomReq struct {
	RoomName  string `form:"room_name" valid:"required~房间名称不能为空"`
	DeviceId  string `form:"device_id" valid:"required~设备ID不能为空"`
	Thumbnail string `form:"thumbnail" valid:"required~房间图片URl不能为空"`
}

// @Summary 新增房间
// @tags Room
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param room_name formData string true "房间名称"
// @Param device_id formData string true "设备ID"
// @Param thumbnail formData string true "房间图片URl"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /room/add [post]
func AddRoom(ctx *gin.Context) {
	var code int
	var req AddRoomReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("add room should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("add room valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	roomService := room.NewRoomService()
	if roomService.CheckIsExitsRoomName(req.RoomName) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "房间名称已经存在", "data": ""})
		return
	}

	DeviceService := device.NewDeviceService()
	if !DeviceService.CheckIsExitsDeviceId(req.DeviceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备不存在", "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	//新增房间信息
	if !roomService.AddRoom(req.RoomName, req.DeviceId, req.Thumbnail, string(merchantId)) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "新增房间失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "新增房间成功", "data": ""})
	return
}

type UpdRoomReq struct {
	RoomId    string `form:"room_id" valid:"required~房间ID不能为空"`
	RoomName  string `form:"room_name" valid:"required~房间名称不能为空"`
	DeviceId  string `form:"device_id" valid:"required~设备ID不能为空"`
	Thumbnail string `form:"thumbnail" valid:"required~房间图片URl不能为空"`
}

// @Summary 更新房间
// @tags Room
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param room_id formData string true "房间ID"
// @Param room_name formData string true "房间名称"
// @Param device_id formData string true "设备ID"
// @Param thumbnail formData string true "房间图片URl"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /room/upd [put]
func UpdRoom(ctx *gin.Context) {
	var code int
	var req UpdRoomReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("upd room should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("upd room valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	roomService := room.NewRoomService()
	if !roomService.CheckIsExitsRoomId(req.RoomId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "房间不存在", "data": ""})
		return
	}

	DeviceService := device.NewDeviceService()
	if !DeviceService.CheckIsExitsDeviceId(req.DeviceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备不存在", "data": ""})
		return
	}

	if !roomService.UpdRoom(req.RoomId, req.RoomName, req.DeviceId, req.Thumbnail) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "更新房间失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "更新房间成功", "data": ""})
	return
}

type DelRoomReq struct {
	RoomId string `form:"room_id" valid:"required~房间ID不能为空"`
}

// @Summary 删除房间
// @tags Room
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param room_id formData string true "房间ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /room/del [delete]
func Delroom(ctx *gin.Context) {
	var code int
	var req DelRoomReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("delete room should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("delete room valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	roomService := room.NewRoomService()
	if !roomService.CheckIsExitsRoomId(req.RoomId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "房间不存在", "data": ""})
		return
	}

	if !roomService.DelRoom(req.RoomId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "删除房间失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "删除房间成功", "data": ""})
	return
}
