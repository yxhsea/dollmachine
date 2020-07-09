package noop

import (
	"errors"

	"github.com/micro/go-micro/registry"
)

type noopWatcher struct {
	exit chan bool
}

func (n *noopWatcher) Next() (*registry.Result, error) {
	select {
	case <-n.exit:
		return nil, errors.New("watcher stopped")
	}
}

func (n *noopWatcher) Stop() {
	select {
	case <-n.exit:
		return
	default:
		close(n.exit)
	}
}
