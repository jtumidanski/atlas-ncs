package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// KiruEreve is located in Empress' Road - Sky Ferry (130000210)
type KiruEreve struct {
}

func (r KiruEreve) NPCId() uint32 {
	return npc.KiruEreve
}

func (r KiruEreve) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r KiruEreve) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmm... The winds are favorable. Are you thinking of leaving ereve and going somewhere else? This ferry sails to Orbis on the Ossyria Continent, Have you taking care of everything you needed to in Ereve? If you happen to be headed toward ").
		BlueText().AddText("Orbis").
		BlackText().AddText(" i can take you there. What do you day? Are you going to go to Orbis?").NewLine().
		OpenItem(0).BlueText().AddText("Orbis (1000 mesos)").CloseItem().NewLine()
	return SendListSelectionExit(l, c, m.String(), r.ChooseDestination, r.OhWell)
}

func (r KiruEreve) ChooseDestination(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.VictoriaIsland
	}
	return nil
}

func (r KiruEreve) VictoriaIsland(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 1000) {
		return r.NotEnoughMeso(l, c)
	}
	return r.Process(l, c)
}

func (r KiruEreve) OhWell(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("If you're not interested, then oh well...")
	return SendOk(l, c, m.String())
}

func (r KiruEreve) NotEnoughMeso(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmm... Are you sure you have ").
		BlueText().AddText("1000").
		BlackText().AddText(" Mesos? Check your Inventory and make sure you have enough. You must pay the fee or I can't let you get on...")
	return SendNext(l, c, m.String(), Exit())
}

func (r KiruEreve) Process(l logrus.FieldLogger, c Context) State {
	err := character.GainMeso(l)(c.CharacterId, -1000)
	if err != nil {
		l.WithError(err).Errorf("Unable to process payment by character %d.", c.CharacterId)
	}
	return Warp(_map.EmpressRoadToOrbis)(l, c)
}
