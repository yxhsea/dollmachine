package ff_mqtt_v1

import (
	log "github.com/sirupsen/logrus"
	"dollmachine/dollmqtt/ff_common/ff_json"
	"dollmachine/dollmqtt/ff_common/ff_mqttmsg"
	"dollmachine/dollmqtt/ff_config/ff_vars"
	"time"
)

//设备上线处理
func StartUpPkgHandle(startUpPkgStr string) {
	var StartUpPkg *ff_mqttmsg.HeartPkg
	ff_json.Unmarshal(startUpPkgStr, &StartUpPkg)
	log.Debug("订阅到的上线心跳包", startUpPkgStr)
	go updateDeviceExt(StartUpPkg, 1)
	//TODO::设置房间上线处理
	updateRoomStatus(StartUpPkg.DeviceID, 1)
	return
}

//设备掉线处理
func WillDownPkgHandle(willDownPkgStr string) {
	var WillDownPkg *ff_mqttmsg.HeartPkg
	ff_json.Unmarshal(willDownPkgStr, &WillDownPkg)
	log.Debug("订阅到的掉线心跳包", willDownPkgStr)
	go updateDeviceExt(WillDownPkg, 2)
	//TODO::设置房间掉线处理
	updateRoomStatus(WillDownPkg.DeviceID, 2)
	return
}

//更新房间状态
func updateRoomStatus(deviceId string, status int) {
	dbr := ff_vars.Dbr
	roomDeviceInfo, err := dbr.Table("mch_room_device").Fields("room_id").Where("device_id", "=", deviceId).First()
	log.Info("Get mch_room_device room_id : ", dbr.LastSql())
	if err != nil {
		log.Errorf("Error: %s", err.Error())
		return
	}

	nowTime := time.Now().Unix()
	data := map[string]interface{}{
		"update_at": nowTime,
		"status":    status,
	}
	_, err = dbr.Table("mch_room").Data(data).Where("room_id", "=", roomDeviceInfo["room_id"]).Update()
	log.Info("Update mch_room status : ", dbr.LastSql())
	if err != nil {
		log.Errorf("Error: %s", err.Error())
	}
}
