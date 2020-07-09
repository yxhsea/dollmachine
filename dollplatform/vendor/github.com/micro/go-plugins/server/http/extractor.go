package http

import (
	"strconv"
	"strings"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"

	"github.com/micro/util/go/lib/addr"
)

func serviceDef(opts server.Options) *registry.Service {
	var advt, host string
	var port int

	if len(opts.Advertise) > 0 {
		advt = opts.Advertise
	} else {
		advt = opts.Address
	}

	parts := strings.Split(advt, ":")
	if len(parts) > 1 {
		host = strings.Join(parts[:len(parts)-1], ":")
		port, _ = strconv.Atoi(parts[len(parts)-1])
	} else {
		host = parts[0]
	}

	addr, err := addr.Extract(host)
	if err != nil {
		addr = host
	}

	node := &registry.Node{
		Id:       opts.Name + "-" + opts.Id,
		Address:  addr,
		Port:     port,
		Metadata: opts.Metadata,
	}

	node.Metadata["server"] = "http"

	return &registry.Service{
		Name:    opts.Name,
		Version: opts.Version,
		Nodes:   []*registry.Node{node},
	}
}
