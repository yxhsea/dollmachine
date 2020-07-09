// Package noop is a registry which does nothing
package noop

import (
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry"
)

type noopRegistry struct {
	options registry.Options
}

func init() {
	cmd.DefaultRegistries["noop"] = NewRegistry
}

func (m *noopRegistry) Options() registry.Options {
	return m.options
}

func (m *noopRegistry) GetService(service string) ([]*registry.Service, error) {
	return nil, nil
}

func (m *noopRegistry) ListServices() ([]*registry.Service, error) {
	return nil, nil
}

func (m *noopRegistry) Register(s *registry.Service, opts ...registry.RegisterOption) error {
	return nil
}

func (m *noopRegistry) Deregister(s *registry.Service) error {
	return nil
}

func (m *noopRegistry) Watch(opts ...registry.WatchOption) (registry.Watcher, error) {
	return &noopWatcher{exit: make(chan bool)}, nil
}

func (m *noopRegistry) String() string {
	return "noop"
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	var options registry.Options
	for _, o := range opts {
		o(&options)
	}
	return &noopRegistry{options}
}
