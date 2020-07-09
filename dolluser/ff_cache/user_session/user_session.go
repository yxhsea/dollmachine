package user_session

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_common/ff_hash"
	"dollmachine/dolluser/ff_common/ff_json"
	"dollmachine/dolluser/ff_common/ff_random"
	"dollmachine/dolluser/ff_config/ff_vars"
	"time"
)

type UserSession struct {
	Token      string `json:"token"`
	OpenId     string `json:"open_id"`
	UserId     int64  `json:"user_id"`
	NickName   string `json:"nick_name"`
	MerchantId int64  `json:"merchant_id"`
}

func NewUserSession() *UserSession {
	return &UserSession{}
}

func (p *UserSession) getCachePrefixSessionKey(token string) string {
	return "DollUser:user:session:info:" + token
}

func (p *UserSession) getSessionKeyExpired() time.Duration {
	return time.Minute * 30
}

func (p *UserSession) getToken() string {
	sid := ff_hash.MD5EncodeToString(time.Now().String() + ff_random.KrandAll(64))
	return sid
}

func (p *UserSession) SetUserSession(userKey *UserSession) bool {
	token := p.getToken()
	userKey.Token = token
	conn := ff_vars.RedisConn
	err := conn.Set(p.getCachePrefixSessionKey(token), ff_json.MarshalToStringNoError(userKey), p.getSessionKeyExpired()).Err()
	if err != nil {
		logrus.Errorf("Set user session failure. Error : %v", err)
		return false
	}
	return true
}

func (p *UserSession) GetUserSession(token string) *UserSession {
	key := p.getCachePrefixSessionKey(token)
	conn := ff_vars.RedisConn
	valStr, err := conn.Get(key).Result()
	if err != nil {
		logrus.Errorf("Get user session by token failure. Error : %v", err)
		return nil
	}

	var userSession *UserSession
	ff_json.Unmarshal(valStr, &userSession)
	return userSession
}

func (p *UserSession) CheckIsExitsByToken(token string) bool {
	key := p.getCachePrefixSessionKey(token)
	conn := ff_vars.RedisConn
	valInt64, err := conn.Exists(key).Result()
	if err != nil {
		logrus.Errorf("Check token is exits failure. Error : %v", err)
		return false
	}
	if valInt64 > 0 {
		return true
	}
	return false
}
