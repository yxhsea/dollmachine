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

type ApplyForWithDrawReq struct {
	Account        string `form:"account" valid:"required~账号不能为空"`
	AccountName    string `form:"account_name" valid:"required~账号名称不能为空"`
	Bank           string `form:"bank" valid:"required~所属银行不能为空"`
	SubBank        string `form:"sub_bank" valid:"required~所属支行不能为空"`
	BenefitData    string `form:"benefit_data" valid:"required~分润数据不能为空"`
	Invoice        string `form:"invoice" valid:"required~发票不能为空"`
	Amount         string `form:"amount" valid:"required~提现金额不能为空"`
	ExpressCompany string `form:"express_company" valid:"required~快递公司不能为空"`
	ExpressNo      string `form:"express_no" valid:"required~快递单号不能为空"`
}

// @Summary 申请提现
// @tags Settle
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param account formData string true "账号"
// @Param account_name formData string true "账号名称"
// @Param bank formData string true "所属银行"
// @Param sub_bank formData string true "所属支行"
// @Param benefit_data formData string true "分润数据"
// @Param invoice formData string true "发票"
// @Param amount formData string true "提现金额"
// @Param express_company formData string true "快递公司"
// @Param express_no formData string true "快递单号"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /settle/apply/withdraw [post]
func ApplyForWithDraw(ctx *gin.Context) {
	var code int
	var req ApplyForWithDrawReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("apply withdraw should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("apply withdraw valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	//判断是否已到允许申请时间
	if !checkApplyTimeRightOrError() {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "未到申请时间", "data": ""})
		return
	}

	mchSession := ff_header.NewFFHeader(ctx).GetRequestMerchantSession()
	merchantId := mchSession.MerchantId
	merchantName := mchSession.MerchantName
	staffName := mchSession.StaffName
	staffPhone := mchSession.StaffPhone

	settleService := settle.NewSettleService()
	flag := settleService.ApplyForWithDraw(merchantId, merchantName, staffName, staffPhone, req.Account, req.AccountName, req.Bank, req.SubBank,
		req.BenefitData, req.Invoice, req.Amount, req.ExpressCompany, req.ExpressNo)
	if !flag {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "申请提现失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "申请提现成功", "data": ""})
	return
}

// 必须要25号以后才能申请上个月的提现
func checkApplyTimeRightOrError() bool {
	nowTime := time.Now().Unix()
	startTime := now.New(time.Unix(nowTime, 0)).BeginningOfMonth()
	NeedTime := startTime.AddDate(0, 1, 24)
	if NeedTime.Unix() < time.Now().Unix() {
		return true
	}
	return false
}

type GetApplyDrawRecordListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取申请提现记录列表
// @tags Settle
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /settle/apply/withdraw/list [get]
func GetApplyDrawRecordList(ctx *gin.Context) {
	var code int
	var req GetMonthDetailListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get apply draw record list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get apply draw record list valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	settleService := settle.NewSettleService()
	applyDrawRecordList, page, err := settleService.GetApplyDrawRecordList(req.Offset, req.PageSize, req.TotalSize, merchantId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取申请提现记录列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取申请提现记录列表成功", "data": applyDrawRecordList, "page": page})
	return
}
