package ff_vars

import (
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/garyburd/redigo/redis"
	"github.com/gohouse/gorose"
)

var Dbr gorose.Connection
var Err error
var RedisConn *redis.Pool

var MqttClient mqtt.Client
var UrlMqtt string
var TopicMqtt string
var QosMqtt int64
