# Template Service

This is the Template service

Generated with

```
micro new github.com/micro/examples/template/srv --namespace=go.micro --alias=template --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.template
- Type: srv
- Alias: template

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./template-srv
```

Build a docker image
```
make docker
```