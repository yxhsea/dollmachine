// Package consul provides a Consul registry. Implementation https://godoc.org/github.com/micro/go-micro/registry/consul
package consul

import (
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

func NewRegistry(opts ...registry.Option) registry.Registry {
	return consul.NewRegistry(opts...)
}
