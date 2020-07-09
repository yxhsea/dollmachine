# Round Robin

An example of using a round robin client wrapper with the greeter application. 

## Contents

- api.go - a modified version of the greeter api to include roundrobin

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

### Micro

```
go get github.com/micro/micro
```

## Run 

### Greeter Service

Run multiple copies of the greeter

```
cd ../greeter
go run srv/main.go
```

### Greeter API

```
go run api.go
```

### Micro API

```
micro api
```

### Call API

```shell
curl  http://localhost:8080/greeter/say/hello?name=John
```

