package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// AthenaPierceDemo is located in Maple Road : Split Road of Destiny (1020000)
type AthenaPierceDemo struct {
}

func (r AthenaPierceDemo) NPCId() uint32 {
	return npc.AthenaPierceDemo
}

func (r AthenaPierceDemo) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return r.BowmanIntroduction(l, c)
}

func (r AthenaPierceDemo) BowmanIntroduction(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Bowmen are blessed with dexterity and power, taking charge of long-distance attacks, providing support for those at the front line of the battle. Very adept at using landscape as part of the arsenal.")
	return script.SendNext(l, c, m.String(), r.Demo)
}

func (r AthenaPierceDemo) Demo(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Bowman?")
	return script.SendYesNo(l, c, m.String(), r.DoDemo, r.SeeMeAgain)
}

func (r AthenaPierceDemo) DoDemo(l logrus.FieldLogger, c script.Context) script.State {
	npc.LockUI(l)(c.CharacterId)
	return script.WarpById(_map.BowmanDemo, 0)(l, c)
}

func (r AthenaPierceDemo) SeeMeAgain(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Bowman, come see me again.")
	return script.SendNext(l, c, m.String(), script.Exit())
}
