package dns

import (
	"context"

	"github.com/micro/go-micro/selector"
)

type domainKey struct{}

// Domain sets the dns domain for a service
func Domain(d string) selector.Option {
	return func(o *selector.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, domainKey{}, d)
	}
}
