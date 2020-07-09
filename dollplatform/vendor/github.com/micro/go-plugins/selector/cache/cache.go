// Package cache is a caching selector.
package cache

import (
	"github.com/micro/go-micro/selector"
	"github.com/micro/go-micro/selector/cache"
)

/*
	Cache selector is a client side load balancer for go-micro.
	This selector uses the registry Watcher to cache Selected services.
	It uses random hashed load balancing to balance requests across services.
	Implementation here https://godoc.org/github.com/micro/go-micro/selector/cache
	We add a link here for completeness
*/

func NewSelector(opts ...selector.Option) selector.Selector {
	return cache.NewSelector(opts...)
}
