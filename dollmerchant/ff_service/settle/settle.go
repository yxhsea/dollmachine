package settle

import (
	"context"
	"errors"
	"fmt"
	"github.com/Unknwon/com"
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_common/ff_page"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"dollmachine/dollmerchant/ff_model/draw_record"
	"dollmachine/dollmerchant/ff_model/exchange_integral"
	"dollmachine/dollmerchant/ff_model/pmt_recharge"
	"dollmachine/dollmerchant/ff_model/usr_user"
	UniqueId "dollmachine/dollmerchant/proto/unique_id"
	"time"
)

type SettleService struct {
}

func NewSettleService() *SettleService {
	return &SettleService{}
}

func (p *SettleService) GetUserRechargeList(offset, pageSize, totalSize string, merchantId, StartTime, EndTime int64) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()

	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	pmtRechargeDao := pmt_recharge.NewPmtRecharge()
	userRechargeList, err := pmtRechargeDao.GetPmtRechargeListByUserRecharge(Offset, PageSize, "*", merchantId, StartTime, EndTime)
	if err != nil || len(userRechargeList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}

	//补齐用户名
	usrDao := usr_user.NewUsrUser()
	for _, v := range userRechargeList {
		usrInfo := usrDao.GetUsrUserOne(fmt.Sprint(v["user_id"]), "nick_name")
		v["user_name"] = usrInfo["nick_name"]
	}

	count, _ := pmtRechargeDao.GetPmtRechargeListTotalCountByUserRecharge(merchantId, StartTime, EndTime)
	page.SetTotalSize(count)
	return userRechargeList, page, nil
}

func (p *SettleService) GetUserIntegralList(offset, pageSize, totalSize string, merchantId, StartTime, EndTime int64) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()

	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	exchangeIntegralDao := exchange_integral.NewExchangeIntegral()
	userIntegralList, err := exchangeIntegralDao.GetExchangeIntegralList(Offset, PageSize, "*", merchantId, StartTime, EndTime)
	if err != nil || len(userIntegralList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}

	//补齐用户名
	usrDao := usr_user.NewUsrUser()
	for _, v := range userIntegralList {
		usrInfo := usrDao.GetUsrUserOne(fmt.Sprint(v["user_id"]), "nick_name")
		v["user_name"] = usrInfo["nick_name"]
	}

	count, _ := exchangeIntegralDao.GetExchangeIntegralListTotalCount(merchantId, StartTime, EndTime)
	page.SetTotalSize(count)
	return userIntegralList, page, nil
}

func (p *SettleService) GetCurrentMonthDetail(merchantId, StartTime, EndTime int64) (map[string]interface{}, error) {
	pmtRechargeDao := pmt_recharge.NewPmtRecharge()
	monthDetail, err := pmtRechargeDao.GetCurrentMonthDetail(merchantId, StartTime, EndTime)
	if err != nil {
		return nil, errors.New("暂无数据")
	}
	return monthDetail, nil
}

func (p *SettleService) GetMonthDetailList(offset, pageSize, totalSize string, merchantId int64) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()

	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	drawRecordDao := draw_record.NewDrawRecord()
	drawRecordList, err := drawRecordDao.GetDrawRecordList(Offset, PageSize, "*", merchantId)
	if err != nil || len(drawRecordList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}

	count, _ := drawRecordDao.GetDrawRecordListTotalCount(merchantId)
	page.SetTotalSize(count)
	return drawRecordList, page, nil
}

func (p *SettleService) ApplyForWithDraw(merchantId int64, merchantName string, staffName string, staffPhone string, account, accountName, bank,
	subBank, benefitData, invoice, amount, expressCompany, expressNo string) bool {

	//生成DrawId
	cli := UniqueId.NewGenerateUniqueIdService("go.micro.srv.unique_id", ff_vars.RpcSrv.Client())
	rsp, err := cli.GenerateUniqueId(context.TODO(), &UniqueId.UniqueIdRequest{Key: "drawId"})
	if err != nil {
		logrus.Errorf("Generate DrawId error: %v", err)
		return false
	}
	DrawId := rsp.Value

	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	_, err = dbr.Table("draw_record").Data(map[string]interface{}{
		"draw_id":         DrawId,
		"merchant_id":     merchantId,
		"merchant_name":   merchantName,
		"account_name":    accountName,
		"account":         account,
		"bank":            bank,
		"sub_bank":        subBank,
		"benefit_data":    benefitData,
		"invoice":         invoice,
		"amount":          amount,
		"express_no":      expressNo,
		"express_company": expressCompany,
		"apply_name":      staffName,
		"apply_phone":     staffPhone,
		"apply_status":    1,

		"status":     1,
		"updated_at": nowTime,
		"created_at": nowTime,
	}).Insert()
	logrus.Debugf("ApplyForWithDraw. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("ApplyForWithDraw. Error : %v", err)
		return false
	}
	return true
}

func (p *SettleService) GetApplyDrawRecordList(offset, pageSize, totalSize string, merchantId int64) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()

	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	drawRecordDao := draw_record.NewDrawRecord()
	drawRecordList, err := drawRecordDao.GetDrawRecordList(Offset, PageSize, "*", merchantId)
	if err != nil || len(drawRecordList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}

	count, _ := drawRecordDao.GetDrawRecordListTotalCount(merchantId)
	page.SetTotalSize(count)
	return drawRecordList, page, nil
}
