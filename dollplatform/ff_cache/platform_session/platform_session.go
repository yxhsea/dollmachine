package platform_session

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollplatform/ff_common/ff_hash"
	"dollmachine/dollplatform/ff_common/ff_json"
	"dollmachine/dollplatform/ff_common/ff_random"
	"dollmachine/dollplatform/ff_config/ff_vars"
	"time"
)

type PlatformSession struct {
	Token      string `json:"token"`
	ManagerId  int64  `json:"manager_id"`
	RoleId     int64  `json:"role_id"`
	LoginToken string `json:"login_token"`
	NickName   string `json:"nick_name"`
	Rules      string `json:"rules"`
}

func NewPlatformSession() *PlatformSession {
	return &PlatformSession{}
}

func (p *PlatformSession) getCachePrefixSessionKey(token string) string {
	return "dollPlatform:platform:session:info:" + token
}

func (p *PlatformSession) getSessionKeyExpired() time.Duration {
	return time.Minute * 30
}

func (p *PlatformSession) getToken() string {
	sid := ff_hash.MD5EncodeToString(time.Now().String() + ff_random.KrandAll(64))
	return sid
}

func (p *PlatformSession) SetPlatformSession(platformKey *PlatformSession) bool {
	token := p.getToken()
	platformKey.Token = token
	conn := ff_vars.RedisConn
	err := conn.Set(p.getCachePrefixSessionKey(token), ff_json.MarshalToStrNoErr(platformKey), p.getSessionKeyExpired()).Err()
	if err != nil {
		logrus.Errorf("Set user session failure. Error : %v", err)
		return false
	}
	return true
}

func (p *PlatformSession) GetPlatformSession(token string) *PlatformSession {
	key := p.getCachePrefixSessionKey(token)
	conn := ff_vars.RedisConn
	valStr, err := conn.Get(key).Result()
	if err != nil {
		logrus.Errorf("Get user session by token failure. Error : %v", err)
		return nil
	}

	var platformSession *PlatformSession
	ff_json.Unmarshal(valStr, &platformSession)
	return platformSession
}

func (p *PlatformSession) CheckIsExitsByToken(token string) bool {
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

func (p *PlatformSession) DeleteToken(token string) bool {
	key := p.getCachePrefixSessionKey(token)
	conn := ff_vars.RedisConn
	valInt64, err := conn.Del(key).Result()
	if err != nil {
		logrus.Errorf("Delete token failure. Error : %v", err)
		return false
	}
	if valInt64 > 0 {
		return true
	}
	return false
}
