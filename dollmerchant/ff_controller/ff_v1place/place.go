package ff_v1place

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_common/ff_header"
	"dollmachine/dollmerchant/ff_service/place"
	"net/http"
)

type GetPlaceInfoReq struct {
	PlaceId string `form:"place_id" valid:"required~投放地点ID不能为空,numeric~投放地点ID应为数字"`
}

// @Summary 获取投放地点
// @tags Place
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param Place_id query string true "投放地点ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /place/info [get]
func GetPlaceInfo(ctx *gin.Context) {
	var code int
	var req GetPlaceInfoReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get Place info should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get Place info valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	PlaceService := place.NewPlaceService()
	if !PlaceService.CheckIsExitsPlaceId(req.PlaceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "投放地点不存在", "data": ""})
		return
	}

	PlaceInfo, err := PlaceService.GetPlaceInfo(req.PlaceId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取投放地点失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取投放地点成功", "data": PlaceInfo})
	return
}

type GetPlaceListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取投放地点列表
// @tags Place
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /place/list [get]
func GetPlaceList(ctx *gin.Context) {
	var code int
	var req GetPlaceListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get Place list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get Place list valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	PlaceService := place.NewPlaceService()
	PlaceList, page, err := PlaceService.GetPlaceList(req.Offset, req.PageSize, req.TotalSize)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取投放地点列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取投放地点列表成功", "data": PlaceList, "page": page})
	return
}

type AddPlaceReq struct {
	Name string `form:"name" valid:"required~投放地点名称不能为空"`
}

// @Summary 新增投放地点
// @tags Place
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param name formData string true "投放地点名称"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /place/add [post]
func AddPlace(ctx *gin.Context) {
	var code int
	var req AddPlaceReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("add Place should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("add Place valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	PlaceService := place.NewPlaceService()
	if PlaceService.CheckIsExitsPlaceName(req.Name) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "投放地点已经存在", "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	//新增投放地点信息
	if !PlaceService.AddPlace(req.Name, string(merchantId)) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "新增投放地点失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "新增投放地点成功", "data": ""})
	return
}

type UpdPlaceReq struct {
	PlaceId string `form:"place_id" valid:"required~投放地点ID不能为空"`
	Name    string `form:"name" valid:"required~投放地点名称不能为空"`
}

// @Summary 更新投放地点
// @tags Place
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param Place_id formData string true "投放地点ID"
// @Param name formData string true "投放地点名称"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /place/upd [put]
func UpdPlace(ctx *gin.Context) {
	var code int
	var req UpdPlaceReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("upd Place should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("upd Place valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	PlaceService := place.NewPlaceService()
	if !PlaceService.CheckIsExitsPlaceId(req.PlaceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "投放地点不存在", "data": ""})
		return
	}

	if !PlaceService.UpdPlace(req.PlaceId, req.Name) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "更新投放地点失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "更新投放地点成功", "data": ""})
	return
}

type DelPlaceReq struct {
	PlaceId string `form:"place_id" valid:"required~投放地点ID不能为空"`
}

// @Summary 删除投放地点
// @tags Place
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param place_id formData string true "投放地点ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /place/del [delete]
func DelPlace(ctx *gin.Context) {
	var code int
	var req DelPlaceReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("delete Place should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("delete Place valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	PlaceService := place.NewPlaceService()
	if !PlaceService.CheckIsExitsPlaceId(req.PlaceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "投放地点不存在", "data": ""})
		return
	}

	if PlaceService.CheckIsBindPlaceId(req.PlaceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "投放地点被绑定", "data": ""})
		return
	}

	if !PlaceService.DelPlace(req.PlaceId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "删除投放地点失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "删除投放地点成功", "data": ""})
	return
}
