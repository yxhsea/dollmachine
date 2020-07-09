package mch_room_device

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

type MchRoomDevice struct {
}

func NewMchRoomDevice() *MchRoomDevice {
	return &MchRoomDevice{}
}

func (p *MchRoomDevice) getTableName() string {
	return "mch_room_device"
}

func (p *MchRoomDevice) CheckIsExitsByRoomId(roomId int64) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("room_id", roomId).Count(1)
	logrus.Debugf("CheckIsExitsByRoomId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsByRoomId. Error : %v ", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *MchRoomDevice) GetMchRoomDeviceInfoByRoomId(roomId int64, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("room_id", roomId).First()
	logrus.Debugf("GetMchRoomDeviceInfoByRoomId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchRoomDeviceInfoByRoomId. Error : %v", err)
		return nil, err
	}
	return one, nil
}
