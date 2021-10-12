package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// KiruEreve is located in Empress' Road - Sky Ferry (130000210)
type KiruEreve struct {
}

func (r KiruEreve) NPCId() uint32 {
	return npc.KiruEreve
}

func (r KiruEreve) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r KiruEreve) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmm... The winds are favorable. Are you thinking of leaving ereve and going somewhere else? This ferry sails to Orbis on the Ossyria Continent, Have you taking care of everything you needed to in Ereve? If you happen to be headed toward ").
		BlueText().AddText("Orbis").
		BlackText().AddText(" i can take you there. What do you day? Are you going to go to Orbis?").NewLine().
		OpenItem(0).BlueText().AddText("Orbis (1000 mesos)").CloseItem().NewLine()
	return script.SendListSelectionExit(l, span, c, m.String(), r.ChooseDestination, r.OhWell)
}

func (r KiruEreve) ChooseDestination(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.VictoriaIsland
	}
	return nil
}

func (r KiruEreve) VictoriaIsland(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 1000) {
		return r.NotEnoughMeso(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r KiruEreve) OhWell(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("If you're not interested, then oh well...")
	return script.SendOk(l, span, c, m.String())
}

func (r KiruEreve) NotEnoughMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmm... Are you sure you have ").
		BlueText().AddText("1000").
		BlackText().AddText(" Mesos? Check your Inventory and make sure you have enough. You must pay the fee or I can't let you get on...")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r KiruEreve) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainMeso(l, span)(c.CharacterId, -1000)
	return script.Warp(_map.EmpressRoadToOrbis)(l, span, c)
}
