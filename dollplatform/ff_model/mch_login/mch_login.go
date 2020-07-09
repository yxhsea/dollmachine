package mch_login

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollplatform/ff_config/ff_vars"
)

type MchLogin struct {
}

func NewMchLogin() *MchLogin {
	return &MchLogin{}
}

func (p *MchLogin) getTableName() string {
	return "mch_login"
}

func (p *MchLogin) CheckIsExitsMchLogin(loginToken string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("login_token", loginToken).Count(1)
	logrus.Debugf("CheckIsExitsMchLogin. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsMchLogin. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
