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
	s.addConversation(APileOfFlowers{})
	s.addConversation(APileOfHerbs{})
	s.addConversation(ArwenTheFairy{})
	s.addConversation(AthenaPierce{})
	s.addConversation(AthenaPierceDemo{})
	s.addConversation(BigHeadward{})
	s.addConversation(Brittany{})
	s.addConversation(Casey{})
	s.addConversation(Chef{})
	s.addConversation(Cherry{})
	s.addConversation(Cloy{})
	s.addConversation(CornerOfTheMagicLibrary{})
	s.addConversation(CornerOfTheMagicLibrary2{})
	s.addConversation(DancesWithBalrog{})
	s.addConversation(DancesWithBalrogDemo{})
	s.addConversation(DarkLordDemo{})
	s.addConversation(GrendelTheReallyOld{})
	s.addConversation(GrendelTheReallyOldDemo{})
	s.addConversation(Heena{})
	s.addConversation(HenesysForest{})
	s.addConversation(HenesysForest2{})
	s.addConversation(Jane{})
	s.addConversation(Joel{})
	s.addConversation(KyrinDemo{})
	s.addConversation(Louis{})
	s.addConversation(Luke{})
	s.addConversation(MarTheFairy{})
	s.addConversation(MrGoldstein{})
	s.addConversation(MrSmith{})
	s.addConversation(MrThunder{})
	s.addConversation(MsTan{})
	s.addConversation(Natalie{})
	s.addConversation(Pason{})
	s.addConversation(Phil{})
	s.addConversation(PowerBForeBowman{})
	s.addConversation(PowerBForeEllinia{})
	s.addConversation(PowerBForeHenesys{})
	s.addConversation(PowerBForeMagician{})
	s.addConversation(PowerBForePerion{})
	s.addConversation(PowerBForeWarrior{})
	s.addConversation(Purin{})
	s.addConversation(Rain{})
	s.addConversation(RegularCabEllinia{})
	s.addConversation(RegularCabHenesys{})
	s.addConversation(RegularCabKerningCity{})
	s.addConversation(RegularCabLithHarbor{})
	s.addConversation(RegularCabPerion{})
	s.addConversation(Robin{})
	s.addConversation(Rooney{})
	s.addConversation(Sera{})
	s.addConversation(Shane{})
	s.addConversation(Shanks{})
	s.addConversation(SmallTreeStump{})
	s.addConversation(TrainerBartos{})
	s.addConversation(TrainerFrod{})
	s.addConversation(Vicious{})
	s.addConversation(VIPCabEllinia{})
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
