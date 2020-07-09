package unique_id

import "dollmachine/dolluser/ff_config/ff_vars"

type UniqueId struct {
}

func NewUniqueId() *UniqueId {
	return &UniqueId{}
}

func (p *UniqueId) getCachePrefixKey(sid string) string {
	return "DollUnique:unique:id:" + sid
}

func (p *UniqueId) incrBy(prefix int64, key string, step int64) (int64, error) {
	valInt64, err := ff_vars.RedisConn.IncrBy(key, step).Result()
	if err != nil {
		return 0, err
	}
	return prefix + valInt64, nil
}

func (p *UniqueId) GetRechargeId() (int64, error) {
	var prefix int64 = ff_vars.OrderPrefix * 10000000000000
	return p.incrBy(prefix, p.getCachePrefixKey("rechargeId"), 1)
}
