package ff_vars

import (
	"github.com/go-redis/redis"
	"github.com/gohouse/gorose"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro"
)

var RedisConn *redis.Client
var Upgrader websocket.Upgrader
var DbConn gorose.Connection
var RpcSrv micro.Service
