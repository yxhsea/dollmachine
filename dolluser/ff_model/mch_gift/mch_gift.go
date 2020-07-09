package mch_gift

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

type MchGift struct {
}

func NewMchGift() *MchGift {
	return &MchGift{}
}

func (p *MchGift) getTableName() string {
	return "mch_gift"
}

func (p *MchGift) CheckIsExitsByGiftId(giftId int64) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("gift_id", giftId).Count(1)
	logrus.Debugf("CheckIsExitsByGiftId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsByGiftId. Error : %v ", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *MchGift) GetMchGiftInfoByGiftId(giftId int64, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("gift_id", giftId).First()
	logrus.Debugf("GetMchGiftInfoByGiftId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchGiftInfoByGiftId. Error : %v", err)
		return nil, err
	}
	return one, nil
}
