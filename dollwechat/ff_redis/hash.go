package ff_redis

import (
	"github.com/garyburd/redigo/redis"
)

type Hash struct {

}

func NewHash() *Hash{
	return &Hash{}
}

func (h *Hash) HSet(conn redis.Conn, key string, field string, value string) (bool, error){
	valueInt, err := redis.Int(conn.Do("HSET", key, field, value))
	if err != nil || valueInt == 0 {
		return false, err
	}
	return true, nil
}

func (h *Hash) HMSet(conn redis.Conn, keyFieldValue ...interface{}) (bool, error){
	valueStr, err := conn.Do("HMSET", keyFieldValue...)
	if err != nil || valueStr != "OK"{
		return false, err
	}
	return true, nil
}

func (h *Hash) HGet(conn redis.Conn, key string, field string) (string, error){
	valueStr, err := redis.String(conn.Do("HGET", key, field))
	if err != nil {
		return "", err
	}
	return valueStr, nil
}

func (h *Hash) HGetAll(conn redis.Conn, key string) (map[string]string, error) {
	valueMap, err := redis.StringMap(conn.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}
	return valueMap, nil
}

func (h *Hash) HDel(conn redis.Conn, key string, field string) (bool, error) {
	valueInt, err := redis.Int(conn.Do("HDEL", key, field))
	if err != nil || valueInt == 0 {
		return false, err
	}
	return true, nil
}

func (h *Hash) HExists(conn redis.Conn, key string, field string) (bool, error) {
	valueInt, err := redis.Int(conn.Do("HEXISTS", key, field))
	if err != nil || valueInt == 0 {
		return false, err
	}
	return true, nil
}

func (h *Hash) HLen(conn redis.Conn, key string) (int, error) {
	valueInt, err := redis.Int(conn.Do("HLEN", key))
	if err != nil || valueInt == 0 {
		return 0, err
	}
	return valueInt, nil
}