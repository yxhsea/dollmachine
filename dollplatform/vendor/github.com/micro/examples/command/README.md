# Command

This is an example of a bot command as a microservice

## Contents

- main.go - is the main definition of the service and handler

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
go run main.go
```

## Call the service

Run the [bot](https://micro.mu/docs/bot.html) and send message `command`
