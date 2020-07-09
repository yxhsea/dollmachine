// Package mdns provides a multicast DNS registry. Implementation https://godoc.org/github.com/micro/go-micro/registry/mdns
package mdns

import (
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/mdns"
)

func NewRegistry(opts ...registry.Option) registry.Registry {
	return mdns.NewRegistry(opts...)
}
