package plf_login

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollplatform/ff_config/ff_vars"
)

type PlfLogin struct {
}

func NewPlfLogin() *PlfLogin {
	return &PlfLogin{}
}

func (p *PlfLogin) getTableName() string {
	return "plf_login"
}

func (p *PlfLogin) GetLoginByLoginToken(loginToken string, fields string) (map[string]interface{}, error){
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("login_token", loginToken).First()
	logrus.Debugf("GetLoginByLoginToken. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetLoginByLoginToken. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *PlfLogin) CheckIsExitsByLoginToken(loginToken string) bool {
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

