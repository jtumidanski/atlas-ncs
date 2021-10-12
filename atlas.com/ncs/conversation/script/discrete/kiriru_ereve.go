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

// KiriruEreve is located in Empress' Road - Sky Ferry (130000210)
type KiriruEreve struct {
}

func (r KiriruEreve) NPCId() uint32 {
	return npc.KiriruEreve
}

func (r KiriruEreve) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r KiriruEreve) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Eh, Hello...again. Do you want to leave Ereve and go somewhere else? If so, you've come to the right place. I operate a ferry that goes from ").
		BlueText().AddText("Ereve").
		BlackText().AddText(" to ").
		BlueText().AddText("Victoria Island").
		BlackText().AddText(", I can take you to ").
		BlueText().AddText("Victoria Island").
		BlackText().AddText(" if you want... You'll have to pay a fee of ").
		BlueText().AddText("1000").
		BlackText().AddText(" Mesos.").NewLine().
		OpenItem(0).BlueText().AddText("Victoria Island (1000 mesos)").CloseItem().NewLine()
	return script.SendListSelectionExit(l, span, c, m.String(), r.ChooseDestination, r.OhWell)
}

func (r KiriruEreve) ChooseDestination(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.VictoriaIsland
	}
	return nil
}

func (r KiriruEreve) VictoriaIsland(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasMeso(l, span)(c.CharacterId, 1000) {
		return r.NotEnoughMeso(l, span, c)
	}
	return r.Process(l, span, c)
}

func (r KiriruEreve) OhWell(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("If you're not interested, then oh well...")
	return script.SendOk(l, span, c, m.String())
}

func (r KiriruEreve) NotEnoughMeso(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmm... Are you sure you have ").
		BlueText().AddText("1000").
		BlackText().AddText(" Mesos? Check your Inventory and make sure you have enough. You must pay the fee or I can't let you get on...")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r KiriruEreve) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainMeso(l, span)(c.CharacterId, -1000)
	return script.Warp(_map.EmpressRoadToEllinia)(l, span, c)
}
