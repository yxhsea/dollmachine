package merchant_session

import (
	"github.com/sirupsen/logrus"
	"dollmachine/dollmerchant/ff_common/ff_hash"
	"dollmachine/dollmerchant/ff_common/ff_json"
	"dollmachine/dollmerchant/ff_common/ff_random"
	"dollmachine/dollmerchant/ff_config/ff_vars"
	"time"
)

type MerchantSession struct {
	Token        string `json:"token"`
	MerchantId   int64  `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
	StaffId      int64  `json:"staff_id"`
	StaffName    string `json:"staff_name"`
	StaffPhone   string `json:"staff_phone"`
	RoleId       int64  `json:"role_id"`
	LoginToken   string `json:"login_token"`
	NickName     string `json:"nick_name"`
	Rules        string `json:"rules"`
}

func NewMerchantSession() *MerchantSession {
	return &MerchantSession{}
}

func (p *MerchantSession) getCachePrefixSessionKey(token string) string {
	return "DollMerchant:merchant:session:info:" + token
}

func (p *MerchantSession) getSessionKeyExpired() time.Duration {
	return time.Minute * 30
}

func (p *MerchantSession) getToken() string {
	sid := ff_hash.MD5EncodeToString(time.Now().String() + ff_random.KrandAll(64))
	return sid
}

func (p *MerchantSession) SetMerchantSession(MerchantKey *MerchantSession) bool {
	token := p.getToken()
	MerchantKey.Token = token
	conn := ff_vars.RedisConn
	err := conn.Set(p.getCachePrefixSessionKey(token), ff_json.MarshalToStrNoErr(MerchantKey), p.getSessionKeyExpired()).Err()
	if err != nil {
		logrus.Errorf("Set merchant session failure. Error : %v", err)
		return false
	}
	return true
}

func (p *MerchantSession) GetMerchantSession(token string) *MerchantSession {
	key := p.getCachePrefixSessionKey(token)
	conn := ff_vars.RedisConn
	valStr, err := conn.Get(key).Result()
	if err != nil {
		logrus.Errorf("Get merchant session by token failure. Error : %v", err)
		return nil
	}

	var MerchantSession *MerchantSession
	ff_json.Unmarshal(valStr, &MerchantSession)
	return MerchantSession
}

func (p *MerchantSession) CheckIsExitsByToken(token string) bool {
	key := p.getCachePrefixSessionKey(token)
	conn := ff_vars.RedisConn
	valInt64, err := conn.Exists(key).Result()
	if err != nil {
		logrus.Errorf("Check merchant token is exits failure. Error : %v", err)
		return false
	}
	if valInt64 > 0 {
		return true
	}
	return false
}

func (p *MerchantSession) DeleteToken(token string) bool {
	key := p.getCachePrefixSessionKey(token)
	conn := ff_vars.RedisConn
	valInt64, err := conn.Del(key).Result()
	if err != nil {
		logrus.Errorf("Delete merchant token failure. Error : %v", err)
		return false
	}
	if valInt64 > 0 {
		return true
	}
	return false
}
