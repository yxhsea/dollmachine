package pmt_recharge

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type PmtRecharge struct {
}

func NewPmtRecharge() *PmtRecharge {
	return &PmtRecharge{}
}

func (p *PmtRecharge) getTableName() string {
	return "pmt_recharge"
}

func (p *PmtRecharge) GetCurrentMonthDetail(merchantId, StartTime, EndTime int64) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields("sum(amount) sum_amount, sum(server_fee) sum_server_fee, sum(mch_income) sum_mch_income").
		Where("is_paid", 2).Where("merchant_id", merchantId).Where("created_at", "between", []int64{StartTime, EndTime}).First()
	logrus.Debugf("GetCurrentMonthDetail. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetCurrentMonthDetail. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *PmtRecharge) GetPmtRechargeListByUserRecharge(offset int, pageSize int, fields string, merchantId, StartTime, EndTime int64) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	listDbr := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("merchant_id", merchantId)

	if StartTime > 0 && EndTime > 0 {
		listDbr = listDbr.Where("created_at", "between", []int64{StartTime, EndTime})
	}

	list, err := listDbr.Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetPmtRechargeListByUserRecharge. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPmtRechargeListByUserRecharge. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *PmtRecharge) GetPmtRechargeListTotalCountByUserRecharge(merchantId, StartTime, EndTime int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	countDbr := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("merchant_id", merchantId)

	if StartTime > 0 && EndTime > 0 {
		countDbr = countDbr.Where("created_at", "between", []int64{StartTime, EndTime})
	}

	count, err := countDbr.Count(1)
	logrus.Debugf("GetPmtRechargeListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPmtRechargeListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}
