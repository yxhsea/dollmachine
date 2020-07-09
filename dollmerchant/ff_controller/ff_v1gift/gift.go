package ff_v1gift

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_common/ff_header"
	"dollmachine/dollmerchant/ff_service/gift"
	"net/http"
)

type GetGiftInfoReq struct {
	GiftId string `form:"gift_id" valid:"required~礼品ID不能为空,numeric~礼品ID应为数字"`
}

// @Summary 获取礼品
// @tags Gift
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param gift_id query string true "礼品ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /gift/info [get]
func GetGiftInfo(ctx *gin.Context) {
	var code int
	var req GetGiftInfoReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get gift info should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get gift info valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	GiftService := gift.NewGiftService()
	if !GiftService.CheckIsExitsGiftId(req.GiftId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "礼品不存在", "data": ""})
		return
	}

	GiftInfo, err := GiftService.GetGiftInfo(req.GiftId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取礼品失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取礼品成功", "data": GiftInfo})
	return
}

type GetGiftListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取礼品列表
// @tags Gift
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /gift/list [get]
func GetGiftList(ctx *gin.Context) {
	var code int
	var req GetGiftListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get gift list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get gift list valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	GiftService := gift.NewGiftService()
	GiftList, page, err := GiftService.GetGiftList(req.Offset, req.PageSize, req.TotalSize)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取礼品列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取礼品列表成功", "data": GiftList, "page": page})
	return
}

type AddGiftReq struct {
	Name string `form:"name" valid:"required~礼品名称不能为空"`
}

// @Summary 新增礼品
// @tags Gift
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param name formData string true "礼品名称"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /gift/add [post]
func AddGift(ctx *gin.Context) {
	var code int
	var req AddGiftReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("add gift should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("add gift valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	GiftService := gift.NewGiftService()
	if GiftService.CheckIsExitsGiftName(req.Name) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "礼品已经存在", "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	//新增礼品信息
	if !GiftService.AddGift(req.Name, string(merchantId)) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "新增礼品失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "新增礼品成功", "data": ""})
	return
}

type UpdGiftReq struct {
	GiftId string `form:"gift_id" valid:"required~礼品ID不能为空"`
	Name   string `form:"name" valid:"required~礼品名称不能为空"`
}

// @Summary 更新礼品
// @tags Gift
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param gift_id formData string true "礼品ID"
// @Param name formData string true "礼品名称"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /gift/upd [put]
func UpdGift(ctx *gin.Context) {
	var code int
	var req UpdGiftReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("upd gift should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("upd gift valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	GiftService := gift.NewGiftService()
	if !GiftService.CheckIsExitsGiftId(req.GiftId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "礼品不存在", "data": ""})
		return
	}

	if !GiftService.UpdGift(req.GiftId, req.Name) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "更新礼品失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "更新礼品成功", "data": ""})
	return
}

type DelGiftReq struct {
	GiftId string `form:"gift_id" valid:"required~礼品ID不能为空"`
}

// @Summary 删除礼品
// @tags Gift
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param gift_id formData string true "礼品ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /gift/del [delete]
func DelGift(ctx *gin.Context) {
	var code int
	var req DelGiftReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("delete gift should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("delete gift valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	GiftService := gift.NewGiftService()
	if !GiftService.CheckIsExitsGiftId(req.GiftId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "礼品不存在", "data": ""})
		return
	}

	if GiftService.CheckIsBindGiftId(req.GiftId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "礼品被绑定", "data": ""})
		return
	}

	if !GiftService.DelGift(req.GiftId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "删除礼品失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "删除礼品成功", "data": ""})
	return
}
