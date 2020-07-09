package exchange_integral

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type ExchangeIntegral struct {
}

func NewExchangeIntegral() *ExchangeIntegral {
	return &ExchangeIntegral{}
}

func (p *ExchangeIntegral) getTableName() string {
	return "exchange_integral"
}

func (p *ExchangeIntegral) GetExchangeIntegralList(offset int, pageSize int, fields string, merchantId, StartTime, EndTime int64) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	listDbr := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("merchant_id", merchantId)

	if StartTime > 0 && EndTime > 0 {
		listDbr = listDbr.Where("created_at", "between", []int64{StartTime, EndTime})
	}

	list, err := listDbr.Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetExchangeIntegralList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetExchangeIntegralList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *ExchangeIntegral) GetExchangeIntegralListTotalCount(merchantId, StartTime, EndTime int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	countDbr := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("merchant_id", merchantId)

	if StartTime > 0 && EndTime > 0 {
		countDbr = countDbr.Where("created_at", "between", []int64{StartTime, EndTime})
	}

	count, err := countDbr.Count(1)
	logrus.Debugf("GetExchangeIntegralListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetExchangeIntegralListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}
