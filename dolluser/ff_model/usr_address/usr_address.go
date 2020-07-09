package usr_address

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

type UsrAddress struct {
}

func NewUsrAddress() *UsrAddress {
	return &UsrAddress{}
}

func (p *UsrAddress) getTableName() string {
	return "usr_addr"
}

func (p *UsrAddress) CheckIsExitsByAddressId(addressId int64) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("address_id", addressId).Count(1)
	logrus.Debugf("CheckIsExitsByAddressId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsByAddressId. Error : %v ", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *UsrAddress) GetAddressInfoByAddressId(addressId int64, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("address_id", addressId).First()
	logrus.Debugf("GetAddressInfoByAddressId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetAddressInfoByAddressId. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *UsrAddress) GetUserLastAddressByUserId(userId int64, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("user_id", userId).Limit(1).First()
	logrus.Debugf("Query user address. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query user address failure. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *UsrAddress) AddAddressInfo(addressInfo map[string]interface{}) (int64, error) {
	dbr := ff_vars.DbConn.GetInstance()
	addressId, err := dbr.Table(p.getTableName()).Data(addressInfo).Insert()
	logrus.Debugf("AddAddressInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("AddAddressInfo. Error : %v", err)
		return 0, err
	}
	return int64(addressId), nil
}

func (p *UsrAddress) UpdateAddressInfo(addressInfo map[string]interface{}, addressId int64) error {
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table(p.getTableName()).Data(addressInfo).Update()
	logrus.Debugf("UpdateAddressInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("UpdateAddressInfo. Error : %v", err)
		return err
	}
	return nil
}
