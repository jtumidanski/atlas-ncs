package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/party"
	"github.com/sirupsen/logrus"
)

// AssistantBlue is located in 
type AssistantBlue struct {
}

func (r AssistantBlue) NPCId() uint32 {
	return npc.AssistantBlue
}

func (r AssistantBlue) Initial(l logrus.FieldLogger, c script.Context) script.State {
	party.Warp(l)(c.CharacterId, _map.SpiegelmannsOffice)
	//TODO cancel lobby
	return script.Exit()(l, c)
}
