package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// StrangeLookingStatue is located in Hidden Street - Puppeteer's Secret Passage (910510100)
type StrangeLookingStatue struct {
}

func (r StrangeLookingStatue) NPCId() uint32 {
	return npc.StrangeLookingStatue
}

func (r StrangeLookingStatue) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r StrangeLookingStatue) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Will you exit this trial?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r StrangeLookingStatue) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.PuppeteersHidingPlace, 2)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.PuppeteersHidingPlace, c.NPCId)
	}
	return nil
}
