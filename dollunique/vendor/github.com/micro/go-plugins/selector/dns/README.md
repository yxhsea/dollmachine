# DNS Selector

The dns selector looks up services via dns SRV records

- SRV Target is used as the service Node Id and Port
- The default domain is `micro.local` e.g foo becomes foo.micro.local

## Usage

```go
selector := dns.NewSelector()

service := micro.NewService(
	micro.Selector(selector),
)
```

Specify lookup domain

```go
dns.NewSelector(
	dns.Domain("example.com"),
)
```
