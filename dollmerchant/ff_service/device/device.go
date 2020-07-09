package device

import (
	"context"
	"errors"
	"github.com/Unknwon/com"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_page"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"dollmachine/dollmerchant/ff_model/mch_device"
	"dollmachine/dollmerchant/ff_model/mch_room_device"
	UniqueId "dollmachine/dollmerchant/proto/unique_id"
	"time"
)

type DeviceService struct {
}

func NewDeviceService() *DeviceService {
	return &DeviceService{}
}

func (p *DeviceService) GetDeviceInfo(DeviceId string) (map[string]interface{}, error) {
	DeviceDao := mch_device.NewMchDevice()
	one, err := DeviceDao.GetMchDeviceInfo(DeviceId, "*")
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (p *DeviceService) GetDeviceList(offset, pageSize, totalSize string) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()
	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	DeviceDao := mch_device.NewMchDevice()
	DeviceList, err := DeviceDao.GetMchDeviceList(Offset, PageSize, "*")
	if err != nil || len(DeviceList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}
	count, _ := DeviceDao.GetMchDeviceListTotalCount()
	page.SetTotalSize(count)
	return DeviceList, page, nil
}

func (p *DeviceService) BindDevice(deviceMac, deviceName, deviceTypeId, deviceTypeName, placeId, placeName, giftId, giftName, giftStock, line, merchantId string) bool {
	//生成DeviceId
	cli := UniqueId.NewGenerateUniqueIdService("go.micro.srv.unique_id", ff_vars.RpcSrv.Client())
	rsp, err := cli.GenerateUniqueId(context.TODO(), &UniqueId.UniqueIdRequest{Key: "deviceId"})
	if err != nil {
		logrus.Errorf("Generate DeviceId error: %v", err)
		return false
	}
	DeviceId := rsp.Value

	nowTime := time.Now().Unix()

	//开启事务
	dbr := ff_vars.DbConn.GetInstance()
	dbr.Begin()

	//insert mch_device
	_, err = dbr.Table("mch_device").Data(map[string]interface{}{
		"device_id":   DeviceId,
		"merchant_id": merchantId,
		"name":        deviceName,
		"nick_name":   deviceName,
		"mac":         deviceMac,
		"line":        line,

		"device_type_id":   deviceTypeId,
		"device_type_name": deviceTypeName,
		"place_id":         placeId,
		"place_name":       placeName,
		"gift_id":          giftId,
		"gift_name":        giftName,
		"gift_stock":       giftStock,

		"status":     1,
		"created_at": nowTime,
		"updated_at": nowTime,
	}).Insert()
	logrus.Debugf("Insert mch_device. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Insert mch_device. Error : %v", err)
		return false
	}

	//insert mch_device_ext
	_, err = dbr.Table("mch_device_ext").Data(map[string]interface{}{
		"device_id":    DeviceId,
		"merchant_id":  merchantId,
		"mac":          deviceMac,
		"line":         line,
		"machine_type": 1,

		"status":     1,
		"created_at": nowTime,
		"updated_at": nowTime,
	}).Insert()
	logrus.Debugf("Insert mch_device_ext. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Insert mch_device_ext. Error : %v", err)
		return false
	}

	//insert mch_device_gift
	_, err = dbr.Table("mch_device_gift").Data(map[string]interface{}{
		"device_id":   DeviceId,
		"merchant_id": merchantId,
		"gift_id":     giftId,
		"gift_stock":  giftStock,

		"status":     1,
		"created_at": nowTime,
		"updated_at": nowTime,
	}).Insert()
	logrus.Debugf("Insert mch_device_gift. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Insert mch_device_gift. Error : %v", err)
		return false
	}

	//提交事务
	dbr.Commit()

	return true
}

func (p *DeviceService) UpdDevice(deviceId, deviceName, deviceTypeId, deviceTypeName, placeId, placeName, giftId, giftName, giftStock, line string) bool {
	nowTime := time.Now().Unix()

	//开启事务
	dbr := ff_vars.DbConn.GetInstance()
	dbr.Begin()

	//update mch_device
	_, err := dbr.Table("mch_device").Data(map[string]interface{}{
		"name":      deviceName,
		"nick_name": deviceName,
		"line":      line,

		"device_type_id":   deviceTypeId,
		"device_type_name": deviceTypeName,
		"place_id":         placeId,
		"place_name":       placeName,
		"gift_id":          giftId,
		"gift_name":        giftName,
		"gift_stock":       giftStock,

		"updated_at": nowTime,
	}).Where("device_id", deviceId).Update()
	logrus.Debugf("Update mch_device. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Update mch_device. Error : %v", err)
		return false
	}

	//update mch_device_ext
	_, err = dbr.Table("mch_device_ext").Data(map[string]interface{}{
		"line":       line,
		"updated_at": nowTime,
	}).Where("device_id", deviceId).Update()
	logrus.Debugf("Update mch_device_ext. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Update mch_device_ext. Error : %v", err)
		return false
	}

	//mch_device_gift
	_, err = dbr.Table("mch_device_gift").Data(map[string]interface{}{
		"gift_id":    giftId,
		"gift_stock": giftStock,

		"updated_at": nowTime,
	}).Where("device_id", deviceId).Update()
	logrus.Debugf("Update mch_device_gift. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Update mch_device_gift. Error : %v", err)
		return false
	}

	//提交事务
	dbr.Commit()
	return true
}

func (p *DeviceService) UnbindDevice(deviceId string) bool {
	nowTime := time.Now().Unix()

	//开启事务
	dbr := ff_vars.DbConn.GetInstance()
	dbr.Begin()

	//delete mch_device
	_, err := dbr.Table("mch_device").Data(map[string]interface{}{
		"is_delete":  1,
		"deleted_at": nowTime,
	}).Where("device_id", deviceId).Update()
	logrus.Debugf("delete mch_device. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Delete mch_device. Error : %v", err)
		return false
	}

	//delete mch_device_ext
	_, err = dbr.Table("mch_device_ext").Data(map[string]interface{}{
		"is_delete":  1,
		"deleted_at": nowTime,
	}).Where("device_id", deviceId).Update()
	logrus.Debugf("Delete mch_device_ext. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Delete mch_device_ext. Error : %v", err)
		return false
	}

	//delete mch_device_gift
	_, err = dbr.Table("mch_device_gift").Data(map[string]interface{}{
		"is_delete":  1,
		"deleted_at": nowTime,
	}).Where("device_id", deviceId).Update()
	logrus.Debugf("delete mch_device_gift. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("delete mch_device_gift. Error : %v", err)
		return false
	}

	//提交事务
	dbr.Commit()
	return true
}

func (p *DeviceService) CheckIsExitsDeviceId(DeviceId string) bool {
	DeviceDao := mch_device.NewMchDevice()
	return DeviceDao.CheckIsExitsDeviceId(DeviceId)
}

func (p *DeviceService) CheckIsBindDeviceMac(DeviceId string) bool {
	DeviceDao := mch_device.NewMchDevice()
	return DeviceDao.CheckIsBindDeviceMac(DeviceId)
}

func (p *DeviceService) CheckIsExitsDeviceName(name string) bool {
	DeviceDao := mch_device.NewMchDevice()
	return DeviceDao.CheckIsExitsDeviceName(name)
}

func (p *DeviceService) CheckIsBindDeviceId(DeviceId string) bool {
	roomDeviceDao := mch_room_device.NewMchRoomDevice()
	return roomDeviceDao.CheckIsBindDeviceId(DeviceId)
}
