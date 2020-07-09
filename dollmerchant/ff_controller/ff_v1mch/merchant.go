package ff_v1mch

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_common/ff_err"
	"dollmachine/dollmerchant/ff_service/mch"
	"net/http"
)

type GetMchInfoReq struct {
	MerchantId string `form:"merchant_id" valid:"required~商户名Id不能为空,numeric~商户ID应为数字"`
}

// @Summary 获取商户信息
// @tags Merchant
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
