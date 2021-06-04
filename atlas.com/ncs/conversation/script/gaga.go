package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type Gaga struct {
}

func (r Gaga) NPCId() uint32 {
	return npc.Gaga
}

func (r Gaga) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hey, traveler! I am ").
		ShowNPC(npc.Gaga).
		AddText(", and my job is to recruit travelers like you, who are eager for new challenges daily. Right now, my team is holding contests that thoroughly tests the mental and physical capabilities of adventurers like you.")
	return SendNext(l, c, m.String(), r.BossFights)
}

func (r Gaga) BossFights(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("These contests involve ").
		BlueText().AddText("sequential boss fights").
		BlackText().AddText(", with some resting spots between some sections. These will require some strategy time and enough supplies at hand, as they are not common fights.")
	return SendNextPrevious(l, c, m.String(), r.Confirm, r.Initial)
}

func (r Gaga) Confirm(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("If you feel you are powerful enough, you can join others like you at where we are hosting the contests of power. ... So, what is your decision? Will you come to where the contests are being held right now?")
	return SendAcceptDecline(l, c, m.String(), r.VeryWell, Exit())
}

func (r Gaga) VeryWell(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Very well. Remember, there you can assemble a team or take on the fighting on your own, it's up to you. Good luck!")
	return SendOkTrigger(l, c, m.String(), r.Warp)
}

func (r Gaga) Warp(l logrus.FieldLogger, c Context) State {
	character.SaveLocation(l)(c.CharacterId, "BOSSPQ")
	err := npc.WarpByName(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ExclusiveTrainingCenter, "out00")
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.ExclusiveTrainingCenter, c.NPCId)
	}
	return Exit()(l, c)
}
