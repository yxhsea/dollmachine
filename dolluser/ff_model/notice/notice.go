package notice

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

type Notice struct {
}

func NewNotice() *Notice {
	return &Notice{}
}

func (p *Notice) getTableName() string {
	return "notice"
}

func (p *Notice) CheckIsExitsByNoticeId(noticeId int64) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("notice_id", noticeId).Count(1)
	logrus.Debugf("CheckIsExitsByNoticeId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsByNoticeId. Error : %v ", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *Notice) GetNoticeInfoByNoticeId(noticeId int64, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("notice_id", noticeId).First()
	logrus.Debugf("GetNoticeInfoByNoticeId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetNoticeInfoByNoticeId. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *Notice) GetNoticeListByMerchantId(merchantId int64, offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("merchant_id", merchantId).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetNoticeListByMerchantId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetNoticeListByMerchantId. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *Notice) GetNoticeListTotalCount(merchantId int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("merchant_id", merchantId).Count(1)
	logrus.Debugf("GetNoticeListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetNoticeListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}
