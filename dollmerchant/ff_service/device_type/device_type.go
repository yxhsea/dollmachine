package device_type

import (
	"context"
	"errors"
	"github.com/Unknwon/com"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_page"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"dollmachine/dollmerchant/ff_model/mch_device_type"
	UniqueId "dollmachine/dollmerchant/proto/unique_id"
	"time"
)

type DeviceTypeService struct {
}

func NewDeviceTypeService() *DeviceTypeService {
	return &DeviceTypeService{}
}

func (p *DeviceTypeService) GetDeviceTypeInfo(deviceTypeId string) (map[string]interface{}, error) {
	deviceTypeDao := mch_device_type.NewMchDeviceType()
	one, err := deviceTypeDao.GetMchDeviceTypeInfo(deviceTypeId, "*")
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (p *DeviceTypeService) GetDeviceTypeList(offset, pageSize, totalSize string) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()
	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	deviceTypeDao := mch_device_type.NewMchDeviceType()
	deviceTypeList, err := deviceTypeDao.GetMchDeviceTypeList(Offset, PageSize, "*")
	if err != nil || len(deviceTypeList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}
	count, _ := deviceTypeDao.GetMchDeviceTypeListTotalCount()
	page.SetTotalSize(count)
	return deviceTypeList, page, nil
}

func (p *DeviceTypeService) AddDeviceType(name string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()

	//deviceTypeId, _ := unique_id.NewUniqueId().GetDeviceTypeId()

	//生成deviceTypeId
	cli := UniqueId.NewGenerateUniqueIdService("go.micro.srv.unique_id", ff_vars.RpcSrv.Client())
	rsp, err := cli.GenerateUniqueId(context.TODO(), &UniqueId.UniqueIdRequest{Key: "merchantId"})
	if err != nil {
		logrus.Errorf("Generate merchantId error: %v", err)
		return false
	}
	deviceTypeId := rsp.Value

	_, err = dbr.Table("mch_device_type").Data(map[string]interface{}{
		"device_type_id": deviceTypeId,
		"merchant_id":    0,
		"name":           name,
		"nick_name":      name,
		"status":         1,
		"created_at":     nowTime,
		"updated_at":     nowTime,
	}).Insert()
	logrus.Debugf("Insert device type. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Insert device type. Error : %v", err)
		return false
	}
	return true
}

func (p *DeviceTypeService) UpdDeviceType(deviceTypeId, name string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("mch_device_type").Data(map[string]interface{}{
		"name":       name,
		"nick_name":  name,
		"updated_at": nowTime,
	}).Where("device_type_id", deviceTypeId).Update()
	logrus.Debugf("Update device type. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Update device type. Error : %v", err)
		return false
	}
	return true
}

func (p *DeviceTypeService) CheckIsExitsDeviceTypeId(deviceTypeId string) bool {
	deviceTypeDao := mch_device_type.NewMchDeviceType()
	return deviceTypeDao.CheckIsExitsDeviceTypeId(deviceTypeId)
}

func (p *DeviceTypeService) CheckIsExitsDeviceTypeName(name string) bool {
	deviceTypeDao := mch_device_type.NewMchDeviceType()
	return deviceTypeDao.CheckIsExitsDeviceTypeName(name)
}
