# Hello World

This is hello world using micro

## Contents

- main.go - is the main definition of the service, handler and client
- proto - contains the protobuf definition of the API

## Dependencies

Install the following

- [consul](https://www.consul.io/intro/getting-started/install.html)
- [micro](https://github.com/micro/micro)
- [protoc-gen-micro](https://github.com/micro/protoc-gen-micro)

## Run Service

```shell
go run main.go
```

## Query Service

```
micro query greeter Greeter.Hello '{"name": "John"}'
```
