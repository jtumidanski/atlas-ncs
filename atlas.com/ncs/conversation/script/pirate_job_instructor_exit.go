package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// PirateJobInstructorExit is located in Hidden Street - Pirate Test Room (108000500)
type PirateJobInstructorExit struct {
}

func (r PirateJobInstructorExit) NPCId() uint32 {
	return npc.PirateJobInstructorExit
}

func (r PirateJobInstructorExit) Initial(l logrus.FieldLogger, c Context) State {
	//TODO should be a better implementation than this
	if c.MapId != _map.PirateTestRoom1 && c.MapId != _map.PirateTestRoom2 {
		return r.Error(l, c)
	}
	return r.Warp(l, c)
}

func (r PirateJobInstructorExit) Error(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Error. Please report this.")
	return SendNext(l, c, m.String(), r.ErrorWarp)
}

func (r PirateJobInstructorExit) ErrorWarp(l logrus.FieldLogger, c Context) State {
	character.RemoveAll(l)(c.CharacterId, item.PotentPowerCrystal)
	character.RemoveAll(l)(c.CharacterId, item.PotentWindCrystal)
	return r.Warp(l, c)
}

func (r PirateJobInstructorExit) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.NavigationRoom, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.NavigationRoom, c.NPCId)
	}
	return Exit()(l, c)
}
