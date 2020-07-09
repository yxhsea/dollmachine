package nats

import (
	"context"

	"github.com/micro/go-micro/transport"
	"github.com/nats-io/nats"
)

type optionsKey struct{}

// Options allow to inject a nats.Options struct for configuring
// the nats connection
func Options(nopts nats.Options) transport.Option {
	return func(o *transport.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, optionsKey{}, nopts)
	}
}
