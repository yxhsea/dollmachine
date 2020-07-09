package trade

import (
	"context"
	"errors"
	"fmt"
	"github.com/Unknwon/com"
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_cache/unique_id"
	"dollmachine/dolluser/ff_cache/user_session"
	"dollmachine/dolluser/ff_common/ff_convert"
	"dollmachine/dolluser/ff_common/ff_paygate/ff_fuyou/ff_basepay"
	"dollmachine/dolluser/ff_common/ff_paygate/ff_payconf"
	"dollmachine/dolluser/ff_common/ff_paygate/ff_payment"
	"dollmachine/dolluser/ff_config/ff_vars"
	UniqueId "dollmachine/dolluser/proto/unique_id"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type TradeService struct {
}

func NewTradeService() *TradeService {
	return &TradeService{}
}

//充值支付
func (p *TradeService) AddRecharge(userKey *user_session.UserSession, amount float64, coin int, deviceId int64, clientIp string, merchantId int64, fromTag string, tradeType int) (*ff_basepay.CreateRechargeResponse, error) {
	var err error
	var payGate ff_basepay.IPayGate
	switch tradeType {
	case 1:
		payGate, err = ff_payment.GetPayment(ff_payconf.WECHATPAY_JSPAY_FUYOU, ff_vars.NotifyUrl, "")
		if err != nil {
			return nil, err
		}
	case 2:
		payGate, err = ff_payment.GetPayment(ff_payconf.WechatPay_Qr_FuYou, ff_vars.NotifyUrl, "")
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("交易类型不存在")
	}

	//设备信息
	dbr := ff_vars.DbConn.GetInstance()
	deviceInfo, err := dbr.Table("mch_device").Where("is_delete", "=", 0).Where("device_id", "=", deviceId).First()
	logrus.Debugf("LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Error : %v", err)
		return nil, err
	}

	//商户信息
	dbr = ff_vars.DbConn.GetInstance()
	mchInfo, err := dbr.Table("mch_merchant").Where("is_delete", "=", 0).Where("merchant_id", "=", deviceInfo["merchant_id"]).First()
	logrus.Debugf("LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Error : %v", err)
		return nil, err
	}

	//用户钱包Id
	userCoinWallet, err := dbr.Table("coin_wallet").
		Fields("coin_wallet_id").
		Where("is_delete", "=", 0).
		Where("merchant_id", "=", merchantId).
		Where("user_id", "=", userKey.UserId).
		First()

	//商户收入、平台收入、充值分成比例
	mchIncome, plfIncome, rechargeRate, err := p.getMchAndPlfIncome(amount, merchantId)
	if err != nil {
		return nil, err
	}

	//生成交易ID
	rechargeId, ccErr := unique_id.NewUniqueId().GetRechargeId()
	if ccErr != nil {
		return nil, ccErr
	}

	nowTime := time.Now().Unix()
	rechargeInfo := map[string]interface{}{
		"recharge_id": rechargeId, //交易Id

		"user_name":      userKey.NickName,                             //用户名
		"user_id":        userKey.UserId,                               //用户Id
		"user_gender":    3,                                            //用户性别
		"coin_wallet_id": fmt.Sprint(userCoinWallet["coin_wallet_id"]), //用户钱包Id
		"client_ip":      ff_convert.StringIpToInt(clientIp),           //创建交易者的ip

		"merchant_name": mchInfo["nick_name"],   //商户名称
		"merchant_id":   mchInfo["merchant_id"], //商户Id

		"is_paid":     1, //是否已经支付，1|未支付，2|已支付
		"is_refunded": 1, //是否退款，1|未退款，2|部分退款，3|全额退款
		"is_bonus":    1, //是否赠送的，1|不是，2|是

		"channel": payGate.GetPayConfMeta().Channel, //支付通道
		"subject": "游戏充值" + fmt.Sprint(rechargeId),  //支付标题

		"transaction_no":  "",     //支付渠道返回的交易流水号
		"currency":        "cny",  //三位 ISO 货币代码，人民币为 cny
		"amount":          amount, //交易金额
		"amount_settle":   0,      //清算金额，预留
		"amount_refunded": 0,      //已退款总金额
		"coin":            coin,   //用户充值金币

		"place_id":         deviceInfo["place_id"],         //场地Id
		"place_name":       deviceInfo["place_name"],       //场地名称
		"device_id":        deviceInfo["device_id"],        //设备Id
		"device_name":      deviceInfo["device_name"],      //设备名称
		"device_type_id":   deviceInfo["device_type_id"],   //设备类型Id
		"device_type_name": deviceInfo["device_type_name"], //设备类型名称

		"from_tag":             "",                //充值来源
		"server_fee":           amount * 6 / 1000, //服务费
		"mch_income":           mchIncome,         //商户收入
		"plf_income":           plfIncome,         //平台收入
		"exchange_amount_rate": rechargeRate,      //分成比例
		"exchange_rate":        0,                 //兑换比例，1元等于几币

		"status":     1,
		"created_at": nowTime,
		"updated_at": nowTime,
		"expired_at": nowTime + 60*60*2, //过期时间
		"paid_at":    0,                 //订单支付完成时间，用 Unix 时间戳表示
	}

	recharge := &ff_basepay.PmtRecharge{}
	recharge.Subject = fmt.Sprint(rechargeInfo["subject"])
	recharge.RechargeId = rechargeId
	recharge.Currency = "cny"
	recharge.Amount = amount
	recharge.ClientIp = ff_convert.StringIpToInt(clientIp)
	recharge.CreatedAt = nowTime
	recharge.Channel = payGate.GetPayConfMeta().Channel
	recharge.AddnInf = fmt.Sprint(rechargeId)

	/*usrLogin, err := dbr.Table("usr_login").Fields("fy_openid").Where("user_id", "=", userKey.UserId).First()
	if err != nil {
		logrus.Debug(dbr.LastSql())
		return nil, ff_res.NewCCErr("获取富友Openid失败!", ff_res.ErrorCode0ServiceError, ff_res.ErrorCode1TypePlatformApiErr, ff_res.ErrorCode2TypeInvalidParameter, "db", ff_res.ErrorCode3TypeSystemBusy, "获取富友Openid.Err", ",err:", err.Error(), ",merchantId:", "")
	}

	//富友Openid
	var fyOpenid string
	if fmt.Sprint(usrLogin["fy_openid"]) == "" {
		fyOpenid = "ooIeqsza0qteX1OidgsOAIjxS1Ws"
	} else {
		fyOpenid = fmt.Sprint(usrLogin["fy_openid"])
	}
	fmt.Println("____fyopenid : ", fyOpenid)*/

	meta := ff_basepay.RechargeExtMeta{}
	meta.OpenId = ""
	meta.SubOpenId = userKey.OpenId
	meta.SubAppId = "wxc7d98f96c6bb4a79"

	//调用充值方法
	createRechargeResponse, err := payGate.RechargeCreate(recharge, &ff_basepay.PmtRechargeExt{Meta: dbr.JsonEncode(meta)})
	if err != nil {
		return nil, err
	}

	rechargeExtInfo := map[string]interface{}{
		"recharge_id": rechargeId,
		"meta":        dbr.JsonEncode(meta),
		"additional":  payGate.GetPayConfMeta().MchId,
		"credential":  createRechargeResponse.Credential,
		"extra":       dbr.JsonEncode(createRechargeResponse.SnapShot),
		"created_at":  nowTime,
		"updated_at":  nowTime,
		"status":      1,
	}

	//开启事务
	dbr = ff_vars.DbConn.GetInstance()
	dbr.Begin()

	_, err = dbr.Table("pmt_recharge").Data(rechargeInfo).Insert()
	logrus.Debugf("Inset pmt_recharge lastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Insert pmt_recharge. Error : %v", err)
		dbr.Rollback()
		return nil, err
	}

	_, err = dbr.Table("pmt_recharge_ext").Data(rechargeExtInfo).Insert()
	logrus.Debugf("Insert pmt_recharge_ext lastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Insert pmt_recharge_ext. Error : %v", err)
		dbr.Rollback()
		return nil, err
	}

	//提交事务
	dbr.Commit()

	//交易日志
	addFyTradeRecordLog(map[string]interface{}{
		"fy_trade_record_id": rechargeId,
		"channel":            payGate.GetPayConfMeta().Channel,
		"subject":            "游戏充值" + fmt.Sprint(rechargeId),
		"credential":         createRechargeResponse.Credential,
		"extra":              dbr.JsonEncode(createRechargeResponse.SnapShot),
		"additional":         payGate.GetPayConfMeta().MchId,
		"meta":               dbr.JsonEncode(meta),
		"client_ip":          ff_convert.StringIpToInt(clientIp),
		"is_paid":            1,
		"created_at":         nowTime,
		"updated_at":         nowTime,
	})

	return createRechargeResponse, nil
}

//邮费支付
func (p *TradeService) AddExpressRecharge(userKey *user_session.UserSession, amount float64, clientIp string, tradeType int, exchangeRecordId int64) (*ff_basepay.CreateRechargeResponse, error) {
	var err error
	var payGate ff_basepay.IPayGate
	switch tradeType {
	case 1:
		payGate, err = ff_payment.GetPayment(ff_payconf.WECHATPAY_JSPAY_FUYOU, ff_vars.NotifyEspUrl, "")
		if err != nil {
			return nil, err
		}
	case 2:
		payGate, err = ff_payment.GetPayment(ff_payconf.WechatPay_Qr_FuYou, ff_vars.NotifyEspUrl, "")
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("交易类型不存在")
	}

	//生成交易ID
	rechargeId, ccErr := unique_id.NewUniqueId().GetRechargeId()
	if ccErr != nil {
		return nil, ccErr
	}

	nowTime := time.Now().Unix()
	recharge := &ff_basepay.PmtRecharge{}
	recharge.Subject = "邮费支付" + fmt.Sprint(rechargeId)
	recharge.RechargeId = rechargeId
	//recharge.RechargeBakId = fmt.Sprint(1142) + ff_random.KrandNum(2) + "|" + fmt.Sprint(exchangeRecordId)
	recharge.AddnInf = fmt.Sprint(exchangeRecordId)
	recharge.Currency = "cny"
	recharge.Amount = amount
	recharge.ClientIp = ff_convert.StringIpToInt(clientIp)
	recharge.CreatedAt = nowTime
	recharge.Channel = payGate.GetPayConfMeta().Channel

	/*usrLogin, err := dbr.Table("usr_login").Fields("fy_openid").Where("user_id", "=", userKey.UserId).First()
	if err != nil {
		logrus.Debug(dbr.LastSql())
		return nil, ff_res.NewCCErr("获取富友Openid失败!", ff_res.ErrorCode0ServiceError, ff_res.ErrorCode1TypePlatformApiErr, ff_res.ErrorCode2TypeInvalidParameter, "db", ff_res.ErrorCode3TypeSystemBusy, "获取富友Openid.Err", ",err:", err.Error(), ",merchantId:", "")
	}

	//富友Openid
	var fyOpenid string
	if fmt.Sprint(usrLogin["fy_openid"]) == "" {
		fyOpenid = "ooIeqsza0qteX1OidgsOAIjxS1Ws"
	} else {
		fyOpenid = fmt.Sprint(usrLogin["fy_openid"])
	}
	fmt.Println("____fyopenid : ", fyOpenid)*/

	meta := ff_basepay.RechargeExtMeta{}
	meta.OpenId = ""
	meta.SubOpenId = userKey.OpenId
	meta.SubAppId = "wxc7d98f96c6bb4a79"

	dbr := ff_vars.DbConn.GetInstance()
	//调用充值方法
	createRechargeResponse, err := payGate.RechargeCreate(recharge, &ff_basepay.PmtRechargeExt{Meta: dbr.JsonEncode(meta)})
	if err != nil {
		return nil, err
	}

	rechargeExtInfo := map[string]interface{}{
		"recharge_id": rechargeId,
		"record_id":   exchangeRecordId,
		"meta":        dbr.JsonEncode(meta),
		"additional":  payGate.GetPayConfMeta().MchId,
		"credential":  createRechargeResponse.Credential,
		"extra":       dbr.JsonEncode(createRechargeResponse.SnapShot),
		"created_at":  nowTime,
		"updated_at":  nowTime,
		"status":      1,
	}

	dbr = ff_vars.DbConn.GetInstance()
	Cnt, err := dbr.Table("express_pay_ext").Where("recharge_id", "=", rechargeId).Count(1)
	logrus.Debugf("Query express_pay_ext through recharge id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query express_pay_ext through recharge id. Error : %v", err)
		return nil, errors.New("查询失败")
	}

	//开启事务
	dbr = ff_vars.DbConn.GetInstance()
	dbr.Begin()
	if Cnt <= 0 {
		_, err = dbr.Table("express_pay_ext").Data(rechargeExtInfo).Insert()
		logrus.Debugf("Insert express record. LastSql : %v", dbr.LastSql)
		if err != nil {
			dbr.Rollback()
			logrus.Errorf("Insert express record. Error : %v", err)
			return nil, errors.New("添加邮费交易扩展记录失败")
		}
	}
	_, err = dbr.Table("award_record").Data(map[string]interface{}{"express_fee": amount, "updated_at": nowTime}).Where("record_id", "=", exchangeRecordId).Update()
	logrus.Debugf("Update award record through record id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Update award record failure through record id. Error : %v", err)
		dbr.Rollback()
		return nil, errors.New("更新记录失败")
	}
	//提交事务
	dbr.Commit()

	dbr = ff_vars.DbConn.GetInstance()
	Cnt, _ = dbr.Table("fy_trade_record_log").Where("fy_trade_record_id", "=", rechargeId).Count(1)
	if Cnt <= 0 {
		//交易日志
		addFyTradeRecordLog(map[string]interface{}{
			"fy_trade_record_id": rechargeId,
			"channel":            payGate.GetPayConfMeta().Channel,
			"subject":            "邮费支付" + fmt.Sprint(rechargeId),
			"credential":         createRechargeResponse.Credential,
			"extra":              dbr.JsonEncode(createRechargeResponse.SnapShot),
			"additional":         payGate.GetPayConfMeta().MchId,
			"meta":               dbr.JsonEncode(meta),
			"client_ip":          ff_convert.StringIpToInt(clientIp),
			"is_paid":            1,
			"created_at":         nowTime,
			"updated_at":         nowTime,
		})
	}
	return createRechargeResponse, nil
}

//充值支付异步通知
func (p *TradeService) NotifyRecharge(notifyStr string) (*ff_basepay.QueryRechargeResponse, error) {
	payGate, err := ff_payment.GetPayment("", "", "")
	if err != nil {
		return nil, err
	}

	notifyStr, err = url.PathUnescape(notifyStr)
	if err != nil {
		return nil, errors.New("notifyStr格式有误")
	}
	//再次解 去除+等信息
	notifyStr, err = url.QueryUnescape(notifyStr)
	if err != nil {
		return nil, errors.New("notifyStr格式有误")
	}

	notifyStr = strings.Replace(notifyStr, "req=", "", 1)

	queryRechargeResponse, err := payGate.ParseRechargeNotify(notifyStr)
	if err != nil {
		return nil, err
	}
	if queryRechargeResponse.RechargeId == "" {
		return nil, errors.New("notifyStr格式有误")
	}

	dbr := ff_vars.DbConn.GetInstance()
	recharge, err := dbr.Table("pmt_recharge").Where("recharge_id", queryRechargeResponse.RechargeId).First()
	logrus.Debugf("Query pmt_recharge through recharge id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query pmt_recharge failure through recharge id. Error : %v", err)
		return nil, err
	}
	isPaid, _ := com.StrTo(fmt.Sprint(recharge["is_paid"])).Int()

	//如果已经支付成功，不需要再操作
	if isPaid == ff_basepay.RechargeIsPaidYesPaid {
		return queryRechargeResponse, nil
	}

	dbr = ff_vars.DbConn.GetInstance()
	rechargeExt, err := dbr.Table("pmt_recharge_ext").Where("recharge_id", queryRechargeResponse.RechargeId).First()
	logrus.Debugf("Query pmt_recharge_ext through recharge id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query pmt_recharge_ext failure through recharge id. Error : %v", err)
		return nil, err
	}

	var payConf *ff_payconf.PayConf
	payConf = ff_payconf.NewPayConf()
	payConfMeta, err := payConf.GetPayConfMeta(fmt.Sprint(rechargeExt["additional"]))
	if err != nil {
		return nil, err
	}
	payGate.SetPayConfMeta(payConfMeta)
	//验证数据是否来自官方
	_, err = payGate.ValidRechargeNotify(notifyStr)
	if err != nil {
		return nil, err
	}
	if queryRechargeResponse.IsPaid == ff_basepay.RechargeIsPaidYesPaid {
		return p.notifyDealWith(queryRechargeResponse.TransactionNo, queryRechargeResponse.RechargeId)
	}

	return queryRechargeResponse, nil
}

//邮费支付异步通知
func (p *TradeService) NotifyExpress(notifyStr string) (*ff_basepay.QueryRechargeResponse, error) {
	payGate, err := ff_payment.GetPayment("", "", "")
	if err != nil {
		return nil, err
	}

	notifyStr, err = url.PathUnescape(notifyStr)
	if err != nil {
		return nil, errors.New("notifyStr格式有误")
	}
	//再次解 去除+等信息
	notifyStr, err = url.QueryUnescape(notifyStr)
	if err != nil {
		return nil, errors.New("notifyStr格式有误")
	}

	notifyStr = strings.Replace(notifyStr, "req=", "", 1)

	queryRechargeResponse, err := payGate.ParseRechargeNotify(notifyStr)
	if err != nil {
		return nil, err
	}
	if queryRechargeResponse.RechargeId == "" {
		return nil, errors.New("notifyStr格式有误")
	}

	dbr := ff_vars.DbConn.GetInstance()
	exchangeRecordId := queryRechargeResponse.AddnInf
	rechargeInfo, err := dbr.Table("award_record").Fields("is_pay").Where("record_id", "=", exchangeRecordId).First()
	logrus.Debugf("Query award record through record id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query award record failure through record id. Error : %v", err)
		return nil, err
	}

	isPaid, _ := strconv.Atoi(fmt.Sprint(rechargeInfo["is_pay"]))
	//如果已经支付成功，不需要再操作
	if isPaid == 1 {
		return queryRechargeResponse, nil
	}

	dbr = ff_vars.DbConn.GetInstance()
	expressExt, err := dbr.Table("express_pay_ext").Fields("additional").Where("recharge_id", "=", queryRechargeResponse.RechargeId).First()
	if err != nil {
		return nil, err
	}

	var payConf *ff_payconf.PayConf
	payConf = ff_payconf.NewPayConf()
	payConfMeta, err := payConf.GetPayConfMeta(fmt.Sprint(expressExt["additional"]))
	if err != nil {
		return nil, err
	}
	payGate.SetPayConfMeta(payConfMeta)
	//验证数据是否来自官方
	_, err = payGate.ValidRechargeNotify(notifyStr)
	if err != nil {
		return nil, err
	}

	if queryRechargeResponse.IsPaid == ff_basepay.RechargeIsPaidYesPaid {
		return p.notifyExpressDealWith(queryRechargeResponse.TransactionNo, queryRechargeResponse.RechargeId, exchangeRecordId)
	}

	return queryRechargeResponse, nil
}

//商户和平台的收入
func (p *TradeService) getMchAndPlfIncome(amount float64, merchantId int64) (float64, float64, float64, error) {
	dbr := ff_vars.DbConn.GetInstance()
	mchInfo, err := dbr.Table("mch_merchant").Fields("recharge_rate,recharge_rate,exchange_rate").Where("is_delete", "=", 0).Where("merchant_id", "=", merchantId).First()
	logrus.Debugf("Query merchant through merchant id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query merchant failure through merchant id. Error : %v", err)
		return 0, 0, 0, err
	}
	rechargeRate, _ := strconv.ParseFloat(fmt.Sprint(mchInfo["recharge_rate"]), 64)
	mchIncome := rechargeRate * amount
	plfIncome := (1 - rechargeRate) * amount
	return mchIncome, plfIncome, rechargeRate, nil
}

//充值异步回调业务处理
func (p *TradeService) notifyDealWith(transactionNo string, rechargeId string) (*ff_basepay.QueryRechargeResponse, error) {
	dbr := ff_vars.DbConn.GetInstance()
	//拉取充值订单记录
	rechargeInfo, err := dbr.Table("pmt_recharge").Where("recharge_id", "=", rechargeId).First()
	logrus.Debugf("Query pmt_recharge through recharge id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query pmt_recharge failure through recharge id. Error : %v", err)
		return nil, err
	}

	//拉取用户旧数据
	dbr = ff_vars.DbConn.GetInstance()
	coinWallet, err := dbr.Table("coin_wallet").Fields("coin").Where("coin_wallet_id", "=", rechargeInfo["coin_wallet_id"]).First()
	logrus.Debugf("Query user coin wallet through coin_wallet_id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query user coin wallet through coin_wallet_id. Error : %v", err)
		return nil, errors.New("NotifyRecharge.用户旧金币获取")
	}

	//开启事务
	dbr = ff_vars.DbConn.GetInstance()
	dbr.Begin()

	nowTime := time.Now().Unix()
	//更新订单状态
	dataUpd := map[string]interface{}{
		"is_paid":        ff_basepay.RechargeIsPaidYesPaid,
		"transaction_no": transactionNo,
		"paid_at":        nowTime,
		"updated_at":     nowTime,
	}
	_, err = dbr.Table("pmt_recharge").Data(dataUpd).Where("recharge_id", "=", rechargeId).Update()
	logrus.Debugf("Update pmt_recharge through recharge id", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Update pmt_recharge failure through recharge id", err)
		dbr.Rollback()
		return nil, errors.New("NotifyRecharge.异步回调支付状态更新")
	}

	//生成flowRecordId
	cli := UniqueId.NewGenerateUniqueIdService("go.micro.srv.unique_id", ff_vars.RpcSrv.Client())
	rsp, err := cli.GenerateUniqueId(context.TODO(), &UniqueId.UniqueIdRequest{Key: "flowRecordId"})
	if err != nil {
		logrus.Errorf("Generate merchantId error: %v", err)
		return nil, err
	}
	flowId := rsp.Value

	flowRecordInfo := map[string]interface{}{
		"flow_id":       flowId,
		"merchant_id":   rechargeInfo["merchant_id"],
		"merchant_name": rechargeInfo["merchant_name"],
		"user_id":       rechargeInfo["user_id"],
		"amount":        rechargeInfo["amount"],
		"coin":          rechargeInfo["coin"],
		"trade_type":    1,
		"status":        1,
		"created_at":    nowTime,
		"updated_at":    nowTime,
	}
	_, err = dbr.Table("flow_record").Data(flowRecordInfo).Insert()
	logrus.Debugf("Insert flow record. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Insert flow record failure. Error : %v", err)
		dbr.Rollback()
		return nil, errors.New("NotifyRecharge.添加流水记录")
	}

	//用户金币计算
	oldCoin, _ := strconv.ParseInt(fmt.Sprint(coinWallet["coin"]), 10, 64)
	addCoin, _ := strconv.ParseInt(fmt.Sprint(rechargeInfo["coin"]), 10, 64)
	totalCoin := oldCoin + addCoin

	//更新用户钱包
	_, err = dbr.Table("coin_wallet").Data(map[string]interface{}{"coin": totalCoin}).Where("coin_wallet_id", "=", rechargeInfo["coin_wallet_id"]).Update()
	logrus.Debugf("Update coin wallet through coin wallet id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Up")
		dbr.Rollback()
		return nil, errors.New("更新用户钱包失败")
	}

	//提交事务
	dbr.Commit()

	RechargeId, _ := strconv.ParseInt(rechargeId, 10, 64)
	//交易日志
	updateFyTreadRecordLog(map[string]interface{}{
		"is_paid": 2,
		"paid_at": time.Now().Unix(),
	}, RechargeId)

	return nil, nil
}

//邮费支付异步回调处理
func (p *TradeService) notifyExpressDealWith(transactionNo string, rechargeId string, exchangeRecordId string) (*ff_basepay.QueryRechargeResponse, error) {
	//兑奖记录
	dbr := ff_vars.DbConn.GetInstance()
	awardRecord, err := dbr.Table("award_record").Fields("merchant_id,user_id,express_fee").Where("record_id", "=", exchangeRecordId).First()
	logrus.Debugf("Query award record through record id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query award record through record id. Error : %v", err)
		return nil, errors.New("兑奖记录查询失败")
	}

	//商户名称
	dbr = ff_vars.DbConn.GetInstance()
	merchantInfo, err := dbr.Table("mch_merchant").Fields("nick_name").Where("merchant_id", "=", awardRecord["merchant_id"]).First()
	logrus.Debugf("Query merchant through record id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query merchant failure through record id. Error : %v", err)
		return nil, errors.New("商户信息查询失败")
	}

	//生成flowRecordId
	cli := UniqueId.NewGenerateUniqueIdService("go.micro.srv.unique_id", ff_vars.RpcSrv.Client())
	rsp, err := cli.GenerateUniqueId(context.TODO(), &UniqueId.UniqueIdRequest{Key: "flowRecordId"})
	if err != nil {
		logrus.Errorf("Generate merchantId error: %v", err)
		return nil, err
	}
	flowId := rsp.Value

	nowTime := time.Now().Unix()
	flowRecord := map[string]interface{}{
		"flow_id":       flowId,
		"merchant_id":   awardRecord["merchant_id"],
		"merchant_name": merchantInfo["nick_name"],
		"user_id":       awardRecord["user_id"],
		"trade_type":    3,
		"amount":        awardRecord["express_fee"],
		"status":        1,
		"created_at":    nowTime,
		"updated_at":    nowTime,
	}

	//开启事务
	dbr = ff_vars.DbConn.GetInstance()
	dbr.Begin()
	_, err = dbr.Table("flow_record").Data(flowRecord).Insert()
	logrus.Debugf("Insert flow record. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Insert flow record failure. Error : %v", err)
		dbr.Rollback()
		return nil, errors.New("添加流水记录失败")
	}

	awardRecordUpd := map[string]interface{}{
		"is_pay":     1,
		"pay_at":     nowTime,
		"updated_at": nowTime,
	}
	_, err = dbr.Table("award_record").Data(awardRecordUpd).Where("record_id", "=", exchangeRecordId).Update()
	logrus.Debugf("Update award record. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Update award record failure. Error : %v", err)
		dbr.Rollback()
		return nil, errors.New("更新记录失败")
	}

	//提交事务
	dbr.Commit()

	RechargeId, _ := strconv.ParseInt(rechargeId, 10, 64)
	//交易日志
	updateFyTreadRecordLog(map[string]interface{}{
		"is_paid": 2,
		"paid_at": time.Now().Unix(),
	}, RechargeId)

	return nil, nil
}

//充值支付状态查询
func (p *TradeService) QueryRecharge(userKey *user_session.UserSession, rechargeId string) (*ff_basepay.QueryRechargeResponse, error) {
	payGate, err := ff_payment.GetPayment("", "", "")
	if err != nil {
		return nil, err
	}

	recharge := &ff_basepay.PmtRecharge{}
	recharge.RechargeId, _ = strconv.ParseInt(rechargeId, 10, 64)
	queryRechargeResponse, err := payGate.RechargeQuery(recharge, &ff_basepay.PmtRechargeExt{})
	if err != nil {
		return nil, err
	}

	dbr := ff_vars.DbConn.GetInstance()
	rechargeInfo, err := dbr.Table("pmt_recharge").Fields("is_paid").Where("recharge_id", "=", rechargeId).First()
	logrus.Debugf("Query pmt_recharge through recharge id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query pmt_recharge through recharge id. Error : %v", err)
		return nil, err
	}

	isPaid, _ := strconv.Atoi(fmt.Sprint(rechargeInfo["is_paid"]))
	//如果已经支付成功，不需要再操作
	if isPaid == ff_basepay.RechargeIsPaidYesPaid {
		return queryRechargeResponse, nil
	}

	if queryRechargeResponse.IsPaid == ff_basepay.RechargeIsPaidYesPaid {
		return p.notifyDealWith(queryRechargeResponse.TransactionNo, queryRechargeResponse.RechargeId)
	}

	return queryRechargeResponse, nil
}

//邮费支付状态查询
func (p *TradeService) QueryRechargeExpress(userKey *user_session.UserSession, rechargeId string) (*ff_basepay.QueryRechargeResponse, error) {
	payGate, err := ff_payment.GetPayment("", "", "")
	if err != nil {
		return nil, err
	}

	recharge := &ff_basepay.PmtRecharge{}
	recharge.RechargeId, _ = strconv.ParseInt(rechargeId, 10, 64)
	queryRechargeResponse, err := payGate.RechargeQuery(recharge, &ff_basepay.PmtRechargeExt{})
	exchangeRecordId := queryRechargeResponse.AddnInf
	if err != nil {
		return nil, err
	}

	dbr := ff_vars.DbConn.GetInstance()
	rechargeInfo, err := dbr.Table("award_record").Fields("is_pay").Where("record_id", "=", exchangeRecordId).First()
	logrus.Debugf("Query award record through record id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Query award record through record id. Error : %v", err)
		return nil, err
	}

	isPaid, _ := strconv.Atoi(fmt.Sprint(rechargeInfo["is_pay"]))
	//如果已经支付成功，不需要再操作
	if isPaid == 1 {
		return queryRechargeResponse, nil
	}

	if queryRechargeResponse.IsPaid == ff_basepay.RechargeIsPaidYesPaid {
		return p.notifyExpressDealWith(queryRechargeResponse.TransactionNo, queryRechargeResponse.RechargeId, exchangeRecordId)
	}

	return queryRechargeResponse, nil
}

//添加富友交易记录日志
func addFyTradeRecordLog(mapData map[string]interface{}) {
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("fy_trade_record_log").Data(mapData).Insert()
	logrus.Debugf("[AddFyTradeRecordLog] get lastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("[AddFyTradeRecordLog] Error : %v ", err)
	}
}

//更新富友交易记录支付状态
func updateFyTreadRecordLog(mapData map[string]interface{}, fyTradeRecordId int64) {
	dbr := ff_vars.DbConn.GetInstance()
	_, err := dbr.Table("fy_trade_record_log").Data(mapData).Where("fy_trade_record_id", "=", fyTradeRecordId).Update()
	logrus.Debugf("[updateFyTreadRecordLog] get LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("[updateFyTreadRecordLog] Error : %v ", err)
	}
}
