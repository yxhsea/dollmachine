package ff_vars

import (
	"github.com/go-redis/redis"
	"github.com/gohouse/gorose"
	"github.com/micro/go-micro"
	"github.com/qiniu/api.v7/storage"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

var RedisConn *redis.Client
var DbConn gorose.Connection
var RpcSrv micro.Service

var MpAppId string
var MpAppSecret string
var MpAccessTokenServer core.AccessTokenServer
var MpWechatClient *core.Client
var PcAppId string
var PcAppSecret string
var PcAccessTokenServer core.AccessTokenServer
var PcWechatClient *core.Client
var MiniUserAppId string
var MiniUserAppSecret string
var MiniUserAccessTokenServer core.AccessTokenServer
var MiniUserWechatClient *core.Client

var QiNiuAccessKey string
var QiNiuSecretKey string
var QiNiuZone *storage.Zone
