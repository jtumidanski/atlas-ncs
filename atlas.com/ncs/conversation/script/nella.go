package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Nella is located in Kerning PQ
type Nella struct {
}

func (r Nella) NPCId() uint32 {
	return npc.Nella
}

func (r Nella) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.FirstAccompanimentExit {
		return r.ExitTheExit(l, c)
	}
	if c.MapId == _map.FirstAccompanimentBonus {
		return r.ExitBonus(l, c)
	}
	return r.ExitQuestMaps(l, c)
}

func (r Nella) ExitQuestMaps(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Once you leave the map, you'll have to restart the whole quest if you want to try it again.  Do you still want to leave this map?")
	return SendYesNo(l, c, m.String(), r.WarpToExit, Exit())
}

func (r Nella) ExitBonus(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Are you ready to leave this map?")
	return SendYesNo(l, c, m.String(), r.WarpToExit, Exit())
}

func (r Nella) ExitTheExit(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("To return back to the city, follow this way.")
	return SendNext(l, c, m.String(), r.WarpToKerning)
}

func (r Nella) WarpToKerning(l logrus.FieldLogger, c Context) State {
	err := npc.Warp(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.KerningCity)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.KerningCity, c.NPCId)
	}
	return Exit()(l, c)
}

func (r Nella) WarpToExit(l logrus.FieldLogger, c Context) State {
	err := npc.WarpByName(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.FirstAccompanimentExit, "st00")
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.FirstAccompanimentExit, c.NPCId)
	}
	return Exit()(l, c)
}
