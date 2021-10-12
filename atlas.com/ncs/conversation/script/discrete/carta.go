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

// Carta is located in Aqua Road - Carta's Cave (230040001)
type Carta struct {
}

func (r Carta) NPCId() uint32 {
	return npc.Carta
}

func (r Carta) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 6301) {
		return r.DoNotFoolAround(l, span, c)
	}
	if !character.HasItem(l, span)(c.CharacterId, item.MiniaturePianus) {
		return r.MustPossessItem(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r Carta) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.MiniaturePianus, -1)
	return script.WarpById(_map.WarpedDimension, 0)(l, span, c)
}

func (r Carta) MustPossessItem(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("In order to open the crack of dimension you will have to possess one piece of Miniature Pianus. Those could be gained by defeating a Pianus.")
	return script.SendOk(l, span, c, m.String())
}

func (r Carta) DoNotFoolAround(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I'm ").
		BlueText().AddText("Carta the sea-witch.").
		BlackText().AddText(" Don't fool around with me, as I'm known for my habit of turning people into worms.")
	return script.SendOk(l, span, c, m.String())
}
