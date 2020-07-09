package mch_device_gift

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type MchDeviceGift struct {
}

func NewMchDeviceGift() *MchDeviceGift {
	return &MchDeviceGift{}
}

func (p *MchDeviceGift) getTableName() string {
	return "mch_device_gift"
}

func (p *MchDeviceGift) GetMchDeviceGiftList(offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetMchDeviceGiftList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchDeviceGiftList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *MchDeviceGift) GetMchDeviceGiftListTotalCount() (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Count(1)
	logrus.Debugf("GetMchDeviceGiftListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchDeviceGiftListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *MchDeviceGift) GetMchDeviceGiftInfo(DeviceId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("device_id", DeviceId).First()
	logrus.Debugf("GetMchDeviceGiftInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchDeviceGiftInfo. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchDeviceGift) CheckIsExitsGiftTypeName(name string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("name", name).Count(1)
	logrus.Debugf("CheckIsExitsGiftTypeName. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsGiftTypeName. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *MchDeviceGift) CheckIsExitsDeviceId(DeviceId string) bool {
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

//判断该礼品是否被绑定
func (p *MchDeviceGift) CheckIsBindGiftId(giftId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("gift_id", giftId).Count(1)
	logrus.Debugf("CheckIsBindGiftId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsBindGiftId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
