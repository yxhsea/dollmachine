package mch_merchant

import (
	"github.com/gogap/logrus"
	"dollmachine/dollplatform/ff_config/ff_vars"
)

type MchMerchant struct {
}

func NewMchMerchant() *MchMerchant {
	return &MchMerchant{}
}

func (p *MchMerchant) getTableName() string {
	return "mch_merchant"
}

func (p *MchMerchant) GetMchList(offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetMchList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *MchMerchant) GetMchListTotalCount() (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Count(1)
	logrus.Debugf("GetMchListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *MchMerchant) GetMchInfo(merchantId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("merchant_id", merchantId).First()
	logrus.Debugf("GetMchInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchInfo. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchMerchant) CheckIsExitsMchName(name string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("name", name).Count(1)
	logrus.Debugf("CheckIsExitsMchName. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsMchName. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *MchMerchant) CheckIsExitsMchId(merchantId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("merchant_id", merchantId).Count(1)
	logrus.Debugf("CheckIsExitsMchId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsMchId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
