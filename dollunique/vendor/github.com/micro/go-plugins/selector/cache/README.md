# Cache Selector

The cache selector is a client side load balancer for go-micro. This selector uses the registry Watcher to cache selected services. 
It defaults random hashed strategy for load balancing requests across services. 

The implementation is at [github.com/micro/go-micro/selector/cache](https://godoc.org/github.com/micro/go-micro/selector/cache) but 
we add a link here for completeness.
