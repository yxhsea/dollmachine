# Stream

This is an example of a streaming service and two clients, a streaming rpc client and a client using websockets.

## Contents

- server - is the service
- client - is the rpc client
- web - is the websocket client

## Prereqs

Micro services need a discovery system so they can find each other. Micro uses consul by default but
its easily swapped out with etcd, kubernetes, or various other systems. We'll run consul for convenience.

Install consul
```shell
brew install consul
```

Alternative instructions - [https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

Run Consul

```shell
consul agent -dev
```

## Run the example

Run the service

```shell
go run server/main.go
```

Run the client

```shell
go run client/main.go
```

Run the micro web reverse proxy for the websocket client

``` shell
micro web
```

Run the websocket client

```shell
cd web # must be in the web directory to serve static files.
go run main.go
```

Visit http://localhost:8082/stream and send a request!

And that's all there is to it.
