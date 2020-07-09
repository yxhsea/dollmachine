package ff_v1flow

import (
	"github.com/Unknwon/com"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_common/ff_err"
	"dollmachine/dolluser/ff_common/ff_header"
	"dollmachine/dolluser/ff_service/flow"
	"net/http"
)

type FlowListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取流水列表
// @tags flow
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /flow/list [get]
func GetFlowList(ctx *gin.Context) {
	var code int
	var FlowListReq FlowListReq
	err := ctx.ShouldBind(&FlowListReq)
	if err != nil {
		logrus.Errorf("GetFlowList should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&FlowListReq)
	if err != nil {
		logrus.Errorf("GetFlowList valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestUserSession().MerchantId
	offset, _ := com.StrTo(FlowListReq.Offset).Int()
	pageSize, _ := com.StrTo(FlowListReq.PageSize).Int()
	totalSize, _ := com.StrTo(FlowListReq.TotalSize).Int()

	flowService := flow.NewFlowService()
	FlowList, page, err := flowService.GetFlowList(merchantId, offset, pageSize, totalSize)
	if err != nil {
		logrus.Errorf("Query Flow list information failure. Error : %v", err)
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": FlowList, "page": page})
	return
}
