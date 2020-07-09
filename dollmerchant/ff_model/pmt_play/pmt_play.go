package pmt_play

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type PmtPlay struct {
}

func NewPmtPlay() *PmtPlay {
	return &PmtPlay{}
}

func (p *PmtPlay) getTableName() string {
	return "pmt_play"
}

func (p *PmtPlay) GetPmtPlayListByOnline(offset int, pageSize int, fields string, cashStatus int, userId int64, merchantId int64) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	listDbr := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("merchant_id", merchantId).Where("line", 1)

	if cashStatus == 1 { //已核销
		if userId > 0 {
			listDbr.Where("is_award", 1).Where("status", 1).Where("user_id", userId)
		} else {
			listDbr.Where("is_award", 1).Where("status", 1)
		}
	} else if cashStatus == 2 { //未核销
		if userId > 0 {
			listDbr.Where("is_award", 1).Where("status", 0).Where("user_id", userId)
		} else {
			listDbr.Where("is_award", 1).Where("status", 0)
		}
	} else { //全部
		if userId > 0 {
			listDbr.Where("is_award", 1).Where("user_id", userId)
		} else {
			listDbr.Where("is_award", 1)
		}
	}

	list, err := listDbr.Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetPmtPlayList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPmtPlayList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *PmtPlay) GetPmtPlayListTotalCountByOnline(cashStatus int, userId int64, merchantId int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	countDbr := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("merchant_id", merchantId).Where("line", 1)

	if cashStatus == 1 { //已核销
		if userId > 0 {
			countDbr.Where("is_award", 1).Where("status", 1).Where("user_id", userId)
		} else {
			countDbr.Where("is_award", 1).Where("status", 1)
		}
	} else if cashStatus == 2 { //未核销
		if userId > 0 {
			countDbr.Where("is_award", 1).Where("status", 0).Where("user_id", userId)
		} else {
			countDbr.Where("is_award", 1).Where("status", 0)
		}
	} else { //全部
		if userId > 0 {
			countDbr.Where("is_award", 1).Where("user_id", userId)
		} else {
			countDbr.Where("is_award", 1)
		}
	}
	count, err := countDbr.Count(1)
	logrus.Debugf("GetPmtPlayListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPmtPlayListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *PmtPlay) GetPmtPlayListByOffline(offset int, pageSize int, fields string, cashStatus int, userId int64, merchantId int64) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	listDbr := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("merchant_id", merchantId).Where("line", 2)

	if cashStatus == 1 { //已核销
		if userId > 0 {
			listDbr.Where("is_award", 1).Where("status", 1).Where("user_id", userId)
		} else {
			listDbr.Where("is_award", 1).Where("status", 1)
		}
	} else if cashStatus == 2 { //未核销
		if userId > 0 {
			listDbr.Where("is_award", 1).Where("status", 0).Where("user_id", userId)
		} else {
			listDbr.Where("is_award", 1).Where("status", 0)
		}
	} else { //全部
		if userId > 0 {
			listDbr.Where("is_award", 1).Where("user_id", userId)
		} else {
			listDbr.Where("is_award", 1)
		}
	}

	list, err := listDbr.Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetPmtPlayList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPmtPlayList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *PmtPlay) GetPmtPlayListTotalCountByOffline(cashStatus int, userId int64, merchantId int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	countDbr := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("merchant_id", merchantId).Where("line", 2)

	if cashStatus == 1 { //已核销
		if userId > 0 {
			countDbr.Where("is_award", 1).Where("status", 1).Where("user_id", userId)
		} else {
			countDbr.Where("is_award", 1).Where("status", 1)
		}
	} else if cashStatus == 2 { //未核销
		if userId > 0 {
			countDbr.Where("is_award", 1).Where("status", 0).Where("user_id", userId)
		} else {
			countDbr.Where("is_award", 1).Where("status", 0)
		}
	} else { //全部
		if userId > 0 {
			countDbr.Where("is_award", 1).Where("user_id", userId)
		} else {
			countDbr.Where("is_award", 1)
		}
	}
	count, err := countDbr.Count(1)
	logrus.Debugf("GetPmtPlayListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPmtPlayListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}
