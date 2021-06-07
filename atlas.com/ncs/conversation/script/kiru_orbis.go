package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// KiruOrbis is located in Orbis - Station (200000161)
type KiruOrbis struct {
}

func (r KiruOrbis) NPCId() uint32 {
	return npc.KiruOrbis
}

func (r KiruOrbis) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("This ship will head towards ").
		BlueText().AddText("Ereve").
		BlackText().AddText(", an island where you'll find crimson leaves soaking up the sun, the gentle breeze that glides past the stream, and the Empress of Maple Cygnus. If you're interested in joining the Cygnus Knights, Then you should definitely pay a visit here. Are you interested in visiting Ereve?, The Trip will cost you ").
		BlueText().AddText("1000").
		BlackText().AddText(" Mesos").NewLine().
		OpenItem(0).BlueText().AddText(" Ereve (1000 mesos)").CloseItem()
	return SendListSelectionExit(l, c, m.String(), r.DestinationSelection, r.LetMeKnow)
}

func (r KiruOrbis) LetMeKnow(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("OK. If you ever change your mind, please let me know.")
	return SendOk(l, c, m.String())
}

func (r KiruOrbis) DestinationSelection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Validate
	}
	return nil
}

func (r KiruOrbis) Validate(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 1000) {
		return r.NotEnoughMeso(l, c)
	}
	return r.Process(l, c)
}

func (r KiruOrbis) NotEnoughMeso(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hmm... Are you sure you have ").
		BlueText().AddText("1000").
		BlackText().AddText(" Mesos? Check your Inventory and make sure you have enough. You must pay the fee or I can't let you get on...")
	return SendOk(l, c, m.String())
}

func (r KiruOrbis) Process(l logrus.FieldLogger, c Context) State {
	err := character.GainMeso(l)(c.CharacterId, -1000)
	if err != nil {
		l.WithError(err).Errorf("Unable to process payment by character %d.", c.CharacterId)
	}
	return Warp(_map.EmpressRoadToEreveFromOrbis)(l, c)
}
