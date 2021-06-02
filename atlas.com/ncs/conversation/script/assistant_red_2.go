package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/party"
	"github.com/sirupsen/logrus"
)

// AssistantRed2 is located in 
type AssistantRed2 struct {
}

func (r AssistantRed2) NPCId() uint32 {
	return npc.AssistantRed2
}

func (r AssistantRed2) Initial(l logrus.FieldLogger, c Context) State {
	party.WarpById(l)(c.CharacterId, _map.SpiegelmannsOffice2, 4)
	//TODO cancel lobby
	return Exit()(l, c)
}
