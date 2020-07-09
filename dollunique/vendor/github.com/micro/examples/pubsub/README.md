# PubSub

This is an example of pubsub via the client/server interfaces.

PubSub at the client/server level works much like RPC but for async comms. It uses the same encoding but 
rather than using the transport interface it uses the broker for messaging. This includes the ability 
to encode metadata into context which is passed through with messages.

## Contents

- srv - contains the service
- cli - contains the client

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
