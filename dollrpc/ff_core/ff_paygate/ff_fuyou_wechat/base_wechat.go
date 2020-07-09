package ff_fuyou_wechat

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/qiniu/iconv"
	"dollmachine/dollrpc/ff_core/ff_common/ff_convert"
	"dollmachine/dollrpc/ff_core/ff_common/ff_json"
	"net/url"
	"sort"
	"strings"
	"time"
)

type BaseWechat struct {
}

func NewBaseWechat() *BaseWechat {
	return &BaseWechat{}
}

func (p *BaseWechat) Post(sendMap map[string]interface{}, configStr string, postUrl string, mapArr *map[string]interface{}) error {
	var array []string
	sendMap["sign"], _ = p.sign(sendMap)
	sendMap["configs"] = append(array, configStr)
	fmt.Println("__sign : ", sendMap["sign"])

	reqData := strings.Replace(strings.Replace(strings.Replace(ff_json.MarshalToStringNoError(sendMap), "[\"", "[", -1), "\"]", "]", -1), "\\", "", -1)
	fmt.Println(reqData)

	//encode后再encode
	value := url.Values{}
	value.Add("req", reqData)

	request := gorequest.New()
	resp, body, errs := request.Post(postUrl).Timeout(30 * time.Second).Send(value.Encode()).End()
	if errs != nil {
		return errors.New("请求失败，网络异常, ff_fuyou_wechat.BaseWechat.Post.StatusCode.失败 ,sendMap : " + ff_json.MarshalToStringNoError(sendMap) + ",postUrl : " + postUrl + ",Error : " + ff_json.MarshalToStringNoError(errs))
	}
	//fmt.Println("resp", resp.Status, body)
	if resp.StatusCode != 200 {
		return errors.New("请求失败，状态码不对, ff_fuyou_wechat.BaseWechat.Post.StatusCode.失败, sendMap : " + ff_json.MarshalToStringNoError(sendMap) + ", postUrl : " + postUrl + ", resp.StatusCode : " + ff_convert.Int64ToStr(int64(resp.StatusCode)))
	}

	err := json.Unmarshal([]byte(body), &mapArr)
	if err != nil {
		return errors.New("body转换map失败")
	}

	return nil
}

func (p *BaseWechat) sign(sendMap map[string]interface{}) (string, error) {
	sortStr := p.getMapSort(sendMap)

	fmt.Println("sortStr : ", sortStr)

	//转义成gbk之后再加密
	cd, err := iconv.Open("gbk", "utf-8")
	if err != nil {
		return "", errors.New("密钥转码失败, ff_fuyou_wechat.BaseWechat.Sign.Open.失败, buf : " + sortStr + ", Error : " + err.Error())
	}
	gbkBufStr := cd.ConvString(sortStr)
	cd.Close()

	data := []byte(gbkBufStr)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制

	return strings.ToUpper(string(md5str)), nil
}

func (p *BaseWechat) getMapSort(sortMap map[string]interface{}) string {
	keys := make([]string, len(sortMap))
	i := 0
	for k := range sortMap {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	var buf bytes.Buffer
	for _, k := range keys {
		if k != "sign" && !strings.HasPrefix(k, "reserved") {
			buf.WriteString(k)
			buf.WriteString("=")
			buf.WriteString(fmt.Sprint(sortMap[k]))
			buf.WriteString("&")
		}
	}
	bufStr := buf.String() + "key=12345678901234567890123456789012"
	return bufStr
}
