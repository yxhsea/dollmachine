package ff_v1device_type

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gogap/logrus"
	"dollmachine/dollplatform/ff_common/ff_err"
	"dollmachine/dollplatform/ff_service/device_type"
	"net/http"
)

type GetDeviceTypeInfoReq struct {
	DeviceTypeId string `form:"device_type_id" valid:"required~设备类型ID不能为空,numeric~设备类型ID应为数字"`
}

// @Summary 获取设备类型
// @tags deviceType
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param device_type_id query string true "设备类型ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /dev_type/info [get]
func GetDeviceTypeInfo(ctx *gin.Context) {
	var code int
	var req GetDeviceTypeInfoReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get device type info should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get device type info valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	deviceTypeService := device_type.NewDeviceTypeService()
	if !deviceTypeService.CheckIsExitsDeviceTypeId(req.DeviceTypeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备类型不存在", "data": ""})
		return
	}

	deviceTypeInfo, err := deviceTypeService.GetDeviceTypeInfo(req.DeviceTypeId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取设备类型失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取设备类型成功", "data": deviceTypeInfo})
	return
}

type GetDeviceTypeListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取设备类型列表
// @tags deviceType
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /dev_type/list [get]
func GetDeviceTypeList(ctx *gin.Context) {
	var code int
	var req GetDeviceTypeListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get device type list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get device type list valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	deviceTypeService := device_type.NewDeviceTypeService()
	deviceTypeList, page, err := deviceTypeService.GetDeviceTypeList(req.Offset, req.PageSize, req.TotalSize)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取设备类型列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取设备类型列表成功", "data": deviceTypeList, "page": page})
	return
}

type AddDeviceTypeReq struct {
	Name string `form:"name" valid:"required~设备类型名称不能为空"`
}

// @Summary 新增设备类型
// @tags deviceType
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param name formData string true "设备类型名称"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /dev_type/add [post]
func AddDeviceType(ctx *gin.Context) {
	var code int
	var req AddDeviceTypeReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("add device type should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("add device type valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	deviceTypeService := device_type.NewDeviceTypeService()
	if deviceTypeService.CheckIsExitsDeviceTypeName(req.Name) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备类型已经存在", "data": ""})
		return
	}

	//新增商户信息
	if !deviceTypeService.AddDeviceType(req.Name) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "新增设备类型失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "新增设备类型成功", "data": ""})
	return
}

type UpdDeviceTypeReq struct {
	DeviceTypeId string `form:"device_type_id" valid:"required~设备类型ID不能为空"`
	Name         string `form:"name" valid:"required~设备类型名称不能为空"`
}

// @Summary 更新商户
// @tags deviceType
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param device_type_id formData string true "设备类型ID"
// @Param name formData string true "设备类型名称"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /dev_type/upd [put]
func UpdDeviceType(ctx *gin.Context) {
	var code int
	var req UpdDeviceTypeReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("upd device type should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("upd device type valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	deviceTypeService := device_type.NewDeviceTypeService()
	if !deviceTypeService.CheckIsExitsDeviceTypeId(req.DeviceTypeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "设备类型不存在", "data": ""})
		return
	}

	if !deviceTypeService.UpdDeviceType(req.DeviceTypeId, req.Name) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "更新设备类型失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "更新设备类型成功", "data": ""})
	return
}
