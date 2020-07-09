package suggest

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

type Suggest struct {
}

func NewSuggest() *Suggest {
	return &Suggest{}
}

func (p *Suggest) getTableName() string {
	return "suggest"
}

func (p *Suggest) CheckIsExitsBySuggestId(SuggestId int64) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("suggest_id", SuggestId).Count(1)
	logrus.Debugf("CheckIsExitsBySuggestId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsBySuggestId. Error : %v ", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *Suggest) GetSuggestInfoBySuggestId(SuggestId int64, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("suggest_id", SuggestId).First()
	logrus.Debugf("GetSuggestInfoBySuggestId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetSuggestInfoBySuggestId. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *Suggest) GetSuggestListByMerchantIdAndUserId(merchantId, userId int64, offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("merchant_id", merchantId).Where("user_id", userId).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetSuggestListByMerchantId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetSuggestListByMerchantId. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *Suggest) GetSuggestListTotalCountByMerchantIdAndUserId(merchantId, userId int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("merchant_id", merchantId).Where("user_id", userId).Count(1)
	logrus.Debugf("GetSuggestListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetSuggestListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *Suggest) AddSuggestInfo(suggestInfo map[string]interface{}) (int64, error) {
	dbr := ff_vars.DbConn.GetInstance()
	suggestId, err := dbr.Table(p.getTableName()).Data(suggestInfo).Insert()
	logrus.Debugf("AddSuggestInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("AddSuggestInfo. Error : %v", err)
		return 0, err
	}
	return int64(suggestId), nil
}
