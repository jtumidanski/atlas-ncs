package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Louis is located in Hidden Street - The Forest of Patience (101000100, 101000101, 101000102, 101000103, and 101000104)
type Louis struct {
}

func (r Louis) NPCId() uint32 {
	return npc.Louis
}

func (r Louis) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Return(l, span, c)
}

func (r Louis) Return(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Would you like to return to Ellinia?")
	return script.SendYesNo(l, span, c, m.String(), script.WarpById(_map.Ellinia, 0), script.Exit())
}
