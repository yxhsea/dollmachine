# Sidecar

The sidecar provides a HTTP api, uses dynamic routes, service discovery and load balancing. This and the micro api 
are a powerful alternative to the grpc-gateway.

## Deps

Service discovery is required for all services. Default is Consul or MDNS. You can also use plugins from 
[micro/plugins](https://github.com/micro/go-plugins).

### MDNS

Use the flag `--registry=mdns`

### Consul

```
brew install consul
consul agent -dev
```

## Usage

Check out the micro toolkit

```
git clone https://github.com/micro/micro
cd github.com/micro/micro
```

Create a file `plugin.go` for the grpc plugin

plugin.go

```
package main

import (
	_ "github.com/micro/go-plugins/client/grpc"
)
```

Build the toolkit with the grpc plugin

```
go build -i -o micro ./main.go ./plugin.go
```

Run the sidecar

```
micro --client=grpc sidecar
```

Run the greeter service

```
go run ../greeter/srv/main.go
```

Curl the service

```
curl -H 'Content-Type: application/json' -d '{"name": "john"}' http://localhost:8081/greeter/say/hello
```
