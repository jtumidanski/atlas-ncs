package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// BowmanJobInstructorExit is located in Hidden Street - Ant Tunnel For Bowman (108000100)
type BowmanJobInstructorExit struct {
}

func (r BowmanJobInstructorExit) NPCId() uint32 {
	return npc.BowmanJobInstructorExit
}

func (r BowmanJobInstructorExit) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItems(l, span)(c.CharacterId, item.DarkMarble, 30) {
		return r.CollectMarbles(l, span, c)
	}
	return r.Passed(l, span, c)
}

func (r BowmanJobInstructorExit) CollectMarbles(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You will have to collect me ").
		BlueText().AddText("30 ").ShowItemName1(item.DarkMarble).
		BlackText().AddText(". Good luck.").NewLine().
		OpenItem(0).BlueText().AddText("I would like to leave").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.ExitSelection)
}

func (r BowmanJobInstructorExit) ExitSelection(_ int32) script.StateProducer {
	return script.WarpById(_map.TheRoadToTheDungeon, 9)
}

func (r BowmanJobInstructorExit) Passed(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Ohhhhh.. you collected all 30 Dark Marbles!! It should have been difficult... just incredible! Alright. You've passed the test and for that, I'll reward you ").
		BlueText().AddText("The Proof of a Hero").
		BlackText().AddText(". Take that and go back to Henesys.")
	return script.SendNext(l, span, c, m.String(), r.Reward)
}

func (r BowmanJobInstructorExit) Reward(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.RemoveAll(l)(c.CharacterId, item.DarkMarble)
	quest.Complete(l)(c.CharacterId, 100001)
	quest.Start(l)(c.CharacterId, 100002)
	character.GainItem(l, span)(c.CharacterId, item.ProofOfHero, 1)
	return script.WarpById(_map.TheRoadToTheDungeon, 9)(l, span, c)
}
