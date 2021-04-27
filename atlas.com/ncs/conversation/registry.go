package conversation

import (
	"atlas-ncs/conversation/script"
	"errors"
	"sync"
)

type Pair struct {
	ctx script.Context
	ns  script.State
}

type Registry struct {
	registry map[uint32]Pair
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
	s := &Registry{make(map[uint32]Pair), sync.RWMutex{}}
	return s
}

func (s *Registry) GetPreviousContext(characterId uint32) (*Pair, error) {
	s.mutex.RLock()
	if val, ok := s.registry[characterId]; ok {
		s.mutex.RUnlock()
		return &val, nil
	}
	s.mutex.RUnlock()
	return nil, errors.New("unable to previous context")
}

func (s *Registry) SetContext(characterId uint32, ctx script.Context, ns script.State) {
	s.mutex.Lock()
	s.registry[characterId] = Pair{ctx, ns}
	s.mutex.Unlock()
}

func (s *Registry) ClearContext(characterId uint32) {
	s.mutex.Lock()
	delete(s.registry, characterId)
	s.mutex.Unlock()
}
