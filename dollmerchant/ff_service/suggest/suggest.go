package suggest

import (
	"errors"
	"github.com/Unknwon/com"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_page"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"dollmachine/dollmerchant/ff_model/suggest"
	"time"
)

type SuggestService struct {
}

func NewSuggestService() *SuggestService {
	return &SuggestService{}
}

func (p *SuggestService) GetSuggestInfo(SuggestId string) (map[string]interface{}, error) {
	SuggestDao := suggest.NewSuggest()
	one, err := SuggestDao.GetSuggestInfo(SuggestId, "*")
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (p *SuggestService) GetSuggestList(offset, pageSize, totalSize string) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()
	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	SuggestDao := suggest.NewSuggest()
	SuggestList, err := SuggestDao.GetSuggestList(Offset, PageSize, "*")
	if err != nil || len(SuggestList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}
	count, _ := SuggestDao.GetSuggestListTotalCount()
	page.SetTotalSize(count)
	return SuggestList, page, nil
}

func (p *SuggestService) SuggestReply(suggestId, reply string, oprUid int64, oprUName string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("Suggest").Data(map[string]interface{}{
		"reply":          reply,
		"last_opr_uid":   oprUid,
		"last_opr_uname": oprUName,
		"reply_time":     nowTime,
		"updated_at":     nowTime,
	}).Where("suggest_id", suggestId).Update()
	logrus.Debugf("Update Suggest. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Update Suggest. Error : %v", err)
		return false
	}
	return true
}

func (p *SuggestService) CheckIsExitsSuggestId(SuggestId string) bool {
	SuggestDao := suggest.NewSuggest()
	return SuggestDao.CheckIsExitsSuggestId(SuggestId)
}
