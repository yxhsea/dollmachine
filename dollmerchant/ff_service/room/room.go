package room

import (
	"context"
	"errors"
	"github.com/Unknwon/com"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_page"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"dollmachine/dollmerchant/ff_model/mch_room"
	UniqueId "dollmachine/dollmerchant/proto/unique_id"
	"time"
)

type RoomService struct {
}

func NewRoomService() *RoomService {
	return &RoomService{}
}

func (p *RoomService) GetRoomInfo(RoomId string) (map[string]interface{}, error) {
	RoomDao := mch_room.NewMchRoom()
	one, err := RoomDao.GetMchRoomInfo(RoomId, "*")
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (p *RoomService) GetRoomList(offset, pageSize, totalSize string) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()
	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	RoomDao := mch_room.NewMchRoom()
	RoomList, err := RoomDao.GetMchRoomList(Offset, PageSize, "*")
	if err != nil || len(RoomList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}
	count, _ := RoomDao.GetMchRoomListTotalCount()
	page.SetTotalSize(count)
	return RoomList, page, nil
}

func (p *RoomService) AddRoom(roomName, deviceId, thumbnail, merchantId string) bool {
	//生成RoomId
	cli := UniqueId.NewGenerateUniqueIdService("go.micro.srv.unique_id", ff_vars.RpcSrv.Client())
	rsp, err := cli.GenerateUniqueId(context.TODO(), &UniqueId.UniqueIdRequest{Key: "roomId"})
	if err != nil {
		logrus.Errorf("Generate RoomId error: %v", err)
		return false
	}
	RoomId := rsp.Value

	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	//开启事务
	dbr.Begin()

	//insert mch_room
	_, err = dbr.Table("mch_room").Data(map[string]interface{}{
		"room_id":     RoomId,
		"merchant_id": merchantId,
		"name":        roomName,
		"nick_name":   roomName,
		"thumbnail":   thumbnail,

		"status":     1,
		"created_at": nowTime,
		"updated_at": nowTime,
	}).Insert()
	logrus.Debugf("Insert mch_room. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Insert mch_room. Error : %v", err)
		return false
	}

	//insert mch_room_device
	_, err = dbr.Table("mch_room_device").Data(map[string]interface{}{
		"room_id":     RoomId,
		"device_id":   deviceId,
		"merchant_id": merchantId,

		"status":     1,
		"created_at": nowTime,
		"updated_at": nowTime,
	}).Insert()
	logrus.Debugf("Insert mch_room_device. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Insert mch_room_device. Error : %v", err)
		return false
	}

	//提交事务
	dbr.Commit()
	return true
}

func (p *RoomService) UpdRoom(RoomId, roomName, deviceId, thumbnail string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	//开启事务
	dbr.Begin()

	//Update mch_room
	_, err := dbr.Table("mch_room").Data(map[string]interface{}{
		"room_id":   RoomId,
		"name":      roomName,
		"nick_name": roomName,
		"thumbnail": thumbnail,

		"status":     1,
		"created_at": nowTime,
		"updated_at": nowTime,
	}).Insert()
	logrus.Debugf("Update mch_room. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Update mch_room. Error : %v", err)
		return false
	}

	//Update mch_room_device
	_, err = dbr.Table("mch_room_device").Data(map[string]interface{}{
		"room_id":   RoomId,
		"device_id": deviceId,

		"status":     1,
		"created_at": nowTime,
		"updated_at": nowTime,
	}).Insert()
	logrus.Debugf("Update mch_room_device. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Update mch_room_device. Error : %v", err)
		return false
	}

	//提交事务
	dbr.Commit()
	return true
}

func (p *RoomService) DelRoom(RoomId string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	//开启事务
	dbr.Begin()

	//Delete mch_room
	_, err := dbr.Table("mch_room").Data(map[string]interface{}{
		"is_delete":  1,
		"deleted_at": nowTime,
	}).Where("room_id", RoomId).Update()
	logrus.Debugf("Delete mch_room. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Delete mch_room. Error : %v", err)
		return false
	}

	//Delete mch_room_device
	_, err = dbr.Table("mch_room_device").Data(map[string]interface{}{
		"is_delete":  1,
		"deleted_at": nowTime,
	}).Where("room_id", RoomId).Update()
	logrus.Debugf("Delete mch_room_device. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Delete mch_room_device. Error : %v", err)
		return false
	}

	//提交事务
	dbr.Commit()
	return true
}

func (p *RoomService) CheckIsExitsRoomId(RoomId string) bool {
	RoomDao := mch_room.NewMchRoom()
	return RoomDao.CheckIsExitsRoomId(RoomId)
}

func (p *RoomService) CheckIsExitsRoomName(name string) bool {
	RoomDao := mch_room.NewMchRoom()
	return RoomDao.CheckIsExitsRoomName(name)
}
