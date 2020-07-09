package ff_live

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	DeviceId     string
	WsVideoOne   *websocket.Conn
	WsVideoTwo   *websocket.Conn
	SendQueueOne chan []byte
	SendQueueTwo chan []byte
}

func (Client *Client) WriteVideoOne(msgType int, payload []byte) error {
	writeTimeout := 30 * time.Second

	Client.WsVideoOne.SetWriteDeadline(time.Now().Add(writeTimeout))
	return Client.WsVideoOne.WriteMessage(msgType, payload)
}

func (Client *Client) SendVideoOneChan() {
	ticker := time.NewTicker(30 * time.Second)
	defer func() {
		ticker.Stop()
		Client.WsVideoOne.Close()
	}()
loop:
	for {
		select {
		case message, ok := <-Client.SendQueueOne:
			if ok {
				err := Client.WriteVideoOne(websocket.BinaryMessage, message)
				if err != nil {
					//logrus.WithFields(logrus.Fields{"device_id": Client.DeviceId}).Errorf("%d 号摄像头推流失败~~", 1)
					break loop
				}
			} else {
				break loop
			}
		case <-ticker.C:
		}
	}
}

func (Client *Client) ListenOne() {
	defer func() {
		Client.WsVideoOne.Close()
	}()

	Client.WsVideoOne.SetReadLimit(maxMessageSize)
	Client.WsVideoOne.SetReadDeadline(time.Now().Add(pongWait * 60))
	Client.WsVideoOne.SetPongHandler(func(string) error {
		Client.WsVideoOne.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, _, err := Client.WsVideoOne.ReadMessage()
		if err != nil {
			//fmt.Println("Client.WsVideoOne.ReadMessage", err.Error())
			logrus.Errorf("Client.WsVideoOne.ReadMessage %s", err.Error())
			break
		}
	}
}

func (Client *Client) WriteVideoTwo(msgType int, payload []byte) error {
	writeTimeout := 30 * time.Second

	Client.WsVideoTwo.SetWriteDeadline(time.Now().Add(writeTimeout))
	return Client.WsVideoTwo.WriteMessage(msgType, payload)
}

func (Client *Client) SendVideoTwoChan() {
	ticker := time.NewTicker(30 * time.Second)
	defer func() {
		ticker.Stop()
		Client.WsVideoTwo.Close()
	}()
loop:
	for {
		select {
		case message, ok := <-Client.SendQueueTwo:
			if ok {
				err := Client.WriteVideoTwo(websocket.BinaryMessage, message)
				if err != nil {
					//logrus.WithFields(logrus.Fields{"device_id": Client.DeviceId}).Errorf("%d 号摄像头推流失败~~", 2)
				}
			} else {
				break loop
			}
		case <-ticker.C:
		}
	}
}

func (Client *Client) ListenTwo() {
	defer func() {
		Client.WsVideoTwo.Close()
	}()

	Client.WsVideoTwo.SetReadLimit(maxMessageSize)
	Client.WsVideoTwo.SetReadDeadline(time.Now().Add(pongWait * 60))
	Client.WsVideoTwo.SetPongHandler(func(string) error {
		Client.WsVideoTwo.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, _, err := Client.WsVideoTwo.ReadMessage()
		if err != nil {
			//fmt.Println("Client.WsVideoTwo.ReadMessage", err.Error())
			logrus.Errorf("Client.WsVideoTwo.ReadMessage %s", err.Error())
			break
		}
	}
}
