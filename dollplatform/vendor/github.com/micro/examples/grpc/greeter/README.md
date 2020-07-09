# Greeter Service

An example Go-Micro based GRPC service

## What's here?

- **srv** - a GRPC greeter service
- **cli** - a GRPC client that calls the service once

Run Service
```
$ go run srv/main.go --registry=mdns
2016/11/03 18:41:22 Listening on [::]:55194
2016/11/03 18:41:22 Broker Listening on [::]:55195
2016/11/03 18:41:22 Registering node: go.micro.srv.greeter-1e200612-a1f5-11e6-8e84-68a86d0d36b6
```

Test Service
```
$ go run cli/main.go --registry=mdns
Hello John
```

