package ff_redis

import (
	"gogs.yetanotherpay.tk/yxhsea/DollWechat/ff_config/ff_vars"
	"github.com/garyburd/redigo/redis"
)

type UniqueId struct {}

func NewUniqueId() *UniqueId {
	return &UniqueId{}
}

func (p *UniqueId) getCachePrefixKey(sid string) string {
	return "dollmachine:unique:id:" + sid
}

func (p *UniqueId) incrBy(prefix int64, key string, step int64) (int64, error) {
	conn := ff_vars.RedisConn.Get()
	defer conn.Close()
	ValueInt64, err := redis.Int64(conn.Do("INCRBY", key, step))
	if err != nil {
		return 0, err
	}
	return prefix + ValueInt64, nil
}

func (p *UniqueId) GetUserId() (int64, error) {
	var prefix int64 = 100*1000000000 + 238160
	return p.incrBy(prefix, p.getCachePrefixKey("userid"), 1)
}