package ff_mqtt_v1

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"dollmachine/dollmqtt/ff_common/ff_json"
	"dollmachine/dollmqtt/ff_common/ff_mqttmsg"
	"dollmachine/dollmqtt/ff_config/ff_vars"
	"strconv"
	"time"
	"dollmachine/dollmqtt/ff_redis"
	"dollmachine/dollmqtt/ff_config/ff_const"
)

//处理心跳包
func HeartPkgHandle(heartPkgStr string) {
	var HearPkg *ff_mqttmsg.HeartPkg
	ff_json.Unmarshal(heartPkgStr, &HearPkg)
	//log.Debug("订阅到的心跳包", heartPkgStr)

	heartValue := map[string]interface{}{
		"device_id" : HearPkg.DeviceID,
		"timestamp" : time.Now().Unix(),
	}
	conn := ff_vars.RedisConn.Get()
	defer conn.Close()
	_, err := ff_redis.NewString().Set(conn,ff_const.DeviceHeartPkg + ":" + HearPkg.DeviceID, ff_json.MarshalToStringNoError(heartValue), 3600 * 24 * 30)
	if err != nil {
		log.Errorf("set device %v heartpkg fail Error : %v ", HearPkg.DeviceID, err.Error())
	}

	//go updateDeviceExt(HearPkg, 1)
	return
}

//处理游戏结果数据上报
func GameOverHandle(deviceId string, PlayOverPkgStr string) {
	//解析游戏结果包数据
	var PlayOverPkg *ff_mqttmsg.PlayOverPkg
	ff_json.Unmarshal(PlayOverPkgStr, &PlayOverPkg)
	log.Debug("订阅到的游戏结果数据包", deviceId, PlayOverPkgStr)

	//游戏中奖结果标识
	var IsAward int
	var AwardTime int64

	if PlayOverPkg.IsAward == 1 { //游戏结果命中,更新当前设备数据为命中礼物
		IsAward = 1
		AwardTime = int64(time.Now().Unix())

		//设备信息
		deviceGiftInfo := getDeviceGiftInfo(deviceId)
		giftId, _ := strconv.ParseInt(fmt.Sprint(deviceGiftInfo["gift_id"]), 10, 64)
		giftName := getGiftName(giftId)

		//更新游戏礼品信息
		updatePlayGiftInfo(giftId, giftName, deviceId)

		//减少设备礼品库存
		reduceDeviceGiftStock(deviceId)
	} else {
		IsAward = 2
	}

	//更新游戏结果
	updatePlayInfo(IsAward, AwardTime, deviceId)
}

//更新设备在线信息
func updateDeviceExt(HearPkg *ff_mqttmsg.HeartPkg, status int) {
	nowTime := time.Now().Unix()
	deviceExtInfo := map[string]interface{}{
		"status":         1,
		"mac":            HearPkg.Mac,
		"camera_num":     HearPkg.CameraNum,
		"camera_working": HearPkg.CameraWorking,
		"updated_at":     nowTime,
		"is_online":      status,
		"sys_version":    HearPkg.Version,
	}
	dbr := ff_vars.Dbr
	_, err := dbr.Table("mch_device_ext").Data(deviceExtInfo).Where("device_id", "=", HearPkg.DeviceID).Update()
	log.Debug("Update deviceExt : ", dbr.LastSql())
	if err != nil {
		log.Errorf("Error: %s", err.Error())
	}
}

//更新游戏结果信息
func updatePlayInfo(IsAward int, AwardTime int64, deviceId string) {
	nowTime := time.Now().Unix()
	sqlStr := fmt.Sprintf("update pmt_play set updated_at=%d,is_award=%d,award_time=%d where device_id=%s order by created_at desc limit 1", nowTime, IsAward, AwardTime, deviceId)
	dbr := ff_vars.Dbr
	_, err := dbr.Execute(sqlStr)
	log.Debug("[updatePlayInfo] Get LastSql : ", dbr.LastSql())
	if err != nil {
		log.Errorf("[updatePlayInfo] Error: %s", err.Error())
	}
}

//更新游戏礼品信息
func updatePlayGiftInfo(giftId int64, giftName string, deviceId string) {
	nowTime := time.Now().Unix()
	sqlStr := fmt.Sprintf("update pmt_play set updated_at=%d,gift_id=%d,gift_name='%s' WHERE device_id = %s order by created_at desc LIMIT 1", nowTime, giftId, giftName, deviceId)
	dbr := ff_vars.Dbr
	_, err := dbr.Execute(sqlStr)
	log.Debug("[updatePlayGiftInfo] Get LastSql : ", dbr.LastSql())
	if err != nil {
		log.Errorf("[updatePlayGiftInfo] Error: %s", err.Error())
	}
}

//获取设备礼品信息
func getDeviceGiftInfo(deviceId string) map[string]interface{} {
	dbr := ff_vars.Dbr
	deviceInfo, err := dbr.Table("mch_device_gift").Fields("gift_id").Where("device_id", "=", deviceId).First()
	log.Debug("[getDeviceGiftInfo] Get LastSql : ", dbr.LastSql())
	if err != nil {
		log.Errorf("[getDeviceGiftInfo] Error: %s", err.Error())
		return nil
	}
	return deviceInfo
}

//减少设备礼品库存
func reduceDeviceGiftStock(deviceId string) {
	var err error
	dbr := ff_vars.Dbr

	//开启事务
	dbr.Begin()
	sqlStr1 := fmt.Sprintf("update mch_device_gift set gift_stock = gift_stock - 1 where device_id = %s and gift_stock > 0", deviceId)
	_, err = dbr.Execute(sqlStr1)
	log.Debug("[reduceDeviceGiftStock] Get LastSql : ", dbr.LastSql())
	if err != nil {
		//回滚
		dbr.Rollback()
		log.Errorf("[reduceDeviceGiftStock] Error: %s", err.Error())
		return
	}

	sqlStr2 := fmt.Sprintf("update mch_device set gift_stock = gift_stock - 1 where device_id = %s and gift_stock > 0", deviceId)
	_, err = dbr.Execute(sqlStr2)
	log.Debug("[reduceDeviceGiftStock]Get LastSql : ", dbr.LastSql())
	if err != nil {
		//回滚
		dbr.Rollback()
		log.Errorf("[reduceDeviceGiftStock] Error: %s", err.Error())
		return
	}
	//提交事务
	dbr.Commit()
}

//礼品名称
func getGiftName(giftId int64) string{
	dbr := ff_vars.Dbr
	giftInfo, err := dbr.Table("mch_gift").Fields("nick_name").Where("gift_id","=",giftId).First()
	log.Debug("[getGiftName] get LastSql : ", dbr.LastSql())
	if err != nil {
		log.Errorf("[getGiftName] Error : %s", err.Error())
	}
	return fmt.Sprint(giftInfo["nick_name"])
}