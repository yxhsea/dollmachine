package gift

import (
	"context"
	"errors"
	"github.com/Unknwon/com"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_page"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"dollmachine/dollmerchant/ff_model/mch_device_gift"
	"dollmachine/dollmerchant/ff_model/mch_gift"
	UniqueId "dollmachine/dollmerchant/proto/unique_id"
	"time"
)

type GiftService struct {
}

func NewGiftService() *GiftService {
	return &GiftService{}
}

func (p *GiftService) GetGiftInfo(GiftId string) (map[string]interface{}, error) {
	GiftDao := mch_gift.NewMchGift()
	one, err := GiftDao.GetMchGiftInfo(GiftId, "*")
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (p *GiftService) GetGiftList(offset, pageSize, totalSize string) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()
	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	GiftDao := mch_gift.NewMchGift()
	GiftList, err := GiftDao.GetMchGiftList(Offset, PageSize, "*")
	if err != nil || len(GiftList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}
	count, _ := GiftDao.GetMchGiftListTotalCount()
	page.SetTotalSize(count)
	return GiftList, page, nil
}

func (p *GiftService) AddGift(name, merchantId string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()

	//生成GiftId
	cli := UniqueId.NewGenerateUniqueIdService("go.micro.srv.unique_id", ff_vars.RpcSrv.Client())
	rsp, err := cli.GenerateUniqueId(context.TODO(), &UniqueId.UniqueIdRequest{Key: "giftId"})
	if err != nil {
		logrus.Errorf("Generate GiftId error: %v", err)
		return false
	}
	GiftId := rsp.Value

	_, err = dbr.Table("mch_gift").Data(map[string]interface{}{
		"gift_id":     GiftId,
		"merchant_id": merchantId,
		"name":        name,
		"nick_name":   name,
		"status":      1,
		"created_at":  nowTime,
		"updated_at":  nowTime,
	}).Insert()
	logrus.Debugf("Insert gift. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Insert gift. Error : %v", err)
		return false
	}
	return true
}

func (p *GiftService) UpdGift(GiftId, name string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("mch_gift").Data(map[string]interface{}{
		"name":       name,
		"nick_name":  name,
		"updated_at": nowTime,
	}).Where("gift_id", GiftId).Update()
	logrus.Debugf("Update gift. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Update gift. Error : %v", err)
		return false
	}
	return true
}

func (p *GiftService) DelGift(GiftId string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("mch_gift").Data(map[string]interface{}{
		"is_delete":  1,
		"deleted_at": nowTime,
	}).Where("gift_id", GiftId).Update()
	logrus.Debugf("delete gift. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("delete gift. Error : %v", err)
		return false
	}
	return true
}

func (p *GiftService) CheckIsExitsGiftId(GiftId string) bool {
	GiftDao := mch_gift.NewMchGift()
	return GiftDao.CheckIsExitsGiftId(GiftId)
}

func (p *GiftService) CheckIsExitsGiftName(name string) bool {
	GiftDao := mch_gift.NewMchGift()
	return GiftDao.CheckIsExitsGiftName(name)
}

func (p *GiftService) CheckIsBindGiftId(GiftId string) bool {
	deviceGiftDao := mch_device_gift.NewMchDeviceGift()
	return deviceGiftDao.CheckIsBindGiftId(GiftId)
}
