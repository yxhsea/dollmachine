package suggest

import (
	"github.com/pkg/errors"
	"dollmachine/dolluser/ff_common/ff_page"
	"dollmachine/dolluser/ff_model/suggest"
	"time"
)

type SuggestService struct {
}

func NewSuggestService() *SuggestService {
	return &SuggestService{}
}

func (p *SuggestService) GetSuggestInfo(SuggestId int64) (map[string]interface{}, error) {
	//意见信息
	suggestDao := suggest.NewSuggest()
	if !suggestDao.CheckIsExitsBySuggestId(SuggestId) {
		return nil, errors.New("记录不存在")
	}
	suggestInfo, err := suggestDao.GetSuggestInfoBySuggestId(SuggestId, "*")
	if err != nil {
		return nil, err
	}
	return suggestInfo, nil
}

func (p *SuggestService) GetSuggestList(merchantId, userId int64, offset int, pageSize int, totalSize int) ([]map[string]interface{}, *ff_page.Page, error) {
	page := ff_page.NewPage(offset, pageSize, totalSize)
	suggestDao := suggest.NewSuggest()
	suggestList, err := suggestDao.GetSuggestListByMerchantIdAndUserId(merchantId, userId, offset, pageSize, "*")
	if err != nil || len(suggestList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}

	count, _ := suggestDao.GetSuggestListTotalCountByMerchantIdAndUserId(merchantId, userId)
	page.SetTotalSize(count)
	return suggestList, page, nil
}

func (p *SuggestService) AddSuggest(merchantId, userId int64, sgType int, content, contactTel string) (int64, error) {
	suggestDao := suggest.NewSuggest()
	nowTime := time.Now().Unix()
	suggestId, err := suggestDao.AddSuggestInfo(map[string]interface{}{
		"merchant_id": merchantId,
		"user_id":     userId,
		"content":     content,
		"type":        sgType,
		"contact_tel": contactTel,
		"created_at":  nowTime,
		"updated_at":  nowTime,
	})
	if err != nil {
		return 0, err
	}
	return suggestId, nil
}
