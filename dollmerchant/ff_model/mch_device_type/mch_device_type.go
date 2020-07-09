package mch_device_type

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type MchDeviceType struct {
}

func NewMchDeviceType() *MchDeviceType {
	return &MchDeviceType{}
}

func (p *MchDeviceType) getTableName() string {
	return "mch_device_type"
}

func (p *MchDeviceType) GetMchDeviceTypeList(offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetMchDeviceTypeList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchDeviceTypeList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *MchDeviceType) GetMchDeviceTypeListTotalCount() (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Count(1)
	logrus.Debugf("GetMchDeviceTypeListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchDeviceTypeListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *MchDeviceType) GetMchDeviceTypeInfo(deviceTypeId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("device_type_id", deviceTypeId).First()
	logrus.Debugf("GetMchDeviceTypeInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchDeviceTypeInfo. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchDeviceType) CheckIsExitsDeviceTypeName(name string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("name", name).Count(1)
	logrus.Debugf("CheckIsExitsDeviceTypeName. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsDeviceTypeName. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *MchDeviceType) CheckIsExitsDeviceTypeId(deviceTypeId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("device_type_id", deviceTypeId).Count(1)
	logrus.Debugf("CheckIsExitsDeviceTypeId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsDeviceTypeId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
