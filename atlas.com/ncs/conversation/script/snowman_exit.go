package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// SnowmanExit is located in 
type SnowmanExit struct {
}

func (r SnowmanExit) NPCId() uint32 {
	return npc.SnowmanExit
}

func (r SnowmanExit) Initial(l logrus.FieldLogger, c Context) State {
	area := c.MapId % 10
	if area > 0 {
		m := message.NewBuilder().
			AddText("Do you wish to leave this place?")
		return SendYesNo(l, c, m.String(), WarpById(c.MapId+1, 0), Exit())
	} else {
		m := message.NewBuilder().
			AddText("Do you wish to return to ").
			BlueText().AddText("Happyville").
			BlackText().AddText("?")
		return SendYesNo(l, c, m.String(), r.WarpToHappyville, Exit())
	}
}

func (r SnowmanExit) WarpToHappyville(l logrus.FieldLogger, c Context) State {
	return Warp(_map.Happyville)(l, c)
}
