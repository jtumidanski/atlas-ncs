package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// MagicianJobInstructorExit is located in Hidden Street - Magician's Tree Dungeon (108000200)
type MagicianJobInstructorExit struct {
}

func (r MagicianJobInstructorExit) NPCId() uint32 {
	return npc.MagicianJobInstructorExit
}

func (r MagicianJobInstructorExit) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.HasItems(l)(c.CharacterId, item.DarkMarble, 30) {
		return r.CollectMarbles(l, c)
	}
	return r.Passed(l, c)
}

func (r MagicianJobInstructorExit) CollectMarbles(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You will have to collect me ").
		BlueText().AddText("30 ").ShowItemName1(item.DarkMarble).
		BlackText().AddText(". Good luck.").NewLine().
		OpenItem(0).BlueText().AddText("I would like to leave").CloseItem()
	return script.SendListSelection(l, c, m.String(), r.ExitSelection)
}

func (r MagicianJobInstructorExit) ExitSelection(_ int32) script.StateProducer {
	return r.WarpExit
}

func (r MagicianJobInstructorExit) WarpExit(l logrus.FieldLogger, c script.Context) script.State {
	return script.WarpById(_map.TheForestNorthOfEllinia, 1)(l, c)
}

func (r MagicianJobInstructorExit) Passed(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Ohhhhh.. you collected all 30 Dark Marbles!! It should have been difficult... just incredible! Alright. You've passed the test and for that, I'll reward you ").
		BlueText().AddText("The Proof of a Hero").
		BlackText().AddText(". Take that and go back to Ellinia.")
	return script.SendNext(l, c, m.String(), r.Reward)
}

func (r MagicianJobInstructorExit) Reward(l logrus.FieldLogger, c script.Context) script.State {
	character.RemoveAll(l)(c.CharacterId, item.DarkMarble)
	character.CompleteQuest(l)(c.CharacterId, 100007)
	character.StartQuest(l)(c.CharacterId, 100008)
	character.GainItem(l)(c.CharacterId, item.ProofOfHero, 1)
	return r.WarpExit(l, c)
}
