package gobreaker

import (
	"github.com/micro/go-micro/client"
	"github.com/sony/gobreaker"

	"context"
)

type clientWrapper struct {
	cb *gobreaker.CircuitBreaker
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	_, err := c.cb.Execute(func() (interface{}, error) {
		cerr := c.Client.Call(ctx, req, rsp, opts...)
		return nil, cerr
	})
	return err
}

// NewClientWrapper takes a *gobreaker.CircuitBreaker and returns a client Wrapper.
func NewClientWrapper(cb *gobreaker.CircuitBreaker) client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{cb, c}
	}
}
