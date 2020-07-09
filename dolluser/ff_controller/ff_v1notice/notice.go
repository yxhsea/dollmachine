package ff_v1Notice

import (
	"github.com/Unknwon/com"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_common/ff_err"
	"dollmachine/dolluser/ff_common/ff_header"
	"dollmachine/dolluser/ff_service/notice"
	"net/http"
)

type Notice struct {
	NoticeId string `form:"notice_id" valid:"required~房间ID不能为空,numeric~房间ID必须是数字"`
}

// @Summary 获取公告信息
// @tags notice
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param notice_id query string true "公告ID"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /notice/info [get]
func GetNoticeInfo(ctx *gin.Context) {
	var code int
	var reqNotice Notice
	err := ctx.ShouldBind(&reqNotice)
	if err != nil {
		logrus.Errorf("GetNoticeInfo should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&reqNotice)
	if err != nil {
		logrus.Errorf("GetNoticeInfo valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	noticeId, _ := com.StrTo(reqNotice.NoticeId).Int64()
	//查询公告信息
	userService := notice.NewNoticeService()
	NoticeInfo, err := userService.GetNoticeInfo(noticeId)
	if err != nil {
		logrus.Errorf("Query Notice information failure. Error : %v", err)
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": NoticeInfo})
	return
}

type NoticeListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取公告列表
// @tags notice
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /notice/list [get]
func GetNoticeList(ctx *gin.Context) {
	var code int
	var NoticeListReq NoticeListReq
	err := ctx.ShouldBind(&NoticeListReq)
	if err != nil {
		logrus.Errorf("GetNoticeList should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&NoticeListReq)
	if err != nil {
		logrus.Errorf("GetNoticeList valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestUserSession().MerchantId
	offset, _ := com.StrTo(NoticeListReq.Offset).Int()
	pageSize, _ := com.StrTo(NoticeListReq.PageSize).Int()
	totalSize, _ := com.StrTo(NoticeListReq.TotalSize).Int()
	NoticeService := notice.NewNoticeService()
	NoticeList, page, err := NoticeService.GetNoticeList(merchantId, offset, pageSize, totalSize)
	if err != nil {
		logrus.Errorf("Query Notice list information failure. Error : %v", err)
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": NoticeList, "page": page})
	return
}
