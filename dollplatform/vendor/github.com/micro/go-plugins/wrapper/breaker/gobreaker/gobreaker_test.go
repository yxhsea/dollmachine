package gobreaker

import (
	"testing"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/mock"
	"github.com/micro/go-micro/selector"
	"github.com/sony/gobreaker"

	"context"
)

func TestBreaker(t *testing.T) {
	// setup
	r := mock.NewRegistry()
	s := selector.NewSelector(selector.Registry(r))

	c := client.NewClient(
		// set the selector
		client.Selector(s),
		// add the breaker wrapper
		client.Wrap(NewClientWrapper(
			gobreaker.NewCircuitBreaker(gobreaker.Settings{}),
		)),
	)

	req := c.NewRequest("test.service", "Test.Method", map[string]string{
		"foo": "bar",
	}, client.WithContentType("application/json"))

	var rsp map[string]interface{}

	// Force to point of trip
	for i := 0; i < 6; i++ {
		c.Call(context.TODO(), req, rsp)
	}

	err := c.Call(context.TODO(), req, rsp)
	if err == nil {
		t.Error("Expecting tripped breaker, got nil error")
	}

	if err.Error() != "circuit breaker is open" {
		t.Errorf("Expecting tripped breaker, got %v", err)
	}
}
