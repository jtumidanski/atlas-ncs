package topic

import (
	"errors"
	"sync"
)

type Registry struct {
	topics map[string]*Model
	lock   sync.RWMutex
}

var once sync.Once
var registry *Registry

func GetRegistry() *Registry {
	once.Do(func() {
		registry = &Registry{
			topics: make(map[string]*Model, 0),
			lock:   sync.RWMutex{},
		}
	})
	return registry
}

func (r *Registry) Get(token string) (*Model, error) {
	r.lock.RLock()
	if val, ok := r.topics[token]; ok {
		r.lock.RUnlock()
		return val, nil
	} else {
		r.lock.RUnlock()
		r.lock.Lock()
		if val, ok = r.topics[token]; ok {
			r.lock.Unlock()
			return val, nil
		}
		t, err := RequestTopic(token)
		if err != nil {
			r.lock.Unlock()
			return nil, errors.New("unable to locate topic for token")
		}
		r.topics[token] = t
		r.lock.Unlock()
		return t, nil
	}
}
