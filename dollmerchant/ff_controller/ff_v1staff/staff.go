package ff_v1staff

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_common/ff_header"
	"dollmachine/dollmerchant/ff_service/staff"
	"net/http"
)

type GetStaffInfoReq struct {
	StaffId string `form:"staff_id" valid:"required~职工ID不能为空,numeric~职工ID应为数字"`
}

// @Summary 获取职工
// @tags Staff
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param staff_id query string true "职工ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /staff/info [get]
func GetStaffInfo(ctx *gin.Context) {
	var code int
	var req GetStaffInfoReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get staff info should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get staff info valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	staffService := staff.NewStaffService()
	if !staffService.CheckIsExitsStaffId(req.StaffId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "职工不存在", "data": ""})
		return
	}

	staffInfo, err := staffService.GetStaffInfo(req.StaffId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取职工失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取职工成功", "data": staffInfo})
	return
}

type GetStaffListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取职工列表
// @tags Staff
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /staff/list [get]
func GetStaffList(ctx *gin.Context) {
	var code int
	var req GetStaffListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get staff list should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get staff list valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	staffService := staff.NewStaffService()
	staffList, page, err := staffService.GetStaffList(req.Offset, req.PageSize, req.TotalSize, merchantId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取职工列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取职工列表成功", "data": staffList, "page": page})
	return
}

type AddStaffReq struct {
	LoginToken string `form:"login_token" valid:"required~职工账号不能为空"`
	StaffName  string `form:"staff_name" valid:"required~职工名称"`
	Password   string `form:"password" valid:"required~职工密码不能为空"`
	RoleId     string `form:"role_id" valid:"required~职工角色不能为空"`
}

// @Summary 新增职工
// @tags Staff
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param login_token formData string true "职工账号"
// @Param staff_name formData string true "职工名称"
// @Param password formData string true "职工密码"
// @Param role_id formData string true "职工角色ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /staff/add [post]
func AddStaff(ctx *gin.Context) {
	var code int
	var req AddStaffReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("add staff should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("add staff valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	staffService := staff.NewStaffService()
	merchantId := ff_header.NewFFHeader(ctx).GetRequestMerchantSession().MerchantId
	//新增职工信息
	if !staffService.AddStaff(req.LoginToken, req.StaffName, req.Password, req.RoleId, merchantId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "新增职工失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "新增职工成功", "data": ""})
	return
}

type UpdStaffReq struct {
	StaffId    string `form:"staff_id" valid:"required~职工ID不能为空"`
	StaffName  string `form:"staff_name" valid:"required~职工名称不能为空"`
	LoginToken string `form:"login_token" valid:"required~职工账号不能为空"`
	RoleId     string `form:"role_id" valid:"required~职工角色不能为空"`
}

// @Summary 更新职工
// @tags Staff
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param staff_id formData string true "职工ID"
// @Param login_token formData string true "职工账号"
// @Param staff_name formData string true "职工名称"
// @Param role_id formData string true "职工角色ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /staff/upd [put]
func UpdStaff(ctx *gin.Context) {
	var code int
	var req UpdStaffReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("upd staff should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("upd staff valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	staffService := staff.NewStaffService()
	if !staffService.CheckIsExitsStaffId(req.StaffId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "职工不存在", "data": ""})
		return
	}

	if !staffService.UpdStaff(req.StaffId, req.StaffName, req.LoginToken, req.RoleId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "更新职工失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "更新职工成功", "data": ""})
	return
}

type PwdStaffReq struct {
	StaffId     string `form:"staff_id" valid:"required~职工ID不能为空"`
	NewPassword string `form:"new_password" valid:"required~新密码不能为空"`
	OldPassword string `form:"old_password" valid:"required~原密码不能为空"`
}

// @Summary 更新职工密码
// @tags Staff
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param staff_id formData string true "职工ID"
// @Param new_password formData string true "新密码"
// @Param old_password formData string true "原密码"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /staff/pwd [put]
func PwdStaff(ctx *gin.Context) {
	var code int
	var req PwdStaffReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("upd staff should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("upd staff valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	staffService := staff.NewStaffService()
	if !staffService.CheckIsExitsStaffId(req.StaffId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "职工不存在", "data": ""})
		return
	}

	if !staffService.CheckIsCorrectPassword(req.StaffId, req.OldPassword) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "原密码不正确", "data": ""})
		return
	}

	if !staffService.PwdStaff(req.StaffId, req.NewPassword, req.OldPassword) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "更新职工失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "更新职工成功", "data": ""})
	return
}

type DelStaffReq struct {
	StaffId string `form:"staff_id" valid:"required~职工ID不能为空"`
}

// @Summary 删除职工
// @tags Staff
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param staff_id formData string true "职工ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /staff/del [delete]
func DelStaff(ctx *gin.Context) {
	var code int
	var req DelStaffReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("delete staff should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("delete staff valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	staffService := staff.NewStaffService()
	if !staffService.CheckIsExitsStaffId(req.StaffId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "职工不存在", "data": ""})
		return
	}

	if !staffService.DelStaff(req.StaffId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "删除职工失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "删除职工成功", "data": ""})
	return
}
