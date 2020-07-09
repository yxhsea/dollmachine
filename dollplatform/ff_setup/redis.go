package ff_setup

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"dollmachine/dollplatform/ff_config/ff_vars"
)

func SetupRedis(host string, auth string) {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: auth,
		DB:       3,
	})
	_, err := client.Ping().Result()
	if err != nil {
		logrus.Errorf("Redis init fail, Error : %v ", err.Error())
	}
	ff_vars.RedisConn = client
}
