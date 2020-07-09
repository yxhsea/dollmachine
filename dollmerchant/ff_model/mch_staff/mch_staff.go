package mch_staff

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type MchStaff struct {
}

func NewMchStaff() *MchStaff {
	return &MchStaff{}
}

func (p *MchStaff) getTableName() string {
	return "mch_staff"
}

func (p *MchStaff) GetMchStaffByStaffId(StaffId int64, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("staff_id", StaffId).First()
	logrus.Debugf("GetMchStaffByStaffId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchStaffByStaffId. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchStaff) GetMchStaffList(offset int, pageSize int, fields string, merchantId int64) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetMchStaffList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchStaffList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *MchStaff) GetMchStaffListTotalCount(merchantId int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Count(1)
	logrus.Debugf("GetMchStaffListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchStaffListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *MchStaff) GetMchStaffInfo(StaffId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("staff_id", StaffId).First()
	logrus.Debugf("GetMchStaffInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetMchStaffInfo. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *MchStaff) CheckIsExitsStaffId(StaffId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("staff_id", StaffId).Count(1)
	logrus.Debugf("CheckIsExitsStaffId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsStaffId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
