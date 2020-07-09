package game

import (
	"errors"
	"dollmachine/dolluser/ff_common/ff_page"
	"dollmachine/dolluser/ff_model/pmt_play"
)

type GameService struct {
}

func NewGameService() *GameService {
	return &GameService{}
}

func (p *GameService) GetGameList(merchantId, userId int64, offset int, pageSize int, totalSize int) ([]map[string]interface{}, *ff_page.Page, error) {
	page := ff_page.NewPage(offset, pageSize, totalSize)
	gameDao := pmt_play.NewPmtPlay()
	gameList, err := gameDao.GetPlayListByMchIdAndUserId(merchantId, userId, offset, pageSize, "*")
	if err != nil || len(gameList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}

	count, _ := gameDao.GetPlayListTotalCountByMchIdAndUserId(merchantId, userId)
	page.SetTotalSize(count)
	return gameList, page, nil
}
