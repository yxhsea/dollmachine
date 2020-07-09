package ff_v1suggest

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_common/ff_header"
	"dollmachine/dollmerchant/ff_service/suggest"
	"net/http"
)

type GetSuggestInfoReq struct {
	SuggestId string `form:"suggest_id" valid:"required~意见ID不能为空,numeric~意见ID应为数字"`
}

// @Summary 获取意见
// @tags Suggest
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param suggest_id query string true "意见ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /suggest/info [get]
func GetSuggestInfo(ctx *gin.Context) {
	var code int
	var req GetSuggestInfoReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get Suggest info should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get Suggest info valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	SuggestService := suggest.NewSuggestService()
	if !SuggestService.CheckIsExitsSuggestId(req.SuggestId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "意见不存在", "data": ""})
		return
	}

	SuggestInfo, err := SuggestService.GetSuggestInfo(req.SuggestId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取意见失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取意见成功", "data": SuggestInfo})
	return
}

type GetSuggestListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取意见列表
// @tags Suggest
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /suggest/list [get]
func GetSuggestList(ctx *gin.Context) {
	var code int
	var req GetSuggestListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get Suggest list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get Suggest list valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	SuggestService := suggest.NewSuggestService()
	SuggestList, page, err := SuggestService.GetSuggestList(req.Offset, req.PageSize, req.TotalSize)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取意见列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取意见列表成功", "data": SuggestList, "page": page})
	return
}

type SuggestReplyReq struct {
	SuggestId string `form:"Suggest_id" valid:"required~意见ID不能为空"`
	Reply     string `form:"reply" valid:"required~意见回复不能为空"`
}

// @Summary 意见回复
// @tags Suggest
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param suggest_id formData string true "意见ID"
// @Param reply formData string true "意见回复"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /suggest/reply [post]
func SuggestReply(ctx *gin.Context) {
	var code int
	var req SuggestReplyReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("upd Suggest should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("upd Suggest valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	SuggestService := suggest.NewSuggestService()
	if !SuggestService.CheckIsExitsSuggestId(req.SuggestId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "意见不存在", "data": ""})
		return
	}

	mchSession := ff_header.NewFFHeader(ctx).GetRequestMerchantSession()
	oprUid := mchSession.StaffId
	oprUName := mchSession.StaffName
	if !SuggestService.SuggestReply(req.SuggestId, req.Reply, oprUid, oprUName) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "更新意见失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "更新意见成功", "data": ""})
	return
}
