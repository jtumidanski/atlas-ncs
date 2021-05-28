package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// BowmanJobInstructorExit is located in Hidden Street - Ant Tunnel For Bowman (108000100)
type BowmanJobInstructorExit struct {
}

func (r BowmanJobInstructorExit) NPCId() uint32 {
	return npc.BowmanJobInstructorExit
}

func (r BowmanJobInstructorExit) Initial(l logrus.FieldLogger, c Context) State {
	if !character.HasItems(l)(c.CharacterId, item.DarkMarble, 30) {
		return r.CollectMarbles(l, c)
	}
	return r.Passed(l, c)
}

func (r BowmanJobInstructorExit) CollectMarbles(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You will have to collect me ").
		BlueText().AddText("30 ").ShowItemName1(item.DarkMarble).
		BlackText().AddText(". Good luck.").NewLine().
		OpenItem(0).BlueText().AddText("I would like to leave").CloseItem()
	return SendListSelection(l, c, m.String(), r.ExitSelection)
}

func (r BowmanJobInstructorExit) ExitSelection(_ int32) StateProducer {
	return r.WarpExit
}

func (r BowmanJobInstructorExit) WarpExit(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.TheRoadToTheDungeon, 9)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.TheRoadToTheDungeon, c.NPCId)
	}
	return Exit()(l, c)
}

func (r BowmanJobInstructorExit) Passed(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Ohhhhh.. you collected all 30 Dark Marbles!! It should have been difficult... just incredible! Alright. You've passed the test and for that, I'll reward you ").
		BlueText().AddText("The Proof of a Hero").
		BlackText().AddText(". Take that and go back to Henesys.")
	return SendNext(l, c, m.String(), r.Reward)
}

func (r BowmanJobInstructorExit) Reward(l logrus.FieldLogger, c Context) State {
	character.RemoveAll(l)(c.CharacterId, item.DarkMarble)
	character.CompleteQuest(l)(c.CharacterId, 100001)
	character.StartQuest(l)(c.CharacterId, 100002)
	character.GainItem(l)(c.CharacterId, item.ProofOfHero, 1)
	return r.WarpExit(l, c)
}