# Service

This is an example of creating a micro service.

## Contents

- main.go - is the main definition of the service, handler and client
- proto - contains the protobuf definition of the API
- wrapper - demonstrates the use of Client and Server Wrappers

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
consul agent -dev -advertise=127.0.0.1
```

## Run the example

Run the service

```shell
go run main.go
```

Run the client

```shell
go run main.go --run_client
```

And that's all there is to it.
