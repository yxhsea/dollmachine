package pmt_play

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

type PmtPlay struct {
}

func NewPmtPlay() *PmtPlay {
	return &PmtPlay{}
}

func (p *PmtPlay) getTableName() string {
	return "pmt_play"
}

func (p *PmtPlay) CheckIsExitsByPlayId(playId int64) bool {
	dbr := ff_vars.DbConn.GetInstance()
	valInt, err := dbr.Table(p.getTableName()).Where("play_id", playId).Count(1)
	logrus.Debugf("CheckIsExitsByPlayId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("CheckIsExitsByPlayId. Error : %v ", err)
		return false
	}
	if valInt > 0 {
		return true
	}
	return false
}

func (p *PmtPlay) GetPlayInfoByPlayId(playId int64, fields string) (map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	one, err := dbr.Table(p.getTableName()).Fields(fields).Where("play_id", playId).First()
	logrus.Debugf("GetPmtPlayInfoByPlayId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPmtPlayInfoByPlayId. Error : %v", err)
		return nil, err
	}
	return one, nil
}

func (p *PmtPlay) GetPlayListByMerchantId(merchantId int64, offset int, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("merchant_id", merchantId).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetPlayListByMerchantId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPlayListByMerchantId. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *PmtPlay) GetPlayListTotalCountByMerchantId(merchantId int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("merchant_id", merchantId).Count(1)
	logrus.Debugf("GetPlayListTotalCount. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPlayListTotalCount. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *PmtPlay) GetPlayListByMchIdAndUserId(merchantId, userId int64, offset, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).Fields(fields).Where("merchant_id", merchantId).Where("user_id", userId).Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetPlayListByMchIdAndUserId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPlayListByMchIdAndUserId. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *PmtPlay) GetPlayListTotalCountByMchIdAndUserId(merchantId, userId int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).Where("merchant_id", merchantId).Where("user_id", userId).Count(1)
	logrus.Debugf("GetPlayListTotalCountByMchIdAndUserId. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPlayListTotalCountByMchIdAndUserId. Error : %v", err)
		return 0, err
	}
	return count, nil
}

func (p *PmtPlay) GetPlayListByMchIdAndUIDAndAward(merchantId, userId int64, offset, pageSize int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).
		Fields(fields).
		Where("merchant_id", merchantId).
		Where("user_id", userId).
		Where("is_award", 1).
		Offset(offset).Limit(pageSize).Get()
	logrus.Debugf("GetPlayListByMchIdAndUIDAndAward. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPlayListByMchIdAndUIDAndAward. Error : %v", err)
		return nil, err
	}
	return list, nil
}

func (p *PmtPlay) GetPlayListTotalCountByMchIdAndUIDAndAward(merchantId, userId int64) (int, error) {
	dbr := ff_vars.DbConn.GetInstance()
	count, err := dbr.Table(p.getTableName()).
		Where("merchant_id", merchantId).
		Where("user_id", userId).
		Where("is_award", 1).
		Count(1)
	logrus.Debugf("GetPlayListTotalCountByMchIdAndUIDAndAward. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPlayListTotalCountByMchIdAndUIDAndAward. Error : %v", err)
		return 0, err
	}
	return count, nil
}

//查询某一台设备上的中奖记录
func (p *PmtPlay) GetPlayListByMchIdAndDevIdAndAward(merchantId, deviceId int64, limit int, fields string) ([]map[string]interface{}, error) {
	dbr := ff_vars.DbConn.GetInstance()
	list, err := dbr.Table(p.getTableName()).
		Fields(fields).
		Where("merchant_id", merchantId).
		Where("device_id", deviceId).
		Where("is_award", 1).
		Order("award_time desc").Limit(limit).Get()
	logrus.Debugf("GetPlayListByMchIdAndDevIdAndAward. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("GetPlayListByMchIdAndDevIdAndAward. Error : %v", err)
		return nil, err
	}
	return list, nil
}
