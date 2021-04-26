package conversation

import (
	"atlas-ncs/conversation/script"
	"errors"
	"sync"
)

type Registry struct {
	registry map[uint32]script.Context
	mutex    sync.RWMutex
}

var once sync.Once
var registry *Registry

func GetRegistry() *Registry {
	once.Do(func() {
		registry = initRegistry()
	})
	return registry
}

func initRegistry() *Registry {
	s := &Registry{make(map[uint32]script.Context), sync.RWMutex{}}
	return s
}

func (s *Registry) GetPreviousContext(characterId uint32) (*script.Context, error) {
	s.mutex.RLock()
	if val, ok := s.registry[characterId]; ok {
		s.mutex.RUnlock()
		return &val, nil
	}
	s.mutex.RUnlock()
	return nil, errors.New("unable to previous context")
}

func (s *Registry) SetContext(characterId uint32, c script.Context) {
	s.mutex.Lock()
	s.registry[characterId] = c
	s.mutex.Unlock()
}
