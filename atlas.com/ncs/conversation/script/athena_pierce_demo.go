package script

import (
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

func (r AthenaPierceDemo) Initial(l logrus.FieldLogger, c Context) State {
	return r.BowmanIntroduction(l, c)
}

func (r AthenaPierceDemo) BowmanIntroduction(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Bowmen are blessed with dexterity and power, taking charge of long-distance attacks, providing support for those at the front line of the battle. Very adept at using landscape as part of the arsenal.")
	return SendNext(l, c, m.String(), r.Demo)
}

func (r AthenaPierceDemo) Demo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Bowman?")
	return SendYesNo(l, c, m.String(), r.DoDemo, r.SeeMeAgain)
}

func (r AthenaPierceDemo) DoDemo(l logrus.FieldLogger, c Context) State {
	npc.LockUI(l)(c.CharacterId)

	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.BowmanDemo, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.BowmanDemo, c.NPCId)
	}
	return nil
}

func (r AthenaPierceDemo) SeeMeAgain(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Bowman, come see me again.")
	return SendNext(l, c, m.String(), Exit())
}
