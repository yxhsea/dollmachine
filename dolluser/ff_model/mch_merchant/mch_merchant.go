package mch_merchant

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

type MchMerchant struct {
}

func NewMchMerchant() *MchMerchant {
	return &MchMerchant{}
}

func (u *MchMerchant) getTableName() string {
	return "mch_merchant"
}

func (u *MchMerchant) CheckIsExitsByMerchantId(merchantId int64) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(u.getTableName()).Where("merchant_id", merchantId).Count(1)
	logrus.Debugf("Check merchant id is exits. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Check merchant id is exits failure. Error : %v ", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
