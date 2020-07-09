package mch_device

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

type MchDevice struct {
}

func NewMchDevice() *MchDevice {
	return &MchDevice{}
}

func (p *MchDevice) getTableName() string {
	return "mch_device"
}

func (p *MchDevice) CheckIsExitsByDeviceId(deviceId int64) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("device_id", deviceId).Count(1)
	logrus.Debugf("CheckIsExitsByDeviceId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsByDeviceId. Error : %v ", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *MchDevice) GetMchDeviceInfoByDeviceId(deviceId int64, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("device_id", deviceId).First()
	logrus.Debugf("GetMchDeviceInfoByDeviceId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchDeviceInfoByDeviceId. Error : %v", err)
		return nil, err
	}
	return one, nil
}
