package pre_create

import "dollmachine/dollrpc/ff_core/ff_pay_test/common_msg"

/**
术语及定义
交易金额: 金额以分为单位，不带小数点
必填:M 必填,O 选填
字段类型: A 字符 N 数字
*/

type PreCreateRequestMsg struct {
	common_msg.CommonRequestMsg
	OrderType        string `json:"order_type"`         //M	A	20	订单类型:ALIPAY , WECHAT, QQ(QQ钱包)，UNIONPAY(银联二维码)，BESTPAY(翼支付)
	NotifyUrl        string `json:"notify_url"`         //M	A	256	通知地址, 接收富友异步通知回调地址，通知url必须为直接可访问的url，不能携带参数
	ReservedSubAppid string `json:"reserved_sub_appid"` //O	A	32	子商户公众号id(后期拓展字段，请先不要填写)
	ReservedLimitPay string `json:"reserved_limit_pay"` //O	A	32	限制支付,no_credit:不能使用信用卡
}

type PreCreateResponseMsg struct {
	common_msg.CommonResponseMsg
	OrderType string `json:"order_type"` //M	A	20	订单类型:ALIPAY , WECHAT, QQ(QQ钱包), UNIONPAY
	SessionId string `json:"session_id"` //O	A	64	预支付交易会话标识, 富友返回支付宝生成的预支付回话标识，用于后续接口调用中使用，该值有效期为2小时
	QrCode    string `json:"qr_code"`    //O	A	64	二维码链接
}
