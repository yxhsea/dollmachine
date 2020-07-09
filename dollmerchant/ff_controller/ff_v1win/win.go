package ff_v1win

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_common/ff_header"
	"dollmachine/dollmerchant/ff_service/win"
	"net/http"
)

type GetWinOnlineListReq struct {
	Offset     string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize  string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize   string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
	CashStatus string `form:"cash_status" valid:"required~cash_status不能为空,numeric~cash_status必须是数字"`
	UserId     string `form:"user_id" valid:"required~user_id不能为空,numeric~user_id必须是数字"`
}

// @Summary 获取线上中奖列表
// @tags Win
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Param cash_status query string true "核销状态 0|全部 1|已核销 2|未核销"
// @Param user_id query string true "用户ID 0|全部"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /win/online/list [get]
func GetWinOnlineList(ctx *gin.Context) {
	var code int
	var req GetWinOnlineListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get Win online list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get Win online valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	WinService := win.NewWinService()
	WinList, page, err := WinService.GetWinOnlineList(req.Offset, req.PageSize, req.TotalSize, req.CashStatus, req.UserId, merchantId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取线上中奖列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取线上中奖列表成功", "data": WinList, "page": page})
	return
}

type GetWinOfflineListReq struct {
	Offset     string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize  string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize   string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
	CashStatus string `form:"cash_status" valid:"required~cash_status不能为空,numeric~cash_status必须是数字"`
	UserId     string `form:"user_id" valid:"required~user_id不能为空,numeric~user_id必须是数字"`
}

// @Summary 获取线下中奖列表
// @tags Win
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Param cash_status query string true "核销状态 0|全部 1|已核销 2|未核销"
// @Param user_id query string true "用户ID 0|全部"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /win/offline/list [get]
func GetWinOfflineList(ctx *gin.Context) {
	var code int
	var req GetWinOfflineListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get Win offline list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get Win offline valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	WinService := win.NewWinService()
	WinList, page, err := WinService.GetWinOfflineList(req.Offset, req.PageSize, req.TotalSize, req.CashStatus, req.UserId, merchantId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取线下中奖列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取线下中奖列表成功", "data": WinList, "page": page})
	return
}

type GetWinExpressListReq struct {
	Offset     string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize  string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize   string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
	SendStatus string `form:"send_status" valid:"required~send_status不能为空,numeric~send_status必须是数字"`
	UserId     string `form:"user_id" valid:"required~user_id不能为空,numeric~user_id必须是数字"`
}

// @Summary 获取快递发货列表
// @tags Win
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Param send_status query string true "发货状态 0|全部 1|已发货 2|未发货"
// @Param user_id query string true "用户ID 0|全部"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /win/express/list [get]
func GetWinExpressList(ctx *gin.Context) {
	var code int
	var req GetWinExpressListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get Win express list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get Win express valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	WinService := win.NewWinService()
	WinList, page, err := WinService.GetWinExpressList(req.Offset, req.PageSize, req.TotalSize, req.SendStatus, req.UserId, merchantId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取快递发货列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取快递发货列表成功", "data": WinList, "page": page})
	return
}

type WinCashReq struct {
	PlayId     string `form:"play_id" valid:"required~play_id不能为空,numeric~play_id必须是数字"`
	CashStatus string `form:"cash_status" valid:"required~cash_status不能为空,numeric~cash_status必须是数字"`
}

// @Summary 核销奖品
// @tags Win
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param play_id formData string true "游戏记录ID"
// @Param cash_status formData string true "1|核销"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /win/cash [post]
func WinCash(ctx *gin.Context) {
	var code int
	var req WinCashReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("Cash prize should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("Cash prize express valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	mchSession := ff_header.NewFFHeader(ctx).GetRequestMerchantSession()
	oprUid := mchSession.StaffId
	oprUName := mchSession.StaffName
	WinService := win.NewWinService()
	if !WinService.WinCash(req.PlayId, req.CashStatus, oprUid, oprUName) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "核销奖品失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "核销奖品成功", "data": ""})
	return
}

type WinSendReq struct {
	RecordId       string `form:"record_id" valid:"required~record_id不能为空,numeric~record_id必须是数字"`
	ExpressCompany string `form:"express_company" valid:"required~快递公司名称不能为空"`
	ExpressNo      string `form:"express_no" valid:"required~快递单号不能为空"`
}

// @Summary 快递奖品
// @tags Win
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param record_id formData string true "兑奖记录ID"
// @Param express_company formData string true "快递公司名称"
// @Param express_no formData string true "快递单号"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /win/send [post]
func WinSend(ctx *gin.Context) {
	var code int
	var req WinSendReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("Send prize should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("Send prize express valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	WinService := win.NewWinService()
	if !WinService.WinSend(req.RecordId, req.ExpressCompany, req.ExpressNo) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "发货失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "发货成功", "data": ""})
	return
}
