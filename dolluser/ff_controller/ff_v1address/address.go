package ff_v1address

import (
	"github.com/Unknwon/com"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_common/ff_err"
	"dollmachine/dolluser/ff_common/ff_header"
	"dollmachine/dolluser/ff_service/address"
	"net/http"
)

type AddressInfo struct {
	AddressId string `form:"address_id" valid:"required~地址ID不能为空,numeric~地址ID必须是数字"`
}

// @Summary 获取地址信息
// @tags address
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param address_id query string true "地址ID"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /address/info [get]
func GetAddressInfo(ctx *gin.Context) {
	var code int
	var reqAddress AddressInfo
	err := ctx.ShouldBind(&reqAddress)
	if err != nil {
		logrus.Errorf("GetAddressInfo should bind params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	//校验Params参数
	_, err = govalidator.ValidateStruct(&reqAddress)
	if err != nil {
		logrus.Errorf("GetAddressInfo valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	//参数类型转换
	addressId, _ := com.StrTo(reqAddress.AddressId).Int64()

	//查询地址信息
	addressService := address.NewAddressService()
	addressInfo, err := addressService.GetAddressInfo(addressId)
	if err != nil {
		logrus.Errorf("Query Address information failure. Error : %v", err)
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": addressInfo})
	return
}

type Address struct {
	AddressId string `form:"address_id" valid:"-"`
	FullName  string `form:"full_name" valid:"required~姓名不能为空"`
	Phone     string `form:"phone" valid:"required~手机号不能为空,numeric~手机号必须是数字"`
	Province  string `form:"province" valid:"required~省份不能为空"`
	City      string `form:"city" valid:"required~城市不能为空"`
	District  string `form:"district" valid:"required~地区不能为空"`
	Address   string `form:"address" valid:"required~详细地址不能为空"`
}

// @Summary 添加地址
// @tags address
// @Produce  json
// @Param Token header string true "Token令牌"
// @Param address_id formData int false "地址ID"
// @Param full_name formData string true "姓名"
// @Param phone formData int true "手机号"
// @Param province formData string true "省份"
// @Param city formData string true "城市"
// @Param district formData string true "地区"
// @Param address formData string true "详细地址"
// @Success 200 {string} json "{"code":200,"data":"","msg":"ok"}"
// @Router /address/add [post]
func AddAddress(ctx *gin.Context) {
	var code int
	var addressInfo Address
	ctx.ShouldBind(&addressInfo)

	//校验Params参数
	_, err := govalidator.ValidateStruct(&addressInfo)
	if err != nil {
		logrus.Errorf("Add address valid params failure. Error : %v", err)
		code = ff_err.INVALID_PARAMS
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": err.Error(), "data": ""})
		return
	}

	//参数类型转换
	addressId, _ := com.StrTo(addressInfo.AddressId).Int64()
	userId := ff_header.NewFFHeader(ctx).GetRequestUserSession().UserId
	phone, _ := com.StrTo(addressInfo.Phone).Int64()

	addressService := address.NewAddressService()
	addrId, err := addressService.AddAddress(addressId, userId, phone, addressInfo.FullName,
		addressInfo.Province, addressInfo.City, addressInfo.District, addressInfo.Address)
	if err != nil {
		logrus.Errorf("Add Address information failure. Error : %v", err)
		code = ff_err.ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": ""})
		return
	}

	code = ff_err.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": ff_err.GetMsg(code), "data": map[string]interface{}{"address_id": addrId}})
	return
}
