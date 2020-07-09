package ff_setup

import (
	"dollmachine/dollwechat/ff_router"
	"github.com/gin-gonic/gin"
	"github.com/gogap/logrus"
)

func SetupServer(host string, mode string) error {
	router := gin.New()
	ff_router.SetHttpRouter(router, mode)
	ff_router.SetWxRouter()
	err := router.Run(host)
	if err != nil {
		logrus.Errorf("Http Server boot failure. Error : %v", err)
		return err
	}
	return nil
}
