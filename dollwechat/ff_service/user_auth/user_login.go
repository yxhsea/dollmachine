package user_auth

import (
	"dollmachine/dollwechat/ff_redis"
	log "github.com/sirupsen/logrus"
	"bytes"
	"fmt"
	"strconv"
	"dollmachine/dollwechat/ff_common/ff_hash"
	"dollmachine/dollwechat/ff_common/ff_random"
	"time"
	"dollmachine/dollwechat/ff_config/ff_vars"
)

type UserInfo struct {
	OpenId       string `json:"openid"`    // 用户的标识, 对当前公众号唯一
	Nickname     string `json:"nickname"`  // 用户的昵称
	Sex          int    `json:"sex"`       // 用户的性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
	City         string `json:"city"`      // 用户所在城市
	Province     string `json:"province"`  // 用户所在省份
	Country      string `json:"country"`   // 用户所在国家
	HeadImageURL string `json:"headimgurl"` //头像
	UnionId      string `json:"unionid,omitempty"` // 只有在用户将公众号绑定到微信开放平台帐号后, 才会出现该字段.
	FromTag      string `json:"from_tag"`
}

type UserKey struct {
	SessionId    string `json:"session_id"`
	OpenId       string `json:"open_id"`
	SessionKey   string `json:"session_key"`
	UserId       int64  `json:"user_id"`
	NickName     string `json:"nick_name"`
	UserGender   int    `json:"user_gender"`
	MerchantId   int64  `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	UserType     int    `json:"user_type"`
}

type UserSession struct {
}

func NewUserSession() *UserSession {
	return &UserSession{}
}

func (p *UserSession) getCachePrefixSessionKey(sid string) string {
	return "dollmachine:user:session:info:" + sid
}

func (p *UserSession) getSessionKeyExpired() int64 {
	return 60 * 30 * 10
}

func (p *UserSession) IsExist(sid string) (bool, error) {
	sidKey := p.getCachePrefixSessionKey(sid)
	conn := ff_vars.RedisConn.Get()
	defer conn.Close()
	return ff_redis.NewString().Exists(conn, sidKey)
}

func AddOrUpdateUserLogin(userInfo *UserInfo) map[string]interface{}{
	dbr := ff_vars.Dbr
	var userLoginId int64
	if userInfo.UnionId != "" {
		Cnts, err := dbr.Table("usr_login").Where("union_id", "=", userInfo.UnionId).Count(1)
		log.Debug("[Query user openid failure through union_id] lastSql ", dbr.LastSql())
		if err != nil {
			log.Errorf("Query user openid failure through union_id. Error : %s ", err.Error())
		}
		if Cnts > 0 {
			usrLogin, err := dbr.Table("usr_login").Fields("user_id").Where("union_id", "=", userInfo.UnionId).First()
			if err != nil {
				log.Errorf("Query user info failure through union_id. Error : %s", err.Error())
			}
			userLoginId, _ = strconv.ParseInt(fmt.Sprint(usrLogin["user_id"]),10,64)
		}else{
			userLoginId, _ = ff_redis.NewUniqueId().GetUserId()
		}
	}

	//开启事务
	dbr.Begin()
	nowTime := time.Now().Unix()
	usrLoginSqlStr := addLogin(userInfo,userLoginId,nowTime, nowTime)
	_, err := dbr.Execute(usrLoginSqlStr)
	log.Debug("[add or update the data table usr_login] lastSql : ", dbr.LastSql())
	if err != nil {
		dbr.Rollback()
		log.Errorf("add or update the data table usr_login failure. Error : %s ", err.Error())
	}
	usrUserSqlStr := addUser(userInfo, userLoginId, nowTime, nowTime)
	_, err = dbr.Execute(usrUserSqlStr)
	log.Debug("[add or update the data table usr_user ] lastSql : ", dbr.LastSql())
	if err != nil {
		dbr.Rollback()
		log.Errorf("add or update the data table usr_login failure. Error : %s ", err.Error())
	}
	//提交事务
	dbr.Commit()

	//生成session_id
	sid := ff_hash.MD5EncodeToString(time.Now().String() + ff_random.KrandAll(64))
	userKey := &UserKey{
		UserId: userLoginId,
		SessionId:     sid,
		SessionKey: "",
		OpenId:     userInfo.OpenId,
		NickName:   userInfo.Nickname,
		UserGender: userInfo.Sex,
		MerchantId: 0,
		MerchantName: "",
		UserType:   1,
	}

	//将用户session写入redis
	newUserSession := NewUserSession()
	sidKey :=  newUserSession.getCachePrefixSessionKey(sid)
	conn := ff_vars.RedisConn.Get()
	defer conn.Close()
	_, err = ff_redis.NewString().Set(conn, sidKey, dbr.JsonEncode(userKey), newUserSession.getSessionKeyExpired(),)
	if err != nil {
		log.Errorf("User session information write cache failure. Error : %s ", err.Error())
	}

	return map[string]interface{}{"session_id":sid,"user_id":userLoginId,"nick_name":userInfo.Nickname,"avatar":userInfo.HeadImageURL}
}

func addLogin(userInfo *UserInfo, userId int64, createdAt int64, updatedAt int64) string {
	var buf bytes.Buffer
	buf.WriteString("INSERT INTO ")
	buf.WriteString("usr_login")
	buf.WriteString(" (login_token,union_id,login_secret,login_salt,login_type,user_id,created_at,updated_at,status,from_tag,gender,country,province,city,lat,lng,telephone) VALUES ")
	buf.WriteString("('%v', '%v', '%v', '%v', %v, %v, %v, %v, %v, '%v', %v, '%v', '%v', '%v', '%v', '%v', '%v')")
	buf.WriteString(" ON DUPLICATE KEY UPDATE updated_at=VALUES(updated_at),union_id=VALUES(union_id),gender=VALUES(gender),country=VALUES(country),province=VALUES(province),city=VALUES(city)")
	return fmt.Sprintf(buf.String(),userInfo.OpenId, userInfo.UnionId, userInfo.OpenId, userInfo.OpenId,1,userId,createdAt,updatedAt,1,userInfo.FromTag,userInfo.Sex,userInfo.Country,userInfo.Province,userInfo.City,"","","")
}

func addUser(userInfo *UserInfo, userId int64, createdAt int64, updatedAt int64) string{
	var buf bytes.Buffer
	buf.WriteString("INSERT INTO ")
	buf.WriteString("usr_user ")
	buf.WriteString("(user_id,merchant_id,name,nick_name,avatar,user_type,created_at,updated_at,status) VALUES ")
	buf.WriteString("(%v, %v, '%v', '%v', '%v', %v, %v, %v, %v) ")
	buf.WriteString("ON DUPLICATE KEY UPDATE updated_at=VALUES(updated_at),name=VALUES(name),nick_name=VALUES(nick_name),avatar=VALUES(avatar)")
	return fmt.Sprintf(buf.String(),userId,0,userInfo.Nickname,userInfo.Nickname,userInfo.HeadImageURL,1,createdAt,updatedAt,1)
}
