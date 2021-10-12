package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/event"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// GrandpaMoonBunny is located in Space Mine   - Back of Space Mine (922241100), Hidden Street - Lunar World (922230000), Space Gaga   - Moon Corner (922240200), Space Gaga   - Rescue Gaga! (922240000), and Space Mine   - Space Mine (922241000)
type GrandpaMoonBunny struct {
}

func (r GrandpaMoonBunny) NPCId() uint32 {
	return npc.GrandpaMoonBunny
}

func (r GrandpaMoonBunny) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if c.MapId == _map.MoonCorner {
		m := message.NewBuilder().AddText("Did you have something to say...? ").NewLine().
			OpenItem(0).BlueText().AddText("I want to rescue Gaga.").CloseItem().NewLine().
			OpenItem(1).BlueText().AddText("I want to go to the Space Mine.").CloseItem()
		return script.SendListSelectionExit(l, span, c, m.String(), r.Selection, r.AShame)
	} else if c.MapId >= 922240000 && c.MapId <= 922240019 {
		m := message.NewBuilder().AddText("Don't worry if you fail. You'll have 3 chances. Do you still want to give up?")
		return script.SendYesNo(l, span, c, m.String(), r.GiveUp, script.Exit())
	} else if c.MapId >= 922240100 && c.MapId <= 922240119 {
		m := message.NewBuilder().AddText("You went through so much trouble to rescue Gaga, but it looks like we're back to square one. Let's go back now.")
		return script.SendNext(l, span, c, m.String(), r.GiveUp)
	}
	return script.Exit()(l, span, c)
}

func (r GrandpaMoonBunny) AShame(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("That's a shame, come back when your ready.")
	return script.SendOk(l, span, c, m.String())
}

func (r GrandpaMoonBunny) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Welcome
	case 1:
		return r.SpaceMine
	}
	return nil
}

func (r GrandpaMoonBunny) Welcome(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Welcome! I heard what happened from Baby Moon Bunny I'm glad you came since I was Planning on requesting some help. Gaga is a friend of mine who has helped me before and often stops by to say hello. Unfortunately, he was kidnapped by aliens.")
	return script.SendNext(l, span, c, m.String(), r.GoRescue)
}

func (r GrandpaMoonBunny) SpaceMine(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("At the Space Mine, you can find special ores called ").
		BlueText().AddText("Krypto Crystals").
		BlackText().AddText(" that contains the mysterious power of space. ").
		BlueText().AddText("Krypto Crystals").
		BlackText().AddText(" are usually emerald in color, but will turn brown if hit with the Spaceship's ").
		BlueText().AddText("Space Beam").
		BlackText().AddText(". Remember, in order to thwart this alien conspiracy, ").
		BlueText().AddText("10 Brown Krypto Crystal's").
		BlackText().AddText(" and ").
		BlueText().AddText("10 Emerald Krypto Crystal's").
		BlackText().AddText(" are needed. But since even ").
		BlueText().AddText("1 Krypto Crystal").
		BlackText().AddText(" can be of help, bring me as many as possible. Oh, and one more thing! The Space Mines are protected by the Space Mateons. They are extremely strong due to the power of the ").
		BlueText().AddText("Krypto Crystals").
		BlackText().AddText(", so don't try to defeat them. Simply concentrate on quickly collecting the crystals.")
	return script.SendYesNoExit(l, span, c, m.String(), r.NotCoded, r.AShame, r.AShame)
}

func (r GrandpaMoonBunny) NotCoded(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	//TODO figure this out
	m := message.NewBuilder().AddText("This is not coded yet.")
	return script.SendOk(l, span, c, m.String())
}

func (r GrandpaMoonBunny) GoRescue(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("If we just leave Gaga with the aliens, something terrible will happen to him! I'll let you borrow a spaceship that the Moon Bunnies use for traveling so that you can rescue Gaga.").
		BlueText().AddText(" Although he might appear a bit indecisive, slow, and immature at times").
		BlackText().AddText(", he's really a nice young man. Do you want to go rescue him now?")
	return script.SendYesNoExit(l, span, c, m.String(), r.StartRescueGaga, r.AShame, r.AShame)
}

func (r GrandpaMoonBunny) StartRescueGaga(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	ok := event.StartEvent(l)(c.CharacterId, "RescueGaga")
	if !ok {
		m := message.NewBuilder().AddText("There is currently someone in this map, come back later.")
		return script.SendOk(l, span, c, m.String())
	}
	return script.Exit()(l, span, c)
}

func (r GrandpaMoonBunny) GiveUp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.MoonCorner, 0)(l, span, c)
}
