package discrete

import (
	"atlas-ncs/conversation/script"
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

func (r SnowmanExit) Initial(l logrus.FieldLogger, c script.Context) script.State {
	area := c.MapId % 10
	if area > 0 {
		m := message.NewBuilder().
			AddText("Do you wish to leave this place?")
		return script.SendYesNo(l, c, m.String(), script.WarpById(c.MapId+1, 0), script.Exit())
	} else {
		m := message.NewBuilder().
			AddText("Do you wish to return to ").
			BlueText().AddText("Happyville").
			BlackText().AddText("?")
		return script.SendYesNo(l, c, m.String(), r.WarpToHappyville, script.Exit())
	}
}

func (r SnowmanExit) WarpToHappyville(l logrus.FieldLogger, c script.Context) script.State {
	return script.Warp(_map.Happyville)(l, c)
}
