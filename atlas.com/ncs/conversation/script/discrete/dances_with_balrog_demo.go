package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
    "atlas-ncs/npc"
    "atlas-ncs/npc/message"
    "github.com/sirupsen/logrus"
)

// DancesWithBalrogDemo is located in Maple Road : Split Road of Destiny (1020000)
type DancesWithBalrogDemo struct {
}

func (r DancesWithBalrogDemo) NPCId() uint32 {
	return npc.DancesWithBalrogDemo
}

func (r DancesWithBalrogDemo) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return r.WarriorIntroduction(l, c)
}

func (r DancesWithBalrogDemo) WarriorIntroduction(l logrus.FieldLogger, c script.Context) script.State {
    m := message.NewBuilder().
        AddText("Warriors possess an enormous power with stamina to back it up, and they shine the brightest in melee combat situation. Regular attacks are powerful to begin with, and armed with complex skills, the job is perfect for explosive attacks.")
    return script.SendNext(l, c, m.String(), r.Demo)
}

func (r DancesWithBalrogDemo) Demo(l logrus.FieldLogger, c script.Context) script.State {
    m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Warrior?")
    return script.SendYesNo(l, c, m.String(), r.DoDemo, r.SeeMeAgain)
}

func (r DancesWithBalrogDemo) DoDemo(l logrus.FieldLogger, c script.Context) script.State {
    npc.LockUI(l)(c.CharacterId)
    return script.WarpById(_map.WarriorDemo, 0)(l, c)
}

func (r DancesWithBalrogDemo) SeeMeAgain(l logrus.FieldLogger, c script.Context) script.State {
    m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Warrior, come see me again.")
    return script.SendNext(l, c, m.String(), script.Exit())
}
