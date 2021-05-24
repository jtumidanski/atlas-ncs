package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// GrendelTheReallyOldDemo is located in Maple Road : Split Road of Destiny (1020000)
type GrendelTheReallyOldDemo struct {
}

func (r GrendelTheReallyOldDemo) NPCId() uint32 {
	return npc.GrendelTheReallyOldDemo
}

func (r GrendelTheReallyOldDemo) Initial(l logrus.FieldLogger, c Context) State {
	return r.MagicianIntroduction(l, c)
}

func (r GrendelTheReallyOldDemo) MagicianIntroduction(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Magicians are armed with flashy element-based spells and secondary magic that aids party as a whole. After the 2nd job adv., the elemental-based magic will provide ample amount of damage to enemies of opposite element.")
	return SendNext(l, c, m.String(), r.Demo)
}

func (r GrendelTheReallyOldDemo) Demo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Magician?")
	return SendYesNo(l, c, m.String(), r.DoDemo, r.SeeMeAgain)
}

func (r GrendelTheReallyOldDemo) DoDemo(l logrus.FieldLogger, c Context) State {
	npc.Processor(l).LockUI()

	err := npc.Processor(l).WarpById(c.WorldId, c.ChannelId, c.CharacterId, _map.MagicianDemo, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.MagicianDemo, c.NPCId)
	}
	return nil
}

func (r GrendelTheReallyOldDemo) SeeMeAgain(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Magician, come see me again.")
	return SendNext(l, c, m.String(), Exit())
}