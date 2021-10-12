package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// MarkTheToySoldier is located in Hidden Street - Doll's House (922000010)
type MarkTheToySoldier struct {
}

func (r MarkTheToySoldier) NPCId() uint32 {
	return npc.MarkTheToySoldier
}

func (r MarkTheToySoldier) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 3230) {
		return r.ThankYou(l, span, c)
	}
	if !character.HasItem(l, span)(c.CharacterId, item.Pendulum) {
		return r.YouHaveNot(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r MarkTheToySoldier) ThankYou(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Thank you for finding the pendulum. Are you ready to return to Eos Tower?")
	return script.SendYesNo(l, span, c, m.String(), r.Warp, script.Exit())
}

func (r MarkTheToySoldier) YouHaveNot(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You haven't found the pendulum yet. Do you want to go back to Eos Tower?")
	return script.SendYesNo(l, span, c, m.String(), r.Warp, script.Exit())
}

func (r MarkTheToySoldier) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	quest.Complete(l)(c.CharacterId, 3230)
	character.GainItem(l, span)(c.CharacterId, item.Pendulum, -1)
	return r.ThankYou(l, span, c)
}

func (r MarkTheToySoldier) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.EosTower100thFloor, 4)(l, span, c)
}
