# Named selector

The named selector returns service name as the node address.

Why? This is useful where you want to offload discovery and load balancing to the transport itself e.g message bus.

## Example

When a service uses a message bus such as NATS for transport it will use its service name as the address.
This will force any instance of the service to subscribe to a topic with its own service name.

When a request is made to this service NATS will handle discovery and load balancing, essentially offloading 
this concern from micro to the message bus itself.

## Usage

```go
selector := named.NewSelector()

service := micro.NewService(
	micro.Selector(selector),
)
```
