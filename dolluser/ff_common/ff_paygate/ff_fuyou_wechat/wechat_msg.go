package ff_fuyou_wechat

//渠道商户号查询字段定义 请求
type WxChnlMchntGetReq struct {
	TraceNo int64  `json:"trace_no"` //唯一流水号
	MchntCd int64  `json:"mchnt_cd"` //商户号
	InsCd   string `json:"ins_cd"`   //机构号
	Sign    string `json:"sign"`     //签名
}

//渠道商户号查询字段定义 响应
type WxChnlMchntGetRes struct {
	TraceNo       int64  `json:"trace_no"`        //唯一流水号
	RetCode       int64  `json:"ret_code"`        //响应结果
	RetMsg        string `json:"ret_msg"`         //响应对应中文解释
	LinkMchntCd   string `json:"link_mchnt_cd"`   //挂靠一级商户号
	ChnlMchntInfo string `json:"chnl_mchnt_info"` //渠道商户信息
}

//微信参数配置字段定义
type WechatConfigSetReq struct {
	TraceNo        int64  `json:"trace_no"`        //唯一流水号
	InsCd          string `json:"ins_cd"`          //机构号
	AgencyType     int    `json:"agency_type"`     //代理商类型(0:微众一般类，1:绿洲)
	Sign           string `json:"sign"`            //签名
	MchntCd        string `json:"mchnt_cd"`        //富友商户号
	JsapiPath      string `json:"jsapi_path"`      //子商户公众账号JS API支付授权目录
	SubAppid       string `json:"sub_appid"`       //子商户SubAPPID
	SubscribeAppid string `json:"subscribe_appid"` //子商户推荐关注公众账号APPID
}
