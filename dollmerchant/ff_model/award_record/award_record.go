package award_record

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type AwardRecord struct {
}

func NewAwardRecord() *AwardRecord {
	return &AwardRecord{}
}

func (p *AwardRecord) getTableName() string {
	return "award_record"
}

func (p *AwardRecord) GetAwardRecordListByExpress(offset int, pageSize int, fields string, sendStatus int, userId, merchantId int64) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	listDbr := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("merchant_id", merchantId)

	if sendStatus == 1 { //已发货
		if userId > 0 {
			listDbr.Where("is_send", 1).Where("user_id", userId)
		} else {
			listDbr.Where("is_send", 1)
		}
	} else if sendStatus == 2 { //未发货
		if userId > 0 {
			listDbr.Where("is_send", 0).Where("user_id", userId)
		} else {
			listDbr.Where("is_send", 0)
		}
	} else { //全部
		if userId > 0 {
			listDbr.Where("user_id", userId)
		}
	}

	list, err := listDbr.Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetAwardRecordListByExpress. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetAwardRecordListByExpress. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *AwardRecord) GetAwardRecordListTotalCountByExpress(sendStatus int, userId, merchantId int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	countDbr := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("merchant_id", merchantId)

	if sendStatus == 1 { //已发货
		if userId > 0 {
			countDbr.Where("is_send", 1).Where("user_id", userId)
		} else {
			countDbr.Where("is_send", 1)
		}
	} else if sendStatus == 2 { //未发货
		if userId > 0 {
			countDbr.Where("is_send", 0).Where("user_id", userId)
		} else {
			countDbr.Where("is_send", 0)
		}
	} else { //全部
		if userId > 0 {
			countDbr.Where("user_id", userId)
		}
	}

	count, err := countDbr.Count(1)
	logrus.Debugf("GetAwardRecordListTotalCountByExpress. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetAwardRecordListTotalCountByExpress. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *AwardRecord) GetAwardRecordInfo(RecordId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("record_id", RecordId).First()
	logrus.Debugf("GetAwardRecordInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetAwardRecordInfo. Error : %v", err)
		return nil, err
	}
	return one, nil
}
