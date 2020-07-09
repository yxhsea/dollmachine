package ff_basepay

const (
	RechargeStatusCreated = 1
	RechargeStatusPaid    = 2
	RechargeIsPaidNoPaid  = 1
	RechargeIsPaidYesPaid = 2
	RechargeUserPayIng    = 3
	RechargeUserPayFail   = 4
)

type PreCreateResp struct {
	ResultCode             string `json:"result_code"`               //错误代码, 000000成功,其他详细参见错误列表
	ResultMsg              string `json:"result_msg"`                //错误代码描述
	InsCd                  string `json:"ins_cd"`                    //机构号,接入机构在富友的唯一代码
	MchntCd                string `json:"mchnt_cd"`                  //商户号, 富友分配的商户号
	TermId                 string `json:"term_id"`                   //终端号
	RandomStr              string `json:"random_str"`                //随机字符串
	Sign                   string `json:"sign"`                      //签名, 详见签名生成算法
	OrderType              string `json:"order_type"`                //订单类型:ALIPAY , WECHAT,JD(京东钱包),QQ(QQ钱包), UNIONPAY
	SessionId              string `json:"session_id"`                //预支付交易会话标识, 富友返回支付宝生成的预支付回话标识，用于后续接口调用中使用，该值有效期为2小时
	QrCode                 string `json:"qr_code"`                   //二维码链接
	ReservedFyOrderNo      string `json:"reserved_fy_order_no"`      //富友生成的订单号,需要商户与商户订单号进行关联
	ReservedFyTraceNo      string `json:"reserved_fy_trace_no"`      //富友系统内部追踪号
	ReservedChannelOrderId string `json:"reserved_channel_order_id"` //条码流水号，用户账单二维码对应的流水
}

type WxPreCreateResp struct {
	ResultCode             string `json:"result_code"`               //M 错误代码, 000000成功,其他详细参见错误列表
	ResultMsg              string `json:"result_msg"`                //M 错误代码描述
	InsCd                  string `json:"ins_cd"`                    //M 机构号,接入机构在富友的唯一代码
	MchntCd                string `json:"mchnt_cd"`                  //M 商户号, 富友分配的商户号
	TermId                 string `json:"term_id"`                   //O 终端号
	RandomStr              string `json:"random_str"`                //M	随机字符串
	Sign                   string `json:"sign"`                      //M	签名, 详见签名生成算法
	SubMerId               string `json:"sub_mer_id"`                //	支付通道对应的子商户识别码
	SessionId              string `json:"session_id"`                //O	预支付交易会话标识, 富友返回支付宝生成的预支付回话标识，用于后续接口调用中使用，该值有效期为2小时
	QrCode                 string `json:"qr_code"`                   //O	二维码链接
	SubAppid               string `json:"sub_appid"`                 //O	子商户公众号id
	SubOpenid              string `json:"sub_openid"`                //O	子商户用户标识
	SdkAppid               string `json:"sdk_appid"`                 //M	公众号id
	SdkTimestamp           string `json:"sdk_timestamp"`             //M	时间戳，自1970年1月1日 0点0分0秒以来的秒数
	SdkNoncestr            string `json:"sdk_noncestr"`              //M	随机字符串
	SdkPackage             string `json:"sdk_package"`               //M	订单性情扩展字符串
	SdkSigntype            string `json:"sdk_signtype"`              //O	签名方式, trade_type为JSAPI时才返回
	SdkPaysign             string `json:"sdk_paysign"`               //M	签名
	SdkPartnerid           string `json:"sdk_partnerid"`             //O	trade_type为APP时才返回
	ReservedFyOrderNo      string `json:"reserved_fy_order_no"`      //M	富友生成的订单号,需要商户与商户订单号进行关联
	ReservedFySettleDt     string `json:"reserved_fy_settle_dt"`     //M	富友清算日
	ReservedTransactionId  string `json:"reserved_transaction_id"`   //O	渠道交易流水号, trade_type为FWC时返回
	ReservedFyTraceNo      string `json:"reserved_fy_trace_no"`      //M	富友系统内部追踪号
	ReservedPayInfo        string `json:"reserved_pay_info"`         //O	支付参数
	ReservedChannelOrderId string `json:"reserved_channel_order_id"` //M	条码流水号，用户账单二维码对应的流水
}

type CommonQueryResp struct {
	InsCd                  string `json:"ins_cd"`
	MchntCd                string `json:"mchnt_cd"`
	RandomStr              string `json:"random_str"`
	ResultCode             string `json:"result_code"`
	ResultMsg              string `json:"result_msg"`
	TermId                 string `json:"term_id"`
	OrderType              string `json:"order_type"`
	TransStat              string `json:"trans_stat"`     //SUCCESS—支付成功 REFUND—转入退款 NOTPAY—未支付 CLOSED—已关闭 REVOKED—已撤销（刷卡支付） USERPAYING--用户支付中  PAYERROR--支付失败(其他原因，如银行返回失败)
	OrderAmt               string `json:"order_amt"`      //订单金额, 单位为分
	TransactionId          string `json:"transaction_id"` //渠道订单号
	MchntOrderNo           string `json:"mchnt_order_no"` //商户订单号, 商户系统的订单号，与请求一致
	AddnInf                string `json:"addn_inf"`
	ReservedFySettleDt     string `json:"reserved_fy_settle_dt"`
	ReservedCouponFee      string `json:"reserved_coupon_fee"`
	ReservedBuyerLogonId   string `json:"reserved_buyer_logon_id"`
	ReservedFundBillList   string `json:"reserved_fund_bill_list"`
	ReservedFyOrderNo      string `json:"reserved_fy_order_no"`
	ReservedFyTradeNo      string `json:"reserved_fy_trace_no"` //富友系统内部追踪号
	ReservedFyTermId       string `json:"reserved_fy_term_id"`
	ReservedChannelOrderId string `json:"reserved_channel_order_id"`
}

type RechargeNotifyResp struct {
	ResultCode             string `json:"result_code"`
	ResultMsg              string `json:"result_msg"`
	InsCd                  string `json:"ins_cd"`
	MchntCd                string `json:"mchnt_cd"`
	TermId                 string `json:"term_id"`
	RandomStr              string `json:"random_str"`
	OrderAmt               string `json:"order_amt"` //订单金额, 单位为分
	SettleOrderAmt         string `json:"settle_order_amt"`
	CurrType               string `json:"curr_type"`
	TransactionId          string `json:"transaction_id"` //渠道订单号
	MchntOrderNo           string `json:"mchnt_order_no"` //商户订单号, 商户系统的订单号，与请求一致
	OrderType              string `json:"order_type"`
	TxnFinTs               string `json:"txn_fin_ts"` //支付完成时间, 订单支付时间，格式为yyyyMMddHHmmss
	ReservedFySettleDt     string `json:"reserved_fy_settle_dt"`
	ReservedCouponFee      string `json:"reserved_coupon_fee"`
	ReservedBuyerLogonId   string `json:"reserved_buyer_logon_id"`
	ReservedFundBillList   string `json:"reserved_fund_bill_list"`
	ReservedFyTradeNo      string `json:"reserved_fy_trace_no"` //富友系统内部追踪号
	ReservedChannelOrderId string `json:"reserved_channel_order_id"`
	ReservedAddnInf        string `json:"reserved_addn_inf"` //附加数据
}

type RechargeExtMeta struct {
	OpenId    string `json:"open_id"`
	SubOpenId string `json:"sub_open_id"`
	SubAppId  string `json:"sub_app_id"`
}

type PmtRecharge struct {
	RechargeId     int64   `json:"m_recharge_id"`
	IsPaid         int64   `json:"m_is_paid"`
	IsRefunded     int64   `json:"m_is_refunded"`
	IsBonus        int64   `json:"m_is_bonus"`
	Channel        string  `json:"m_channel"`
	Subject        string  `json:"m_subject"`
	ExpiredAt      int64   `json:"m_expired_at"`
	TransactionNo  string  `json:"m_transaction_no"`
	Currency       string  `json:"m_currency"`
	Amount         float64 `json:"m_amount"`
	AmountSettle   float64 `json:"m_amount_settle"`
	AmountRefunded float64 `json:"m_amount_refunded"`
	ExchangeRate   int64   `json:"m_exchange_rate"`
	PaidAt         int64   `json:"m_paid_at"`
	ClientIp       int64   `json:"m_client_ip"`
	CreatedAt      int64   `json:"m_created_at"`
	AddnInf        string  `json:"addn_inf"`
	RechargeBakId  string  `db:"recharge_bak_id" json:"m_recharge_bak_id"` //交易备用字段
}

type CreateRechargeResponse struct {
	Channel    string            `json:"m_channel"`
	Credential string            `json:"m_credential"` // 需要返回给前端的内容
	Action     string            `json:"m_action"`     // 返回类型标识 告知前端需要执行什么操作 例如 QrCode二维码形式 WxCreatePay调起微信支付
	RechargeId int64             `json:"m_recharge_id"`
	SnapShot   map[string]string `json:"-"`
}

type PmtRechargeExt struct {
	RechargeId int64  `json:"recharge_id"`
	Credential string `json:"credential"`
	Meta       string `json:"meta"`
	Extra      string `json:"extra"`
	Additional string `json:"additional"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
	Status     int    `json:"status"`
}

type QueryRechargeResponse struct {
	RechargeId    string            `json:"m_recharge_id"`
	Amount        float64           `json:"m_amount"`
	PaidAt        int64             `json:"m_paid_at"`
	TransactionNo string            `json:"-"`
	OrderNo       string            `json:"m_order_no"`
	Status        int               `json:"m_status"`
	IsPaid        int               `json:"is_paid"`
	SnapShot      map[string]string `json:"-"`
	RespContent   string            `json:"-"`
	RespHeader    string            `json:"-"`
	AddnInf       string            `json:"addn_inf"`
}
