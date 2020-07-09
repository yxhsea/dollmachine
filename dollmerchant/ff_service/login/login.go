package login

import (
	"fmt"
	"github.com/Unknwon/com"
	"dollmachine/dollmerchant/ff_cache/merchant_session"
	"dollmachine/dollmerchant/ff_common/ff_password"
	"dollmachine/dollmerchant/ff_model/mch_login"
	"dollmachine/dollmerchant/ff_model/mch_role"
	"dollmachine/dollmerchant/ff_model/mch_staff"
)

type LoginService struct {
}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (p *LoginService) CheckIsExitsByLoginToken(loginToken string) bool {
	mchLoginDao := mch_login.NewMchLogin()
	if !mchLoginDao.CheckIsExitsByLoginToken(loginToken) {
		return false
	}
	return true
}

func (p *LoginService) CheckPassword(loginToken, passWord string) bool {
	mchLoginDao := mch_login.NewMchLogin()
	loginInfo, err := mchLoginDao.GetLoginByLoginToken(loginToken, "login_secret, login_salt, merchant_id")
	if err != nil {
		return false
	}
	loginSecret, _ := loginInfo["login_secret"]
	loginSalt, _ := loginInfo["login_salt"]

	if !ff_password.Check(fmt.Sprint(loginSecret), fmt.Sprint(loginSalt), passWord) {
		return false
	}
	return true
}

func (p *LoginService) SignIn(loginToken, passWord string) (*merchant_session.MerchantSession, bool) {
	mchLoginDao := mch_login.NewMchLogin()
	loginInfo, err := mchLoginDao.GetLoginByLoginToken(loginToken, "merchant_id, staff_id")
	merchantId, _ := loginInfo["merchant_id"]
	merchantIdInt64, _ := com.StrTo(fmt.Sprint(merchantId)).Int64()
	staffId, _ := loginInfo["staff_id"]
	staffIdInt64, _ := com.StrTo(fmt.Sprint(staffId)).Int64()

	mchStaffDao := mch_staff.NewMchStaff()
	staffInfo, err := mchStaffDao.GetMchStaffByStaffId(staffIdInt64, "nick_name, role_id, phone")
	if err != nil {
		return nil, false
	}
	staffName, _ := staffInfo["nick_name"]
	staffPhone, _ := staffInfo["phone"]
	roleId, _ := staffInfo["role_id"]
	roleIdInt64, _ := com.StrTo(fmt.Sprint(roleId)).Int64()

	mchRoleDao := mch_role.NewMchRole()
	roleInfo, err := mchRoleDao.GetMchRoleByRoleId(roleIdInt64, "rules")
	if err != nil {
		return nil, false
	}
	rules, _ := roleInfo["rules"]

	mchUserSession := &merchant_session.MerchantSession{
		MerchantId:   merchantIdInt64,
		MerchantName: "",
		StaffId:      staffIdInt64,
		StaffName:    fmt.Sprint(staffName),
		StaffPhone:   fmt.Sprint(staffPhone),
		RoleId:       roleIdInt64,
		LoginToken:   loginToken,
		Rules:        fmt.Sprint(rules),
	}

	mchSessionCache := merchant_session.NewMerchantSession()
	flag := mchSessionCache.SetMerchantSession(mchUserSession)
	if !flag {
		return nil, false
	}

	return mchUserSession, true
}

func (p *LoginService) SignOut(token string) bool {
	mchSessionCache := merchant_session.NewMerchantSession()
	if !mchSessionCache.CheckIsExitsByToken(token) {
		return true
	}
	if !mchSessionCache.DeleteToken(token) {
		return false
	}
	return true
}
