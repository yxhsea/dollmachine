package ff_setup

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
	"github.com/sirupsen/logrus"
	"dollmachine/dolluser/ff_config/ff_vars"
)

func SetupMysql(host, port, user, password, dbname, charset string, poolnum int){
	//Db配置
	var DbConfig = map[string]interface{}{
		"Default": "mysql_dev",
		"SetMaxOpenConns": poolnum,
		"SetMaxIdleConns": 10,
		"Connections": map[string]map[string]string{
			"mysql_dev": map[string]string{
				"host":     host,
				"username": user,
				"password": password,
				"port":     port,
				"database": dbname,
				"charset":  charset,
				"protocol": "tcp",
				"prefix":   "",
				"driver":   "mysql",
			},
		},
	}

	//dbr连接资源句柄
	dbr, err := gorose.Open(DbConfig)
	if err != nil {
		logrus.Errorf("Mysql init fail, Error : %v ", err.Error())
	}
	ff_vars.DbConn = dbr
}