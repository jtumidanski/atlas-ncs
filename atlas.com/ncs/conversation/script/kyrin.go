package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Kyrin is located in Maple Road : Split Road of Destiny (1020000)
type Kyrin struct {
}

func (r Kyrin) NPCId() uint32 {
	return npc.KyrinDemo
}

func (r Kyrin) Initial(l logrus.FieldLogger, c Context) State {
	return r.PirateIntroduction(l, c)
}

func (r Kyrin) PirateIntroduction(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Pirates are blessed with outstanding dexterity and power, utilizing their guns for long-range attacks while using their power on melee combat situations. Gunslingers use elemental-based bullets for added damage, while Infighters transform to a different being for maximum effect.")
	return SendNext(l, c, m.String(), r.Demo)
}

func (r Kyrin) Demo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Pirate?")
	return SendYesNo(l, c, m.String(), r.DoDemo, r.SeeMeAgain)
}

func (r Kyrin) DoDemo(l logrus.FieldLogger, c Context) State {
	npc.Processor(l).LockUI()

	err := npc.Processor(l).Warp(c.WorldId, c.ChannelId, c.CharacterId, _map.PirateDemo, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.PirateDemo, c.NPCId)
	}
	return nil
}

func (r Kyrin) SeeMeAgain(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Pirate, come see me again.")
	return SendNext(l, c, m.String(), Exit())
}
