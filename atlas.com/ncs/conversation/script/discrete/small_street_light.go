package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// SmallStreetLight is located in Kerning City Subway - Line 1 <Area 4> (103000105)
type SmallStreetLight struct {
}

func (r SmallStreetLight) NPCId() uint32 {
	return npc.SmallStreetLight
}

func (r SmallStreetLight) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("A small focus of light lighting in the immersive darkness.")
	return script.SendOk(l, span, c, m.String())
}
