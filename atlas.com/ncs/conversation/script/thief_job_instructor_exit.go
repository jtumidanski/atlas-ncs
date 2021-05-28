package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// ThiefJobInstructorExit is located in Hidden Street - Thief's Construction Site (108000400)
type ThiefJobInstructorExit struct {
}

func (r ThiefJobInstructorExit) NPCId() uint32 {
	return npc.ThiefJobInstructorExit
}

func (r ThiefJobInstructorExit) Initial(l logrus.FieldLogger, c Context) State {
	if !character.HasItems(l)(c.CharacterId, item.DarkMarble, 30) {
		return r.CollectMarbles(l, c)
	}
	return r.Passed(l, c)
}

func (r ThiefJobInstructorExit) CollectMarbles(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You will have to collect me ").
		BlueText().AddText("30 ").ShowItemName1(item.DarkMarble).
		BlackText().AddText(". Good luck.").NewLine().
		OpenItem(0).BlueText().AddText("I would like to leave").CloseItem()
	return SendListSelection(l, c, m.String(), r.ExitSelection)
}

func (r ThiefJobInstructorExit) ExitSelection(_ int32) StateProducer {
	return r.WarpExit
}

func (r ThiefJobInstructorExit) WarpExit(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ConstructionSiteNorthOfKerningCity, 9)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.ConstructionSiteNorthOfKerningCity, c.NPCId)
	}
	return Exit()(l, c)
}

func (r ThiefJobInstructorExit) Passed(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Ohhhhh.. you collected all 30 Dark Marbles!! It should have been difficult... just incredible! Alright. You've passed the test and for that, I'll reward you ").
		BlueText().AddText("The Proof of a Hero").
		BlackText().AddText(". Take that and go back to Kerning City.")
	return SendNext(l, c, m.String(), r.Reward)
}

func (r ThiefJobInstructorExit) Reward(l logrus.FieldLogger, c Context) State {
	character.RemoveAll(l)(c.CharacterId, item.DarkMarble)
	character.CompleteQuest(l)(c.CharacterId, 100010)
	character.StartQuest(l)(c.CharacterId, 100011)
	character.GainItem(l)(c.CharacterId, item.ProofOfHero, 1)
	return r.WarpExit(l, c)
}