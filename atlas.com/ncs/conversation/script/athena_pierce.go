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

func (r AthenaPierce) Initial(l logrus.FieldLogger, c Context) State {
	return r.BowmanIntroduction(l, c)
}

func (r AthenaPierce) BowmanIntroduction(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Bowmen are blessed with dexterity and power, taking charge of long-distance attacks, providing support for those at the front line of the battle. Very adept at using landscape as part of the arsenal.")
	return SendNext(l, c, m.String(), r.Demo)
}

func (r AthenaPierce) Demo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Bowman?")
	return SendYesNo(l, c, m.String(), r.DoDemo, r.SeeMeAgain)
}

func (r AthenaPierce) DoDemo(l logrus.FieldLogger, c Context) State {
	npc.Processor(l).LockUI()

	mapId := uint32(1020300)
	err := npc.Processor(l).Warp(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
	}
	return nil
}

func (r AthenaPierce) SeeMeAgain(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Bowman, come see me again.")
	return SendNext(l, c, m.String(), Exit())
}
