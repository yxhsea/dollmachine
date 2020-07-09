package staff

import (
	"context"
	"errors"
	"fmt"
	"github.com/Unknwon/com"
	"github.com/gogap/logrus"
	"dollmachine/dollmerchant/ff_common/ff_page"
	"dollmachine/dollmerchant/ff_common/ff_password"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"dollmachine/dollmerchant/ff_model/mch_login"
	"dollmachine/dollmerchant/ff_model/mch_role"
	"dollmachine/dollmerchant/ff_model/mch_staff"
	UniqueId "dollmachine/dollmerchant/proto/unique_id"
	"time"
)

type StaffService struct {
}

func NewStaffService() *StaffService {
	return &StaffService{}
}

func (p *StaffService) GetStaffInfo(staffId string) (map[string]interface{}, error) {
	staffDao := mch_staff.NewMchStaff()
	one, err := staffDao.GetMchStaffInfo(staffId, "*")
	if err != nil {
		return nil, err
	}
	return one, nil
}

func (p *StaffService) GetStaffList(offset, pageSize, totalSize string, merchantId int64) ([]map[string]interface{}, *ff_page.Page, error) {
	Offset, _ := com.StrTo(offset).Int()
	PageSize, _ := com.StrTo(pageSize).Int()
	TotalSize, _ := com.StrTo(totalSize).Int()
	page := ff_page.NewPage(Offset, PageSize, TotalSize)
	staffDao := mch_staff.NewMchStaff()
	staffList, err := staffDao.GetMchStaffList(Offset, PageSize, "*", merchantId)
	if err != nil || len(staffList) == 0 {
		return nil, nil, errors.New("暂无数据")
	}
	count, _ := staffDao.GetMchStaffListTotalCount(merchantId)
	page.SetTotalSize(count)
	return staffList, page, nil
}

func (p *StaffService) AddStaff(loginToken, staffName, password, roleId string, merchantId int64) bool {
	roleDao := mch_role.NewMchRole()
	if !roleDao.CheckIsExitsRoleId(roleId) {
		return false
	}

	//生成staffId
	cli := UniqueId.NewGenerateUniqueIdService("go.micro.srv.unique_id", ff_vars.RpcSrv.Client())
	rsp, err := cli.GenerateUniqueId(context.TODO(), &UniqueId.UniqueIdRequest{Key: "staffId"})
	if err != nil {
		logrus.Errorf("Generate staffId error: %v", err)
		return false
	}
	staffId := rsp.Value

	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	//开启事务
	dbr.Begin()
	_, err = dbr.Table("mch_staff").Data(map[string]interface{}{
		"staff_id":    staffId,
		"merchant_id": merchantId,
		"role_id":     roleId,
		"name":        staffName,
		"nick_name":   staffName,

		"status":     1,
		"created_at": nowTime,
		"updated_at": nowTime,
	}).Insert()
	logrus.Debugf("Insert staff. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Insert staff. Error : %v", err)
		return false
	}

	loginSalt := ff_password.GenerateSalt()
	loginSecret := ff_password.GeneratePassword(password, loginSalt)
	_, err = dbr.Table("mch_login").Data(map[string]interface{}{
		"staff_id":     staffId,
		"merchant_id":  merchantId,
		"login_token":  loginToken,
		"login_secret": loginSecret,
		"login_salt":   loginSalt,

		"status":     1,
		"updated_at": nowTime,
		"created_at": nowTime,
	}).Insert()
	logrus.Debugf("Insert mch_login. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Insert mch_login. Error : %v", err)
	}

	//提交事务
	dbr.Commit()
	return true
}

func (p *StaffService) UpdStaff(staffId, name, loginToken, roleId string) bool {
	roleDao := mch_role.NewMchRole()
	if !roleDao.CheckIsExitsRoleId(roleId) {
		return false
	}

	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	//开启事务
	dbr.Begin()

	//update mch_staff
	_, err := dbr.Table("mch_staff").Data(map[string]interface{}{
		"name":       name,
		"nick_name":  name,
		"role_id":    roleId,
		"updated_at": nowTime,
	}).Where("staff_id", staffId).Update()
	logrus.Debugf("Update staff. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Update staff. Error : %v", err)
		return false
	}

	//update mch_login
	_, err = dbr.Table("mch_login").Data(map[string]interface{}{
		"login_token": loginToken,
		"updated_at":  nowTime,
	}).Where("staff_id", staffId).Update()
	logrus.Debugf("Update mch_login. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Update mch_login. Error : %v", err)
		return false
	}

	//提交事务
	dbr.Commit()
	return true
}

//修改密码
func (p *StaffService) PwdStaff(staffId, newPassword, oldPassword string) bool {
	loginSalt := ff_password.GenerateSalt()
	loginSecret := ff_password.GeneratePassword(newPassword, loginSalt)

	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	//update mch_login
	_, err := dbr.Table("mch_login").Data(map[string]interface{}{
		"login_secret": loginSecret,
		"login_salt":   loginSalt,
		"updated_at":   nowTime,
	}).Where("staff_id", staffId).Update()
	logrus.Debugf("Update mch_login. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("Update mch_login. Error : %v", err)
		return false
	}

	return true
}

//判断密码是否正确
func (p *StaffService) CheckIsCorrectPassword(staffId, oldPassword string) bool {
	mchLoginDao := mch_login.NewMchLogin()
	loginInfo, _ := mchLoginDao.GetLoginByStaffId(staffId, "login_secret, login_salt")
	if !ff_password.Check(fmt.Sprint(loginInfo["LoginSecret"]), fmt.Sprint(loginInfo["LoginSalt"]), oldPassword) {
		return false
	}
	return true
}

func (p *StaffService) DelStaff(staffId string) bool {
	nowTime := time.Now().Unix()
	dbr := ff_vars.DbConn.GetInstance()
	//开启事务
	dbr.Begin()

	//delete mch_staff
	_, err := dbr.Table("mch_staff").Data(map[string]interface{}{
		"is_delete":  1,
		"deleted_at": nowTime,
	}).Where("staff_id", staffId).Update()
	logrus.Debugf("delete mch_staff. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("delete mch_staff. Error : %v", err)
		return false
	}

	//delete mch_login
	_, err = dbr.Table("mch_login").Data(map[string]interface{}{
		"is_delete":  1,
		"deleted_at": nowTime,
	}).Where("staff_id", staffId).Update()
	logrus.Debugf("delete mch_login. LastSql : %v", dbr.LastSql)
	if err != nil {
		dbr.Rollback()
		logrus.Errorf("delete mch_login. Error : %v", err)
		return false
	}

	//提交事务
	dbr.Commit()
	return true
}

func (p *StaffService) CheckIsExitsStaffId(staffId string) bool {
	staffDao := mch_staff.NewMchStaff()
	return staffDao.CheckIsExitsStaffId(staffId)
}
