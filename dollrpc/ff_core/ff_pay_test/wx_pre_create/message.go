package wx_pre_create

import "dollmachine/dollrpc/ff_core/ff_pay_test/common_msg"

/**
术语及定义
交易金额: 金额以分为单位，不带小数点
必填:M 必填,O 选填
字段类型: A 字符 N 数字
*/

type WxPreCreateRequestMsg struct {
	common_msg.CommonRequestMsg
	ProductId        string `json:"product_id"`         //O	A	32	商品标识
	NotifyUrl        string `json:"notify_url"`         //M	A	256	通知地址, 接收富友异步通知回调地址，通知url必须为直接可访问的url，不能携带参数 主扫时必填
	LimitPay         string `json:"limit_pay"`          //O	A	32	限制支付,no_credit:不能使用信用卡
	TradeType        string `json:"trade_type"`         //M	A	16	JSAPI--公众号支付、APP--app支付、FWC--支付宝服务窗、QQ--QQ页面支付、QQJSAPI--QQJSAPI、LETPAY-小程序
	Openid           string `json:"openid"`             //O	A	128	用户标识(微信公众号服务商模式必填，其他不填)
	SubOpenid        string `json:"sub_openid"`         //O	A	128	子商户用户标识 支付宝服务窗为用户buyer_id（此场景必填） 微信公众号为用户的openid(非服务商模式必填)
	SubAppid         string `json:"sub_appid"`          //O	A	32	子商户公众号id, 微信交易为商户的app_id(非服务商模式必填)
	ReservedTxnBonus string `json:"reserved_txn_bonus"` //0	N	16	积分抵扣金额,单位为分
}

type WxPreCreateResponseMsg struct {
	common_msg.CommonResponseMsg
	SubMerId              string `json:"sub_mer_id"`              //				支付通道对应的子商户识别码
	SessionId             string `json:"session_id"`              //O	A	64	Trade_type为APP--app支付时，prepayid字段取此字段值
	QrCode                string `json:"qr_code"`                 //O	A	64	二维码链接
	SubAppid              string `json:"sub_appid"`               //O	A	32	子商户公众号id
	SubOpenid             string `json:"sub_openid"`              //O	A	128	子商户用户标识
	SdkAppid              string `json:"sdk_appid"`               //M	A	16	公众号id
	SdkTimestamp          string `json:"sdk_timestamp"`           //M	A	32	时间戳，自1970年1月1日 0点0分0秒以来的秒数
	SdkNoncestr           string `json:"sdk_noncestr"`            //M	A	32	随机字符串
	SdkPackage            string `json:"sdk_package"`             //M	A	128	订单性情扩展字符串
	SdkSigntype           string `json:"sdk_signtype"`            //O	A	32	签名方式, trade_type为JSAPI、LETPAY时才返回
	SdkPaysign            string `json:"sdk_paysign"`             //M	A	64	签名
	SdkPartnerid          string `json:"sdk_partnerid"`           //O	A	32	trade_type为APP时才返回
	ReservedFySettleDt    string `json:"reserved_fy_settle_dt"`   //O	A	8	富友清算日
	ReservedTransactionId string `json:"reserved_transaction_id"` //O	A	64 渠道交易流水号, trade_type为FWC时返回（用于调起支付）
	ReservedPayInfo       string `json:"reserved_pay_info"`       //O	A	不定长	支付参数
}
