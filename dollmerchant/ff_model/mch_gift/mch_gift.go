package mch_gift

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type MchGift struct {
}

func NewMchGift() *MchGift {
	return &MchGift{}
}

func (p *MchGift) getTableName() string {
	return "mch_gift"
}

func (p *MchGift) GetMchGiftList(offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetMchGiftList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchGiftList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *MchGift) GetMchGiftListTotalCount() (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Count(1)
	logrus.Debugf("GetMchGiftListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchGiftListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *MchGift) GetMchGiftInfo(GiftId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("gift_id", GiftId).First()
	logrus.Debugf("GetMchGiftInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchGiftInfo. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchGift) CheckIsExitsGiftName(name string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("name", name).Count(1)
	logrus.Debugf("CheckIsExitsGiftName. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsGiftName. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *MchGift) CheckIsExitsGiftId(GiftId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("gift_id", GiftId).Count(1)
	logrus.Debugf("CheckIsExitsGiftId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsGiftId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

//判断该礼品类型是否被绑定
func (p *MchGift) CheckIsBindGiftTypeId(GiftTypeId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("gift_type_id", GiftTypeId).Count(1)
	logrus.Debugf("CheckIsBindGiftTypeId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsBindGiftTypeId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
