package user

import (
	"fmt"
	"dollmachine/dolluser/ff_cache/user_session"
	"dollmachine/dolluser/ff_model/coin_wallet"
	"dollmachine/dolluser/ff_model/usr_address"
	"dollmachine/dolluser/ff_model/usr_login"
	"dollmachine/dolluser/ff_model/usr_user"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (p *UserService) CheckIsExitsByUserId(userId int64) bool {
	usrLogin := usr_login.NewUsrLogin()
	usrUser := usr_user.NewUsrUser()
	if !usrLogin.CheckIsExitsByUserId(userId) || !usrUser.CheckIsExitsByUserId(userId) {
		return false
	}
	return true
}

func (p *UserService) UserLogin(userId, merchantId int64) (*user_session.UserSession, bool) {
	usrLogin := usr_login.NewUsrLogin()
	usrUser := usr_user.NewUsrUser()

	usrLoginOne := usrLogin.GetUsrLoginOne(userId, "login_token")
	openId, _ := usrLoginOne["login_token"]

	usrUserOne := usrUser.GetUsrUserOne(userId, "nick_name")
	nickName, _ := usrUserOne["nick_name"]

	userSession := &user_session.UserSession{
		UserId:     userId,
		OpenId:     fmt.Sprint(openId),
		NickName:   fmt.Sprint(nickName),
		MerchantId: merchantId,
	}
	userSessionCache := user_session.NewUserSession()
	flag := userSessionCache.SetUserSession(userSession)
	if !flag {
		return nil, false
	}

	return userSession, true
}

func (p *UserService) GetUserInfo(userId int64) map[string]interface{} {
	//用户钱包
	coinWallet := coin_wallet.NewCoinWallet()
	coinWalletOne := coinWallet.GetCoinWalletOneByUserId(userId, "coin,integral")
	coin, _ := coinWalletOne["coin"]
	integral, _ := coinWalletOne["integral"]

	//用户信息
	usrUser := usr_user.NewUsrUser()
	usrUserOne := usrUser.GetUsrUserOne(userId, "nick_name,avatar")
	nickName, _ := usrUserOne["nick_name"]
	avatar, _ := usrUserOne["avatar"]

	//用户login信息
	usrLogin := usr_login.NewUsrLogin()
	usrLoginOne := usrLogin.GetUsrLoginOne(userId, "login_token")
	openId, _ := usrLoginOne["login_token"]

	//用户地址信息
	usrAddress := usr_address.NewUsrAddress()
	usrAddressOne, _ := usrAddress.GetUserLastAddressByUserId(userId, "full_name,phone,province,city,district,address")
	fullName, _ := usrAddressOne["full_name"]
	phone, _ := usrAddressOne["phone"]
	province, _ := usrAddressOne["province"]
	city, _ := usrAddressOne["city"]
	district, _ := usrAddressOne["district"]
	address, _ := usrAddressOne["address"]

	userInfo := map[string]interface{}{
		"user_id":   userId,
		"coin":      coin,
		"integral":  integral,
		"nick_name": nickName,
		"avatar":    avatar,
		"open_id":   openId,
		"full_name": fullName,
		"phone":     phone,
		"province":  province,
		"city":      city,
		"district":  district,
		"address":   address,
	}
	return userInfo
}
