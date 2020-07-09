package common_msg

type CommonRequestMsg struct {
	Version              string `json:"version"`                //M	A	8	1.0
	InsCd                string `json:"ins_cd"`                 //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd              string `json:"mchnt_cd"`               //M	A	15	商户号, 富友分配给二级商户的商户号
	TermId               string `json:"term_id"`                //M	A	8	终端号(没有真实终端号统一填88888888)
	RandomStr            string `json:"random_str"`             //M	A	32	随机字符串
	Sign                 string `json:"sign"`                   //M	A	32	签名, 详见签名生成算法
	GoodsDes             string `json:"goods_des"`              //M	A	128	商品描述, 商品或支付单简要描述
	GoodsDetail          string `json:"goods_detail"`           //O	A	600	商品详情, 商品名称明细
	AddnInf              string `json:"addn_inf"`               //O	A	100	附加数据
	TermIp               string `json:"term_ip"`                //M	A	16	终端IP
	CurrType             string `json:"curr_type"`              //O	A	3	货币类型,默认人民币：CNY
	OrderAmt             string `json:"order_amt"`              //M	N	16	总金额, 订单总金额，单位为分
	MchntOrderNo         string `json:"mchnt_order_no"`         //M	A	30	商户订单号, 商户系统内部的订单号（5到30个字符、 只能包含字母数字,区分大小写)
	GoodsTag             string `json:"goods_tag"`              //O	A	32	商品标记
	TxnBeginTs           string `json:"txn_begin_ts"`           //M	A	14	交易起始时间, 订单生成时间，格式为yyyyMMddHHmmss
	ReservedExpireMinute string `json:"reserved_expire_minute"` //M	N	8	交易关闭时间,如不设置，默认填0单位:分钟 最大值为1440  例如:1天=1440
	ReservedFyTermSn     string `json:"reserved_fy_term_sn"`    //O	A	20	终端序列号
	ReservedFyTermId     string `json:"reserved_fy_term_id"`    //O	A	20	富友终端号(如果不是用的富友的POS终端，此字段千万不要填，不然会影响清算）
	ReservedFyTermType   string `json:"reserved_fy_term_type"`  //O	A	1	0:其他 1:富友终端 2:POS机 3:台卡 4:PC软件
}

type CommonResponseMsg struct {
	ResultCode             string `json:"result_code"`               //M	A	16	错误代码, 000000成功,其他详细参见错误列表
	ResultMsg              string `json:"result_msg"`                //M	A	128	错误代码描述
	InsCd                  string `json:"ins_cd"`                    //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd                string `json:"mchnt_cd"`                  //M	A	15	商户号, 富友分配的商户号
	TermId                 string `json:"term_id"`                   //O	A	8	终端号
	RandomStr              string `json:"random_str"`                //M	A	32	随机字符串
	Sign                   string `json:"sign"`                      //M	A	32	签名, 详见签名生成算法
	ReservedFyOrderNo      string `json:"reserved_fy_order_no"`      //O	A	30	富友生成的订单号,需要商户与商户订单号进行关联
	ReservedFyTraceNo      string `json:"reserved_fy_trace_no"`      //M	A	12	富友系统内部追踪号
	ReservedChannelOrderId string `json:"reserved_channel_order_id"` //O	A	64	条码流水号，用户账单二维码对应的流水
}

/*-------------------------------------------------订单查询start--------------------------------------------------------*/
//订单查询 请求报文
type CommonQueryRequestMsg struct {
	Version      string `json:"version"`        //M	A	8	1.0
	InsCd        string `json:"ins_cd"`         //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd      string `json:"mchnt_cd"`       //M	A	15	富友分配的商户号
	TermId       string `json:"term_id"`        //M	A	8	终端号(没有真实终端号统一填88888888)
	OrderType    string `json:"order_type"`     //M	A	20	订单类型:ALIPAY (统一下单、条码支付、服务窗支付), WECHAT(统一下单、条码支付，公众号支付),QQ(QQ钱包),UNIONPAY,BESTPAY(翼支付)
	MchntOrderNo string `json:"mchnt_order_no"` //M	A	30	商户订单号, 商户系统内部的订单号（5到30个字符、 只能包含字母数字,区分大小写)
	RandomStr    string `json:"random_str"`     //M	A	32	随机字符串，不长于32位
	Sign         string `json:"sign"`           //M	A	128	签名，详见签名生成算法
}

//订单查询 响应报文
type CommonQueryResponseMsg struct {
	ResultCode             string `json:"result_code"`               //M	A	16	错误代码, 000000成功,其他详细参见错误列表
	ResultMsg              string `json:"result_msg"`                //M	A	128	错误代码描述
	InsCd                  string `json:"ins_cd"`                    //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd                string `json:"mchnt_cd"`                  //M	A	15	商户号, 富友分配的商户号
	TermId                 string `json:"term_id"`                   //O	A	8	终端号
	RandomStr              string `json:"random_str"`                //M	A	32	随机字符串
	Sign                   string `json:"sign"`                      //M	A	32	签名, 详见签名生成算法
	BuyerId                string `json:"buyer_id"`                  //O	A	128	用户在商户的id
	OrderType              string `json:"order_type"`                //M	A	20	订单类型:ALIPAY , WECHAT
	TransStat              string `json:"trans_stat"`                //M	A	32	交易状态 SUCCESS—支付成功 REFUND—转入退款 NOTPAY—未支付 CLOSED—已关闭（交易以关闭） REVOKED—已撤销（刷卡支付） USERPAYING--用户支付中 PAYERROR--支付失败(其他原因，如银行返回失败)
	OrderAmt               string `json:"order_amt"`                 //M	N	16	订单金额, 单位为分
	TransactionId          string `json:"transaction_id"`            //M	A	64	渠道订单号
	MchntOrderNo           string `json:"mchnt_order_no"`            //M	A	30	商户订单号, 商户系统内部的订单号（5到30个字符、 只能包含字母数字,区分大小写)
	AddnInf                string `json:"addn_inf"`                  //O	A	128	附加数据
	ReservedFySettleDt     string `json:"reserved_fy_settle_dt"`     //M	A	8	富友清算日
	ReservedCouponFee      string `json:"reserved_coupon_fee"`       //O	A	10	优惠金额（分）
	ReservedBuyerLogonId   string `json:"reserved_buyer_logon_id"`   //O	A	128	买家在渠道登录账号
	ReservedFundBillList   string `json:"reserved_fund_bill_list"`   //O	A	不定长	支付宝交易资金渠道,详细渠道
	ReservedFyTraceNo      string `json:"reserved_fy_trace_no"`      //M	A	12	富友系统内部追踪号
	ReservedChannelOrderId string `json:"reserved_channel_order_id"` //O	A	64	条码流水号，用户账单二维码对应的流水
	ReservedFyTermId       string `json:"reserved_fy_term_id"`       //O	A	20	富友终端号(如果不是用的富友的POS终端，此字段千万不要填，不然会影响清算）
	ReservedIsCredit       string `json:"reserved_is_credit"`        //O	A	8	1表示信用卡或者花呗，0表示其他(非信用方式) 不填，表示未知
}

/*-------------------------------------------------订单查询end----------------------------------------------------------*/

/*-------------------------------------------------退款申请start--------------------------------------------------------*/
//退款申请 请求报文
type CommonRefundRequestMsg struct {
	Version          string `json:"version"`             //M	A	8	1.0
	InsCd            string `json:"ins_cd"`              //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd          string `json:"mchnt_cd"`            //M	A	15	富友分配的商户号
	TermId           string `json:"term_id"`             //M	A	8	终端号(没有真实终端号统一填88888888)
	MchntOrderNo     string `json:"mchnt_order_no"`      //M	A	30	商户订单号, 商户系统内部的订单号（5到30个字符、 只能包含字母数字,区分大小写)
	RandomStr        string `json:"random_str"`          //M	A	32	随机字符串，不长于32位
	Sign             string `json:"sign"`                //M	A	128	签名，详见签名生成算法
	OrderType        string `json:"order_type"`          //M	A	20	订单类型:ALIPAY (统一下单、条码支付、服务窗支付), WECHAT(统一下单、条码支付，公众号支付),QQ(QQ钱包),UNIONPAY,BESTPAY(翼支付)
	RefundOrderNo    string `json:"refund_order_no"`     //M	A	30	商户退款单号（5到30个字符、 只能包含字母数字或者下划线，区分大小写）
	TotalAmt         string `json:"total_amt"`           //M	N	16	总金额
	RefundAmt        string `json:"refund_amt"`          //M	N	16	退款金额
	OperatorId       string `json:"operator_id"`         //O	A	32	操作员
	ReservedFyTermId string `json:"reserved_fy_term_id"` //O	A	20	富友终端号(如果不是用的富友的POS终端，此字段千万不要填，不然会影响清算）
}

//退款申请 响应报文
type CommonRefundResponseMsg struct {
	ResultCode         string `json:"result_code"`           //M	A	16	错误代码, 000000成功,其他详细参见错误列表
	ResultMsg          string `json:"result_msg"`            //O	A	128	返回信息, 返回错误原因
	InsCd              string `json:"ins_cd"`                //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd            string `json:"mchnt_cd"`              //M	A	15	商户号, 富友分配的商户号
	TermId             string `json:"term_id"`               //O	A	8	终端号
	RandomStr          string `json:"random_str"`            //M	A	32	随机字符串
	Sign               string `json:"sign"`                  //M	A	32	签名, 详见签名生成算法
	OrderType          string `json:"order_type"`            //M	A	20	订单类型:ALIPAY , WECHAT, QQ(QQ钱包), UNIONPAY
	MchntOrderNo       string `json:"mchnt_order_no"`        //M	A	30	商户订单号, 商户系统内部的订单号（5到30个字符、 只能包含字母数字,区分大小写)
	RefundOrderNo      string `json:"refund_order_no"`       //M	A	30	商户退款单号
	TransactionId      string `json:"transaction_id"`        //O	A	64	渠道交易流水号
	RefundId           string `json:"refund_id"`             //O	A	64	渠道退款流水号
	ReservedFySettleDt string `json:"reserved_fy_settle_dt"` //M	A	8	富友清算日
	ReservedRefundAmt  string `json:"reserved_refund_amt"`   //O	N	16	退款金额
}

/*-------------------------------------------------退款申请end----------------------------------------------------------*/

/*-------------------------------------------------支付结果通知start-----------------------------------------------------*/
//支付结果异步通知 请求报文
type CommonNotifyRequestMsg struct {
	ResultCode             string `json:"result_code"`               //M	A	16	错误代码, 000000成功,其他详细参见错误列表
	ResultMsg              string `json:"result_msg"`                //O	A	128	返回信息, 返回错误原因
	InsCd                  string `json:"ins_cd"`                    //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd                string `json:"mchnt_cd"`                  //M	A	15	商户号, 富友分配的商户号
	TermId                 string `json:"term_id"`                   //O	A	8	终端号,富友分配的终端设备号
	RandomStr              string `json:"random_str"`                //M	A	32	随机字符串
	Sign                   string `json:"sign"`                      //M	A	32	签名, 详见签名生成算法
	UserId                 string `json:"user_id"`                   //O	A	128	用户在商户的id
	OrderAmt               string `json:"order_amt"`                 //M	N	16	订单金额, 单位为分
	SettleOrderAmt         string `json:"settle_order_amt"`          //M	N	16	应结订单金额
	CurrType               string `json:"curr_type"`                 //O	A	3	货币种类
	TransactionId          string `json:"transaction_id"`            //O	A	64	渠道交易流水号
	MchntOrderNo           string `json:"mchnt_order_no"`            //M	A	30	商户订单号, 商户系统内部的订单号（5到30个字符、 只能包含字母数字,区分大小写)）
	OrderType              string `json:"order_type"`                //O	A	20	订单类型:ALIPAY , WECHAT, QQ(QQ钱包), UNIONPAY，BESTPAY(翼支付)
	TxnFinTs               string `json:"txn_fin_ts"`                //M	A	14	支付完成时间, 订单支付时间，格式为yyyyMMddHHmmss
	ReservedFySettleDt     string `json:"reserved_fy_settle_dt"`     //M	A	8	富友清算日
	ReservedCouponFee      string `json:"reserved_coupon_fee"`       //O	A	10	优惠金额（分）
	ReservedBuyerLogonId   string `json:"reserved_buyer_logon_id"`   //O	A	128	买家在渠道登录账号
	ReservedFundBillList   string `json:"reserved_fund_bill_list"`   //O	A	不定长	支付宝交易资金渠道,详细渠道
	ReservedFyTraceNo      string `json:"reserved_fy_trace_no"`      //M	A	12	富友系统内部追踪号
	ReservedChannelOrderId string `json:"reserved_channel_order_id"` //O	A	64	条码流水号，用户账单二维码对应的流水
	ReservedIsCredit       string `json:"reserved_is_credit"`        //O	A	8	1表示信用卡或者花呗，0表示其他(非信用方式) 不填，表示未知
}

//支付结果异步通知 响应报文: 1
/*-------------------------------------------------支付结果通知end-----------------------------------------------------*/

/*--------------------------------------------查询可提现资金信息start---------------------------------------------------*/
//查询可提现资金信息 请求报文
type CommonQueryWithdrawAmtRequestMsg struct {
	InsCd     string `json:"ins_cd"`     //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd   string `json:"mchnt_cd"`   //M	A	15	商户号
	RandomStr string `json:"random_str"` //M	A	32	随机字符串，不长于32位
	Sign      string `json:"sign"`       //M	A	128	签名，详见签名生成算法
}

//查询可提现资金信息 响应报文
type CommonQueryWithdrawAmtResponseMsg struct {
	ReturnCode          string `json:"return_code"`           //M	A	16	返回状态码, SUCCESS/FAIL，此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg           string `json:"return_msg"`            //M	A	128	返回信息, 返回错误原因
	InsCd               string `json:"ins_cd"`                //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd             string `json:"mchnt_cd"`              //M	A	15	商户号
	RandomStr           string `json:"random_str"`            //M	A	32	随机字符串，不长于32位
	Sign                string `json:"sign"`                  //M	A	128	签名
	ResultCode          string `json:"result_code"`           //M	A	16	业务结果, SUCCESS/FAIL
	ErrCode             string `json:"err_code"`              //O	A	32	错误代码, 0000成功,其他详细参见错误列表
	ErrCodeDes          string `json:"err_code_des"`          //O	A	128	错误代码描述
	SettledAmt          string `json:"settled_amt"`           //M	N	16	已结算金额（就是银行已经结算到商户富友账户上的金额）
	NotSettleAmt        string `json:"not_settle_amt"`        //M	N	16	未结算金额（就是银行还未结算到商户富友账户上，如果需要体现，就需要富友垫付资金，所以就需要收取一定的手续费的）
	ReservedBookBalance string `json:"reserved_book_balance"` //M	N	16	账面余额
}

/*--------------------------------------------查询可提现资金信息end-----------------------------------------------------*/

/*--------------------------------------------查询手续费信息start-----------------------------------------------------*/
//查询手续费信息 请求报文
type CommonQueryFeeAmtRequestMsg struct {
	InsCd     string `json:"ins_cd"`     //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd   string `json:"mchnt_cd"`   //M	A	15	商户号
	RandomStr string `json:"random_str"` //	M	A	32	随机字符串，不长于32位
	Sign      string `json:"sign"`       //M	A	128	签名，详见签名生成算法
	Amt       string `json:"amt"`        //M	N	12	待提现金额
}

//查询手续费信息 响应报文
type CommonQueryFeeAmtResponseMsg struct {
	ReturnCode string `json:"return_code"`  //M	A	16	返回状态码, SUCCESS/FAIL，此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg  string `json:"return_msg"`   //M	A	128	返回信息, 返回错误原因
	InsCd      string `json:"ins_cd"`       //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd    string `json:"mchnt_cd"`     //M	A	15	商户号,可选
	RandomStr  string `json:"random_str"`   //M	A	32	随机字符串，不长于32位
	Sign       string `json:"sign"`         //M	A	128	签名
	ResultCode string `json:"result_code"`  //M	A	16	业务结果, SUCCESS/FAIL
	ErrCode    string `json:"err_code"`     //O	A	32	错误代码, 0000成功,其他详细参见错误列表
	ErrCodeDes string `json:"err_code_des"` //O	A	128	错误代码描述
	FeeAmt     string `json:"fee_amt"`      //M	N	16	手续费金额
	FeeDesc    string `json:"fee_desc"`     //O	A	100	手续费描述
}

/*--------------------------------------------查询手续费信息end---------------------------------------------------------*/

/*--------------------------------------------发起提现start-------------------------------------------------------------*/
//发起提现 请求报文
type CommonWithdrawRequestMsg struct {
	InsCd     string `json:"ins_cd"`     //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd   string `json:"mchnt_cd"`   //M	A	15	商户号
	RandomStr string `json:"random_str"` //M	A	32	随机字符串，不长于32位
	Sign      string `json:"sign"`       //M	A	128	签名，详见签名生成算法
	Amt       string `json:"amt"`        //M	N	12	待提现金额
	FeeAmt    string `json:"fee_amt"`    //M	N	16	手续费金额
	TxnType   string `json:"txn_type"`   //M	A	1	1未结提现2 已结提现
}

//发起提现 响应报文
type CommonWithdrawResponseMsg struct {
	ReturnCode string `json:"return_code"`  //M	A	16	返回状态码, SUCCESS/FAIL，此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg  string `json:"return_msg"`   //M	A	128	返回信息, 返回错误原因
	InsCd      string `json:"ins_cd"`       //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd    string `json:"mchnt_cd"`     //M	A	15	商户号,可选
	RandomStr  string `json:"random_str"`   //M	A	32	随机字符串，不长于32位
	Sign       string `json:"sign"`         //M	A	128	签名
	ResultCode string `json:"result_code"`  //M	A	16	业务结果, SUCCESS/FAIL
	ErrCode    string `json:"err_code"`     //O	A	32	错误代码, 0000成功,其他详细参见错误列表
	ErrCodeDes string `json:"err_code_des"` //O	A	128	错误代码描述
}

/*--------------------------------------------发起提现end---------------------------------------------------------------*/

/*--------------------------------------------资金划拨查询start----------------------------------------------------------*/
//资金划拨查询 请求报文
type CommonQueryChnlPayAmtRequestMsg struct {
	InsCd      string `json:"ins_cd"`      //M	A	20	机构号,接入机构在富友的唯一代码
	MchntCd    string `json:"mchnt_cd"`    //M	A	15	商户号
	RandomStr  string `json:"random_str"`  //M	A	32	随机字符串，不长于32位
	Sign       string `json:"sign"`        //M	A	128	签名，详见签名生成算法
	StartDate  string `json:"start_date"`  //M	A	8	查询开始日期（yyyyMMdd）
	EndDate    string `json:"end_date"`    //M	A	8	查询结束日期（yyyyMMdd）
	StartIndex string `json:"start_index"` //M	A	4	分页开始序号
	EndIndex   string `json:"end_index"`   //M	A	4	分页结束序号
}

//资金划拨查询 响应报文
type CommonQueryChnlPayAmtResponseMsg struct {
	MchntList string `json:"mchnt_list"`
	Mchnt     string `json:"mchnt"`       //M	A
	MchntCd   string `json:"mchnt_cd"`    //M	A	15	商户号
	PayAmt    string `json:"pay_amt"`     //M	A	20	金额（分）
	PaySt     string `json:"pay_st"`      //M	A	2	划拨状态1,成功 2,失败 3,划款中 4,已划款,状态未知 5,转入小宝金库
	PayStDesc string `json:"pay_st_desc"` //M	A	64	划拨状态描述
	SettleDt  string `json:"settle_dt"`   //M	A	8	清算日期
	Count     string `json:"count"`       //M	A	10	笔数
	InsCd     string `json:"ins_cd"`      //M	A	20	机构号
	RandomStr string `json:"random_str"`  //M	A	32	随机字符串，不长于32位
	Sign      string `json:"sign"`        //M	A	128	签名，详见签名生成算法
}

/*--------------------------------------------资金划拨查询end-----------------------------------------------------------*/
