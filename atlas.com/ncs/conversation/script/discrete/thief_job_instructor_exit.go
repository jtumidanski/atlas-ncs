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

// ThiefJobInstructorExit is located in Hidden Street - Thief's Construction Site (108000400)
type ThiefJobInstructorExit struct {
}

func (r ThiefJobInstructorExit) NPCId() uint32 {
	return npc.ThiefJobInstructorExit
}

func (r ThiefJobInstructorExit) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.HasItems(l, span)(c.CharacterId, item.DarkMarble, 30) {
		return r.CollectMarbles(l, span, c)
	}
	return r.Passed(l, span, c)
}

func (r ThiefJobInstructorExit) CollectMarbles(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You will have to collect me ").
		BlueText().AddText("30 ").ShowItemName1(item.DarkMarble).
		BlackText().AddText(". Good luck.").NewLine().
		OpenItem(0).BlueText().AddText("I would like to leave").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.ExitSelection)
}

func (r ThiefJobInstructorExit) ExitSelection(_ int32) script.StateProducer {
	return r.WarpExit
}

func (r ThiefJobInstructorExit) WarpExit(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.ConstructionSiteNorthOfKerningCity, 9)(l, span, c)
}

func (r ThiefJobInstructorExit) Passed(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Ohhhhh.. you collected all 30 Dark Marbles!! It should have been difficult... just incredible! Alright. You've passed the test and for that, I'll reward you ").
		BlueText().AddText("The Proof of a Hero").
		BlackText().AddText(". Take that and go back to Kerning City.")
	return script.SendNext(l, span, c, m.String(), r.Reward)
}

func (r ThiefJobInstructorExit) Reward(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.RemoveAll(l)(c.CharacterId, item.DarkMarble)
	quest.Complete(l)(c.CharacterId, 100010)
	quest.Start(l)(c.CharacterId, 100011)
	character.GainItem(l, span)(c.CharacterId, item.ProofOfHero, 1)
	return r.WarpExit(l, span, c)
}