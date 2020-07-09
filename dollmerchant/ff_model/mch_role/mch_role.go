package mch_role

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type MchRole struct {
}

func NewMchRole() *MchRole {
	return &MchRole{}
}

func (p *MchRole) getTableName() string {
	return "mch_role"
}

func (p *MchRole) GetMchRoleByRoleId(roleId int64, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("role_id", roleId).First()
	logrus.Debugf("GetMchRoleByRoleId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchRoleByRoleId. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchRole) CheckIsExitsRoleId(roleId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("role_id", roleId).Count(1)
	logrus.Debugf("CheckIsExitsRoleId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsRoleId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
