package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// PirateJobInstructorExit is located in Hidden Street - Pirate Test Room (108000500)
type PirateJobInstructorExit struct {
}

func (r PirateJobInstructorExit) NPCId() uint32 {
	return npc.PirateJobInstructorExit
}

func (r PirateJobInstructorExit) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	//TODO should be a better implementation than this
	if c.MapId != _map.PirateTestRoom1 && c.MapId != _map.PirateTestRoom2 {
		return r.Error(l, span, c)
	}
	return r.Warp(l, span, c)
}

func (r PirateJobInstructorExit) Error(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Error. Please report this.")
	return script.SendNext(l, span, c, m.String(), r.ErrorWarp)
}

func (r PirateJobInstructorExit) ErrorWarp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.RemoveAll(l)(c.CharacterId, item.PotentPowerCrystal)
	character.RemoveAll(l)(c.CharacterId, item.PotentWindCrystal)
	return r.Warp(l, span, c)
}

func (r PirateJobInstructorExit) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.NavigationRoom, 0)(l, span, c)
}
