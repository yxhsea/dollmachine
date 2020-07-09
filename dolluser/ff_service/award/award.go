package award

import (
	"errors"
	"dollmachine/dolluser/ff_common/ff_page"
	"dollmachine/dolluser/ff_model/pmt_play"
)

type AwardService struct {
}

func NewAwardService() *AwardService {
	return &AwardService{}
}

func (p *AwardService) GetAwardList(merchantId, userId int64, offset int, pageSize int, totalSize int) ([]map[string]interface{}, *ff_page.Page, error) {
	page := ff_page.NewPage(offset, pageSize, totalSize)
	awardDao := pmt_play.NewPmtPlay()
	awardList, err := awardDao.GetPlayListByMchIdAndUIDAndAward(merchantId, userId, offset, pageSize, "*")
	if err != nil || len(awardList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}

	count, _ := awardDao.GetPlayListTotalCountByMchIdAndUIDAndAward(merchantId, userId)
	page.SetTotalSize(count)
	return awardList, page, nil
}
