package ff_fuyou_wechat

import (
	"fmt"
	"dollmachine/dolluser/ff_common/ff_common/ff_json"
	"strings"
	"testing"
)

//action=wechatConfigSet 微信参数配置接口
func TestXmlMap_MarshalXML(t *testing.T) {
	var arr []string
	configStr := ff_json.MarshalToStringNoError(map[string]string{
		"mchntCd":        "0005840F1382995",
		"jsapiPath":      "https://wawafront.tunnel.aioil.cn/dist/",
		"subAppid":       "wxc7d98f96c6bb4a79",
		"subscribeAppid": "wxc7d98f96c6bb4a79",
	})
	arr = append(arr, strings.Replace(configStr, "\"", "", -1))
	req := map[string]interface{}{
		"traceNo":    "1525490056905",
		"agencyType": 0,
		"insCd":      "08M0025639",
		"configs":    arr,
	}

	var mapArr map[string]interface{}
	NewBaseWechat().Post(req, configStr, "http://www-1.fuiou.com:28090/wmp/wxMchntMng.fuiou?action=wechatConfigSet", &mapArr)
	fmt.Println(mapArr)
}

/**
action=wechatConfigGet 微信参数查询接口
*/
func TestBaseWechat_Post2(t *testing.T) {
	var arr []string
	configStr := ff_json.MarshalToStringNoError(map[string]string{"mchntCd": "0005840F1382995"})
	arr = append(arr, strings.Replace(configStr, "\"", "", -1))
	req := map[string]interface{}{
		"traceNo":    "1525490056905",
		"agencyType": 0,
		"insCd":      "08M0025639",
		"configs":    arr,
	}

	var mapArr map[string]interface{}
	NewBaseWechat().Post(req, configStr, "http://www-1.fuiou.com:28090/wmp/wxMchntMng.fuiou?action=wechatConfigGet", &mapArr)
}
