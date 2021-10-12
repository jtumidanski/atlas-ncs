package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Spiruna is located in Orbis - Old Man's House (200050001)
type Spiruna struct {
}

func (r Spiruna) NPCId() uint32 {
	return npc.Spiruna
}

func (r Spiruna) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsCompleted(l)(c.CharacterId, 3034) {
		return r.GoAway(l, span, c)
	}

	return r.RefineDarkCrystal(l, span, c)
}

func (r Spiruna) GoAway(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Go away, I'm trying to meditate.")
	return script.SendOk(l, span, c, m.String())
}

func (r Spiruna) RefineDarkCrystal(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You've been so much of a help to me... If you have any Dark Crystal Ore, I can refine it for you for only ").
		BlueText().AddText(fmt.Sprintf("%d meso", 500000)).
		BlackText().AddText(" each.")
	return script.SendYesNo(l, span, c, m.String(), r.HowMany, script.Exit())
}

func (r Spiruna) HowMany(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Okay, so how many do you want me to make?")
	return script.SendGetNumber(l, span, c, m.String(), r.Validate, 1, 1, 100)
}

func (r Spiruna) Validate(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if !character.HasMeso(l, span)(c.CharacterId, uint32(500000*selection)) {
			return r.NotForFree(l, span, c)
		}

		if !character.HasItems(l, span)(c.CharacterId, item.DarkCrystalOre, uint32(10*selection)) {
			return r.MoreOre(l, span, c)
		}

		if !character.CanHoldAll(l)(c.CharacterId, item.DarkCrystal, uint32(selection)) {
			return r.MoreInventorySpace(l, span, c)
		}

		return r.Process(selection)(l, span, c)
	}
}

func (r Spiruna) NotForFree(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("I'm sorry, but I am NOT doing this for free.")
	return script.SendOk(l, span, c, m.String())
}

func (r Spiruna) MoreOre(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("I need that ore to refine the Crystal. No exceptions..")
	return script.SendOk(l, span, c, m.String())
}

func (r Spiruna) MoreInventorySpace(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Are you having trouble with no empty slots on your inventory? Sort that out first!")
	return script.SendOk(l, span, c, m.String())
}

func (r Spiruna) Process(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		character.GainMeso(l, span)(c.CharacterId, -selection*500000)
		character.GainItem(l, span)(c.CharacterId, item.DarkCrystalOre, -selection*10)
		character.GainItem(l, span)(c.CharacterId, item.DarkCrystal, selection)
		return r.Success(l, span, c)
	}
}

func (r Spiruna) Success(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Use it wisely.")
	return script.SendOk(l, span, c, m.String())
}
