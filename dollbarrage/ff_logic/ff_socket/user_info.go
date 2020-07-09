package ff_socket

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"dollmachine/dollbarrage/ff_config/ff_const"
	"dollmachine/dollbarrage/ff_config/ff_vars"
)

//处理消息
func HandlerMsg(msg []byte) []byte {
	var msgMap map[string]interface{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(msg, &msgMap)

	var msgJson []byte
	if msgMap["user_id"] == "null" {
		replyMsg := map[string]interface{}{
			"user_id":   "",
			"nick_name": "",
			"avatar":    "",
			"msg":       "",
		}
		msgJson, _ = json.Marshal(replyMsg)
	} else {
		oneUser := getUserInfo(fmt.Sprint(msgMap["user_id"]))
		replyMsg := map[string]interface{}{
			"user_id":   msgMap["user_id"],
			"nick_name": oneUser.NickName,
			"avatar":    oneUser.Avatar,
			"msg":       msgMap["msg"],
		}

		msgJson, _ = json.Marshal(replyMsg)
	}

	return []byte(msgJson)
}

type UserInfo struct {
	UserId   string `json:"user_id"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

//获取用户信息
func getUserInfo(userId string) *UserInfo {
	flag := isExist(userId)
	if !flag {
		setUserInfoByCache(userId, ff_vars.DbConn.JsonEncode(getUserInfoByDB(userId)))
	}
	oneUser := getUserInfoByCache(userId)

	var user *UserInfo
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal([]byte(oneUser), &user)
	return user
}

//从数据库中查询用户信息
func getUserInfoByDB(userId string) map[string]interface{} {
	user := ff_vars.DbConn.GetInstance().Table("usr_user")
	oneUser, err := user.Fields("user_id,nick_name,avatar").Where("user_id", userId).Limit(1).First()
	logrus.Debugf("Query user info LastSql : %v", user.LastSql)
	if err != nil {
		logrus.Fatalf("Query user info failure, Error : %v", err.Error())
	}
	return oneUser
}

//从缓存中查询用户信息
func getUserInfoByCache(userId string) string {
	conn := ff_vars.RedisConn
	key := ff_const.CacheDollBarrageUserInfo + ":" + userId
	val, err := conn.Get(key).Result()
	if err != nil {
		logrus.Fatalf("Query user information failure from caching. Error : %v", err.Error())
	}
	return val
}

//缓存用户信息
func setUserInfoByCache(userId string, userInfo string) {
	conn := ff_vars.RedisConn
	key := ff_const.CacheDollBarrageUserInfo + ":" + userId
	err := conn.Set(key, userInfo, 0).Err()
	if err != nil {
		logrus.Fatalf("User information write cache failure. Error : %v", err.Error())
	}
}

//判断是否存在
func isExist(userId string) bool {
	conn := ff_vars.RedisConn
	key := ff_const.CacheDollBarrageUserInfo + ":" + userId
	val, err := conn.Exists(key).Result()
	if err != nil {
		logrus.Fatalf("Query failure. Error : %v", err.Error())
	}
	if val > 0 {
		return true
	}
	return false
}
