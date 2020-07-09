package main

import (
	"context"

	"github.com/micro/go-micro"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, name *string, msg *string) error {
	*msg = "Hello " + *name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
	)
	service.Init()

	// set the handler
	service.Server().Handle(
		service.Server().NewHandler(
			new(Greeter),
		),
	)

	service.Run()
}
