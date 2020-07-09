package plf_role

import (
	"dollmachine/dollplatform/ff_config/ff_vars"
	"github.com/sirupsen/logrus"
)

type PlfRole struct {
}

func NewPlfRole() *PlfRole {
	return &PlfRole{}
}

func (p *PlfRole) getTableName() string {
	return "plf_role"
}

func (p *PlfRole) GetPlfRoleByRoleId(roleId int64, fields string) (map[string]interface{}, error){
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("role_id", roleId).First()
	logrus.Debugf("GetPlfRoleByRoleId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPlfRoleByRoleId. Error : %v", err)
		return nil, err
	}
	return one, nil
}