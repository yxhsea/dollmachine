package ff_logic

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"dollmachine/dollbarrage/ff_config/ff_vars"
	"dollmachine/dollbarrage/ff_logic/ff_socket"
)

func HttpHandler(ctx *gin.Context) {
	ctx.Request.Header.Del("Sec-WebSocket-Protocol")
	conn, err := ff_vars.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logrus.Errorf("WebSocket Upgrade failed Error : %v ", err.Error())
	}
	logrus.Debugf("remote address : %v", conn.RemoteAddr())

	sid, _ := com.StrTo(ctx.Param("sid")).Int64()
	service := ff_socket.NewService(sid)
	client := &ff_socket.Client{
		Service: service,
		Conn:    conn,
		Send:    make(chan []byte, 1024),
	}
	client.Service.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
