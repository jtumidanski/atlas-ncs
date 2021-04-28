package script

import (
    "atlas-ncs/npc"
    "atlas-ncs/npc/message"
    "github.com/sirupsen/logrus"
)

// Kyrin is located in Maple Road : Split Road of Destiny (1020000)
type Kyrin struct {
}

func (r Kyrin) NPCId() uint32 {
	return 10204
}

func (r Kyrin) Initial() StateProducer {
	return r.PirateIntroduction
}

func (r Kyrin) PirateIntroduction(l logrus.FieldLogger, c Context) State {
    conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
    m := message.NewBuilder().
        AddText("Pirates are blessed with outstanding dexterity and power, utilizing their guns for long-range attacks while using their power on melee combat situations. Gunslingers use elemental-based bullets for added damage, while Infighters transform to a different being for maximum effect.")
    conversation.SendNext(m.String())
    return Next(GenericExit, r.Demo)
}

func (r Kyrin) Demo(l logrus.FieldLogger, c Context) State {
    conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
    m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Pirate?")
    conversation.SendYesNo(m.String())
    return YesNo(GenericExit, r.DoDemo, r.SeeMeAgain)
}

func (r Kyrin) DoDemo(l logrus.FieldLogger, c Context) State {
    npc.Processor(l).LockUI()
    npc.Processor(l).Warp(c.WorldId, c.ChannelId, c.CharacterId, 1020500, 0)
    return nil
}

func (r Kyrin) SeeMeAgain(l logrus.FieldLogger, c Context) State {
    conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
    m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Pirate, come see me again.")
    conversation.SendNext(m.String())
    return nil
}
