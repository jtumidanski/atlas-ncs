package script

import (
	"atlas-ncs/event"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Amon is located in 
type Amon struct {
}

func (r Amon) NPCId() uint32 {
	return npc.Amon
}

func (r Amon) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId != _map.ZakumsAltar {
		return r.LeaveNow(l, c)
	}

	if !event.Cleared(l)(c.CharacterId) {
		return r.LeaveNow(l, c)
	}

	return r.Congratulations(l, c)
}

func (r Amon) LeaveNow(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("If you leave now, you'll have to start over. Are you sure you want to leave?")
	return SendYesNo(l, c, m.String(), WarpById(_map.TheDoorToZakum, 0), Exit())
}

func (r Amon) Congratulations(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You guys finally overthrew Zakum, what a superb feat! Congratulations! Are you sure you want to leave now?")
	return SendYesNo(l, c, m.String(), WarpById(_map.TheDoorToZakum, 0), Exit())
}
