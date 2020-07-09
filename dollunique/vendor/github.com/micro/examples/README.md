# Micro Examples  [![Travis CI](https://travis-ci.org/micro/examples.svg?branch=master)](https://travis-ci.org/micro/examples) [![Go Report Card](https://goreportcard.com/badge/micro/examples)](https://goreportcard.com/report/github.com/micro/examples)

This is a repository for micro examples. Feel free to contribute.

## Contents

- [api](api) - Provides API usage examples
- [booking](booking) - A booking.com demo application
- [broker](broker) - A example of using Broker for Publish and Subscribing.
- [client](client) - Usage of the Client package to call a service.
- [command](command) - An example of bot commands as micro services
- [event](event) - Using the API Gateway event handler
- [flags](flags) - Using command line flags with a service
- [form](form) - How to parse a form behind the micro api
- [function](function) - Example of using Function programming model
- [graceful](graceful) - Demonstrates graceful shutdown of a service
- [greeter](greeter) - A complete greeter example (includes python, ruby examples)
- [heartbeat](heartbeat) - Make services heartbeat with discovery for high availability
- [helloworld](helloworld) - Hello world using micro
- [metadata](metadata) - Extracting metadata from context of a request
- [mocking](mocking) - Demonstrate mocking helloworld service
- [noproto](noproto) - Use micro without protobuf or code generation, only go types
- [options](options) - Setting options in the go-micro framework
- [plugins](plugins) - How to use plugins
- [pubsub](pubsub) - Example of using pubsub at the client/server level
- [grpc](grpc) - Examples of how to use [go-grpc](https://github.com/micro/go-grpc)
- [redirect](redirect) - An example of how to http redirect using an API service
- [roundrobin](roundrobin) - A stateful client wrapper for true round robin of requests
- [secure](secure) - Demonstrates use of transport secure option for self signed certs
- [server](server) - Use of the Server package directly to server requests.
- [service](service) - Example of the top level Service in go-micro.
- [sharding](sharding) - An example of how to shard requests or use session affinity
- [shutdown](shutdown) - Demonstrates graceful shutdown via context cancellation
- [stream](stream) - An example of a streaming service and client
- [template](template) - Api, web and srv service templates generated with the 'micro new' command
- [waitgroup](waitgroup) - Demonstrates how to use a waitgroup with a service
- [wrapper](wrapper) - A simple example of using a log wrapper

## Community

Find contributions from the community via the [explorer](https://micro.mu/explore/)

- [go-shopping](https://github.com/autodidaddict/go-shopping) - A sample product with a suite of services
- [shippy](https://github.com/EwanValentine/shippy) - A multi app demo and tutorial

## Dependencies

- [Service Discovery](#service-discovery)
- [Protobuf](#protobuf)

## Service Discovery

All services require service discovery. The default is Consul or MDNS.

### Consul

Install
```
brew install consul
```

Run
```
consul agent -dev
```

### Multicast DNS

Use flag `--registry=mdns` for a zero dependency configuration

## Protobuf

Protobuf is used for code generation of message types and client/hander stubs.

If making changes recompile the protos.

### Install

Install [protoc](https://github.com/google/protobuf) for your environment. Then:

```shell
go get github.com/golang/protobuf/{proto,protoc-gen-go}
```

```shell
go get github.com/micro/protoc-gen-micro
```

### Compile Proto

```shell
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. path/to/proto
```

