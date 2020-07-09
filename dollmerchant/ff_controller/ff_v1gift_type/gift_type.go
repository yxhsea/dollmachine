package ff_v1gift_type

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_service/gift_type"
	"net/http"
)

type GetGiftTypeInfoReq struct {
	GiftTypeId string `form:"gift_type_id" valid:"required~礼品类型ID不能为空,numeric~礼品类型ID应为数字"`
}

// @Summary 获取礼品类型
// @tags GiftType
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param gift_type_id query string true "礼品类型ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /gift_type/info [get]
func GetGiftTypeInfo(ctx *gin.Context) {
	var code int
	var req GetGiftTypeInfoReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get gift type info should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get gift type info valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	GiftTypeService := gift_type.NewGiftTypeService()
	if !GiftTypeService.CheckIsExitsGiftTypeId(req.GiftTypeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "礼品类型不存在", "data": ""})
		return
	}

	GiftTypeInfo, err := GiftTypeService.GetGiftTypeInfo(req.GiftTypeId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取礼品类型失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取礼品类型成功", "data": GiftTypeInfo})
	return
}

type GetGiftTypeListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取礼品类型列表
// @tags GiftType
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /gift_type/list [get]
func GetGiftTypeList(ctx *gin.Context) {
	var code int
	var req GetGiftTypeListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get gift type list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get gift type list valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	GiftTypeService := gift_type.NewGiftTypeService()
	GiftTypeList, page, err := GiftTypeService.GetGiftTypeList(req.Offset, req.PageSize, req.TotalSize)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取礼品类型列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取礼品类型列表成功", "data": GiftTypeList, "page": page})
	return
}

type AddGiftTypeReq struct {
	Name string `form:"name" valid:"required~礼品类型名称不能为空"`
}

// @Summary 新增礼品类型
// @tags GiftType
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param name formData string true "礼品类型名称"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /gift_type/add [post]
func AddGiftType(ctx *gin.Context) {
	var code int
	var req AddGiftTypeReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("add gift type should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("add gift type valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	GiftTypeService := gift_type.NewGiftTypeService()
	if GiftTypeService.CheckIsExitsGiftTypeName(req.Name) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "礼品类型已经存在", "data": ""})
		return
	}

	//新增礼品类型信息
	if !GiftTypeService.AddGiftType(req.Name) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "新增礼品类型失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "新增礼品类型成功", "data": ""})
	return
}

type UpdGiftTypeReq struct {
	GiftTypeId string `form:"gift_type_id" valid:"required~礼品类型ID不能为空"`
	Name       string `form:"name" valid:"required~礼品类型名称不能为空"`
}

// @Summary 更新礼品类型
// @tags GiftType
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param gift_type_id formData string true "礼品类型ID"
// @Param name formData string true "礼品类型名称"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /gift_type/upd [put]
func UpdGiftType(ctx *gin.Context) {
	var code int
	var req UpdGiftTypeReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("upd gift type should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("upd gift type valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	GiftTypeService := gift_type.NewGiftTypeService()
	if !GiftTypeService.CheckIsExitsGiftTypeId(req.GiftTypeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "礼品类型不存在", "data": ""})
		return
	}

	if !GiftTypeService.UpdGiftType(req.GiftTypeId, req.Name) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "更新礼品类型失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "更新礼品类型成功", "data": ""})
	return
}

type DelGiftTypeReq struct {
	GiftTypeId string `form:"gift_type_id" valid:"required~礼品类型ID不能为空"`
}

// @Summary 删除礼品类型
// @tags GiftType
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param gift_type_id formData string true "礼品类型ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /gift_type/del [delete]
func DelGiftType(ctx *gin.Context) {
	var code int
	var req DelGiftTypeReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("delete gift type should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("delete gift type valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	GiftTypeService := gift_type.NewGiftTypeService()
	if !GiftTypeService.CheckIsExitsGiftTypeId(req.GiftTypeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "礼品类型不存在", "data": ""})
		return
	}

	if GiftTypeService.CheckIsBindGiftTypeId(req.GiftTypeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "礼品类型被绑定", "data": ""})
		return
	}

	if !GiftTypeService.DelGiftType(req.GiftTypeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "删除礼品类型失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "删除礼品类型成功", "data": ""})
	return
}
