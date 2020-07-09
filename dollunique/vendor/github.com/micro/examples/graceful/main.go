package main

import (
	"fmt"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

func main() {
	service := micro.NewService()

	service.Server().Init(
		server.Wait(true),
	)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
