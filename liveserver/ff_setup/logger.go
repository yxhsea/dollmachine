package ff_setup

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"github.com/zbindenren/logrus_mail"
	"fmt"
	"os"
	"errors"
	"time"
)

func SetupLogger(loggerMap map[string]interface{}, mailConfig map[string]interface{}) error {
	setupBaseLog(loggerMap["level"].(string))
	if loggerMap["enable"] == true {
		setupFileLog(loggerMap["filepath"].(string), loggerMap["maxdays"].(int64), loggerMap["maxlines"].(int64), loggerMap["maxsize"].(int64))
	}
	if mailConfig["enable"] == true {
		setUpMailLog(mailConfig["host"].(string), mailConfig["port"].(string), mailConfig["from_user"].(string), mailConfig["password"].(string), mailConfig["to_user"].(string))
	}
	return nil
}

func setupBaseLog(logLevel string) error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	return nil
}

//设置文件日志
func setupFileLog(filePath string, maxdays int64, maxlines int64, maxsize int64) error {
	timeDate := time.Now().Format("2006-01-02")
	filePath = filePath + "_" + timeDate + ".log"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return errors.New("Failed to log to file, using default stderr")
	}
	logrus.SetOutput(file)
	return nil
}

//设置邮件同步日志
/**
logrus.WithFields(logrus.Fields{"file": "1234.txt","content": "GG",}).Error("这是测试")
 */
func setUpMailLog(host string, port string, fromUser string, passWord string, toUser string){
	Port, _ := strconv.Atoi(port)
	//首先开启smtp服务，最后两个参数是smtp的用户名和密码
	hook, err := logrus_mail.NewMailAuthHook("LiveServer", host, Port, fromUser, toUser, fromUser, passWord)
	if err != nil {
		fmt.Println("setUpEmail err :",err.Error())
	}
	logrus.AddHook(hook)
}