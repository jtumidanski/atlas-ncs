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

// MarkTheToySoldier is located in Hidden Street - Doll's House (922000010)
type MarkTheToySoldier struct {
}

func (r MarkTheToySoldier) NPCId() uint32 {
	return npc.MarkTheToySoldier
}

func (r MarkTheToySoldier) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.QuestStarted(l)(c.CharacterId, 3230) {
		return r.ThankYou(l, c)
	}
	if !character.HasItem(l)(c.CharacterId, item.Pendulum) {
		return r.YouHaveNot(l, c)
	}
	return r.Process(l, c)
}

func (r MarkTheToySoldier) ThankYou(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Thank you for finding the pendulum. Are you ready to return to Eos Tower?")
	return script.SendYesNo(l, c, m.String(), r.Warp, script.Exit())
}

func (r MarkTheToySoldier) YouHaveNot(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("You haven't found the pendulum yet. Do you want to go back to Eos Tower?")
	return script.SendYesNo(l, c, m.String(), r.Warp, script.Exit())
}

func (r MarkTheToySoldier) Process(l logrus.FieldLogger, c script.Context) script.State {
	character.CompleteQuest(l)(c.CharacterId, 3230)
	character.GainItem(l)(c.CharacterId, item.Pendulum, -1)
	return r.ThankYou(l, c)
}

func (r MarkTheToySoldier) Warp(l logrus.FieldLogger, c script.Context) script.State {
	return script.WarpById(_map.EosTower100thFloor, 4)(l, c)
}
