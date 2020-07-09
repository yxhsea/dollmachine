package mch_login

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type MchLogin struct {
}

func NewMchLogin() *MchLogin {
	return &MchLogin{}
}

func (p *MchLogin) getTableName() string {
	return "mch_login"
}

func (p *MchLogin) GetLoginByLoginToken(loginToken string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("login_token", loginToken).First()
	logrus.Debugf("GetLoginByLoginToken. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetLoginByLoginToken. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchLogin) GetLoginByStaffId(staffId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("staff_id", staffId).First()
	logrus.Debugf("GetLoginByStaffId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetLoginByStaffId. Error : %v", err)
		return nil, err
	}
	return one, nil
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

func (p *MchLogin) CheckIsExitsByLoginToken(loginToken string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("login_token", loginToken).Count(1)
	logrus.Debugf("Check login token is exits. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Check login token is exits failure. Error : %v ", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
