package ff_v1live

import (
	"github.com/sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"dollmachine/liveserver/ff_common/ff_live"
	"dollmachine/liveserver/ff_config/ff_vars"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//服务端接收视频流
func LiveHandle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	deviceArr := strings.Split(ps.ByName("device_id"), "_")
	device, _ := ff_live.GetDevice(deviceArr[0])
	typeId, _ := strconv.Atoi(deviceArr[1])
	chunk := make([]byte, 40*1024)
	for {
		time.Sleep(40 * time.Millisecond)
		n, err := r.Body.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			}
		} else {
			if n > 0 {
				//fmt.Println(device.DeviceId, n, typeId)
				device.Broadcast(chunk[:n], typeId)
			} else {
				break
			}
		}
	}
}

//前端拉取1号摄像头视频流
func LiveVideoOne(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.Header.Del("Sec-WebSocket-Protocol")
	device, _ := ff_live.GetDevice(ps.ByName("device_id"))
	ws, err := ff_vars.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{"device_id": device.DeviceId}).Errorf("Websocket upgrade failed: ", r.RemoteAddr, err.Error())
		return
	}

	client := &ff_live.Client{}
	client.DeviceId = device.DeviceId
	client.WsVideoOne = ws
	client.SendQueueOne = make(chan []byte, 1024*40)
	device.Register <- client

	go client.SendVideoOneChan()

	client.ListenOne()
}

//前端拉取2号摄像头视频流
func LiveVideoTwo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.Header.Del("Sec-WebSocket-Protocol")
	device, _ := ff_live.GetDevice(ps.ByName("device_id"))
	ws, err := ff_vars.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{"device_id": device.DeviceId}).Errorf("Websocket upgrade failed: ", r.RemoteAddr, err.Error())
		return
	}

	client := &ff_live.Client{}
	client.DeviceId = device.DeviceId
	client.WsVideoTwo = ws
	client.SendQueueTwo = make(chan []byte, 1024*40)
	device.Register <- client

	go client.SendVideoTwoChan()

	client.ListenTwo()
}
