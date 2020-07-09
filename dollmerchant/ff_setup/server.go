package ff_setup

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"dollmachine/dollmerchant/ff_config/ff_vars"
)

func SetupRpcServer(host string) {
	service := micro.NewService(
		micro.Registry(consul.NewRegistry(consul.Config(&api.Config{Address: host}))),
	)
	ff_vars.RpcSrv = service
}

func SetupServer(host string) {
	router := gin.Default()
	gin.SetMode("debug")

	//Api文档
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	SetupRouter(router)

	router.Run(host)
}
