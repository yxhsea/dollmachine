package place

import (
	"context"
	"errors"
	"github.com/Unknwon/com"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_page"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"dollmachine/dollmerchant/ff_model/mch_device"
	"dollmachine/dollmerchant/ff_model/mch_place"
	UniqueId "dollmachine/dollmerchant/proto/unique_id"
	"time"
)

type PlaceService struct {
}

func NewPlaceService() *PlaceService {
	return &PlaceService{}
}

func (p *PlaceService) GetPlaceInfo(PlaceId string) (map[string]interface{}, error) {
	PlaceDao := mch_place.NewMchPlace()
	one, err := PlaceDao.GetMchPlaceInfo(PlaceId, "*")
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (p *PlaceService) GetPlaceList(offset, pageSize, totalSize string) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()
	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	PlaceDao := mch_place.NewMchPlace()
	PlaceList, err := PlaceDao.GetMchPlaceList(Offset, PageSize, "*")
	if err != nil || len(PlaceList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}
	count, _ := PlaceDao.GetMchPlaceListTotalCount()
	page.SetTotalSize(count)
	return PlaceList, page, nil
}

func (p *PlaceService) AddPlace(name, merchantId string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()

	//生成PlaceId
	cli := UniqueId.NewGenerateUniqueIdService("go.micro.srv.unique_id", ff_vars.RpcSrv.Client())
	rsp, err := cli.GenerateUniqueId(context.TODO(), &UniqueId.UniqueIdRequest{Key: "placeId"})
	if err != nil {
		logrus.Errorf("Generate PlaceId error: %v", err)
		return false
	}
	PlaceId := rsp.Value

	_, err = dbr.Table("mch_place").Data(map[string]interface{}{
		"place_id":    PlaceId,
		"merchant_id": merchantId,
		"name":        name,
		"nick_name":   name,
		"status":      1,
		"created_at":  nowTime,
		"updated_at":  nowTime,
	}).Insert()
	logrus.Debugf("Insert Place. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Insert Place. Error : %v", err)
		return false
	}
	return true
}

func (p *PlaceService) UpdPlace(PlaceId, name string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("mch_place").Data(map[string]interface{}{
		"name":       name,
		"nick_name":  name,
		"updated_at": nowTime,
	}).Where("place_id", PlaceId).Update()
	logrus.Debugf("Update Place. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Update Place. Error : %v", err)
		return false
	}
	return true
}

func (p *PlaceService) DelPlace(PlaceId string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("mch_place").Data(map[string]interface{}{
		"is_delete":  1,
		"deleted_at": nowTime,
	}).Where("place_id", PlaceId).Update()
	logrus.Debugf("delete Place. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("delete Place. Error : %v", err)
		return false
	}
	return true
}

func (p *PlaceService) CheckIsExitsPlaceId(PlaceId string) bool {
	PlaceDao := mch_place.NewMchPlace()
	return PlaceDao.CheckIsExitsPlaceId(PlaceId)
}

func (p *PlaceService) CheckIsExitsPlaceName(name string) bool {
	PlaceDao := mch_place.NewMchPlace()
	return PlaceDao.CheckIsExitsPlaceName(name)
}

func (p *PlaceService) CheckIsBindPlaceId(PlaceId string) bool {
	deviceDao := mch_device.NewMchDevice()
	return deviceDao.CheckIsBindPlaceId(PlaceId)
}
