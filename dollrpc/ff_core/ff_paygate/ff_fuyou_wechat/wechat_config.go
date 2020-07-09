package ff_fuyou_wechat

import (
	"dollmachine/dollrpc/ff_core/ff_common/ff_json"
	"strings"
)

const (
	MchntCd        = "0005840F1382995"                         //富友商户号
	JsApiPath      = "https://wawafront.tunnel.aioil.cn/dist/" //子商户公众账号JS API支付授权目录
	SubAppid       = "wxc7d98f96c6bb4a79"                      //子商户SubAPPID
	SubscribeAppid = "wxc7d98f96c6bb4a79"                      //子商户推荐关注公众账号APPID
	InsCd          = "08M0025639"                              //机构号
	AgencyType     = 0

	WechatConfigSetUrl = "http://www-1.fuiou.com:28090/wmp/wxMchntMng.fuiou?action=wechatConfigSet"
	WechatConfigGetUrl = "http://www-1.fuiou.com:28090/wmp/wxMchntMng.fuiou?action=wechatConfigGet"
)

//微信参数配置
func WechatConfigSet() {
	var confArr []string
	configStr := ff_json.MarshalToStringNoError(map[string]string{
		"mchntCd":        MchntCd,
		"jsapiPath":      JsApiPath,
		"subAppid":       SubAppid,
		"subscribeAppid": SubscribeAppid,
	})
	confArr = append(confArr, strings.Replace(configStr, "\"", "", -1))
	req := map[string]interface{}{
		"traceNo":    "1525490056905",
		"agencyType": AgencyType,
		"insCd":      InsCd,
		"configs":    confArr,
	}

	var mapArr map[string]interface{}
	NewBaseWechat().Post(req, configStr, WechatConfigSetUrl, &mapArr)
}

//微信参数查询
func WechatConfigGet() {
	var confArr []string
	configStr := ff_json.MarshalToStringNoError(map[string]string{"mchntCd": MchntCd})
	confArr = append(confArr, strings.Replace(configStr, "\"", "", -1))
	req := map[string]interface{}{
		"traceNo":    "1525490056905",
		"agencyType": AgencyType,
		"insCd":      InsCd,
		"configs":    confArr,
	}
	var mapArr map[string]interface{}
	NewBaseWechat().Post(req, configStr, WechatConfigGetUrl, &mapArr)
}
