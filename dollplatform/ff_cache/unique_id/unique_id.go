package unique_id

import "dollmachine/dollplatform/ff_config/ff_vars"

type UniqueId struct {
}

func NewUniqueId() *UniqueId {
	return &UniqueId{}
}

func (p *UniqueId) getCachePrefixKey(sid string) string {
	return "dollPlatform:unique:id:" + sid
}

func (p *UniqueId) incrBy(prefix int64, key string, step int64) (int64, error) {
	valInt64, err := ff_vars.RedisConn.IncrBy(key, step).Result()
	if err != nil {
		return 0, err
	}
	return prefix + valInt64, nil
}

func (p *UniqueId) GetMerchantId() (int64, error) {
	var prefix int64 = 200*1000000000 + 432743
	return p.incrBy(prefix, p.getCachePrefixKey("merchantId"), 1)
}

func (p *UniqueId) GetStaffId() (int64, error) {
	var prefix int64 = 210*1000000000 + 432743
	return p.incrBy(prefix, p.getCachePrefixKey("staffId"), 1)
}

func (p *UniqueId) GetDeviceTypeId() (int64, error) {
	var prefix int64 = 290*1000000 + 29283
	return p.incrBy(prefix, p.getCachePrefixKey("deviceTypeId"), 1)
}
