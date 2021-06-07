package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Sage2 is located in 
type Sage2 struct {
}

func (r Sage2) NPCId() uint32 {
	return npc.Sage2
}

func (r Sage2) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.UpperAscent {
		m := message.NewBuilder().
			AddText("The Crimsonwood Keep lies right ahead, a great feat has been made by you this day, salute to thee. Pass through these woods to enter the gates of the Keep.")
		return SendOk(l, c, m.String())
	}
	m := message.NewBuilder().
		AddText("So far your progress is splendid, good job. However, to make it to the Keep, you must face and accomplish this ordeal, carry on.")
	return SendOk(l, c, m.String())
}
