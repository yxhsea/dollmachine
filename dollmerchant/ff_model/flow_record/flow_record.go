package flow_record

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

type FlowRecord struct {
}

func NewFlowRecord() *FlowRecord {
	return &FlowRecord{}
}

func (p *FlowRecord) getTableName() string {
	return "flow_record"
}

func (p *FlowRecord) GetFlowRecordList(offset int, pageSize int, fields string, merchantId, StartTime, EndTime, UserId int64, IType int) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	listDbr := dbr.Table(p.getTableName()).Fields(fields).Where("is_delete", 0).Where("merchant_id", merchantId)

	if IType > 0 { //1|充值 2|积分兑换 3|运费支付 4|游戏消耗
		if UserId > 0 {
			if StartTime > 0 && EndTime > 0 {
				listDbr = listDbr.Where("trade_type", IType).Where("user_id", UserId).Where("created_at", "between", []int64{StartTime, EndTime})
			} else {
				listDbr = listDbr.Where("trade_type", IType).Where("user_id", UserId)
			}
		} else {
			if StartTime > 0 && EndTime > 0 {
				listDbr = listDbr.Where("trade_type", IType).Where("created_at", "between", []int64{StartTime, EndTime})
			} else {
				listDbr = listDbr.Where("trade_type", IType)
			}
		}
	} else { //全部
		if UserId > 0 {
			if StartTime > 0 && EndTime > 0 {
				listDbr = listDbr.Where("user_id", UserId).Where("created_at", "between", []int64{StartTime, EndTime})
			} else {
				listDbr = listDbr.Where("user_id", UserId)
			}
		} else {
			if StartTime > 0 && EndTime > 0 {
				listDbr = listDbr.Where("created_at", "between", []int64{StartTime, EndTime})
			}
		}
	}

	list, err := listDbr.Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetFlowRecordList. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetFlowRecordList. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *FlowRecord) GetFlowRecordListTotalCount(merchantId, StartTime, EndTime, UserId int64, IType int) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	countDbr := dbr.Table(p.getTableName()).Where("is_delete", 0).Where("merchant_id", merchantId)

	if IType > 0 { //1|充值 2|积分兑换 3|运费支付 4|游戏消耗
		if UserId > 0 {
			if StartTime > 0 && EndTime > 0 {
				countDbr = countDbr.Where("trade_type", IType).Where("user_id", UserId).Where("created_at", "between", []int64{StartTime, EndTime})
			} else {
				countDbr = countDbr.Where("trade_type", IType).Where("user_id", UserId)
			}
		} else {
			if StartTime > 0 && EndTime > 0 {
				countDbr = countDbr.Where("trade_type", IType).Where("created_at", "between", []int64{StartTime, EndTime})
			} else {
				countDbr = countDbr.Where("trade_type", IType)
			}
		}
	} else { //全部
		if UserId > 0 {
			if StartTime > 0 && EndTime > 0 {
				countDbr = countDbr.Where("user_id", UserId).Where("created_at", "between", []int64{StartTime, EndTime})
			} else {
				countDbr = countDbr.Where("user_id", UserId)
			}
		} else {
			if StartTime > 0 && EndTime > 0 {
				countDbr = countDbr.Where("created_at", "between", []int64{StartTime, EndTime})
			}
		}
	}

	count, err := countDbr.Count(1)
	logrus.Debugf("GetFlowRecordListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetFlowRecordListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}
