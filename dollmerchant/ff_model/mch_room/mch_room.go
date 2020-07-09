package mch_room

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type MchRoom struct {
}

func NewMchRoom() *MchRoom {
	return &MchRoom{}
}

func (p *MchRoom) getTableName() string {
	return "mch_room"
}

func (p *MchRoom) GetMchRoomList(offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetMchRoomList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchRoomList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *MchRoom) GetMchRoomListTotalCount() (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Count(1)
	logrus.Debugf("GetMchRoomListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchRoomListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *MchRoom) GetMchRoomInfo(RoomId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("room_id", RoomId).First()
	logrus.Debugf("GetMchRoomInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchRoomInfo. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchRoom) CheckIsExitsRoomName(name string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("name", name).Count(1)
	logrus.Debugf("CheckIsExitsRoomName. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsRoomName. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *MchRoom) CheckIsExitsRoomId(RoomId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("room_id", RoomId).Count(1)
	logrus.Debugf("CheckIsExitsRoomId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsRoomId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
