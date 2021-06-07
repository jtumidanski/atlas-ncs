package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Milla is located in 
type Milla struct {
}

func (r Milla) NPCId() uint32 {
	return npc.Milla
}

func (r Milla) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.EntranceToMVsLair {
		m := message.NewBuilder().
			AddText("Hi, I'm ").
			ShowNPC(npc.Milla).
			AddText(".")
		return SendOk(l, c, m.String())
	}
	if c.MapId == _map.TreasureDungeon {
		m := message.NewBuilder().
			AddText("Hi there, ").
			ShowCharacterName().
			AddText(". This is the MV's treasure room. Use the time you have here to do whatever you want, there are a lot of things to uncover here, actually. Or else you can use the portal here to ").
			RedText().AddText("go back").
			BlackText().AddText(" to the entrance.")
		return SendOk(l, c, m.String())
	}
	m := message.NewBuilder().AddText("Are you sure you want to return? By returning now you are leaving your partners behind, do you really want to do it?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r Milla) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.Warp(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.EntranceToMVsLair)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.EntranceToMVsLair, c.NPCId)
	}
	return Exit()(l, c)
}
