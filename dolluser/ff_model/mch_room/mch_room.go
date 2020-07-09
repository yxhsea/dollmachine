package mch_room

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

type MchRoom struct {
}

func NewMchRoom() *MchRoom {
	return &MchRoom{}
}

func (p *MchRoom) getTableName() string {
	return "mch_room"
}

func (p *MchRoom) CheckIsExitsByRoomId(roomId int64) bool {
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

func (p *MchRoom) GetMchRoomInfoByRoomId(roomId int64, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("room_id", roomId).First()
	logrus.Debugf("GetMchRoomInfoByRoomId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchRoomInfoByRoomId. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchRoom) GetMchRoomListByMerchantId(merchantId int64, offset int, pageSize int) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields("room_id,merchant_id").Where("merchant_id", merchantId).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetMchRoomListByMerchantId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchRoomListByMerchantId. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *MchRoom) GetMchRoomListTotalCount(merchantId int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("merchant_id", merchantId).Count(1)
	logrus.Debugf("GetMchRoomListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchRoomListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}
