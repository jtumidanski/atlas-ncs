package script

import (
	"errors"
	"sync"
)

type Registry struct {
	registry map[uint32]Script
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
	s := &Registry{make(map[uint32]Script)}
	s.addConversation(AthenaPierce{})
	s.addConversation(DancesWithBalrog{})
	s.addConversation(DarkLord{})
	s.addConversation(GrendelTheReallyOld{})
	s.addConversation(Kyrin{})
	s.addConversation(Rain{})
	s.addConversation(Robin{})
	s.addConversation(Shanks{})
	s.addConversation(Phil{})
	return s
}

func (s *Registry) GetScript(npcId uint32) (*Script, error) {
	if val, ok := s.registry[npcId]; ok {
		return &val, nil
	}
	return nil, errors.New("unable to locate script")
}

func (s *Registry) addConversation(handler Script) {
	s.registry[handler.NPCId()] = handler
}
