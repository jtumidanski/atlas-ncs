package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/party"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// AssistantBlue2 is located in 
type AssistantBlue2 struct {
}

func (r AssistantBlue2) NPCId() uint32 {
	return npc.AssistantBlue2
}

func (r AssistantBlue2) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	party.WarpById(l)(c.CharacterId, _map.SpiegelmannsOffice2, 4)
	//TODO cancel lobby
	return script.Exit()(l, span, c)
}
