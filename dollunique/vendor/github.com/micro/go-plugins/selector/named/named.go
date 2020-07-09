// Package named provides a selector which returns the service name as the address
package named

import (
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
)

type namedSelector struct{}

func init() {
	cmd.DefaultSelectors["named"] = NewSelector
}

func (r *namedSelector) Init(opts ...selector.Option) error {
	return nil
}

func (r *namedSelector) Options() selector.Options {
	return selector.Options{}
}

func (r *namedSelector) Select(service string, opts ...selector.SelectOption) (selector.Next, error) {
	node := &registry.Node{
		Id:      service,
		Address: service,
	}

	return func() (*registry.Node, error) {
		return node, nil
	}, nil
}

func (r *namedSelector) Mark(service string, node *registry.Node, err error) {
	return
}

func (r *namedSelector) Reset(service string) {
	return
}

func (r *namedSelector) Close() error {
	return nil
}

func (r *namedSelector) String() string {
	return "named"
}

func NewSelector(opts ...selector.Option) selector.Selector {
	return &namedSelector{}
}
