package usr_login

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

type UsrLogin struct {
}

func NewUsrLogin() *UsrLogin {
	return &UsrLogin{}
}

func (u *UsrLogin) getTableName() string {
	return "usr_login"
}

func (u *UsrLogin) CheckIsExitsByUserId(userId int64) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(u.getTableName()).Where("user_id", userId).Count(1)
	logrus.Debugf("Check user id is exits. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Check userId is exits failure. Error : %v ", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (u *UsrLogin) GetUsrLoginOne(userId int64, fields string) map[string]interface{} {
	dbr := ff_vars.DbConn.GetInstance()
	usrLoginOne, err := dbr.Table(u.getTableName()).Fields(fields).Where("user_id", userId).Limit(1).First()
	logrus.Debugf("Query usr_login LastSql : %v ", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query usr_login failure. Error : %v", err)
		return nil
	}
	return usrLoginOne
}
