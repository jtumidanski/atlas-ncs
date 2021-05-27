package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// CrumblingStatue is located in Hidden Street - The Deep Forest of Patience <Step 1> (105040310) Hidden Street - The Deep Forest of Patience <Step 2> (105040311) Hidden Street - The Deep Forest of Patience <Step 3> (105040312) Hidden Street - The Deep Forest of Patience <Step 4> (105040313) Hidden Street - The Deep Forest of Patience <Step 5> (105040314) Hidden Street - The Deep Forest of Patience <Step 6> (105040315) Hidden Street - The Deep Forest of Patience <Step 7> (105040316)
type CrumblingStatue struct {
}

func (r CrumblingStatue) NPCId() uint32 {
	return npc.CrumblingStatue
}

func (r CrumblingStatue) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Would you like to leave?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r CrumblingStatue) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.Processor(l).WarpById(c.WorldId, c.ChannelId, c.CharacterId, _map.Sleepywood, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.Sleepywood, c.NPCId)
	}
	return nil
}
