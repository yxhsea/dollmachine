# Function

This is an example of creating a micro function. A function is a one time executing service.

## Contents

- main.go - is the main definition of the function
- proto - contains the protobuf definition of the API

## Dependencies

Service discovery is required to resolve names to addresses

Install consul
```shell

brew install consul
```

Run Consul

```shell
consul agent -dev
```

## Install micro

```
go get github.com/micro/micro
```

## Run function

```shell
micro run -r github.com/micro/examples/function
```

## Call function

```shell
micro query greeter Greeter.Hello '{"name": "john"}'
```
