package room

import (
	"errors"
	"fmt"
	"github.com/Unknwon/com"
	"dollmachine/dolluser/ff_common/ff_page"
	"dollmachine/dolluser/ff_model/mch_device"
	"dollmachine/dolluser/ff_model/mch_device_gift"
	"dollmachine/dolluser/ff_model/mch_gift"
	"dollmachine/dolluser/ff_model/mch_room"
	"dollmachine/dolluser/ff_model/mch_room_device"
	"dollmachine/dolluser/ff_model/pmt_play"
	"dollmachine/dolluser/ff_model/usr_user"
)

type RoomService struct {
}

func NewRoomService() *RoomService {
	return &RoomService{}
}

func (p *RoomService) GetRoomInfo(roomId int64) (map[string]interface{}, error) {
	//房间信息
	mchRoomDao := mch_room.NewMchRoom()
	if !mchRoomDao.CheckIsExitsByRoomId(roomId) {
		return nil, errors.New("房间不存在")
	}
	mchRoomInfo, _ := mchRoomDao.GetMchRoomInfoByRoomId(roomId, "merchant_id,nick_name,thumbnail")
	merchantId, _ := mchRoomInfo["merchant_id"]
	roomName, _ := mchRoomInfo["nick_name"]
	roomThumbnail, _ := mchRoomInfo["thumbnail"]

	//房间与设备关联
	mchRoomDeviceDao := mch_room_device.NewMchRoomDevice()
	if !mchRoomDeviceDao.CheckIsExitsByRoomId(roomId) {
		return nil, errors.New("房间未与设备关联")
	}
	mchRoomDeviceInfo, _ := mchRoomDeviceDao.GetMchRoomDeviceInfoByRoomId(roomId, "device_id")
	deviceId, _ := mchRoomDeviceInfo["device_id"]
	deviceIdInt64, _ := com.StrTo(fmt.Sprint(deviceId)).Int64()
	if deviceIdInt64 == 0 {
		return nil, errors.New("设备ID不能为空")
	}

	//设备信息
	mchDeviceDao := mch_device.NewMchDevice()
	if !mchDeviceDao.CheckIsExitsByDeviceId(deviceIdInt64) {
		return nil, errors.New("设备不存在")
	}
	mchDeviceInfo, _ := mchDeviceDao.GetMchDeviceInfoByDeviceId(deviceIdInt64, "nick_name,coin")
	deviceName, _ := mchDeviceInfo["nick_name"]
	coin, _ := mchDeviceInfo["coin"]

	//设备与礼品关联
	mchDeviceGiftDao := mch_device_gift.NewMchRoomDevice()
	if !mchDeviceGiftDao.CheckIsExitsByDeviceId(deviceIdInt64) {
		return nil, errors.New("未设置礼品信息")
	}
	mchDeviceGiftInfo, _ := mchDeviceGiftDao.GetMchDeviceGiftInfoByRoomId(deviceIdInt64, "gift_id,gift_stock")
	giftId, _ := mchDeviceGiftInfo["gift_id"]
	giftIdInt64, _ := com.StrTo(fmt.Sprint(giftId)).Int64()
	if giftIdInt64 == 0 {
		return nil, errors.New("礼品ID不能为空")
	}
	giftStock, _ := mchDeviceGiftInfo["gift_stock"]

	//礼品信息
	mchGiftDao := mch_gift.NewMchGift()
	if !mchGiftDao.CheckIsExitsByGiftId(giftIdInt64) {
		return nil, errors.New("未设置礼品信息")
	}
	mchGiftInfo, _ := mchGiftDao.GetMchGiftInfoByGiftId(giftIdInt64, "nick_name,thumbnail,gift_type_name")
	giftName, _ := mchGiftInfo["nick_name"]
	giftTypeName, _ := mchGiftInfo["gift_type_name"]
	giftThumbnail, _ := mchGiftInfo["thumbnail"]

	roomInfo := map[string]interface{}{
		"room_id":        roomId,
		"merchant_id":    merchantId,
		"room_name":      roomName,
		"room_thumbnail": roomThumbnail,
		"device_id":      deviceId,
		"gift_id":        giftId,
		"gift_stock":     giftStock,
		"gift_name":      giftName,
		"gift_type_name": giftTypeName,
		"gift_thumbnail": giftThumbnail,
		"device_name":    deviceName,
		"coin":           coin,
	}

	return roomInfo, nil
}

func (p *RoomService) GetRoomList(merchantId int64, offset int, pageSize int, totalSize int) ([]map[string]interface{}, *ff_page.Page, error) {
	page := ff_page.NewPage(offset, pageSize, totalSize)
	mchRoomDao := mch_room.NewMchRoom()
	roomList, err := mchRoomDao.GetMchRoomListByMerchantId(merchantId, offset, pageSize)
	if err != nil || len(roomList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}
	for _, val := range roomList {
		roomId, _ := com.StrTo(fmt.Sprint(val["room_id"])).Int64()
		roomInfo, err := p.GetRoomInfo(roomId)
		if err == nil {
			val["room_name"] = roomInfo["room_name"]
			val["room_thumbnail"] = roomInfo["room_thumbnail"]
			val["device_id"] = roomInfo["device_id"]
			val["gift_id"] = roomInfo["gift_id"]
			val["gift_stock"] = roomInfo["gift_stock"]
			val["gift_name"] = roomInfo["gift_name"]
			val["gift_type_name"] = roomInfo["gift_type_name"]
			val["gift_thumbnail"] = roomInfo["gift_thumbnail"]
			val["device_name"] = roomInfo["device_name"]
			val["coin"] = roomInfo["device_name"]
		}
	}

	count, _ := mchRoomDao.GetMchRoomListTotalCount(merchantId)
	page.SetTotalSize(count)
	return roomList, page, nil
}

func (p *RoomService) GetRoomAward(merchantId, roomId int64, limit int) ([]map[string]interface{}, error) {
	//查询设备ID
	roomDeviceInfo, _ := mch_room_device.NewMchRoomDevice().GetMchRoomDeviceInfoByRoomId(roomId, "device_id")
	deviceId, _ := roomDeviceInfo["device_id"]
	devId, _ := com.StrTo(fmt.Sprint(deviceId)).Int64()
	if devId == 0 {
		return nil, errors.New("暂无数据")
	}

	playDao := pmt_play.NewPmtPlay()
	playList, _ := playDao.GetPlayListByMchIdAndDevIdAndAward(merchantId, devId, limit, "user_id, user_name, award_time")
	usrDao := usr_user.NewUsrUser()
	for _, val := range playList {
		userId, _ := com.StrTo(fmt.Sprint(val["user_id"])).Int64()
		avatar, _ := usrDao.GetUsrUserOne(userId, "avatar")["avatar"]
		val["avatar"] = avatar
	}
	return playList, nil
}
