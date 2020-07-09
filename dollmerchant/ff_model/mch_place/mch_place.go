package mch_place

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type MchPlace struct {
}

func NewMchPlace() *MchPlace {
	return &MchPlace{}
}

func (p *MchPlace) getTableName() string {
	return "mch_place"
}

func (p *MchPlace) GetMchPlaceList(offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetMchPlaceList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchPlaceList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *MchPlace) GetMchPlaceListTotalCount() (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Count(1)
	logrus.Debugf("GetMchPlaceListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchPlaceListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *MchPlace) GetMchPlaceInfo(PlaceId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("place_id", PlaceId).First()
	logrus.Debugf("GetMchPlaceInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchPlaceInfo. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchPlace) CheckIsExitsPlaceName(name string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("name", name).Count(1)
	logrus.Debugf("CheckIsExitsPlaceName. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsPlaceName. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *MchPlace) CheckIsExitsPlaceId(PlaceId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("place_id", PlaceId).Count(1)
	logrus.Debugf("CheckIsExitsPlaceId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsPlaceId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

//判断该礼品类型是否被绑定
func (p *MchPlace) CheckIsBindGiftTypeId(GiftTypeId string) bool {
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
