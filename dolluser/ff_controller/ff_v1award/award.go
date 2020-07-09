package ff_v1award

import (
	"github.com/Unknwon/com"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_common/ff_err"
	"dollmachine/dolluser/ff_common/ff_header"
	"dollmachine/dolluser/ff_service/award"
	"net/http"
)

type AwardListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取中奖记录列表
// @tags award
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /award/list [get]
func GetAwardList(ctx *gin.Context) {
	var code int
	var AwardListReq AwardListReq
	err := ctx.ShouldBind(&AwardListReq)
	if err != nil {
		logrus.Errorf("GetAwardList should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&AwardListReq)
	if err != nil {
		logrus.Errorf("GetAwardList valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	//参数类型转换
	merchantId := ff_header.NewFFHeader(ctx).GetRequestUserSession().MerchantId
	userId := ff_header.NewFFHeader(ctx).GetRequestUserSession().UserId
	offset, _ := com.StrTo(AwardListReq.Offset).Int()
	pageSize, _ := com.StrTo(AwardListReq.PageSize).Int()
	totalSize, _ := com.StrTo(AwardListReq.TotalSize).Int()

	awardService := award.NewAwardService()
	awardList, page, err := awardService.GetAwardList(merchantId, userId, offset, pageSize, totalSize)
	if err != nil {
		logrus.Errorf("Query Award list information failure. Error : %v", err)
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": awardList, "page": page})
	return
}
