package ff_base

import (
	"dollmachine/dollwechat/ff_redis"
	"gopkg.in/chanxuehong/wechat.v2/mp/user"
	log "github.com/sirupsen/logrus"
	"time"
	"gopkg.in/chanxuehong/wechat.v2/mp/message/template"
	"dollmachine/dollwechat/ff_config/ff_vars"
	"dollmachine/dollwechat/ff_common/ff_json"
	"dollmachine/dollwechat/ff_service/user_auth"
)

func UserLogin(eventKey string, fromUserName string) string {
	userInfo, err := user.Get(ff_vars.WxMpWechatClient, fromUserName, user.LanguageZhCN)
	if err != nil {
		log.Errorf("get user info Error : %s", err.Error())
		return ""
	}

	userSession := &user_auth.UserInfo{
		OpenId:userInfo.OpenId,
		Nickname:userInfo.Nickname,
		Sex:userInfo.Sex,
		City:userInfo.City,
		Province:userInfo.Province,
		Country:userInfo.Country,
		HeadImageURL:userInfo.HeadImageURL,
		UnionId:userInfo.UnionId,
	}
	//userMap := insertOrUpdateUserLogin(userInfo)
	userMap := user_auth.AddOrUpdateUserLogin(userSession)

	conn := ff_vars.RedisConn.Get()
	defer conn.Close()
	_, err = ff_redis.NewString().Set(conn, eventKey, ff_vars.Dbr.JsonEncode(userMap), 3 * 60)
	if err != nil {
		log.Errorf("Set userLoginKey fail Error : %s ", err.Error())
		return ""
	}

	return userInfo.Nickname
}

func ReplyContent(toUser string, userNickName string, createdAt int64){
	dataMap := map[string]interface{}{
		"first" : map[string]interface{}{
			"value":"您进行了微信扫一扫登录操作",
			"color":"#173177",
		},
		"keyword1":map[string]interface{}{
			"value": userNickName,
			"color":"#173177",
		},
		"keyword2": map[string]interface{}{
			"value":"翻番物娱",
			"color":"#173177",
		},
		"keyword3": map[string]interface{}{
			"value":time.Unix(createdAt, 0).Format("2006-01-02 15:04:05"),
			"color":"#173177",
		},
		"remark":map[string]interface{}{
			"value":"如有疑问，请直接在公众号内回复。",
			"color":"#173177",
		},
	}

	msg := template.TemplateMessage{
		ToUser : toUser,
		TemplateId : "drg5wKvLWgjB0cO2zPxMcv_De7UezGt4gBeAD-JxiBo",
		URL : "",
		Data : []byte(ff_json.MarshalToStringNoError(dataMap)),
	}
	msgId, err := template.Send(ff_vars.WxMpWechatClient, msg)
	if err != nil {
		log.Errorf("scan login send template msg fail : Error : %s ", err.Error())
	}
	log.Infof("scan login send template msgId : %s", msgId)
}