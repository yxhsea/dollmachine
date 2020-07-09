package ff_queue

import (
	"github.com/sirupsen/log"
	"dollmachine/dollmqtt/ff_config/ff_vars"
)

type GameResQueue struct {
}

func (p *GameResQueue) getGameResQueueKey() string {
	return "dollmachine:game:result:queue"
}

func (p *GameResQueue) PushGameResQueue(dataStr string) {
	_, err := ff_vars.RedisConn.Get().Do("rpush", p.getGameResQueueKey(), dataStr)
	if err != nil {
		log.Errorf("[Push queue Game result fail : %s ]", err.Error())
	}
	log.Debug("游戏结果推入队列~")
}
