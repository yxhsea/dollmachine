package draw_record

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type DrawRecord struct {
}

func NewDrawRecord() *DrawRecord {
	return &DrawRecord{}
}

func (p *DrawRecord) getTableName() string {
	return "draw_record"
}

func (p *DrawRecord) GetDrawRecordList(offset int, pageSize int, fields string, merchantId int64) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	listDbr := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("merchant_id", merchantId)
	list, err := listDbr.Order("created_at desc").Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetDrawRecordList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetDrawRecordList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *DrawRecord) GetDrawRecordListTotalCount(merchantId int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	countDbr := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("merchant_id", merchantId)
	count, err := countDbr.Count(1)
	logrus.Debugf("GetDrawRecordListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetDrawRecordListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}
