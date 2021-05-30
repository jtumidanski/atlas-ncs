package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// PuroRien is located in Snow Island - Dangerous Forest (140020300)
type PuroRien struct {
}

func (r PuroRien) NPCId() uint32 {
	return npc.PuroRien
}

func (r PuroRien) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r PuroRien) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Are you trying to leave Rien? Board this ship and I'll take you from ").
		BlueText().AddText("Rien").
		BlackText().AddText(" to ").
		BlueText().AddText("Lith Harbor").
		BlackText().AddText(" and back. for a ").
		BlueText().AddText("fee of 800").
		BlackText().AddText(" Mesos. Would you like to head over to Lith Harbor now? It'll take about a minute to get there.").NewLine().
		OpenItem(0).BlueText().AddText(" Lith Harbor (800 mesos)").CloseItem()
	return SendListSelectionExit(l, c, m.String(), r.Selection, r.LetMeKnow)
}

func (r PuroRien) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Validate
	}
	return nil
}

func (r PuroRien) LetMeKnow(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("OK. If you ever change your mind, please let me know.")
	return SendOk(l, c, m.String())
}

func (r PuroRien) Validate(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 800) {
		return r.NotEnoughMeso(l, c)
	}
	return r.Process(l, c)
}

func (r PuroRien) NotEnoughMeso(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmm... Are you sure you have ").
		BlueText().AddText("800").
		BlackText().AddText(" Mesos? Check your Inventory and make sure you have enough. You must pay the fee or I can't let you get on...")
	return SendOk(l, c, m.String())
}

func (r PuroRien) Process(l logrus.FieldLogger, c Context) State {
	err := character.GainMeso(l)(c.CharacterId, -800)
	if err != nil {
		l.WithError(err).Errorf("Unable to complete meso transaction with %d.", c.CharacterId)
		return nil
	}

	err = npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ToLithHarbor, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to map %d. Refunding mesos.", c.CharacterId, _map.ToLithHarbor)
		err = character.GainMeso(l)(c.CharacterId, 800)
		if err != nil {
			l.WithError(err).Errorf("Error processing refund, %d has lost %d mesos.", c.CharacterId, 800)
		}
	}
	return nil
}