package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Parwen is located in Hidden Street - Authorized Personnel Only (261020401)
type Parwen struct {
}

func (r Parwen) NPCId() uint32 {
	return npc.Parwen
}

func (r Parwen) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3320) || quest.IsCompleted(l)(c.CharacterId, 3320) {
		return script.WarpById(_map.DransLab, 1)(l, span, c)
	}
	m := message.NewBuilder().AddText("uuuuhuk...Why only Ghost are around here?...")
	return script.SendOk(l, span, c, m.String())
}
