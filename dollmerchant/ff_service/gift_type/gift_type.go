package gift_type

import (
	"context"
	"errors"
	"github.com/Unknwon/com"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_page"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"dollmachine/dollmerchant/ff_model/mch_gift"
	"dollmachine/dollmerchant/ff_model/mch_gift_type"
	UniqueId "dollmachine/dollmerchant/proto/unique_id"
	"time"
)

type GiftTypeService struct {
}

func NewGiftTypeService() *GiftTypeService {
	return &GiftTypeService{}
}

func (p *GiftTypeService) GetGiftTypeInfo(GiftTypeId string) (map[string]interface{}, error) {
	GiftTypeDao := mch_gift_type.NewMchGiftType()
	one, err := GiftTypeDao.GetMchGiftTypeInfo(GiftTypeId, "*")
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (p *GiftTypeService) GetGiftTypeList(offset, pageSize, totalSize string) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()
	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	GiftTypeDao := mch_gift_type.NewMchGiftType()
	GiftTypeList, err := GiftTypeDao.GetMchGiftTypeList(Offset, PageSize, "*")
	if err != nil || len(GiftTypeList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}
	count, _ := GiftTypeDao.GetMchGiftTypeListTotalCount()
	page.SetTotalSize(count)
	return GiftTypeList, page, nil
}

func (p *GiftTypeService) AddGiftType(name string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()

	//生成GiftTypeId
	cli := UniqueId.NewGenerateUniqueIdService("go.micro.srv.unique_id", ff_vars.RpcSrv.Client())
	rsp, err := cli.GenerateUniqueId(context.TODO(), &UniqueId.UniqueIdRequest{Key: "gift_type_id"})
	if err != nil {
		logrus.Errorf("Generate giftTypeId error: %v", err)
		return false
	}
	GiftTypeId := rsp.Value

	_, err = dbr.Table("mch_gift_type").Data(map[string]interface{}{
		"gift_type_id": GiftTypeId,
		"merchant_id":  0,
		"name":         name,
		"nick_name":    name,
		"status":       1,
		"created_at":   nowTime,
		"updated_at":   nowTime,
	}).Insert()
	logrus.Debugf("Insert device type. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Insert device type. Error : %v", err)
		return false
	}
	return true
}

func (p *GiftTypeService) UpdGiftType(GiftTypeId, name string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("mch_gift_type").Data(map[string]interface{}{
		"name":       name,
		"nick_name":  name,
		"updated_at": nowTime,
	}).Where("gift_type_id", GiftTypeId).Update()
	logrus.Debugf("Update gift type. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Update gift type. Error : %v", err)
		return false
	}
	return true
}

func (p *GiftTypeService) DelGiftType(GiftTypeId string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("mch_gift_type").Data(map[string]interface{}{
		"is_delete":  1,
		"deleted_at": nowTime,
	}).Where("gift_type_id", GiftTypeId).Update()
	logrus.Debugf("delete gift type. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("delete gift type. Error : %v", err)
		return false
	}
	return true
}

func (p *GiftTypeService) CheckIsExitsGiftTypeId(GiftTypeId string) bool {
	GiftTypeDao := mch_gift_type.NewMchGiftType()
	return GiftTypeDao.CheckIsExitsGiftTypeId(GiftTypeId)
}

func (p *GiftTypeService) CheckIsExitsGiftTypeName(name string) bool {
	GiftTypeDao := mch_gift_type.NewMchGiftType()
	return GiftTypeDao.CheckIsExitsGiftTypeName(name)
}

func (p *GiftTypeService) CheckIsBindGiftTypeId(GiftTypeId string) bool {
	giftDao := mch_gift.NewMchGift()
	return giftDao.CheckIsBindGiftTypeId(GiftTypeId)
}
