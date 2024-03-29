package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// SubwayTrashCan4 is located in Kerning City Subway - Line 1 <Area 1> (103000101)
type SubwayTrashCan4 struct {
}

func (r SubwayTrashCan4) NPCId() uint32 {
	return npc.SubwayTrashCan4
}

func (r SubwayTrashCan4) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Just a trash can sitting there.")
	return script.SendOk(l, span, c, m.String())
}
