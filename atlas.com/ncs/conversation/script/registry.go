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
	s.addConversation(AthenaPierceDemo{})
	s.addConversation(Chef{})
	s.addConversation(Cloy{})
	s.addConversation(DancesWithBalrogDemo{})
	s.addConversation(DarkLordDemo{})
	s.addConversation(GrendelTheReallyOld{})
	s.addConversation(GrendelTheReallyOldDemo{})
	s.addConversation(Heena{})
	s.addConversation(Jane{})
	s.addConversation(KyrinDemo{})
	s.addConversation(MrGoldstein{})
	s.addConversation(Pason{})
	s.addConversation(Rain{})
	s.addConversation(RegularCabEllinia{})
	s.addConversation(RegularCabHenesys{})
	s.addConversation(RegularCabKerningCity{})
	s.addConversation(RegularCabLithHarbor{})
	s.addConversation(RegularCabPerion{})
	s.addConversation(Robin{})
	s.addConversation(Sera{})
	s.addConversation(Shanks{})
	s.addConversation(Phil{})
	s.addConversation(Vicious{})
	s.addConversation(VIPCabLithHarbor{})
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
