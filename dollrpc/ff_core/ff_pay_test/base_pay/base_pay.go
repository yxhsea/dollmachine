package base_pay

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/qiniu/iconv"
	"dollmachine/dollrpc/ff_core/ff_common/ff_convert"
	"dollmachine/dollrpc/ff_core/ff_common/ff_json"
	"io"
	"net/url"
	"sort"
	"strings"
	"time"
)

type BasePay struct {
	PayConfMeta
}

type xmlMap map[string]string

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

//编码成XML
func (m xmlMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	for k, v := range m {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}
	return e.EncodeToken(start.End())
}

//解码XML
func (m *xmlMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = xmlMap{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

//发起支付请求
func (p *BasePay) Post(sendMap map[string]string, postUrl string, respData interface{}) error {
	sign, err := p.Sign(sendMap)
	if err != nil {
		return err
	}
	sendMap["sign"] = sign

	xmlByte, _ := xml.MarshalIndent(xmlMap(sendMap), "", "  ")
	sendXml := "<?xml version=\"1.0\" encoding=\"GBK\" standalone=\"yes\"?>" + string(xmlByte)
	sendXml = strings.Replace(sendXml, "<xmlMap>", "<xml>", 1)
	sendXml = strings.Replace(sendXml, "</xmlMap>", "</xml>", 1)

	//转码成gbk
	cd, err := iconv.Open("gbk", "utf-8")
	if err != nil {
		return fmt.Errorf("数据包编码初始化失败, Post.Open.失败; Error : %s; sendMap : %s; postUrl : %s", err.Error(), ff_json.MarshalToStringNoError(sendMap), postUrl)
	}

	var outbuf [1024]byte
	s1, _, err := cd.Conv([]byte(sendXml), outbuf[:])
	if err != nil {
		return fmt.Errorf("数据包编码转码失败, Post.Open.失败; Error : %s; sendMap : %s; postUrl : %s", err.Error(), ff_json.MarshalToStringNoError(sendMap), postUrl)
	}
	gbkSendXml := string(s1)
	cd.Close()

	//encode后再encode
	value := url.Values{}
	value.Add("req", url.QueryEscape(gbkSendXml))

	request := gorequest.New()
	resp, body, errs := request.Post(postUrl).Timeout(30 * time.Second).Send(value.Encode()).End()
	if errs != nil {
		return fmt.Errorf("订单请求失败,网络异常, Post.StatusCode.失败; sendMap : %s; postUrl : %s; Error : %s", ff_json.MarshalToStringNoError(sendMap), postUrl, ff_json.MarshalToStringNoError(errs))
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("订单请求失败,状态码不对, Post.StatusCode.失败; sendMap : %s; postUrl : %s; StatusCode : %s", ff_json.MarshalToStringNoError(sendMap), postUrl, ff_convert.Int64ToStr(int64(resp.StatusCode)))
	}

	body, err = url.QueryUnescape(body)
	if err != nil {
		return fmt.Errorf("数据包解码失败, Post.QueryUnescape.失败; Error : %s; sendMap : %s; postUrl : %s", err.Error(), ff_json.MarshalToStringNoError(sendMap), postUrl)
	}

	//解析xml
	respMap, err := p.parseRespBody(sendMap, postUrl, body, respData)
	if err != nil {
		return err
	}

	if codeValue, ok := respMap["result_code"]; !ok || codeValue != "000000" {
		showErrMsg := "服务商返回错误"
		if msgValue, ok := respMap["result_msg"]; ok {
			showErrMsg = showErrMsg + ":" + msgValue
		}
		return fmt.Errorf(showErrMsg+"; Post.QueryUnescape.失败; respMap : %s; sendMap : %s; postUrl : %s", ff_json.MarshalToStringNoError(respMap), ff_json.MarshalToStringNoError(sendMap), postUrl)
	}

	//验证xml sign
	_, err = p.Verify(respMap)
	if err != nil {
		return err
	}

	return nil
}

//解析支付请求响应
func (p *BasePay) parseRespBody(sendMap map[string]string, postUrl string, body string, respData interface{}) (map[string]string, error) {
	cd, err := iconv.Open("utf-8", "gbk")
	if err != nil {
		return nil, fmt.Errorf("数据包解码初始化失败, parseRespBody.Open.失败; Error : %s; sendMap : %s; postUrl : %s", err.Error(), ff_json.MarshalToStringNoError(sendMap), postUrl)
	}

	var outbuf [1024]byte
	s1, _, err := cd.Conv([]byte(body), outbuf[:])
	if err != nil {
		return nil, fmt.Errorf("数据包编码转码失败, parseRespBody.Conv.失败; Error : %s; sendMap : %s; postUrl : %s", err.Error(), ff_json.MarshalToStringNoError(sendMap), postUrl)
	}
	utf8Body := string(s1)
	cd.Close()
	utf8Body = strings.Replace(utf8Body, "<?xml version=\"1.0\" encoding=\"GBK\" standalone=\"yes\"?>", "", 1)
	utf8Body = strings.Replace(utf8Body, "<xml>", "<xmlMap>", 1)
	utf8Body = strings.Replace(utf8Body, "</xml>", "</xmlMap>", 1)

	var respMap map[string]string
	xml.Unmarshal([]byte(utf8Body), (*xmlMap)(&respMap))
	ff_json.Unmarshal(ff_json.MarshalToStringNoError(respMap), respData)

	return respMap, nil
}

//获取私钥
func (p *BasePay) getPriKey(in []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(in)
	if block == nil {
		return nil, fmt.Errorf("数据加密密钥出错, in : %s", string(in))
	}
	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("数据加密密钥出错, in : %s; Error : %s", string(in), err.Error())
	} else {
		return pri.(*rsa.PrivateKey), nil
	}
}

//获取公钥
func (p *BasePay) getPubKey(in []byte) (*rsa.PublicKey, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		return nil, fmt.Errorf("数据加密公钥出错, in : %s; Error : %s", string(in), err.Error())
	}
	pub, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("数据加密公钥出错, in : %s; Error : %s", string(in), err.Error())
	} else {
		return pub.(*rsa.PublicKey), nil
	}
}

//Map排序
func (p *BasePay) getMapSort(sortMap map[string]string) string {
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
			buf.WriteString(sortMap[k])
			buf.WriteString("&")
		}
	}
	bufStr := buf.String()
	return bufStr[0:(len(bufStr) - 1)]
}

//验签
func (p *BasePay) Verify(respMap map[string]string) (bool, error) {
	public, err := p.getPubKey(p.PayConfMeta.RsaPublic)
	if err != nil {
		return false, err
	}
	sortStr := p.getMapSort(respMap)

	//转义成gbk之后再加密
	cd, err := iconv.Open("gbk", "utf-8")
	if err != nil {
		return false, fmt.Errorf("密钥转码失败, Verify.Open.失败; Error : %s; buf : %s", err.Error(), sortStr)
	}
	gbkBufStr := cd.ConvString(sortStr)
	cd.Close()

	sign, err := base64.StdEncoding.DecodeString(respMap["sign"])
	if err != nil {
		return false, fmt.Errorf("密钥验证解码失败, Verify.DecodeString.失败; Error : %s; buf : %s; sgin : %s", err.Error(), sortStr, respMap["sign"])
	}

	h := crypto.Hash.New(crypto.MD5)
	h.Write([]byte(gbkBufStr))
	hashed := h.Sum(nil)
	err = rsa.VerifyPKCS1v15(public, crypto.MD5, hashed, sign)
	if err != nil {
		return false, fmt.Errorf("密钥验证失败, Verify.VerifyPKCS1v15.失败; Error : %s; buf : %s; sgin : %s", err.Error(), sortStr, respMap["sign"])
	}

	return true, nil
}

//签名
func (p *BasePay) Sign(sendMap map[string]string) (string, error) {
	private, err := p.getPriKey(p.PayConfMeta.RsaPrivate)
	if err != nil {
		return "", err
	}
	sortStr := p.getMapSort(sendMap)

	//转义成gbk之后再加密
	cd, err := iconv.Open("gbk", "utf-8")
	if err != nil {
		return "", fmt.Errorf("密钥转码失败, Sign.Open.失败; buf : %s; Error : %s", sortStr, err.Error())
	}
	gbkBufStr := cd.ConvString(sortStr)
	cd.Close()

	h := crypto.Hash.New(crypto.MD5)
	h.Write([]byte(gbkBufStr))
	hashed := h.Sum(nil)

	// 进行rsa加密签名
	signedData, err := rsa.SignPKCS1v15(rand.Reader, private, crypto.MD5, hashed)
	if err != nil {
		return "", fmt.Errorf("数据加密密钥出错, getPriKey.失败; Error : %s; sendMap : %s", err.Error(), ff_json.MarshalToStringNoError(sendMap))
	}

	return base64.StdEncoding.EncodeToString(signedData), nil
}
