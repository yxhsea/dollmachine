package ff_vars

import (
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/garyburd/redigo/redis"
	"github.com/gocraft/dbr"
	"github.com/gohouse/gorose"
	"github.com/gorilla/websocket"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

var MysqlConn *dbr.Connection
var Dbr gorose.Connection
var Err error
var RedisConn *redis.Pool

var Upgrader websocket.Upgrader
var MqttClient mqtt.Client
var UrlMqtt string

var WxMpAppId string
var WxMpAppSecret string
var WxOriId string
var WxToken string
var WxEncodedAesKey string
var WxMpAccessTokenServer core.AccessTokenServer
var WxMpWechatClient *core.Client