package suggest

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type Suggest struct {
}

func NewSuggest() *Suggest {
	return &Suggest{}
}

func (p *Suggest) getTableName() string {
	return "suggest"
}

func (p *Suggest) GetSuggestList(offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetSuggestList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetSuggestList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *Suggest) GetSuggestListTotalCount() (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Count(1)
	logrus.Debugf("GetSuggestListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetSuggestListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *Suggest) GetSuggestInfo(SuggestId string, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("suggest_id", SuggestId).First()
	logrus.Debugf("GetSuggestInfo. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetSuggestInfo. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *Suggest) CheckIsExitsSuggestId(SuggestId string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("suggest_id", SuggestId).Count(1)
	logrus.Debugf("CheckIsExitsSuggestId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsSuggestId. Error : %v", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}
