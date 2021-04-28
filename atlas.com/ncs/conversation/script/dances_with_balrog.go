package script

import (
    "atlas-ncs/npc"
    "atlas-ncs/npc/message"
    "github.com/sirupsen/logrus"
)

// DancesWithBalrog is located in Maple Road : Split Road of Destiny (1020000)
type DancesWithBalrog struct {
}

func (r DancesWithBalrog) NPCId() uint32 {
	return 10202
}

func (r DancesWithBalrog) Initial() StateProducer {
	return r.WarriorIntroduction
}

func (r DancesWithBalrog) WarriorIntroduction(l logrus.FieldLogger, c Context) State {
    conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
    m := message.NewBuilder().
        AddText("Warriors possess an enormous power with stamina to back it up, and they shine the brightest in melee combat situation. Regular attacks are powerful to begin with, and armed with complex skills, the job is perfect for explosive attacks.")
    conversation.SendNext(m.String())
    return Next(GenericExit, r.Demo)
}

func (r DancesWithBalrog) Demo(l logrus.FieldLogger, c Context) State {
    conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
    m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Warrior?")
    conversation.SendYesNo(m.String())
    return YesNo(GenericExit, r.DoDemo, r.SeeMeAgain)
}

func (r DancesWithBalrog) DoDemo(l logrus.FieldLogger, c Context) State {
    npc.Processor(l).LockUI()
    npc.Processor(l).Warp(c.WorldId, c.ChannelId, c.CharacterId, 1020100, 0)
    return nil
}

func (r DancesWithBalrog) SeeMeAgain(l logrus.FieldLogger, c Context) State {
    conversation := npc.Processor(l).Conversation(c.CharacterId, c.NPCId)
    m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Warrior, come see me again.")
    conversation.SendNext(m.String())
    return nil
}
