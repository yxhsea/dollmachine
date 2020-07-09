package ff_redis

import (
	"github.com/garyburd/redigo/redis"
)

type String struct {}

func NewString() *String{
	return &String{}
}

func (s *String) Set(conn redis.Conn, key string, value string, expire int64) (bool, error){
	valueStr, err := redis.String(conn.Do("SET", key, value, "EX", expire))
	if err != nil || valueStr != "OK"{
		return false, err
	}
	return true, nil
}

func (s *String) Get(conn redis.Conn, key string) (string, error){
	valueStr, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return valueStr, nil
}

func (s *String) Del(conn redis.Conn, key string) (bool, error){
	valueInt, err := redis.Int(conn.Do("DEL", key))
	if err != nil || valueInt != 1{
		return false, err
	}
	return true, nil
}

func (s *String) Exists(conn redis.Conn, key string) (bool, error) {
	valueInt, err := redis.Int(conn.Do("EXISTS", key))
	if err != nil || valueInt != 1 {
		return false, err
	}
	return true, nil
}
