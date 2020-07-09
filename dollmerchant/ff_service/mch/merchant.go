package mch

import (
	"context"
	"errors"
	"github.com/Unknwon/com"
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_common/ff_page"
	"dollmachine/dollmerchant/ff_common/ff_password"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"dollmachine/dollmerchant/ff_model/mch_login"
	"dollmachine/dollmerchant/ff_model/mch_merchant"
	UniqueId "dollmachine/dollmerchant/proto/unique_id"
	"time"
)

type MchService struct {
}

func NewMchService() *MchService {
	return &MchService{}
}

func (p *MchService) GetMchInfo(merchantId string) (map[string]interface{}, error) {
	mchMerchantDao := mch_merchant.NewMchMerchant()
	one, err := mchMerchantDao.GetMchInfo(merchantId, "*")
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (p *MchService) GetMchList(offset, pageSize, totalSize string) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()
	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	mchMerchantDao := mch_merchant.NewMchMerchant()
	mchList, err := mchMerchantDao.GetMchList(Offset, PageSize, "*")
	if err != nil || len(mchList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}
	count, _ := mchMerchantDao.GetMchListTotalCount()
	page.SetTotalSize(count)
	return mchList, page, nil
}

func (p *MchService) AddMch(name, province, city, district, address, rechargeRate, integralConvertRate, integralRate, contactName, contactPhone string) bool {

	nowTime := time.Now().Unix()

	//开启事务
	dbr := ff_vars.DbConn.GetInstance()
	dbr.Begin()

	//mch_merchant
	//merchantId, _ := unique_id.NewUniqueId().GetMerchantId()

	//生成merchantId
	cli := UniqueId.NewGenerateUniqueIdService("go.micro.srv.unique_id", ff_vars.RpcSrv.Client())
	rsp, err := cli.GenerateUniqueId(context.TODO(), &UniqueId.UniqueIdRequest{Key: "merchantId"})
	if err != nil {
		logrus.Errorf("Generate merchantId error: %v", err)
		return false
	}
	merchantId := rsp.Value

	_, err = dbr.Table("mch_merchant").Data(map[string]interface{}{
		"merchant_id": merchantId,
		"nick_name":   name,
		"name":        name,

		"province":      province,
		"city":          city,
		"address":       address,
		"district":      district,
		"contact_name":  contactName,
		"contact_phone": contactPhone,

		"recharge_rate":         rechargeRate,
		"integral_convert_rate": integralConvertRate,
		"integral_rate":         integralRate,

		"status":     1,
		"updated_at": nowTime,
		"created_at": nowTime,
	}).Insert()
	logrus.Debugf("Insert merchant information. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Insert merchant information failure. Error : %v", err)
		return false
	}

	//mch_login
	//staffId, _ := unique_id.NewUniqueId().GetMerchantId()

	//生成staffId
	rsp, err = cli.GenerateUniqueId(context.TODO(), &UniqueId.UniqueIdRequest{Key: "staffId"})
	if err != nil {
		logrus.Errorf("Generate staffId error: %v", err)
		return false
	}
	staffId := rsp.Value

	password := "88888888"
	loginSalt := ff_password.GenerateSalt()
	loginSecret := ff_password.GeneratePassword(password, loginSalt)
	_, err = dbr.Table("mch_login").Data(map[string]interface{}{
		"staff_id":     staffId,
		"merchant_id":  merchantId,
		"login_token":  contactPhone,
		"login_salt":   loginSalt,
		"login_secret": loginSecret,

		"status":     1,
		"updated_at": nowTime,
		"created_at": nowTime,
	}).Insert()
	logrus.Debugf("Insert merchant login information. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Debugf("Insert merchant login information. Error : %v", err)
		return false
	}

	//mch_staff
	_, err = dbr.Table("mch_staff").Data(map[string]interface{}{
		"staff_id":    staffId,
		"merchant_id": merchantId,
		"role_id":     1,

		"phone":     contactPhone,
		"name":      contactPhone,
		"nick_name": contactPhone,

		"status":     1,
		"updated_at": nowTime,
		"created_at": nowTime,
	}).Insert()
	logrus.Debugf("Insert merchant staff information. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Insert merchant staff information. Error : %v", err)
		return false
	}

	//提交事务
	dbr.Commit()

	return true
}

func (p *MchService) UpdMch(merchantId, name, province, city, district, address, rechargeRate, integralConvertRate, integralRate, contactName, contactPhone string) bool {

	nowTime := time.Now().Unix()

	//开启事务
	dbr := ff_vars.DbConn.GetInstance()
	dbr.Begin()

	//mch_merchant
	_, err := dbr.Table("mch_merchant").Data(map[string]interface{}{
		"nick_name": name,
		"name":      name,

		"province":      province,
		"city":          city,
		"address":       address,
		"district":      district,
		"contact_name":  contactName,
		"contact_phone": contactPhone,

		"recharge_rate":         rechargeRate,
		"integral_convert_rate": integralConvertRate,
		"integral_rate":         integralRate,

		"updated_at": nowTime,
	}).Where("merchant_id", merchantId).Update()
	logrus.Debugf("Update merchant information. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Update merchant information failure. Error : %v", err)
		return false
	}

	//query staff id.
	staffOne, err := dbr.Table("mch_staff").Fields("staff_id").Where("merchant_id", merchantId).Order("created_at desc").First()
	logrus.Debugf("Get staff id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Get staff id failure. Error : %v", err)
		return false
	}

	//mch_login
	_, err = dbr.Table("mch_login").Data(map[string]interface{}{
		"login_token": contactPhone,
		"updated_at":  nowTime,
	}).Where("staff_id", staffOne["staff_id"]).Update()
	logrus.Debugf("Update merchant login information. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Debugf("Update merchant login information. Error : %v", err)
		return false
	}

	//mch_staff
	_, err = dbr.Table("mch_staff").Data(map[string]interface{}{
		"phone":      contactPhone,
		"name":       contactPhone,
		"nick_name":  contactPhone,
		"updated_at": nowTime,
	}).Where("staff_id", staffOne["staff_id"]).Update()
	logrus.Debugf("Update merchant staff information. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Update merchant staff information. Error : %v", err)
		return false
	}

	//提交事务
	dbr.Commit()

	return true
}

func (p *MchService) UpdMchPwd(merchantId, password string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	nowTime := time.Now().Unix()

	//query staff id.
	staffOne, err := dbr.Table("mch_staff").Fields("staff_id").Where("merchant_id", merchantId).Order("created_at desc").First()
	logrus.Debugf("Get staff id. LastSql : %v", dbr.LastSql)
	if err != nil {
		logrus.Errorf("Get staff id failure. Error : %v", err)
		return false
	}

	loginSalt := ff_password.GenerateSalt()
	loginSecret := ff_password.GeneratePassword(password, loginSalt)
	_, err = dbr.Table("mch_login").Data(map[string]interface{}{
		"login_salt":   loginSalt,
		"login_secret": loginSecret,
		"updated_at":   nowTime,
	}).Where("staff_id", staffOne["staff_id"]).Update()
	logrus.Debugf("Insert merchant login information. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Debugf("Insert merchant login information. Error : %v", err)
		return false
	}

	return true
}

func (p *MchService) UpdMchState(merchantId, status string) bool {
	dbr := ff_vars.DbConn.GetInstance()
	nowTime := time.Now().Unix()

	//开启事务
	dbr.Begin()

	//update mch_merchant
	_, err := dbr.Table("mch_merchant").Data(map[string]interface{}{
		"status":     status,
		"updated_at": nowTime,
	}).Where("merchant_id", merchantId).Update()
	logrus.Debugf("Update merchant status. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Update merchant status failure. Error : %v", err)
		return false
	}

	//update mch login
	_, err = dbr.Table("mch_login").Data(map[string]interface{}{
		"status":     status,
		"updated_at": nowTime,
	}).Where("merchant_id", merchantId).Update()
	logrus.Debugf("Update merchant login account status. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Update merchant login account status failure. Error : %v", err)
		return false
	}

	//提交事务
	dbr.Commit()

	return true
}

func (p *MchService) CheckIsExitsMchLogin(loginToken string) bool {
	mchLoginDao := mch_login.NewMchLogin()
	return mchLoginDao.CheckIsExitsMchLogin(loginToken)
}

func (p *MchService) CheckIsExitsMchName(name string) bool {
	mchMerchantDao := mch_merchant.NewMchMerchant()
	return mchMerchantDao.CheckIsExitsMchName(name)
}

func (p *MchService) CheckIsExitsMchId(merchantId string) bool {
	mchMerchantDao := mch_merchant.NewMchMerchant()
	return mchMerchantDao.CheckIsExitsMchId(merchantId)
}
