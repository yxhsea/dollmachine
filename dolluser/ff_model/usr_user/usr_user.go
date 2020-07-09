package usr_user

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

type UsrUser struct {
}

func NewUsrUser() *UsrUser {
	return &UsrUser{}
}

func (u *UsrUser) getTableName() string {
	return "usr_user"
}

func (u *UsrUser) CheckIsExitsByUserId(userId int64) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(u.getTableName()).Where("user_id", userId).Count(1)
	logrus.Debugf("Check user id is exits through table usr_user. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Check user id is exits failure. Error : %v ", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (u *UsrUser) GetUsrUserOne(userId int64, fields string) map[string]interface{} {
	dbr := ff_vars.DbConn.GetInstance()
	usrUserOne, err := dbr.Table(u.getTableName()).Fields(fields).Where("user_id", userId).First()
	logrus.Debugf("Query usr_user LastSql : %v ", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query usr_user failure. Error : %v", err)
		return nil
	}
	return usrUserOne
}
