package ff_basepay

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/qiniu/iconv"
	"dollmachine/dollrpc/ff_core/ff_common/ff_convert"
	"dollmachine/dollrpc/ff_core/ff_common/ff_json"
	"dollmachine/dollrpc/ff_core/ff_common/ff_random"
	"dollmachine/dollrpc/ff_core/ff_paygate/ff_payconf"
	"io"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type BasePay struct {
	PayGate
	FuyouService string
}

func (p *BasePay) SetRechargeNotifyUrl(notifyUrl string) {
	p.RechargeNotifyUrl = notifyUrl
}

func (p *BasePay) SetService(service string) {
	p.FuyouService = service
}

type xmlMap map[string]string

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

// MarshalXML marshals the map to XML, with each key in the map being a
// tag and it's corresponding value being it's contents.
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

// UnmarshalXML unmarshals the XML into a map of string to strings,
// creating a key in the map for each tag and setting it's value to the
// tags contents.
//
// The fact this function is on the pointer of Map is important, so that
// if m is nil it can be initialized, which is often the case if m is
// nested in another xml structurel. This is also why the first thing done
// on the first line is initialize it.
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

func (p *BasePay) getPriKey(in []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(in)
	if block == nil {
		return nil, errors.New("数据加密密钥出错, in : " + string(in))
	}
	pri, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("数据加密密钥出错, in : " + string(in) + ", Error : " + err.Error())
	} else {
		return pri.(*rsa.PrivateKey), nil
	}
}

func (p *BasePay) getPubKey(in []byte) (*rsa.PublicKey, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		return nil, errors.New("数据加密公钥出错, in : " + string(in))
	}
	pub, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return nil, errors.New("数据加密公钥出错, in : " + string(in) + ", Error : " + err.Error())
	} else {
		return pub.(*rsa.PublicKey), nil
	}
}

func (p *BasePay) Post(sendMap map[string]string, postUrl string, respData interface{}) error {
	sign, err := p.Sign(sendMap)
	if err != nil {
		return err
	}
	sendMap["sign"] = sign

	xmlByte, _ := xml.MarshalIndent(xmlMap(sendMap), "", "  ")
	sendXml := "<?xml version=\"1.0\" encoding=\"GBK\" standalone=\"yes\"?>" + string(xmlByte)
	fmt.Println(sendXml)
	sendXml = strings.Replace(sendXml, "<xmlMap>", "<xml>", 1)
	fmt.Println(sendXml)
	sendXml = strings.Replace(sendXml, "</xmlMap>", "</xml>", 1)

	//转码成gbk
	cd, err := iconv.Open("gbk", "utf-8")
	if err != nil {
		return errors.New("数据包编码初始化失败, ff_fuyou.BasePay.Post.Open.失败, Error : " + err.Error() + ",sendMap:" + ff_json.MarshalToStringNoError(sendMap) + ",postUrl:" + postUrl)
	}
	//gbkSendXml := cd.ConvString(sendXml)

	var outbuf [1024]byte
	s1, _, err := cd.Conv([]byte(sendXml), outbuf[:])
	if err != nil {
		return errors.New("数据包编码转码失败, ff_fuyou.BasePay.Post.Open.失败, Error : " + err.Error() + ",sendMap : " + ff_json.MarshalToStringNoError(sendMap) + ",postUrl:" + postUrl)
	}
	gbkSendXml := string(s1)

	cd.Close()

	//encode后再encode
	value := url.Values{}
	value.Add("req", url.QueryEscape(gbkSendXml))

	fmt.Println("gbkSendXml", gbkSendXml)

	request := gorequest.New()
	resp, body, errs := request.Post(postUrl).Timeout(30 * time.Second).Send(value.Encode()).End()
	if errs != nil {
		return errors.New("订单请求失败，网络异常, ff_fuyou.BasePay.Post.StatusCode.失败 ,sendMap : " + ff_json.MarshalToStringNoError(sendMap) + ",postUrl : " + postUrl + ",Error : " + ff_json.MarshalToStringNoError(errs))
	}
	//fmt.Println("resp", resp.Status, body)
	if resp.StatusCode != 200 {
		return errors.New("订单请求失败，状态码不对, ff_fuyou.BasePay.Post.StatusCode.失败, sendMap : " + ff_json.MarshalToStringNoError(sendMap) + ", postUrl : " + postUrl + ", resp.StatusCode : " + ff_convert.Int64ToStr(int64(resp.StatusCode)))
	}

	body, err = url.QueryUnescape(body)
	//fmt.Println("body",body )
	if err != nil {
		return errors.New("数据包解码失败, ff_fuyou.BasePay.Post.QueryUnescape.失败, Error : " + err.Error() + ", sendMap : " + ff_json.MarshalToStringNoError(sendMap) + ", postUrl : " + postUrl)
	}
	//解析xml
	respMap, err := p.parseRespBody(sendMap, postUrl, body, respData)
	if err != nil {
		return err
	}
	fmt.Println("respMap", ff_json.MarshalToStringNoError(respMap))

	if codeValue, ok := respMap["result_code"]; !ok || codeValue != "000000" {
		showErrMsg := "服务商返回错误"
		if msgValue, ok := respMap["result_msg"]; ok {
			showErrMsg = showErrMsg + ":" + msgValue
		}
		return errors.New(showErrMsg + ", ff_fuyou.BasePay.Post.QueryUnescape.失败, respMap : " + ff_json.MarshalToStringNoError(respMap) + ", sendMap : " + ff_json.MarshalToStringNoError(sendMap) + ", postUrl : " + postUrl)
	}

	//验证xml sign
	_, err = p.Verify(respMap)
	if err != nil {
		return err
	}
	return nil
}

func (p *BasePay) parseRespBody(sendMap map[string]string, postUrl string, body string, respData interface{}) (map[string]string, error) {
	cd, err := iconv.Open("utf-8", "gbk")
	if err != nil {
		return nil, errors.New("数据包解码初始化失败, ff_fuyou.BasePay.parseRespBody.Open.失败, Error : " + err.Error() + ", sendMap : " + ff_json.MarshalToStringNoError(sendMap) + ", postUrl : " + postUrl)
	}

	var outbuf [1024]byte
	s1, _, err := cd.Conv([]byte(body), outbuf[:])
	if err != nil {
		return nil, errors.New("数据包编码转码失败, ff_fuyou.BasePay.parseRespBody.Conv.失败, Error : " + err.Error() + ", sendMap : " + ff_json.MarshalToStringNoError(sendMap) + ", postUrl : " + postUrl)
	}
	utf8Body := string(s1)
	//fmt.Println("utf8Body", utf8Body, respData)
	cd.Close()
	utf8Body = strings.Replace(utf8Body, "<?xml version=\"1.0\" encoding=\"GBK\" standalone=\"yes\"?>", "", 1)
	utf8Body = strings.Replace(utf8Body, "<xml>", "<xmlMap>", 1)
	utf8Body = strings.Replace(utf8Body, "</xml>", "</xmlMap>", 1)
	//utf8Body = strings.Replace(utf8Body, "\n", "", -1)

	var respMap map[string]string
	xml.Unmarshal([]byte(utf8Body), (*xmlMap)(&respMap))
	fmt.Println("utf8Bodyutf8Body", utf8Body, respMap)
	ff_json.Unmarshal(ff_json.MarshalToStringNoError(respMap), respData)

	return respMap, nil
}

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

func (p *BasePay) Verify(respMap map[string]string) (bool, error) {
	public, err := p.getPubKey(p.PayConfMeta.RsaPublic)
	if err != nil {
		return false, err
	}
	sortStr := p.getMapSort(respMap)

	//转义成gbk之后再加密
	cd, err := iconv.Open("gbk", "utf-8")
	if err != nil {
		return false, errors.New("密钥转码失败, ff_fuyou.BasePay.Verify.Open.失败, Error : " + err.Error() + ", buf : " + sortStr)
	}
	gbkBufStr := cd.ConvString(sortStr)
	cd.Close()

	sign, err := base64.StdEncoding.DecodeString(respMap["sign"])
	if err != nil {
		return false, errors.New("密钥验证解码失败, ff_fuyou.BasePay.Verify.DecodeString.失败, Error : " + err.Error() + ", buf : " + sortStr + ", sign : " + respMap["sign"])
	}

	h := crypto.Hash.New(crypto.MD5)
	h.Write([]byte(gbkBufStr))
	hashed := h.Sum(nil)
	fmt.Println("sign", sign, respMap["sign"])
	err = rsa.VerifyPKCS1v15(public, crypto.MD5, hashed, sign)
	if err != nil {
		return false, errors.New("密钥验证失败, ff_fuyou.BasePay.Verify.VerifyPKCS1v15.失败, Error : " + err.Error() + ", buf : " + sortStr + ", sign:" + respMap["sign"])
	}
	return true, nil
}

func (p *BasePay) Sign(sendMap map[string]string) (string, error) {
	private, err := p.getPriKey(p.PayConfMeta.RsaPrivate)
	if err != nil {
		return "", err
	}

	sortStr := p.getMapSort(sendMap)

	//转义成gbk之后再加密
	cd, err := iconv.Open("gbk", "utf-8")
	if err != nil {
		return "", errors.New("密钥转码失败, ff_fuyou.BasePay.Sign.Open.失败, buf : " + sortStr + ", Error : " + err.Error())
	}
	gbkBufStr := cd.ConvString(sortStr)
	cd.Close()

	h := crypto.Hash.New(crypto.MD5)
	h.Write([]byte(gbkBufStr))
	hashed := h.Sum(nil)

	// 进行rsa加密签名
	signedData, err := rsa.SignPKCS1v15(rand.Reader, private, crypto.MD5, hashed)
	if err != nil {
		return "", errors.New("数据加密密钥出错, ff_fuyou.BasePay.getPriKey.失败, Error : " + err.Error() + " ,sendMap : " + ff_json.MarshalToStringNoError(sendMap))
	}
	return base64.StdEncoding.EncodeToString(signedData), nil
}

func (p *BasePay) parseAmount(amountStr string) (float64, error) {
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, errors.New("订单金额解析失败, ff_fuyou.BasePay.parseAmount.失败, Error : " + err.Error() + ", amountStr : " + amountStr)
	}
	return amount, nil
}

func (p *BasePay) getRechargeStatus(transStat string) (int, int) {
	//SUCCESS—支付成功 REFUND—转入退款 NOTPAY—未支付 CLOSED—已关闭 REVOKED—已撤销（刷卡支付） USERPAYING--用户支付中  PAYERROR--支付失败(其他原因，如银行返回失败)
	if transStat == "SUCCESS" {
		return RechargeStatusPaid, RechargeIsPaidYesPaid
	} else if transStat == "REFUND" {
		return RechargeStatusCreated, RechargeUserPayFail
	} else if transStat == "NOTPAY" {
		return RechargeStatusCreated, RechargeIsPaidNoPaid
	} else if transStat == "CLOSED" {
		return RechargeStatusCreated, RechargeUserPayFail
	} else if transStat == "REVOKED" {
		return RechargeStatusCreated, RechargeUserPayFail
	} else if transStat == "USERPAYING" {
		return RechargeStatusCreated, RechargeUserPayIng
	} else if transStat == "PAYERROR" {
		return RechargeStatusCreated, RechargeUserPayFail
	} else {
		return RechargeStatusCreated, RechargeUserPayFail
	}
}

func (p *BasePay) RechargeQuery(pmtRecharge *PmtRecharge, pmtRechargeExt *PmtRechargeExt) (*QueryRechargeResponse, error) {
	var sendMap = map[string]string{
		"version":        "1",
		"ins_cd":         p.PayConfMeta.MchPlatform, //接入机构在富友的唯一代码
		"mchnt_cd":       p.PayConfMeta.MchId,       //富友分配给二级商户的商户号
		"term_id":        "dollmachine",             //终端号
		"random_str":     ff_random.KrandAll(16),    //ff_random.KrandAll(16),
		"sign":           "",
		"order_type":     "WECHAT",                                      //订单类型:ALIPAY , WECHAT, JD(京东钱包),QQ(QQ钱包),UNIONPAY
		"mchnt_order_no": strconv.FormatInt(pmtRecharge.RechargeId, 10), //strconv.FormatInt(pmtRecharge.RechargeId, 10), //todo test
	}

	var commonQueryResp CommonQueryResp
	err := p.Post(sendMap, ff_payconf.CommonQuery, &commonQueryResp)
	if err != nil {
		return nil, err
	}

	if commonQueryResp.MchntOrderNo != ff_convert.Int64ToStr(pmtRecharge.RechargeId) {
		return nil, errors.New("查询数据不匹配, ff_fuyou.BasePay.RechargeQuery.失败, sendMap:" + ff_json.MarshalToStringNoError(sendMap) + ", commonQuery : " + ff_json.MarshalToStringNoError(commonQueryResp))
	}

	amount, err := p.parseAmount(commonQueryResp.OrderAmt)
	if err != nil {
		return nil, err
	}

	var queryRechargeResponse = &QueryRechargeResponse{}
	queryRechargeResponse.RechargeId = commonQueryResp.MchntOrderNo
	queryRechargeResponse.Amount = amount
	queryRechargeResponse.TransactionNo = commonQueryResp.TransactionId
	queryRechargeResponse.OrderNo = commonQueryResp.ReservedFyTradeNo
	queryRechargeResponse.Status, queryRechargeResponse.IsPaid = p.getRechargeStatus(commonQueryResp.TransStat)
	queryRechargeResponse.PaidAt = 0
	if queryRechargeResponse.IsPaid == RechargeIsPaidYesPaid {
		queryRechargeResponse.PaidAt = time.Now().Unix()
	}
	queryRechargeResponse.SnapShot = make(map[string]string)
	queryRechargeResponse.SnapShot["resp"] = ff_json.MarshalToStringNoError(commonQueryResp)

	return queryRechargeResponse, nil
}

func (p *BasePay) ParseRechargeNotify(notifyStr string) (*QueryRechargeResponse, error) {
	var rechargeNotifyResp RechargeNotifyResp
	_, err := p.parseRespBody(map[string]string{}, "ParseRechargeNotify", notifyStr, &rechargeNotifyResp)
	if err != nil {
		return nil, err
	}
	fmt.Println("notifyStr start : ", notifyStr, "end")
	//fmt.Println("rechargeNotifyResp", rechargeNotifyResp, &rechargeNotifyResp)
	amount, err := p.parseAmount(rechargeNotifyResp.OrderAmt)
	if err != nil {
		return nil, err
	}

	var queryRechargeResponse = &QueryRechargeResponse{}
	queryRechargeResponse.RechargeId = rechargeNotifyResp.MchntOrderNo
	queryRechargeResponse.Amount = amount
	queryRechargeResponse.TransactionNo = rechargeNotifyResp.TransactionId
	queryRechargeResponse.OrderNo = rechargeNotifyResp.ReservedFyTradeNo
	queryRechargeResponse.Status = RechargeStatusCreated
	queryRechargeResponse.IsPaid = RechargeIsPaidNoPaid
	queryRechargeResponse.PaidAt = 0
	if rechargeNotifyResp.TxnFinTs != "" {
		queryRechargeResponse.PaidAt = time.Now().Unix()
		queryRechargeResponse.Status = RechargeStatusPaid
		queryRechargeResponse.IsPaid = RechargeIsPaidYesPaid
	}
	queryRechargeResponse.SnapShot = make(map[string]string)
	queryRechargeResponse.SnapShot["resp"] = ff_json.MarshalToStringNoError(rechargeNotifyResp)
	queryRechargeResponse.RespContent = "1"
	queryRechargeResponse.RespHeader = "application/text;charset=utf-8"
	return queryRechargeResponse, nil
}

func (p *BasePay) ValidRechargeNotify(notifyStr string) (bool, error) {
	var rechargeNotifyResp RechargeNotifyResp
	respMap, err := p.parseRespBody(map[string]string{}, "ValidRechargeNotify", notifyStr, rechargeNotifyResp)
	if err != nil {
		return false, err
	}
	_, err = p.Verify(respMap)
	if err != nil {
		return false, err
	}
	return true, nil
}
