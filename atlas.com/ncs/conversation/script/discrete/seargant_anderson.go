package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// SeargantAnderson is located in Abandoned Tower
type SeargantAnderson struct {
}

func (r SeargantAnderson) NPCId() uint32 {
	return npc.SeargantAnderson
}

func (r SeargantAnderson) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if c.MapId == _map.AbandonedTowerEndOfJourney {
		return r.ToReturn(l, c)
	} else {
		return r.Confirm(l, c)
	}
}

func (r SeargantAnderson) Confirm(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Once you leave the map, you'll have to restart the whole quest if you want to try it again.  Do you still want to leave this map?")
	return script.SendYesNo(l, c, m.String(), r.WarpExit, script.Exit())
}

func (r SeargantAnderson) WarpExit(l logrus.FieldLogger, c script.Context) script.State {
	return script.WarpById(_map.AbandonedTowerEndOfJourney, 0)(l, c)
}

func (r SeargantAnderson) ToReturn(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("To return back to the recruitment map, follow this way.")
	return script.SendNext(l, c, m.String(), r.WarpToTower)
}

func (r SeargantAnderson) WarpToTower(l logrus.FieldLogger, c script.Context) script.State {
	return script.WarpById(_map.EosTower101stFloor, 0)(l, c)
}
