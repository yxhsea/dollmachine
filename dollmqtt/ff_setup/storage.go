package ff_setup

import (
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
	"dollmachine/dollmqtt/ff_config/ff_vars"
	"time"
	"fmt"
)

func SetupMysql(user, password, host, port, dbname, charset string, poolNum int) error {
	//Db配置
	var DbConfig = map[string]interface{}{
		// Default database configuration
		"Default": "mysql_dev",
		// (Connection pool) Max open connections, default value 0 means unlimit.
		"SetMaxOpenConns": poolNum,
		// (Connection pool) Max idle connections, default value is 1.
		"SetMaxIdleConns": 10,
		// Define the database configuration character "mysql_dev".
		"Connections": map[string]map[string]string{
			"mysql_dev": map[string]string{
				"host":     host,
				"username": user,
				"password": password,
				"port":     port,
				"database": dbname,
				"charset":  charset,
				"protocol": "tcp",
				"prefix":   "",      // Table prefix
				"driver":   "mysql", // Database driver(mysql,sqlite,postgres,oracle,mssql)
			},
		},
	}

	ff_vars.Dbr, ff_vars.Err = gorose.Open(DbConfig)
	if ff_vars.Err != nil {
		return ff_vars.Err
	}

	go func() {
		t1 := time.NewTimer(time.Second * 3600)
		for {
			select {
			case <-t1.C:
				res, err := ff_vars.Dbr.Table("usr_user").Fields("user_id").First()
				if err != nil {
					fmt.Println(err.Error())
				}
				fmt.Println(res)
				t1.Reset(time.Second * 3600)
			}
		}
	}()

	return nil
}

func SetupRedis(host, auth string, poolNum int) error {
	ff_vars.RedisConn = &redis.Pool{
		MaxIdle:     poolNum,
		MaxActive:   poolNum,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if auth != "" {
				if _, err := c.Do("AUTH", auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}
