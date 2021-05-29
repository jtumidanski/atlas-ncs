package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// KiriruEreve is located in Empress' Road - Sky Ferry (130000210)
type KiriruEreve struct {
}

func (r KiriruEreve) NPCId() uint32 {
	return npc.KiriruEreve
}

func (r KiriruEreve) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r KiriruEreve) Hello(l logrus.FieldLogger, c Context) State {
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
	return SendListSelectionExit(l, c, m.String(), r.ChooseDestination, r.OhWell)
}

func (r KiriruEreve) ChooseDestination(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.VictoriaIsland
	}
	return nil
}

func (r KiriruEreve) VictoriaIsland(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 1000) {
		return r.NotEnoughMeso(l, c)
	}
	return r.Process(l, c)
}

func (r KiriruEreve) OhWell(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("If you're not interested, then oh well...")
	return SendOk(l, c, m.String())
}

func (r KiriruEreve) NotEnoughMeso(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmm... Are you sure you have ").
		BlueText().AddText("1000").
		BlackText().AddText(" Mesos? Check your Inventory and make sure you have enough. You must pay the fee or I can't let you get on...")
	return SendNext(l, c, m.String(), Exit())
}

func (r KiriruEreve) Process(l logrus.FieldLogger, c Context) State {
	err := character.GainMeso(l)(c.CharacterId, -1000)
	if err != nil {
		l.WithError(err).Errorf("Unable to process payment by character %d.", c.CharacterId)
	}
	err = npc.Warp(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.EmpressRoadToEllinia)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.EmpressRoadToEllinia, c.NPCId)
	}
	return Exit()(l, c)
}
