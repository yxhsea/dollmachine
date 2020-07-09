// Package whitelist is a micro plugin for whitelisting service requests
package whitelist

import (
	"net/http"
	"strings"

	"github.com/micro/cli"
	"github.com/micro/go-micro/client"
	"github.com/micro/micro/plugin"
)

type whitelist struct {
	services map[string]bool
}

func (w *whitelist) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "rpc_whitelist",
			Usage:  "Comma separated whitelist of allowed services for RPC calls",
			EnvVar: "RPC_WHITELIST",
		},
	}
}

func (w *whitelist) Commands() []cli.Command {
	return nil
}

func (w *whitelist) Handler() plugin.Handler {
	return func(h http.Handler) http.Handler {
		return h
	}
}

func (w *whitelist) Init(ctx *cli.Context) error {
	if whitelist := ctx.String("rpc_whitelist"); len(whitelist) > 0 {
		client.DefaultClient = newClient(strings.Split(whitelist, ",")...)
	}
	return nil
}

func (w *whitelist) String() string {
	return "whitelist"
}

func NewRPCWhitelist(services ...string) plugin.Plugin {
	list := make(map[string]bool)

	for _, service := range services {
		list[service] = true
	}

	return &whitelist{
		services: list,
	}
}
