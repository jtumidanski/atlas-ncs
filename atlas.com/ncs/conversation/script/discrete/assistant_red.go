package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/party"
	"github.com/sirupsen/logrus"
)

type AssistantRed struct {
}

func (r AssistantRed) NPCId() uint32 {
	return npc.AssistantRed
}

func (r AssistantRed) Initial(l logrus.FieldLogger, c script.Context) script.State {
	party.Warp(l)(c.CharacterId, _map.SpiegelmannsOffice)
	//TODO cancel lobby
	return script.Exit()(l, c)
}
