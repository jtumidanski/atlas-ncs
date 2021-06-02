package script

import (
	"atlas-ncs/character"
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

func (r Moose) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 6180) && character.QuestProgressInt(l)(c.CharacterId, 6180, 9300096) < 200 {
		return r.ConfirmEntrance(l, c)
	} else {
		return r.OnlyAssignedPersonnel(l, c)
	}
}

func (r Moose) ConfirmEntrance(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Pay attention: during the time you stay inside the training ground make sure you ").
		BlueText().AddText("have equipped your ").ShowItemName1(item.SkillEarningShield).
		BlackText().AddText(", it is of the utmost importance. Are you ready to proceed to the training area?")
	return SendYesNo(l, c, m.String(), r.Validate, Exit())
}

func (r Moose) OnlyAssignedPersonnel(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Only assigned personnel can access the training ground.")
	return SendOk(l, c, m.String())
}

func (r Moose) Validate(l logrus.FieldLogger, c Context) State {
	if !character.HasEquipped(l)(c.CharacterId, item.SkillEarningShield) {
		return r.PleaseEquip(l, c)
	}
	return r.KeepItEquipped(l, c)
}

func (r Moose) PleaseEquip(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Please equip the ").
		RedText().ShowItemName1(item.SkillEarningShield).
		BlackText().AddText(" before entering the training ground.")
	return SendOk(l, c, m.String())
}

func (r Moose) KeepItEquipped(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Have your shield equipped until the end of the quest, or else you will need to start all over again!")
	return SendNext(l, c, m.String(), r.Warp)
}

func (r Moose) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.MoosesPracticeField, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.MoosesPracticeField, c.NPCId)
	}
	return Exit()(l, c)
}
