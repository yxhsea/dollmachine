package notice

import (
	"github.com/pkg/errors"
	"dollmachine/dolluser/ff_common/ff_page"
	"dollmachine/dolluser/ff_model/notice"
)

type NoticeService struct {
}

func NewNoticeService() *NoticeService {
	return &NoticeService{}
}

func (p *NoticeService) GetNoticeInfo(noticeId int64) (map[string]interface{}, error) {
	//公告信息
	noticeDao := notice.NewNotice()
	if !noticeDao.CheckIsExitsByNoticeId(noticeId) {
		return nil, errors.New("公告不存在")
	}
	noticeInfo, err := noticeDao.GetNoticeInfoByNoticeId(noticeId, "*")
	if err != nil {
		return nil, err
	}
	return noticeInfo, nil
}

func (p *NoticeService) GetNoticeList(merchantId int64, offset int, pageSize int, totalSize int) ([]map[string]interface{}, *ff_page.Page, error) {
	page := ff_page.NewPage(offset, pageSize, totalSize)
	noticeDao := notice.NewNotice()
	noticeList, err := noticeDao.GetNoticeListByMerchantId(merchantId, offset, pageSize, "*")
	if err != nil || len(noticeList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}

	count, _ := noticeDao.GetNoticeListTotalCount(merchantId)
	page.SetTotalSize(count)
	return noticeList, page, nil
}
