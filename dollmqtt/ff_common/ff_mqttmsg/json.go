package ff_mqttmsg

import (
	"bytes"
	"dollmachine/dollmqtt/ff_common/ff_json"
)

func DecodeJson(s string) (*BasePkg, error) {
	var basePkg *BasePkg
	var err error

	err = ff_json.Unmarshal(s, &basePkg)
	if err != nil {
		return nil, err
	}

	return basePkg, nil
}

func Encode(basePkg *BasePkg) string {
	return ff_json.MarshalToStringNoError(basePkg)
}

func EncodeErr(imei string, action string, errCode string, errType string, errMsg string) string {
	var ctlPkg *CtlPkg

	ctlPkg = NewCtlPkgErr(imei, action, errCode, errType, errMsg)

	return ff_json.MarshalToStringNoError(ctlPkg)
}

//获取设备mqtt对应的topic
func GetPublishTopic(imei string) string {
	var buf bytes.Buffer
	buf.WriteString("dev/wawaji/")
	buf.WriteString(imei)
	buf.WriteString("/ctl")
	return buf.String()
}
