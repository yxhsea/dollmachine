package notice

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type Notice struct {
}

func NewNotice() *Notice {
	return &Notice{}
}

func (p *Notice) getTableName() string {
	return "notice"
}

func (p *Notice) GetNoticeList(offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetNoticeList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetNoticeList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *Notice) GetNoticeListTotalCount() (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Count(1)
	logrus.Debugf("GetNoticeListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetNoticeListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *Notice) GetNoticeInfo(NoticeId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("notice_id", NoticeId).First()
	logrus.Debugf("GetNoticeInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetNoticeInfo. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *Notice) CheckIsExitsNoticeId(NoticeId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("notice_id", NoticeId).Count(1)
	logrus.Debugf("CheckIsExitsNoticeId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsNoticeId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
