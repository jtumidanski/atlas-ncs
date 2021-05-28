package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DarkLordDemo is located in Maple Road : Split Road of Destiny (1020000)
type DarkLordDemo struct {
}

func (r DarkLordDemo) NPCId() uint32 {
	return npc.DarkLordDemo
}

func (r DarkLordDemo) Initial(l logrus.FieldLogger, c Context) State {
	return r.ThiefIntroduction(l, c)
}

func (r DarkLordDemo) ThiefIntroduction(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Thieves are a perfect blend of luck, dexterity, and power that are adept at the surprise attacks against helpless enemies. A high level of avoidability and speed allows Thieves to attack enemies from various angles.")
	return SendNext(l, c, m.String(), r.Demo)
}

func (r DarkLordDemo) Demo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Thief?")
	return SendYesNo(l, c, m.String(), r.DoDemo, r.SeeMeAgain)
}

func (r DarkLordDemo) DoDemo(l logrus.FieldLogger, c Context) State {
	npc.LockUI(l)(c.CharacterId)

	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ThiefDemo, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.ThiefDemo, c.NPCId)
	}
	return nil
}

func (r DarkLordDemo) SeeMeAgain(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Thief, come see me again.")
	return SendNext(l, c, m.String(), Exit())
}
