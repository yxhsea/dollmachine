// Package dns provides a dns SRV selector
package dns

import (
	"net"

	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
)

type dnsSelector struct {
	options selector.Options
	domain  string
}

var (
	DefaultDomain = "micro.local"
)

func init() {
	cmd.DefaultSelectors["dns"] = NewSelector
}

func (r *dnsSelector) Init(opts ...selector.Option) error {
	for _, o := range opts {
		o(&r.options)
	}

	if r.options.Context != nil {
		d, ok := r.options.Context.Value(domainKey{}).(string)
		if ok {
			r.domain = d
		}
	}

	return nil
}

func (r *dnsSelector) Options() selector.Options {
	return r.options
}

func (r *dnsSelector) Select(service string, opts ...selector.SelectOption) (selector.Next, error) {
	_, srv, err := net.LookupSRV(service, "tcp", r.domain)
	if err != nil {
		return nil, err
	}

	var nodes []*registry.Node
	for _, node := range srv {
		nodes = append(nodes, &registry.Node{
			Id:      node.Target,
			Address: node.Target,
			Port:    int(node.Port),
		})
	}

	services := []*registry.Service{
		&registry.Service{
			Name:  service,
			Nodes: nodes,
		},
	}

	sopts := selector.SelectOptions{
		Strategy: r.options.Strategy,
	}

	for _, opt := range opts {
		opt(&sopts)
	}

	// apply the filters
	for _, filter := range sopts.Filters {
		services = filter(services)
	}

	// if there's nothing left, return
	if len(services) == 0 {
		return nil, selector.ErrNoneAvailable
	}

	return sopts.Strategy(services), nil
}

func (r *dnsSelector) Mark(service string, node *registry.Node, err error) {
	return
}

func (r *dnsSelector) Reset(service string) {
	return
}

func (r *dnsSelector) Close() error {
	return nil
}

func (r *dnsSelector) String() string {
	return "dns"
}

func NewSelector(opts ...selector.Option) selector.Selector {
	options := selector.Options{
		Strategy: selector.Random,
	}

	for _, o := range opts {
		o(&options)
	}

	domain := DefaultDomain

	if options.Context != nil {
		d, ok := options.Context.Value(domainKey{}).(string)
		if ok {
			domain = d
		}
	}

	return &dnsSelector{options: options, domain: domain}
}
