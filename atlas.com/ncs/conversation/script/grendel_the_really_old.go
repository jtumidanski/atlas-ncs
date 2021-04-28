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

func (r GrendelTheReallyOld) Initial() StateProducer {
	return r.MagicianIntroduction
}

func (r GrendelTheReallyOld) MagicianIntroduction(l logrus.FieldLogger, c Context) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("Magicians are armed with flashy element-based spells and secondary magic that aids party as a whole. After the 2nd job adv., the elemental-based magic will provide ample amount of damage to enemies of opposite element.")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.Demo)
}

func (r GrendelTheReallyOld) Demo(l logrus.FieldLogger, c Context) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Magician?")
	conversation.SendYesNo(m.String())
	return YesNo(GenericExit, r.DoDemo, r.SeeMeAgain)
}

func (r GrendelTheReallyOld) DoDemo(l logrus.FieldLogger, c Context) State {
	npc.Processor(l).LockUI()
	npc.Processor(l).Warp(c.WorldId, c.ChannelId, c.CharacterId, 1020200, 0)
	return nil
}

func (r GrendelTheReallyOld) SeeMeAgain(l logrus.FieldLogger, c Context) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Magician, come see me again.")
	conversation.SendNext(m.String())
	return nil
}