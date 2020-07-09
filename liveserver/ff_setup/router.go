package ff_setup

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"dollmachine/liveserver/ff_controller/ff_v1/ff_v1live"
)

func SetupRouter(host string, mode string) error{
	prefixUrl := "/" + mode
	router := httprouter.New()
	router.POST(prefixUrl + "/ws/live/:device_id", ff_v1live.LiveHandle)               //服务端接收视频流
	router.GET(prefixUrl + "/ws/video/one/:device_id", ff_v1live.LiveVideoOne) //前端拉取1号摄像头视频流
	router.GET(prefixUrl + "/ws/video/two/:device_id", ff_v1live.LiveVideoTwo) //前端拉取2号摄像头视频流
	err := http.ListenAndServe(host, router)
	if err != nil {
		return err
	}
	return nil
}