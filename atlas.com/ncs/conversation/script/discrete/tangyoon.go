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

// Tangyoon is located in The Nautilus - Cafeteria (120000103)
type Tangyoon struct {
}

func (r Tangyoon) NPCId() uint32 {
	return npc.Tangyoon
}

func (r Tangyoon) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 2180) {
		return script.Exit()(l, span, c)
	}

	return r.SendToStable(l, span, c)
}

func (r Tangyoon) SendToStable(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Okay, I'll now send you to the stable where my cows are. Watch out for the calves that drink all the milk. You don't want your effort to go to waste.")
	return script.SendNext(l, span, c, m.String(), r.IGetConfused)
}

func (r Tangyoon) IGetConfused(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("It won't be easy to tell at a glance between a calf and a cow. Those calves may only be a month or two old, but they have already grown to the size of their mother. They even look alike...even I get confused at times! Good luck!")
	return script.SendNextPrevious(l, span, c, m.String(), r.Validate, r.SendToStable)
}

func (r Tangyoon) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.CanHold(l)(c.CharacterId, item.MilkJug) {
		return r.FullInventory(l, span, c)
	}
	return r.Award(l, span, c)
}

func (r Tangyoon) FullInventory(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I can't give you the empty bottle because your inventory is full. Please make some room in your Etc window.")
	return script.SendOk(l, span, c, m.String())
}

func (r Tangyoon) Award(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.MilkJug, 1)
	return script.WarpById(_map.Stable, 0)(l, span, c)
}
