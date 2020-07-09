package win

import (
	"errors"
	"github.com/Unknwon/com"
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_common/ff_page"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"dollmachine/dollmerchant/ff_model/award_record"
	"dollmachine/dollmerchant/ff_model/pmt_play"
	"time"
)

type WinService struct {
}

func NewWinService() *WinService {
	return &WinService{}
}

func (p *WinService) GetWinOnlineList(offset, pageSize, totalSize, cashStatus, userId string, merchantId int64) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()

	CashStatus, _ := com.StrTo(cashStatus).Int()
	UserId, _ := com.StrTo(userId).Int64()

	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	pmtPlayDao := pmt_play.NewPmtPlay()
	pmtPlayList, err := pmtPlayDao.GetPmtPlayListByOnline(Offset, PageSize, "*", CashStatus, UserId, merchantId)
	if err != nil || len(pmtPlayList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}

	count, _ := pmtPlayDao.GetPmtPlayListTotalCountByOnline(CashStatus, UserId, merchantId)
	page.SetTotalSize(count)
	return pmtPlayList, page, nil
}

func (p *WinService) GetWinOfflineList(offset, pageSize, totalSize, cashStatus, userId string, merchantId int64) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()

	CashStatus, _ := com.StrTo(cashStatus).Int()
	UserId, _ := com.StrTo(userId).Int64()

	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	pmtPlayDao := pmt_play.NewPmtPlay()
	pmtPlayList, err := pmtPlayDao.GetPmtPlayListByOffline(Offset, PageSize, "*", CashStatus, UserId, merchantId)
	if err != nil || len(pmtPlayList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}

	count, _ := pmtPlayDao.GetPmtPlayListTotalCountByOffline(CashStatus, UserId, merchantId)
	page.SetTotalSize(count)
	return pmtPlayList, page, nil
}

func (p *WinService) GetWinExpressList(offset, pageSize, totalSize, sendStatus, userId string, merchantId int64) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()

	SendStatus, _ := com.StrTo(sendStatus).Int()
	UserId, _ := com.StrTo(userId).Int64()

	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	awardRecordDao := award_record.NewAwardRecord()
	pmtPlayList, err := awardRecordDao.GetAwardRecordListByExpress(Offset, PageSize, "*", SendStatus, UserId, merchantId)
	if err != nil || len(pmtPlayList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}

	count, _ := awardRecordDao.GetAwardRecordListTotalCountByExpress(SendStatus, UserId, merchantId)
	page.SetTotalSize(count)
	return pmtPlayList, page, nil
}

func (p *WinService) WinCash(playId, cashStatus string, oprUid int64, oprUName string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("pmt_play").Data(map[string]interface{}{
		"updated_at":     nowTime,
		"exchange_time":  nowTime,
		"status":         cashStatus,
		"last_opr_uid":   oprUid,
		"last_opr_uname": oprUName,
	}).Where("play_id", playId).Update()
	logrus.Debugf("Cash prize. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Cash prize. Error : %v", err)
		return false
	}
	return true
}

func (p *WinService) WinSend(recordId, expressCompany, expressNo string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("award_record").Data(map[string]interface{}{
		"express_company": expressCompany,
		"express_no":      expressNo,
		"updated_at":      nowTime,
		"is_send":         1,
		"send_at":         time.Now().Unix(),
	}).Where("record_id", "=", recordId).Update()
	logrus.Debugf("Send prize. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Send prize. Error : %v", err)
		return false
	}
	return true
}
