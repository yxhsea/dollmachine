package ff_v1device

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_common/ff_header"
	"dollmachine/dollmerchant/ff_service/device"
	"dollmachine/dollmerchant/ff_service/device_type"
	"dollmachine/dollmerchant/ff_service/gift"
	"dollmachine/dollmerchant/ff_service/place"
	"net/http"
)

type GetDeviceInfoReq struct {
	DeviceId string `form:"device_id" valid:"required~设备ID不能为空,numeric~设备ID应为数字"`
}

// @Summary 获取设备
// @tags Device
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param device_id query string true "设备ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /device/info [get]
func GetDeviceInfo(ctx *gin.Context) {
	var code int
	var req GetDeviceInfoReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get Device info should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get Device info valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	DeviceService := device.NewDeviceService()
	if !DeviceService.CheckIsExitsDeviceId(req.DeviceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备不存在", "data": ""})
		return
	}

	DeviceInfo, err := DeviceService.GetDeviceInfo(req.DeviceId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取设备失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取设备成功", "data": DeviceInfo})
	return
}

type GetDeviceListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取设备列表
// @tags Device
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /device/list [get]
func GetDeviceList(ctx *gin.Context) {
	var code int
	var req GetDeviceListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get Device list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get Device list valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	DeviceService := device.NewDeviceService()
	DeviceList, page, err := DeviceService.GetDeviceList(req.Offset, req.PageSize, req.TotalSize)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取设备列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取设备列表成功", "data": DeviceList, "page": page})
	return
}

type BindDeviceReq struct {
	DeviceMac      string `form:"device_mac" valid:"required~设备Mac地址不能为空"`
	DeviceName     string `form:"device_name" valid:"required~设备名称不能为空"`
	DeviceTypeId   string `form:"device_type_id" valid:"required~设备类型ID不能为空,numeric~设备类型ID必须是数字"`
	DeviceTypeName string `form:"device_type_name" valid:"required~设备类型名称不能为空"`
	PlaceId        string `form:"place_id" valid:"required~投放地点ID不能为空,numeric~投放地点ID必须是数字"`
	PlaceName      string `form:"place_name" valid:"required~投放地点名称不能为空"`
	GiftId         string `form:"gift_id" valid:"required~礼品ID不能为空,numeric~礼品ID必须是数字"`
	GiftName       string `form:"gift_name" valid:"required~礼品名称不能为空"`
	GiftStock      string `form:"gift_stock" valid:"required~礼品库存不能为空,numeric~礼品库存必须是数字"`
	UnitPrice      string `form:"unit_price" valid:"required~单局价格,numeric~单局价格必须是数字"`
	Rate           string `form:"rate" valid:"required~中奖概率,numeric~中奖概率必须是数字"`
	Line           string `form:"line" valid:"required~设备场景值,numeric~设备场景值必须是数字"`
}

// @Summary 绑定设备
// @tags Device
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param device_mac formData string true "设备Mac地址"
// @Param device_name formData string true "设备名称"
// @Param device_type_id formData int true "设备类型ID"
// @Param device_type_name formData string true "设备类型名称"
// @Param place_id formData int true "投放地点ID"
// @Param place_name formData string true "投放地点名称"
// @Param gift_id formData int true "礼品ID"
// @Param gift_name formData string true "礼品名称"
// @Param gift_stock formData int true "礼品库存"
// @Param rate formData int true "中奖概率"
// @Param line formData string true "设备场景值  1|线上 2|线下"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /device/bind [post]
func BindDevice(ctx *gin.Context) {
	var code int
	var req BindDeviceReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("Bind Device should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("Bind Device valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	DeviceService := device.NewDeviceService()
	if DeviceService.CheckIsExitsDeviceName(req.DeviceMac) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备已经被绑定", "data": ""})
		return
	}

	if DeviceService.CheckIsExitsDeviceName(req.DeviceName) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备名称已经存在", "data": ""})
		return
	}

	deviceTypeService := device_type.NewDeviceTypeService()
	if !deviceTypeService.CheckIsExitsDeviceTypeId(req.DeviceTypeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备类型不存在", "data": ""})
		return
	}

	PlaceService := place.NewPlaceService()
	if !PlaceService.CheckIsExitsPlaceId(req.PlaceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "投放地点不存在", "data": ""})
		return
	}

	GiftService := gift.NewGiftService()
	if !GiftService.CheckIsExitsGiftId(req.GiftId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "礼品不存在", "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	//新增设备信息
	flag := DeviceService.BindDevice(req.DeviceMac, req.DeviceName, req.DeviceTypeId, req.DeviceTypeName, req.PlaceId, req.PlaceName, req.GiftId,
		req.GiftName, req.GiftStock, req.Line, string(merchantId))
	if !flag {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "新增设备失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "新增设备成功", "data": ""})
	return
}

type UpdDeviceReq struct {
	DeviceId       string `form:"device_id" valid:"required~设备ID地址不能为空"`
	DeviceName     string `form:"device_name" valid:"required~设备名称不能为空"`
	DeviceTypeId   string `form:"device_type_id" valid:"required~设备类型ID不能为空,numeric~设备类型ID必须是数字"`
	DeviceTypeName string `form:"device_type_name" valid:"required~设备类型名称不能为空"`
	PlaceId        string `form:"place_id" valid:"required~投放地点ID不能为空,numeric~投放地点ID必须是数字"`
	PlaceName      string `form:"place_name" valid:"required~投放地点名称不能为空"`
	GiftId         string `form:"gift_id" valid:"required~礼品ID不能为空,numeric~礼品ID必须是数字"`
	GiftName       string `form:"gift_name" valid:"required~礼品名称不能为空"`
	GiftStock      string `form:"gift_stock" valid:"required~礼品库存不能为空,numeric~礼品库存必须是数字"`
	UnitPrice      string `form:"unit_price" valid:"required~单局价格,numeric~单局价格必须是数字"`
	Rate           string `form:"rate" valid:"required~中奖概率,numeric~中奖概率必须是数字"`
	Line           string `form:"line" valid:"required~设备场景值,numeric~设备场景值必须是数字"`
}

// @Summary 更新设备
// @tags Device
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param device_id formData string true "设备ID地址"
// @Param device_name formData string true "设备名称"
// @Param device_type_id formData int true "设备类型ID"
// @Param device_type_name formData string true "设备类型名称"
// @Param place_id formData int true "投放地点ID"
// @Param place_name formData string true "投放地点名称"
// @Param gift_id formData int true "礼品ID"
// @Param gift_name formData string true "礼品名称"
// @Param gift_stock formData int true "礼品库存"
// @Param rate formData int true "中奖概率"
// @Param line formData string true "设备场景值  1|线上 2|线下"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /device/upd [put]
func UpdDevice(ctx *gin.Context) {
	var code int
	var req UpdDeviceReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("upd Device should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("upd Device valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	DeviceService := device.NewDeviceService()
	if !DeviceService.CheckIsExitsDeviceId(req.DeviceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备不存在", "data": ""})
		return
	}

	if DeviceService.CheckIsExitsDeviceName(req.DeviceName) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备名称已经存在", "data": ""})
		return
	}

	deviceTypeService := device_type.NewDeviceTypeService()
	if !deviceTypeService.CheckIsExitsDeviceTypeId(req.DeviceTypeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备类型不存在", "data": ""})
		return
	}

	PlaceService := place.NewPlaceService()
	if !PlaceService.CheckIsExitsPlaceId(req.PlaceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "投放地点不存在", "data": ""})
		return
	}

	GiftService := gift.NewGiftService()
	if !GiftService.CheckIsExitsGiftId(req.GiftId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "礼品不存在", "data": ""})
		return
	}

	if !DeviceService.UpdDevice(req.DeviceId, req.DeviceName, req.DeviceTypeId, req.DeviceTypeName, req.PlaceId, req.PlaceName, req.GiftId,
		req.GiftName, req.GiftStock, req.Line) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "更新设备失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "更新设备成功", "data": ""})
	return
}

type UnbindDeviceReq struct {
	DeviceId string `form:"device_id" valid:"required~设备ID不能为空,numeric~设备ID必须是数字"`
}

// @Summary 解绑设备
// @tags Device
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param device_id formData string true "设备ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /device/unbind [delete]
func UnbindDevice(ctx *gin.Context) {
	var code int
	var req UnbindDeviceReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("Unbind Device should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("Unbind Device valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	DeviceService := device.NewDeviceService()
	if !DeviceService.CheckIsExitsDeviceId(req.DeviceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备不存在", "data": ""})
		return
	}

	if DeviceService.CheckIsBindDeviceId(req.DeviceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备被绑定", "data": ""})
		return
	}

	if !DeviceService.UnbindDevice(req.DeviceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "解绑设备失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "解绑设备成功", "data": ""})
	return
}
