package ff_mqttmsg

import (
	"fmt"
	"dollmachine/dollmqtt/ff_common/ff_json"
	"dollmachine/dollmqtt/ff_common/ff_mqttmsg"
	"dollmachine/dollmqtt/ff_common/ff_repository/ff_redis/ff_socket"
	"math/rand"
	"testing"
	"time"
)

func TestDecodeJson(t *testing.T) {
	fmt.Println(DecodeJson("{\"action\":\"over\",\"content\":{\"cost\":\"0.2\",\"costInt\":20,\"customerName\":\"gugu\",\"endTime\":\"1970-06-25  19:22:58\",\"endTimeLong\":15160978110,\"id\":0,\"isAward\":false,\"machineId\":\"9999\"},\"imei\":\"9999\",\"status\":\"1\"}"))
}

func TestEncode(t *testing.T) {
	var ctlPkg *CtlPkg
	ctlPkg = NewCtlPkgMove("999", "move", "1")

	fmt.Println("move json", ff_json.MarshalToStringNoError(ctlPkg), ff_mqttmsg.GetPublishTopic("999"))
}

func TestDecodeJson2(t *testing.T) {
	var Base *BasePkg
	str := "{\"action\":\"HEART\",\"content\":{\"cameraNum\":2,\"cameraWorking\":2,\"deviceId\":9999,\"mac\":\"807b85cf1c3b\",\"msg\":\"avio_open2 error() error -113: Could not open \u0027null\u0027\",\"status\":2,\"version\":\"1.0\"},\"imei\":\"9999\",\"status\":\"1\"}"
	//str := "{\"action\":\"HEART\",\"content\":\"content\",\"imei\":\"999\",\"status\":\"1\"}"
	//str := "{\"content\":{\"mac\":\"mac\"},\"action\":\"HEART\",\"imei\":\"999\",\"status\":\"1\"}"
	ff_json.Unmarshal(str, &Base)
	fmt.Println(ff_json.MarshalToStringNoError(Base))
}

func TestEncodeErr(t *testing.T) {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(100)
	fmt.Println(rnd)
}

func TestDecodeJson3(t *testing.T) {
	str := "{\"device_id\":9999,\"is_award\":1,\"time_stamp\":1518261554}"
	fmt.Println(str)
	var gameRes *ff_socket.GameRes
	ff_json.Unmarshal(str, &gameRes)
	fmt.Println(gameRes.DeviceId)
}

type BaseDataPkg struct {
	Action  string `json:"action"`
	Content string `json:"content"`
}

func TestGetPublishTopic(t *testing.T) {
	str := "{\"action\":\"abc\",\"content\":\"ddd\",\"time\":\"ddd\"}"
	var data *BaseDataPkg
	ff_json.Unmarshal(str, &data)
	fmt.Println(data.Action)
	fmt.Println(data.Content)
}
