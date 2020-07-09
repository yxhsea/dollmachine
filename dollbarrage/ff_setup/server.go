package ff_setup

import (
	"github.com/gin-gonic/gin"
	"dollmachine/dollbarrage/ff_logic"
)

func SetupServer(host string) {
	router := gin.Default()
	gin.SetMode("debug")
	router.GET("/ws/:sid", ff_logic.HttpHandler)
	router.Run(host)
}
