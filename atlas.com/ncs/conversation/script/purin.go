package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Purin is located in Victoria Road - Before Takeoff <To Orbis> (101000301)
type Purin struct {
}

func (r Purin) NPCId() uint32 {
	return npc.Purin
}

func (r Purin) Initial(l logrus.FieldLogger, c Context) State {
	return r.DoYouWish(l, c)
}

func (r Purin) DoYouWish(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Do you wish to leave the boat?")
	return SendYesNo(l, c, m.String(), r.ProcessLeave, Exit())
}

func (r Purin) ProcessLeave(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Alright, see you next time. Take care.")
	return SendNext(l, c, m.String(), r.Warp)
}

func (r Purin) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ElliniaStation, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.ElliniaStation, c.NPCId)
	}
	return nil
}
