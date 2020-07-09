package mch_room_device

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type MchRoomDevice struct {
}

func NewMchRoomDevice() *MchRoomDevice {
	return &MchRoomDevice{}
}

func (p *MchRoomDevice) getTableName() string {
	return "mch_room_device"
}

func (p *MchRoomDevice) GetMchRoomDeviceList(offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetMchRoomDeviceList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchRoomDeviceList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *MchRoomDevice) GetMchRoomDeviceListTotalCount() (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Count(1)
	logrus.Debugf("GetMchRoomDeviceListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchRoomDeviceListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *MchRoomDevice) GetMchRoomDeviceInfo(DeviceId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("device_id", DeviceId).First()
	logrus.Debugf("GetMchRoomDeviceInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchRoomDeviceInfo. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchRoomDevice) CheckIsExitsDeviceName(name string) bool {
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

func (p *MchRoomDevice) CheckIsExitsDeviceId(DeviceId string) bool {
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
func (p *MchRoomDevice) CheckIsBindDeviceId(deviceId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("device_id", deviceId).Count(1)
	logrus.Debugf("CheckIsBindDeviceId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsBindDeviceId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
