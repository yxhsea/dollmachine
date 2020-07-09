package login

import (
	"fmt"
	"github.com/Unknwon/com"
	"dollmachine/dollplatform/ff_cache/platform_session"
	"dollmachine/dollplatform/ff_common/ff_password"
	"dollmachine/dollplatform/ff_model/plf_login"
	"dollmachine/dollplatform/ff_model/plf_manager"
	"dollmachine/dollplatform/ff_model/plf_role"
)

type LoginService struct {
}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (p *LoginService) CheckIsExitsByLoginToken(loginToken string) bool {
	plfLoginDao := plf_login.NewPlfLogin()
	if !plfLoginDao.CheckIsExitsByLoginToken(loginToken) {
		return false
	}
	return true
}

func (p *LoginService) CheckPassword(loginToken, passWord string) bool {
	plfLoginDao := plf_login.NewPlfLogin()
	loginInfo, err := plfLoginDao.GetLoginByLoginToken(loginToken, "login_secret, login_salt, manager_id")
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

func (p *LoginService) SignIn(loginToken, passWord string) (*platform_session.PlatformSession, bool) {
	plfLoginDao := plf_login.NewPlfLogin()
	loginInfo, err := plfLoginDao.GetLoginByLoginToken(loginToken, "manager_id")
	managerId, _ := loginInfo["manager_id"]
	managerIdInt64, _ := com.StrTo(fmt.Sprint(managerId)).Int64()

	plfManagerDao := plf_manager.NewPlfManager()
	managerInfo, err := plfManagerDao.GetPlfManagerByManagerId(managerIdInt64, "nick_name, role_id")
	if err != nil {
		return nil, false
	}
	nickName, _ := managerInfo["nick_name"]
	roleId, _ := managerInfo["role_id"]
	roleIdInt64, _ := com.StrTo(fmt.Sprint(roleId)).Int64()

	plfRoleDao := plf_role.NewPlfRole()
	roleInfo, err := plfRoleDao.GetPlfRoleByRoleId(roleIdInt64, "rules")
	if err != nil {
		return nil, false
	}
	rules, _ := roleInfo["rules"]

	plfUserSession := &platform_session.PlatformSession{
		ManagerId:  managerIdInt64,
		RoleId:     roleIdInt64,
		NickName:   fmt.Sprint(nickName),
		LoginToken: loginToken,
		Rules:      fmt.Sprint(rules),
	}

	plfSessionCache := platform_session.NewPlatformSession()
	flag := plfSessionCache.SetPlatformSession(plfUserSession)
	if !flag {
		return nil, false
	}

	return plfUserSession, true
}

func (p *LoginService) SignOut(token string) bool {
	plfSessionCache := platform_session.NewPlatformSession()
	if !plfSessionCache.CheckIsExitsByToken(token) {
		return true
	}
	if !plfSessionCache.DeleteToken(token) {
		return false
	}
	return true
}
