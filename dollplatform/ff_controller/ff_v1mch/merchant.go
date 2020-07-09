package ff_v1mch

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dollplatform/ff_common/ff_err"
	"dollmachine/dollplatform/ff_service/mch"
	"net/http"
)

type GetMchInfoReq struct {
	MerchantId string `form:"merchant_id" valid:"required~商户名Id不能为空,numeric~商户ID应为数字"`
}

// @Summary 获取商户信息
// @tags merchant
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param merchant_id query string true "商户ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /mch/info [get]
func GetMchInfo(ctx *gin.Context) {
	var code int
	var req GetMchInfoReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get mch info should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get mch info valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	mchService := mch.NewMchService()
	if !mchService.CheckIsExitsMchId(req.MerchantId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "商户不存在", "data": ""})
		return
	}

	mchInfo, err := mchService.GetMchInfo(req.MerchantId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取商户信息失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取商户信息成功", "data": mchInfo})
	return
}

type GetMchListReq struct {
	Offset    string `form:"offset" valid:"required~offset不能为空,numeric~offset必须是数字"`
	TotalSize string `form:"total_size" valid:"required~total_size不能为空,numeric~total_size必须是数字"`
	PageSize  string `form:"page_size" valid:"required~page_size不能为空,numeric~page_size必须是数字"`
}

// @Summary 获取商户列表
// @tags merchant
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param offset query string true "偏移量"
// @Param total_size query string true "总条数"
// @Param page_size query string true "一页多少条数据"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /mch/list [get]
func GetMchList(ctx *gin.Context) {
	var code int
	var req GetMchListReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("get list info should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("get mch info valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	mchService := mch.NewMchService()
	mchList, page, err := mchService.GetMchList(req.Offset, req.PageSize, req.TotalSize)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "获取商户列表失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "获取商户列表成功", "data": mchList, "page": page})
	return
}

type AddMchReq struct {
	MchName             string `form:"mch_name" valid:"required~商户名称不能为空"`
	Province            string `form:"province" valid:"required~省份不能为空"`
	City                string `form:"city" valid:"required~城市不能为空"`
	District            string `form:"district" valid:"required~地区不能为空"`
	Address             string `form:"address" valid:"required~详细地址不能为空"`
	RechargeRate        string `form:"recharge_rate" valid:"required~充值分成比例不能为空"`
	IntegralConvertRate string `form:"integral_convert_rate" valid:"required~积分兑换比例不能为空"`
	IntegralRate        string `form:"integral_rate" valid:"required~积分分成比例"`
	ContactName         string `form:"contact_name" valid:"required~商户对接人姓名不能为空"`
	ContactPhone        string `form:"contact_phone" valid:"required~商户对接人联系电话不能为空"`
}

// @Summary 新增商户
// @tags merchant
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param mch_name formData string true "商户名称"
// @Param province formData string true "省份或直辖市"
// @Param city formData string true "城市"
// @Param district formData string true "地区"
// @Param address formData string true "详细地址"
// @Param recharge_rate formData string true "充值分成比例"
// @Param integral_convert_rate formData string true "积分兑换比例"
// @Param integral_rate formData string true "积分分成比例"
// @Param contact_name formData string true "商户对接人姓名"
// @Param contact_phone formData string true "商户对接人联系电话"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /mch/add [post]
func AddMch(ctx *gin.Context) {
	var code int
	var req AddMchReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("add mch should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("add mch valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	mchService := mch.NewMchService()
	if mchService.CheckIsExitsMchLogin(req.ContactPhone) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "手机号已经存在", "data": ""})
		return
	}

	if mchService.CheckIsExitsMchName(req.MchName) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "商户名称已经存在", "data": ""})
		return
	}

	//新增商户信息
	if !mchService.AddMch(req.MchName, req.Province, req.City, req.District, req.Address, req.RechargeRate,
		req.IntegralConvertRate, req.IntegralRate, req.ContactName, req.ContactPhone) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "新增商户失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "新增商户成功", "data": ""})
	return
}

type UpdMchReq struct {
	MerchantId          string `form:"merchant_id" valid:"required~商户ID不能为空"`
	MchName             string `form:"mch_name" valid:"required~商户名称不能为空"`
	Province            string `form:"province" valid:"required~省份不能为空"`
	City                string `form:"city" valid:"required~城市不能为空"`
	District            string `form:"district" valid:"required~地区不能为空"`
	Address             string `form:"address" valid:"required~详细地址不能为空"`
	RechargeRate        string `form:"recharge_rate" valid:"required~充值分成比例不能为空"`
	IntegralConvertRate string `form:"integral_convert_rate" valid:"required~积分兑换比例不能为空"`
	IntegralRate        string `form:"integral_rate" valid:"required~积分分成比例"`
	ContactName         string `form:"contact_name" valid:"required~商户对接人姓名不能为空"`
	ContactPhone        string `form:"contact_phone" valid:"required~商户对接人联系电话不能为空"`
}

// @Summary 更新商户
// @tags merchant
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param merchant_id formData string true "商户ID"
// @Param mch_name formData string true "商户名称"
// @Param province formData string true "省份或直辖市"
// @Param city formData string true "城市"
// @Param district formData string true "地区"
// @Param address formData string true "详细地址"
// @Param recharge_rate formData string true "充值分成比例"
// @Param integral_convert_rate formData string true "积分兑换比例"
// @Param integral_rate formData string true "积分分成比例"
// @Param contact_name formData string true "商户对接人姓名"
// @Param contact_phone formData string true "商户对接人联系电话"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /mch/upd [put]
func UpdMch(ctx *gin.Context) {
	var code int
	var req UpdMchReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("upd mch should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("upd mch valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	mchService := mch.NewMchService()
	if !mchService.CheckIsExitsMchId(req.MerchantId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "商户不存在", "data": ""})
		return
	}

	if !mchService.UpdMch(req.MerchantId, req.MchName, req.Province, req.City, req.District, req.Address, req.RechargeRate,
		req.IntegralConvertRate, req.IntegralRate, req.ContactName, req.ContactPhone) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "更新商户失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "更新商户成功", "data": ""})
	return
}

type UpdMchPwdReq struct {
	MerchantId string `form:"merchant_id" valid:"required~商户ID不能为空"`
	Password   string `form:"password" valid:"required~密码不能为空"`
}

// @Summary 更新商户密码
// @tags merchant
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param merchant_id formData string true "商户ID"
// @Param password formData string true "登录密码"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /mch/pwd [put]
func UpdMchPwd(ctx *gin.Context) {
	var code int
	var req UpdMchPwdReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("upd mch pwd should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("upd mch pwd valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	mchService := mch.NewMchService()
	if !mchService.CheckIsExitsMchId(req.MerchantId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "商户不存在", "data": ""})
		return
	}

	if !mchService.UpdMchPwd(req.MerchantId, req.Password) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "更新商户密码失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "更新商户密码成功", "data": ""})
	return
}

type UpdMchStateReq struct {
	MerchantId string `form:"merchant_id" valid:"required~商户ID不能为空"`
	Status     string `form:"status" valid:"required~状态不能为空"`
}

// @Summary 更新商户状态
// @tags merchant
// @Accept  application/x-www-form-urlencoded
// @Produce  json
// @Param Token header string true "Token"
// @Param merchant_id formData string true "商户ID"
// @Param status formData string true "状态 0|禁用 1|正常"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /mch/state [put]
func UpdMchState(ctx *gin.Context) {
	var code int
	var req UpdMchStateReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("upd mch pwd should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("upd mch pwd valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	mchService := mch.NewMchService()
	if !mchService.CheckIsExitsMchId(req.MerchantId) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "商户不存在", "data": ""})
		return
	}

	if !mchService.UpdMchState(req.MerchantId, req.Status) {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "更新商户状态失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": "更新商户状态成功", "data": ""})
	return
}
