package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Kiriko is located in Hidden Street - The 2nd Drill Hall (108000600)
type Kiriko struct {
}

func (r Kiriko) NPCId() uint32 {
	return npc.Kiriko
}

func (r Kiriko) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Would you like to exit the drill hall?")
	return script.SendYesNo(l, span, c, m.String(), script.WarpById(_map.EntranceToTheDrillHall, 0), script.Exit())
}
