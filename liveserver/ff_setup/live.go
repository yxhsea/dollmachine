package ff_setup

import (
	"github.com/gorilla/websocket"
	"net/http"
	"dollmachine/liveserver/ff_config/ff_vars"
)

func SetupLive(readBuffer int64, writeBuffer int64) error {
	ff_vars.Upgrader = websocket.Upgrader{
		ReadBufferSize:  int(readBuffer),
		WriteBufferSize: int(writeBuffer),
		CheckOrigin:     func(r *http.Request) bool { return true },
		Subprotocols:    []string{"binary"},
	}
	return nil
}
