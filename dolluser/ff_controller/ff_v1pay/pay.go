package ff_v1pay

import (
	"github.com/Unknwon/com"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_common/ff_err"
	"dollmachine/dolluser/ff_common/ff_header"
	"dollmachine/dolluser/ff_service/trade"
	"io/ioutil"
	"net/http"
)

type RechargeQueryReq struct {
	RechargeId string `form:"recharge_id" valid:"required~交易ID不能为空,numeric~交易ID必须是数字"`
}

// @Summary 充值结果查询
// @tags pay
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param recharge_id query string true "交易ID"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /pay/recharge/query [get]
func QueryRecharge(ctx *gin.Context) {
	var code int
	var req RechargeQueryReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("QueryRecharge should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("QueryRecharge valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	tradeService := trade.NewTradeService()
	response, err := tradeService.QueryRecharge(ff_header.NewFFHeader(ctx).GetRequestUserSession(), req.RechargeId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "查询失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": response})
	return
}

type ExpressQueryReq struct {
	RechargeId string `form:"recharge_id" valid:"required~交易ID不能为空,numeric~交易ID必须是数字"`
}

// @Summary 邮费结果查询
// @tags pay
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param recharge_id query string true "交易ID"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /pay/express/query [get]
func QueryExpress(ctx *gin.Context) {
	var code int
	var req ExpressQueryReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("QueryExpress should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("QueryExpress valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	tradeService := trade.NewTradeService()
	response, err := tradeService.QueryRechargeExpress(ff_header.NewFFHeader(ctx).GetRequestUserSession(), req.RechargeId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "查询失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": response})
	return
}

type CreateRechargeReq struct {
	MerchantId string `form:"merchant_id" valid:"required~商户ID不能为空,numeric~商户ID必须是数字"`
	DeviceId   string `form:"device_id" valid:"required~设备ID不能为空,numeric~设备ID必须是数字"`
	Amount     string `form:"amount" valid:"required~金额不能为空"`
	Coin       string `form:"coin" valid:"required~金币不能为空"`
	FromTag    string `form:"from_tag" valid:"-"`
	TradeType  string `form:"trade_type" valid:"required~交易类型"`
}

// @Summary 创建充值订单
// @tags pay
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param merchant_id formData string true "商户ID"
// @Param device_id formData string true "设备ID"
// @Param amount formData string true "金额"
// @Param trade_type formData int true "交易类型 1|公众号 2|扫码"
// @Param from_tag formData string true "标签"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /pay/recharge/create [post]
func CreateRecharge(ctx *gin.Context) {
	var code int
	var req CreateRechargeReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("CreateRecharge should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("CreateRecharge valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	amount, _ := com.StrTo(req.Amount).Float64()
	coin, _ := com.StrTo(req.Coin).Int()
	deviceId, _ := com.StrTo(req.DeviceId).Int64()
	merchantId, _ := com.StrTo(req.MerchantId).Int64()
	tradeType, _ := com.StrTo(req.TradeType).Int()
	clientIp := ctx.Request.RemoteAddr
	tradeService := trade.NewTradeService()
	response, err := tradeService.AddRecharge(ff_header.NewFFHeader(ctx).GetRequestUserSession(), amount, coin, deviceId, clientIp, merchantId, req.FromTag, tradeType)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "创建订单失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": response})
	return
}

type CreateExpressReq struct {
	RecordId  string `form:"record_id" valid:"required~兑换记录ID不能为空,numeric~兑换记录ID必须是数字"`
	Amount    string `form:"amount" valid:"required~金额不能为空"`
	TradeType string `form:"trade_type" valid:"required~交易类型"`
}

// @Summary 创建邮费订单
// @tags pay
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param record_id formData string true "兑换记录ID"
// @Param amount formData string true "金额"
// @Param trade_type formData int true "交易类型 1|公众号 2|扫码"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /pay/express/create [post]
func CreateExpress(ctx *gin.Context) {
	var code int
	var req CreateExpressReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("CreateExpress should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&req)
	if err != nil {
		logrus.Errorf("CreateExpress valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	recordId, _ := com.StrTo(req.RecordId).Int64()
	amount, _ := com.StrTo(req.Amount).Float64()
	tradeType, _ := com.StrTo(req.TradeType).Int()
	clientIp := ctx.Request.RemoteAddr
	tradeService := trade.NewTradeService()
	response, err := tradeService.AddExpressRecharge(ff_header.NewFFHeader(ctx).GetRequestUserSession(), amount, clientIp, tradeType, recordId)
	if err != nil {
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": "创建订单失败", "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": response})
	return
}

//充值支付异步通知
func NotifyRecharge(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logrus.Error("Notify recharge read request body failure.")
		return
	}
	notifyStr := string(body)
	tradeService := trade.NewTradeService()
	_, err = tradeService.NotifyRecharge(notifyStr)
	if err != nil {
		logrus.Error("Notify recharge failure.")
	}
	return
}

//邮费支付异步通知
func NotifyExpress(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logrus.Error("Notify express read request body failure.")
		return
	}
	notifyStr := string(body)
	tradeService := trade.NewTradeService()
	_, err = tradeService.NotifyExpress(notifyStr)
	if err != nil {
		logrus.Error("Notify express failure.")
	}

	return
}
