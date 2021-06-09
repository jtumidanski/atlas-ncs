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

// Moose is located in Leafre - Forest : Crossroad (240010400)
type Moose struct {
}

func (r Moose) NPCId() uint32 {
	return npc.Moose
}

func (r Moose) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.QuestStarted(l)(c.CharacterId, 6180) && character.QuestProgressInt(l)(c.CharacterId, 6180, 9300096) < 200 {
		return r.ConfirmEntrance(l, c)
	} else {
		return r.OnlyAssignedPersonnel(l, c)
	}
}

func (r Moose) ConfirmEntrance(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Pay attention: during the time you stay inside the training ground make sure you ").
		BlueText().AddText("have equipped your ").ShowItemName1(item.SkillEarningShield).
		BlackText().AddText(", it is of the utmost importance. Are you ready to proceed to the training area?")
	return script.SendYesNo(l, c, m.String(), r.Validate, script.Exit())
}

func (r Moose) OnlyAssignedPersonnel(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Only assigned personnel can access the training ground.")
	return script.SendOk(l, c, m.String())
}

func (r Moose) Validate(l logrus.FieldLogger, c script.Context) script.State {
	if !character.HasEquipped(l)(c.CharacterId, item.SkillEarningShield) {
		return r.PleaseEquip(l, c)
	}
	return r.KeepItEquipped(l, c)
}

func (r Moose) PleaseEquip(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please equip the ").
		RedText().ShowItemName1(item.SkillEarningShield).
		BlackText().AddText(" before entering the training ground.")
	return script.SendOk(l, c, m.String())
}

func (r Moose) KeepItEquipped(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Have your shield equipped until the end of the quest, or else you will need to start all over again!")
	return script.SendNext(l, c, m.String(), script.WarpById(_map.MoosesPracticeField, 0))
}
