package mch_device

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type MchDevice struct {
}

func NewMchDevice() *MchDevice {
	return &MchDevice{}
}

func (p *MchDevice) getTableName() string {
	return "mch_device"
}

func (p *MchDevice) GetMchDeviceList(offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetMchDeviceList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchDeviceList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *MchDevice) GetMchDeviceListTotalCount() (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Count(1)
	logrus.Debugf("GetMchDeviceListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchDeviceListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *MchDevice) GetMchDeviceInfo(DeviceId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("device_id", DeviceId).First()
	logrus.Debugf("GetMchDeviceInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchDeviceInfo. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchDevice) CheckIsExitsDeviceName(name string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("name", name).Count(1)
	logrus.Debugf("CheckIsExitsDeviceName. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsDeviceName. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *MchDevice) CheckIsExitsDeviceId(DeviceId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("device_id", DeviceId).Count(1)
	logrus.Debugf("CheckIsExitsDeviceId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsDeviceId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

//判断该设备是否被绑定
func (p *MchDevice) CheckIsBindDeviceMac(deviceMac string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("mac", deviceMac).Count(1)
	logrus.Debugf("CheckIsBindDeviceMac. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsBindDeviceMac. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

//判断该投放地点是否被绑定
func (p *MchDevice) CheckIsBindPlaceId(placeId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("place_id", placeId).Count(1)
	logrus.Debugf("CheckIsBindPlaceId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsBindPlaceId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
