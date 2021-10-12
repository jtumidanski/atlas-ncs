package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Mel is located in Ludibrium - Ludibrium Ticketing Place (220000100)
type Mel struct {
}

func (r Mel) NPCId() uint32 {
	return npc.Mel
}

func (r Mel) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hello, I'm in charge of selling tickets for the ship ride to Orbis Station of Ossyria. The ride to Orbis takes off every 10 minutes, beginning on the hour, and it'll cost you ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 6000)).
		BlackText().AddText(". Are you sure you want to purchase ").
		BlueText().ShowItemName1(item.TicketToOrbisRegular).
		BlackText().AddText("?")
	return script.SendYesNo(l, span, c, m.String(), r.Validate, r.SomeBusiness)
}

func (r Mel) SomeBusiness(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You must have some business to take care of here, right?")
	return script.SendOk(l, span, c, m.String())
}

func (r Mel) Validate(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 6000) || !character.CanHold(l)(c.CharacterId, item.TicketToOrbisRegular) {
		return r.AreYouSure(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r Mel) AreYouSure(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Are you sure you have ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 6000)).
		BlackText().AddText("? If so, then I urge you to check your etc. inventory, and see if it's full or not.")
	return script.SendOk(l, span, c, m.String())
}

func (r Mel) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainMeso(l, span)(c.CharacterId, -6000)
	character.GainItem(l, span)(c.CharacterId, item.TicketToOrbisRegular, 1)
	return script.Exit()(l, span, c)
}
