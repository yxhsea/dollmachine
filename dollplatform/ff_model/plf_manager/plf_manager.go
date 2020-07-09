package plf_manager

import (
	"dollmachine/dollplatform/ff_config/ff_vars"
	"github.com/sirupsen/logrus"
)

type PlfManager struct {
}

func NewPlfManager() *PlfManager {
	return &PlfManager{}
}

func (p *PlfManager) getTableName() string {
	return "plf_manager"
}

func (p *PlfManager) GetPlfManagerByManagerId(managerId int64, fields string) (map[string]interface{}, error){
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("manager_id", managerId).First()
	logrus.Debugf("GetPlfManagerByManagerId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPlfManagerByManagerId. Error : %v", err)
		return nil, err
	}
	return one, nil
}