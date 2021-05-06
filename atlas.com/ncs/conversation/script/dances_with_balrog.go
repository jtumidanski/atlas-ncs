package script

import (
    _map "atlas-ncs/map"
    "atlas-ncs/npc"
    "atlas-ncs/npc/message"
    "github.com/sirupsen/logrus"
)

// DancesWithBalrog is located in Maple Road : Split Road of Destiny (1020000)
type DancesWithBalrog struct {
}

func (r DancesWithBalrog) NPCId() uint32 {
	return npc.DancesWithBalrogDemo
}

func (r DancesWithBalrog) Initial(l logrus.FieldLogger, c Context) State {
	return r.WarriorIntroduction(l, c)
}

func (r DancesWithBalrog) WarriorIntroduction(l logrus.FieldLogger, c Context) State {
    m := message.NewBuilder().
        AddText("Warriors possess an enormous power with stamina to back it up, and they shine the brightest in melee combat situation. Regular attacks are powerful to begin with, and armed with complex skills, the job is perfect for explosive attacks.")
    return SendNext(l, c, m.String(), r.Demo)
}

func (r DancesWithBalrog) Demo(l logrus.FieldLogger, c Context) State {
    m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Warrior?")
    return SendYesNo(l, c, m.String(), r.DoDemo, r.SeeMeAgain)
}

func (r DancesWithBalrog) DoDemo(l logrus.FieldLogger, c Context) State {
    npc.Processor(l).LockUI()

    err := npc.Processor(l).Warp(c.WorldId, c.ChannelId, c.CharacterId, _map.WarriorDemo, 0)
    if err != nil {
        l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.WarriorDemo, c.NPCId)
    }
    return nil
}

func (r DancesWithBalrog) SeeMeAgain(l logrus.FieldLogger, c Context) State {
    m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Warrior, come see me again.")
    return SendNext(l, c, m.String(), Exit())
}
