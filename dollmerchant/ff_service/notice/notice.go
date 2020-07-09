package Notice

import (
	"errors"
	"github.com/Unknwon/com"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_page"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"dollmachine/dollmerchant/ff_model/notice"
	"time"
)

type NoticeService struct {
}

func NewNoticeService() *NoticeService {
	return &NoticeService{}
}

func (p *NoticeService) GetNoticeInfo(NoticeId string) (map[string]interface{}, error) {
	NoticeDao := notice.NewNotice()
	one, err := NoticeDao.GetNoticeInfo(NoticeId, "*")
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (p *NoticeService) GetNoticeList(offset, pageSize, totalSize string) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()
	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	NoticeDao := notice.NewNotice()
	NoticeList, err := NoticeDao.GetNoticeList(Offset, PageSize, "*")
	if err != nil || len(NoticeList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}
	count, _ := NoticeDao.GetNoticeListTotalCount()
	page.SetTotalSize(count)
	return NoticeList, page, nil
}

func (p *NoticeService) AddNotice(title, content, thumb string, merchantId int64) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("notice").Data(map[string]interface{}{
		"merchant_id": merchantId,
		"title":       title,
		"content":     content,
		"thumb":       thumb,

		"status":     1,
		"created_at": nowTime,
		"updated_at": nowTime,
	}).Insert()
	logrus.Debugf("Insert Notice. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Insert Notice. Error : %v", err)
		return false
	}
	return true
}

func (p *NoticeService) UpdNotice(title, content, thumb, noticeId string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("notice").Data(map[string]interface{}{
		"title":      title,
		"content":    content,
		"thumb":      thumb,
		"created_at": nowTime,
	}).Where("notice_id", noticeId).Update()
	logrus.Debugf("Update Notice. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Update Notice. Error : %v", err)
		return false
	}
	return true
}

func (p *NoticeService) DelNotice(noticeId string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("notice").Data(map[string]interface{}{
		"is_delete":  1,
		"deleted_at": nowTime,
	}).Where("notice_id", noticeId).Update()
	logrus.Debugf("delete Notice. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("delete Notice. Error : %v", err)
		return false
	}
	return true
}

func (p *NoticeService) CheckIsExitsNoticeId(NoticeId string) bool {
	noticeDao := notice.NewNotice()
	return noticeDao.CheckIsExitsNoticeId(NoticeId)
}
