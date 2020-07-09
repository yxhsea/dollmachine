package ff_vars

import (
	"github.com/go-redis/redis"
	"github.com/gohouse/gorose"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro"
	"dollmachine/dolluser/ff_common/ff_paygate/ff_payconf"
)

var RedisConn *redis.Client
var Upgrader websocket.Upgrader
var DbConn gorose.Connection
var RpcSrv micro.Service

var OrderPrefix int64
var NotifyEspUrl string
var NotifyUrl string

var PayConf *ff_payconf.PayConf
