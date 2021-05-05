package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// GrendelTheReallyOld is located in Maple Road : Split Road of Destiny (1020000)
type GrendelTheReallyOld struct {
}

func (r GrendelTheReallyOld) NPCId() uint32 {
	return 10201
}

func (r GrendelTheReallyOld) Initial(l logrus.FieldLogger, c Context) State {
	return r.MagicianIntroduction(l, c)
}

func (r GrendelTheReallyOld) MagicianIntroduction(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Magicians are armed with flashy element-based spells and secondary magic that aids party as a whole. After the 2nd job adv., the elemental-based magic will provide ample amount of damage to enemies of opposite element.")
	return SendNext(l, c, m.String(), r.Demo)
}

func (r GrendelTheReallyOld) Demo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Magician?")
	return SendYesNo(l, c, m.String(), r.DoDemo, r.SeeMeAgain)
}

func (r GrendelTheReallyOld) DoDemo(l logrus.FieldLogger, c Context) State {
	npc.Processor(l).LockUI()

	mapId := uint32(1020200)
	err := npc.Processor(l).Warp(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
	}
	return nil
}

func (r GrendelTheReallyOld) SeeMeAgain(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Magician, come see me again.")
	return SendNext(l, c, m.String(), Exit())
}