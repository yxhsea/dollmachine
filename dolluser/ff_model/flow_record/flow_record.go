package flow_record

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

type FlowRecord struct {
}

func NewFlowRecord() *FlowRecord {
	return &FlowRecord{}
}

func (p *FlowRecord) getTableName() string {
	return "flow_record"
}

func (p *FlowRecord) CheckIsExitsByFlowRecordId(FlowId int64) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("flow_id", FlowId).Count(1)
	logrus.Debugf("CheckIsExitsByFlowRecordId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsByFlowRecordId. Error : %v ", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *FlowRecord) GetFlowRecordInfoByFlowId(FlowId int64, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("flow_id", FlowId).First()
	logrus.Debugf("GetFlowRecordInfoByFlowId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetFlowRecordInfoByFlowId. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *FlowRecord) GetFlowRecordListByMerchantId(merchantId int64, offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("merchant_id", merchantId).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetFlowRecordListByMerchantId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetFlowRecordListByMerchantId. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *FlowRecord) GetFlowRecordListTotalCount(merchantId int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("merchant_id", merchantId).Count(1)
	logrus.Debugf("GetFlowRecordListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetFlowRecordListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}
