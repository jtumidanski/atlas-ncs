package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Shane is located in Victoria Road - Ellinia (101000000)
type Shane struct {
}

func (r Shane) NPCId() uint32 {
	return npc.Shane
}

func (r Shane) Initial(l logrus.FieldLogger, c Context) State {
	if !character.AboveLevel(l)(c.CharacterId, 25) {
		return r.LevelRequirement(l, c)
	}

	return r.WouldYouLike(l, c)
}

func (r Shane) LevelRequirement(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You must be a higher level to enter the Forest of Patience.")
	return SendOk(l, c, m.String())
}

func (r Shane) WouldYouLike(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hi, i'm Shane. I can let you into the Forest of Patience for a small fee. Would you like to enter for ").
		BlueText().AddText(fmt.Sprintf("5000")).
		BlackText().AddText(" mesos?")
	return SendYesNo(l, c, m.String(), r.ValidatePayment, r.SeeYouNextTime)
}

func (r Shane) SeeYouNextTime(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Alright, see you next time.")
	return SendOk(l, c, m.String())
}

func (r Shane) ValidatePayment(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 5000) {
		return r.NotEnoughMeso(l, c)
	}

	err := character.GainMeso(l)(c.CharacterId, -5000)
	if err != nil {
		l.WithError(err).Errorf("Unable to process payment for character %d.", c.CharacterId)
	}

	var destination uint32
	if character.QuestStarted(l)(c.CharacterId, 2050) {
		destination = _map.TheForestOfPatienceStep1
	} else if character.QuestStarted(l)(c.CharacterId, 2051) {
		destination = _map.TheForestOfPatienceStep3
	} else if character.MeetsCriteria(l)(c.CharacterId, character.LevelBetweenCriteria(24, 50)) {
		destination = _map.TheForestOfPatienceStep1
	} else if character.IsLevel(l)(c.CharacterId, 50) {
		destination = _map.TheForestOfPatienceStep3
	} else {
		l.Warnf("Unsure where to send player, returning money and warping to the same map.")
		err = character.GainMeso(l)(c.CharacterId, 5000)
		if err != nil {
			l.WithError(err).Errorf("Unable to refund payment for character %d.", c.CharacterId)
		}
		destination = c.MapId
	}

	err = npc.Processor(l).WarpById(c.WorldId, c.ChannelId, c.CharacterId, destination, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, destination, c.NPCId)
	}
	return nil
}

func (r Shane) NotEnoughMeso(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Sorry but it doesn't like you have enough mesos!")
	return SendOk(l, c, m.String())
}