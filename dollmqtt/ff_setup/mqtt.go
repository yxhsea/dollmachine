package ff_setup

import (
	"github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"dollmachine/dollmqtt/ff_common/ff_json"
	"dollmachine/dollmqtt/ff_common/ff_mqttmsg"
	"dollmachine/dollmqtt/ff_config/ff_vars"
	"dollmachine/dollmqtt/ff_mqtt_v1"
)

func SetupMqttClient(broker string, user string, password string, clientId string, store string, topic string, qos int64) error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientId)
	opts.SetUsername(user)
	opts.SetPassword(password)
	opts.SetCleanSession(false)
	if store != ":memory:" {
		opts.SetStore(mqtt.NewFileStore(store))
	}
	//opts.WillQos = 2

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	ff_vars.MqttClient = client
	ff_vars.UrlMqtt = broker
	ff_vars.TopicMqtt = topic
	ff_vars.QosMqtt = qos

	//ff_vars.MqttClient.AddRoute(ff_vars.TopicMqtt,SetupDefaultMqttSubMessageHandler)
	if token := ff_vars.MqttClient.Subscribe(topic, 0, SetupDefaultMqttSubMessageHandler); token.Wait() && token.Error() != nil {
		log.Errorf("Subscribe Error : %v ",token.Error())
		return token.Error()
	}

	return nil
}

func SetupDefaultMqttSubMessageHandler(client mqtt.Client, msg mqtt.Message) {
	//基本数据包解析
	BasePkgStr := string(msg.Payload())
	var BasePkg *ff_mqttmsg.BasePkg
	ff_json.Unmarshal(BasePkgStr, &BasePkg)
	log.Info("订阅到的数据包", BasePkgStr)

	switch BasePkg.Action {
	case "HEART": //心跳包
		ff_mqtt_v1.HeartPkgHandle(ff_json.MarshalToStringNoError(BasePkg.Content))
		return
	case "STARTUP": //开机、上线
		ff_mqtt_v1.StartUpPkgHandle(ff_json.MarshalToStringNoError(BasePkg.Content))
		return
	case "BEGIN": //开始游戏
		return
	case "WILLDOWN": //设备掉线
		ff_mqtt_v1.WillDownPkgHandle(ff_json.MarshalToStringNoError(BasePkg.Content))
		return
	case "OVER": //游戏结束 结果上报
		ff_mqtt_v1.GameOverHandle(BasePkg.DeviceId, ff_json.MarshalToStringNoError(BasePkg.Content))
		//var queue *ff_queue.GameResQueue
		//go queue.PushGameResQueue(BasePkgStr)
		return
	default:
		return
	}
}
