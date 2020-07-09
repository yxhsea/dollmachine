package ff_redis

import (
	"github.com/garyburd/redigo/redis"
)

type List struct {}

func NewList() *List{
	return &List{}
}

func (l *List) LPush(conn redis.Conn, key string, value string) (bool, error){
	valueInt, err := redis.Int(conn.Do("LPUSH", key, value))
	if err != nil || valueInt == 0 {
		return false, err
	}
	return true, nil
}

func (l *List) RPop(conn redis.Conn, key string) (string, error){
	valueStr, err := redis.String(conn.Do("RPOP", key))
	if err != nil {
		return "", err
	}
	return valueStr, nil
}

func (l *List) LRange(conn redis.Conn, key string, startIndex int, lastIndex int) (map[string]string, error){
	valueMap, err := redis.StringMap(conn.Do("LRANGE", startIndex, lastIndex))
	if err != nil || valueMap == nil {
		return nil, err
	}
	return valueMap, nil
}

func (l List) LLen(conn redis.Conn, key string) (int, error){
	valueInt, err := redis.Int(conn.Do("LLEN", key))
	if err != nil {
		return 0, err
	}
	return valueInt, nil
}
