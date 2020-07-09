package ff_v1suggest

import (
	"github.com/Unknwon/com"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_common/ff_err"
	"dollmachine/dolluser/ff_common/ff_header"
	"dollmachine/dolluser/ff_service/suggest"
	"net/http"
)

type SuggestListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取意见、投诉、列表
// @tags suggest
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /suggest/list [get]
func GetSuggestList(ctx *gin.Context) {
	var code int
	var suggestListReq SuggestListReq
	err := ctx.ShouldBind(&suggestListReq)
	if err != nil {
		logrus.Errorf("GetSuggestList should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&suggestListReq)
	if err != nil {
		logrus.Errorf("GetSuggestList valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	//参数类型转换
	merchantId := ff_header.NewFFHeader(ctx).GetRequestUserSession().MerchantId
	userId := ff_header.NewFFHeader(ctx).GetRequestUserSession().UserId
	offset, _ := com.StrTo(suggestListReq.Offset).Int()
	pageSize, _ := com.StrTo(suggestListReq.PageSize).Int()
	totalSize, _ := com.StrTo(suggestListReq.TotalSize).Int()

	suggestService := suggest.NewSuggestService()
	suggestList, page, err := suggestService.GetSuggestList(merchantId, userId, offset, pageSize, totalSize)
	if err != nil {
		logrus.Errorf("Query Suggest list information failure. Error : %v", err)
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": suggestList, "page": page})
	return
}

type Suggest struct {
	Content    string `form:"content" valid:"required~内容不能为空"`
	Type       string `form:"type" valid:"required~类型不能为空,numeric~类型必须是数字"`
	ContactTel string `form:"contact_tel" valid:"required~电话不能为空"`
}

// @Summary 添加意见
// @tags suggest
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param content formData string true "反馈内容"
// @Param type formData int true "反馈类型 1|建议 2|投诉 3|其他"
// @Param contact_tel formData string true "联系电话"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /suggest/add [post]
func AddSuggest(ctx *gin.Context) {
	var code int
	var suggestInfo Suggest
	ctx.ShouldBind(&suggestInfo)

	//校验Params参数
	_, err := govalidator.ValidateStruct(&suggestInfo)
	if err != nil {
		logrus.Errorf("Add Suggest valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	//参数类型转换
	merchantId := ff_header.NewFFHeader(ctx).GetRequestUserSession().MerchantId
	userId := ff_header.NewFFHeader(ctx).GetRequestUserSession().UserId
	sgType, _ := com.StrTo(suggestInfo.Type).Int()

	suggestService := suggest.NewSuggestService()
	_, err = suggestService.AddSuggest(merchantId, userId, sgType, suggestInfo.Content, suggestInfo.ContactTel)
	if err != nil {
		logrus.Errorf("Add Suggest information failure. Error : %v", err)
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
	return
}
