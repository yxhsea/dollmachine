package ff_v1settle

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_common/ff_header"
	"dollmachine/dollmerchant/ff_service/settle"
	"net/http"
	"time"
)

// @Summary 获取当前月结算详情
// @tags Settle
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /settle/current/month/detail [get]
func GetCurrentMonthDetail(ctx *gin.Context) {
	var code int
	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	nowTime := time.Now().Unix()
	StartTime := now.New(time.Unix(nowTime, 0)).BeginningOfMonth().Unix()
	EndTime := now.New(time.Unix(nowTime, 0)).EndOfMonth().Unix()
	if nowTime == 0 {
		StartTime, EndTime = 0, 0
	}

	settleService := settle.NewSettleService()
	currMonthDetailList, err := settleService.GetCurrentMonthDetail(merchantId, StartTime, EndTime)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取当前月结算详情失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取当前月结算详情成功", "data": currMonthDetailList})
	return
}

type GetMonthDetailListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
	Time      string `form:"time" valid:"-"`
}

// @Summary 获取月结算详情列表
// @tags Settle
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Param time query string true "时间范围"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /settle/month/detail/list [get]
func GetMonthDetailList(ctx *gin.Context) {
	var code int
	var req GetMonthDetailListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get month detail list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get month detail list valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	settleService := settle.NewSettleService()
	monthDetailList, page, err := settleService.GetMonthDetailList(req.Offset, req.PageSize, req.TotalSize, merchantId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取月结算详情记录列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取月结算详情记录列表成功", "data": monthDetailList, "page": page})
	return
}
