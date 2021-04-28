package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// AthenaPierce is located in Maple Road : Split Road of Destiny (1020000)
type AthenaPierce struct {
}

func (r AthenaPierce) NPCId() uint32 {
	return 10200
}

func (r AthenaPierce) Initial() StateProducer {
	return r.BowmanIntroduction
}

func (r AthenaPierce) BowmanIntroduction(l logrus.FieldLogger, c Context) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().
		AddText("Bowmen are blessed with dexterity and power, taking charge of long-distance attacks, providing support for those at the front line of the battle. Very adept at using landscape as part of the arsenal.")
	conversation.SendNext(m.String())
	return Next(GenericExit, r.Demo)
}

func (r AthenaPierce) Demo(l logrus.FieldLogger, c Context) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Bowman?")
	conversation.SendYesNo(m.String())
	return YesNo(GenericExit, r.DoDemo, r.SeeMeAgain)
}

func (r AthenaPierce) DoDemo(l logrus.FieldLogger, c Context) State {
	npc.Processor(l).LockUI()
	npc.Processor(l).Warp(c.WorldId, c.ChannelId, c.CharacterId, 1020300, 0)
	return nil
}

func (r AthenaPierce) SeeMeAgain(l logrus.FieldLogger, c Context) State {
	conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
	m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Bowman, come see me again.")
	conversation.SendNext(m.String())
	return nil
}
