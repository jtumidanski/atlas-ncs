package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Shulynch is located in The Nautilus - Training Room (120000104)
type Shulynch struct {
}

func (r Shulynch) NPCId() uint32 {
	return npc.Shulynch
}

func (r Shulynch) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestStarted(l)(c.CharacterId, 6410) {
		return r.AnyBusiness(l, c)
	}
	return r.LetsGoSave(l, c)
}

func (r Shulynch) AnyBusiness(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hey, do you have any business with me?")
	return SendOk(l, c, m.String())
}

func (r Shulynch) LetsGoSave(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Let's go save ").
		RedText().ShowNPC(npc.Delli).
		BlackText().AddText("?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r Shulynch) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.LookingForDelli1, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.LookingForDelli1, c.NPCId)
	}
	return nil
}
