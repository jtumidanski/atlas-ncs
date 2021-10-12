package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// BabyMilkCow1 is located in Hidden Chamber - The Nautilus - Stable (912000100)
type BabyMilkCow1 struct {
}

func (r BabyMilkCow1) NPCId() uint32 {
	return npc.BabyMilkCow1
}

func (r BabyMilkCow1) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.MilkJug) {
		return r.RemainsEmpty(l, span, c)
	}
	if character.HasItem(l, span)(c.CharacterId, item.MilkJugOneThird) {
		return r.DrinkBottle(item.MilkJugOneThird)(l, span, c)
	}
	if character.HasItem(l, span)(c.CharacterId, item.MilkJugTwoThird) {
		return r.DrinkBottle(item.MilkJugTwoThird)(l, span, c)
	}
	return r.DrinkBottle(item.MilkJugFull)(l, span, c)
}

func (r BabyMilkCow1) RemainsEmpty(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The hungry calf is drinking all the milk! The bottle remains empty...")
	return script.SendNext(l, span, c, m.String(), r.EmptyBottle)
}

func (r BabyMilkCow1) EmptyBottle(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The hungry calf isn't interested in the empty bottle.")
	return script.SendNextPrevious(l, span, c, m.String(), script.Exit(), r.RemainsEmpty)
}

func (r BabyMilkCow1) DrinkBottle(itemId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GainItem(l, span)(c.CharacterId, itemId, -1)
		character.GainItem(l, span)(c.CharacterId, item.MilkJug, 1)
		m := message.NewBuilder().
			AddText("The hungry calf is drinking all the milk! The bottle is now empty.")
		return script.SendNext(l, span, c, m.String(), script.Exit())
	}
}
