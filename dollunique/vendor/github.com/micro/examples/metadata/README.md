# Metadata

This is an example of sending metadata/headers.

HTTP headers sent to the micro api will be converted to metadata and forwarded on.

## Contents

- **srv** - an RPC service which extracts metadata
- **cli** - an RPC client that calls the service once

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

## Run Service

Start go.micro.srv.greeter
```shell
go run srv/main.go
```

## Client

Call go.micro.srv.greeter via client
```shell
go run cli/main.go
```

