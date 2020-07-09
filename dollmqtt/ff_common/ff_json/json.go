package ff_json

import (
	"errors"
	"github.com/json-iterator/go"
	"dollmachine/dollmqtt/ff_common/ff_convert"
)

func Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, errors.New("invalid memory address or nil pointer dereference")
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(v)
}

func MarshalToString(v interface{}) (string, error) {
	byte, err := Marshal(v)
	return ff_convert.Byte2Str(byte), err
}

func MarshalToStringNoError(v interface{}) string {
	str, _ := MarshalToString(v)
	return str
}

func Unmarshal(data string, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(ff_convert.Str2ByteByReflect(data), v)
}

func UnmarshalByte(data []byte, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}
