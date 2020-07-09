package ff_v1notice

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_common/ff_header"
	"dollmachine/dollmerchant/ff_service/notice"
	"net/http"
)

type GetNoticeInfoReq struct {
	NoticeId string `form:"notice_id" valid:"required~公告ID不能为空,numeric~公告ID应为数字"`
}

// @Summary 获取公告
// @tags Notice
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param Notice_id query string true "公告ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /notice/info [get]
func GetNoticeInfo(ctx *gin.Context) {
	var code int
	var req GetNoticeInfoReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get Notice info should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get Notice info valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	NoticeService := Notice.NewNoticeService()
	if !NoticeService.CheckIsExitsNoticeId(req.NoticeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "公告不存在", "data": ""})
		return
	}

	NoticeInfo, err := NoticeService.GetNoticeInfo(req.NoticeId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取公告失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取公告成功", "data": NoticeInfo})
	return
}

type GetNoticeListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取公告列表
// @tags Notice
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /notice/list [get]
func GetNoticeList(ctx *gin.Context) {
	var code int
	var req GetNoticeListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get Notice list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get Notice list valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	NoticeService := Notice.NewNoticeService()
	NoticeList, page, err := NoticeService.GetNoticeList(req.Offset, req.PageSize, req.TotalSize)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取公告列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取公告列表成功", "data": NoticeList, "page": page})
	return
}

type AddNoticeReq struct {
	Title   string `form:"title" valid:"required~公告名称不能为空"`
	Content string `form:"content" valid:"-"`
	Thumb   string `form:"thumb" valid:"-"`
}

// @Summary 新增公告
// @tags Notice
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param name formData string true "公告名称"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /notice/add [post]
func AddNotice(ctx *gin.Context) {
	var code int
	var req AddNoticeReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("add Notice should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("add Notice valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	NoticeService := Notice.NewNoticeService()
	//新增公告信息
	if !NoticeService.AddNotice(req.Title, req.Content, req.Thumb, merchantId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "新增公告失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "新增公告成功", "data": ""})
	return
}

type UpdNoticeReq struct {
	NoticeId string `form:"notice_id" valid:"required~公告ID不能为空"`
	Title    string `form:"title" valid:"required~公告名称不能为空"`
	Content  string `form:"content" valid:"-"`
	Thumb    string `form:"thumb" valid:"-"`
}

// @Summary 更新公告
// @tags Notice
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param Notice_id formData string true "公告ID"
// @Param name formData string true "公告名称"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /notice/upd [put]
func UpdNotice(ctx *gin.Context) {
	var code int
	var req UpdNoticeReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("upd Notice should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("upd Notice valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	NoticeService := Notice.NewNoticeService()
	if !NoticeService.CheckIsExitsNoticeId(req.NoticeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "公告不存在", "data": ""})
		return
	}

	if !NoticeService.UpdNotice(req.Title, req.Content, req.Thumb, req.NoticeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "更新公告失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "更新公告成功", "data": ""})
	return
}

type DelNoticeReq struct {
	NoticeId string `form:"notice_id" valid:"required~公告ID不能为空"`
}

// @Summary 删除公告
// @tags Notice
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param Notice_id formData string true "公告ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /notice/del [delete]
func DelNotice(ctx *gin.Context) {
	var code int
	var req DelNoticeReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("delete Notice should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("delete Notice valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	NoticeService := Notice.NewNoticeService()
	if !NoticeService.CheckIsExitsNoticeId(req.NoticeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "公告不存在", "data": ""})
		return
	}

	if !NoticeService.DelNotice(req.NoticeId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "删除公告失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "删除公告成功", "data": ""})
	return
}
