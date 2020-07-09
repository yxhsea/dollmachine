package ff_setup

import (
	"dollmachine/dollbarrage/ff_config/ff_vars"
	"net/http"
	"github.com/gorilla/websocket"
)

func SetupWebSocket(readBuffer int64, writeBuffer int64) error {
	ff_vars.Upgrader = websocket.Upgrader{
		ReadBufferSize:  int(readBuffer),
		WriteBufferSize: int(writeBuffer),
		CheckOrigin:     func(r *http.Request) bool { return true },
		Subprotocols:    []string{"binary"},
	}
	return nil
}
