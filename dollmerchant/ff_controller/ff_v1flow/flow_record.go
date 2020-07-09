package ff_v1flow

import (
	"github.com/Unknwon/com"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/now"
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_common/ff_header"
	"dollmachine/dollmerchant/ff_service/flow"
	"net/http"
	"time"
)

type GetFlowRecordListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
	UserId    string `form:"user_id" valid:"-"`
	Type      string `form:"type" valid:"-"`
	Time      string `form:"time" valid:"-"`
}

// @Summary 获取流水列表
// @tags FlowRecord
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Param user_id query string true "用户ID"
// @Param type query string true "类型 0|全部 1|充值 2|积分兑换 3|运费支付 4|游戏消耗"
// @Param time query string true "时间范围"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /flow/record/list [get]
func GetFlowRecordList(ctx *gin.Context) {
	var code int
	var req GetFlowRecordListReq
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

	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	userId := req.UserId
	IType := req.Type
	Time, _ := com.StrTo(req.Time).Int64()
	StartTime := now.New(time.Unix(Time, 0)).BeginningOfMonth().Unix()
	EndTime := now.New(time.Unix(Time, 0)).EndOfMonth().Unix()
	if Time == 0 {
		StartTime, EndTime = 0, 0
	}
	flowService := flow.NewFlowRecordService()
	flowRecordList, page, err := flowService.GetFlowRecordList(req.Offset, req.PageSize, req.TotalSize, userId, IType, merchantId, StartTime, EndTime)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取流水记录列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取流水记录列表成功", "data": flowRecordList, "page": page})
	return
}
