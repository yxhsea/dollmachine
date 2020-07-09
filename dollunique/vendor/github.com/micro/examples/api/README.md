# API

This repo contains examples for serving microservices via the micro api.

The [micro api](https://github.com/micro/micro/tree/master/api) is an API gateway which serves HTTP and routes to RPC based services. 
In the micro ecosystem we logically separate concerns via architecture and tooling. Read more on buiding an API layer of services 
in the [architecture blog post](https://micro.mu/blog/2016/04/18/micro-architecture.html).

The micro api by default serves the namespace go.micro.api. Our service names include this plus a unique name e.g go.micro.api.example. 
You can change the namespace via the flag `--namespace=`.

The micro api has a number of different handlers which lets you define what kind of API services you want. See examples below. The handler 
can be set via the flag `--handler=`.

## Contents

- default - an api using the default micro api request handler
- proxy - an api using the http proxy handler
- rpc - an api using standard go-micro rpc
- meta - an api which specifies its api endpoints

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

## Request Mapping

### API + RPC

Micro maps http paths to rpc services. The mapping table can be seen below.

The default namespace for the api is **go.micro.api** but you can set your own namespace via `--namespace`.

URLs are mapped as follows:

Path	|	Service	|	Method
----	|	----	|	----
/foo/bar	|	go.micro.api.foo	|	Foo.Bar
/foo/bar/baz	|	go.micro.api.foo	|	Bar.Baz
/foo/bar/baz/cat	|	go.micro.api.foo.bar	|	Baz.Cat

Versioned API URLs can easily be mapped to service names:

Path	|	Service	|	Method
----	|	----	|	----
/foo/bar	|	go.micro.api.foo	|	Foo.Bar
/v1/foo/bar	|	go.micro.api.v1.foo	|	Foo.Bar
/v1/foo/bar/baz	|	go.micro.api.v1.foo	|	Bar.Baz
/v2/foo/bar	|	go.micro.api.v2.foo	|	Foo.Bar
/v2/foo/bar/baz	|	go.micro.api.v2.foo	|	Bar.Baz

### Proxy Mapping

Starting the API with `--handler=proxy` will reverse proxy requests to backend services within the served API namespace (default: go.micro.api). 

Example

Path	|	Service	|	Service Path
---	|	---	|	---
/greeter	|	go.micro.api.greeter	|	/greeter
/greeter/:name	|	go.micro.api.greeter	|	/greeter/:name
