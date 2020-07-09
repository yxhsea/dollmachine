package mch_gift_type

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type MchGiftType struct {
}

func NewMchGiftType() *MchGiftType {
	return &MchGiftType{}
}

func (p *MchGiftType) getTableName() string {
	return "mch_gift_type"
}

func (p *MchGiftType) GetMchGiftTypeList(offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetMchGiftTypeList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchGiftTypeList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *MchGiftType) GetMchGiftTypeListTotalCount() (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Count(1)
	logrus.Debugf("GetMchGiftTypeListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchGiftTypeListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *MchGiftType) GetMchGiftTypeInfo(GiftTypeId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("gift_type_id", GiftTypeId).First()
	logrus.Debugf("GetMchGiftTypeInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchGiftTypeInfo. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchGiftType) CheckIsExitsGiftTypeName(name string) bool {
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

func (p *MchGiftType) CheckIsExitsGiftTypeId(GiftTypeId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("gift_type_id", GiftTypeId).Count(1)
	logrus.Debugf("CheckIsExitsGiftTypeId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsGiftTypeId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
