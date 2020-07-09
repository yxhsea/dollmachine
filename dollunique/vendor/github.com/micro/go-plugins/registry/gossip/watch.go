package gossip

import (
	"errors"

	"github.com/micro/go-micro/registry"
)

type gossipWatcher struct {
	wo   registry.WatchOptions
	next chan *registry.Result
	stop chan bool
}

func newGossipWatcher(ch chan *registry.Result, exit chan bool, opts ...registry.WatchOption) (registry.Watcher, error) {
	var wo registry.WatchOptions
	for _, o := range opts {
		o(&wo)
	}

	stop := make(chan bool)

	go func() {
		<-stop
		close(exit)
	}()

	return &gossipWatcher{
		wo:   wo,
		next: ch,
		stop: stop,
	}, nil
}

func (m *gossipWatcher) Next() (*registry.Result, error) {
	for {
		select {
		case r, ok := <-m.next:
			if !ok {
				return nil, errors.New("result chan closed")
			}
			// check watch options
			if len(m.wo.Service) > 0 && r.Service.Name != m.wo.Service {
				continue
			}
			return r, nil
		case <-m.stop:
			return nil, errors.New("watcher stopped")
		}
	}
}

func (m *gossipWatcher) Stop() {
	select {
	case <-m.stop:
		return
	default:
		close(m.stop)
	}
}
