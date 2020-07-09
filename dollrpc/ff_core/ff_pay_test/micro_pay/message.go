package micro_pay

import "dollmachine/dollrpc/ff_core/ff_pay_test/common_msg"

/**
术语及定义
交易金额: 金额以分为单位，不带小数点
必填:M 必填,O 选填
字段类型: A 字符 N 数字
*/

type MicroPayRequestMsg struct {
	common_msg.CommonRequestMsg
	OrderType        string `json:"order_type"`         //M	A	20	订单类型:ALIPAY , WECHAT, QQ(QQ钱包), UNIONPAY(银联二维码）,BESTPAY(翼支付)
	AuthCode         string `json:"auth_code"`          //M	A	128	扫码支付授权码，设备读取用户的条码或者二维码信息.
	Sence            string `json:"sence"`              //O	A	1	支付场景,默认1；条码支付:1 声波支付:2
	ReservedSubAppid string `json:"reserved_sub_appid"` //O	A	32	子商户公众号id(后期拓展字段，请先不要填写)
	ReservedLimitPay string `json:"reserved_limit_pay"` //O	A	32	限制支付,no_credit:不能使用信用卡
}

type MicroPayResponseMsg struct {
	common_msg.CommonResponseMsg
	OrderType            string `json:"order_type"`              //M	A	20	订单类型:ALIPAY , WECHAT, QQ(QQ钱包), UNIONPAY
	TotalAmount          string `json:"total_amount"`            //M	N	16	订单金额，分为单位的整数
	BuyerId              string `json:"buyer_id"`                //O	A	128	买家在渠道账号
	TransactionId        string `json:"transaction_id"`          //M	A	64	渠道交易流水号
	AddnInf              string `json:"addn_inf"`                //O	A	100	附加数据
	ReservedMchntOrderNo string `json:"reserved_mchnt_order_no"` //M	A	30	商户订单号, 商户系统内部的订单号
	ReservedFySettleDt   string `json:"reserved_fy_settle_dt"`   //M	A	8	富友清算日
	ReservedCouponFee    string `json:"reserved_coupon_fee"`     //O	A	10	优惠金额（分）
	ReservedBuyerLogonId string `json:"reserved_buyer_logon_id"` //O	A	128	买家在渠道登录账号
	ReservedFundBillList string `json:"reserved_fund_bill_list"` //O	A	不定长	支付宝交易资金渠道,详细渠道
	ReservedIsCredit     string `json:"reserved_is_credit"`      //O	A	8	1表示信用卡或者花呗，0表示其他(非信用方式) 不填，表示未知
}
