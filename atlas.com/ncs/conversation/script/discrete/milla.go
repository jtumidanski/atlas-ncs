package discrete

import (
	"atlas-ncs/conversation/script"
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

func (r Milla) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if c.MapId == _map.EntranceToMVsLair {
		m := message.NewBuilder().
			AddText("Hi, I'm ").
			ShowNPC(npc.Milla).
			AddText(".")
		return script.SendOk(l, c, m.String())
	}
	if c.MapId == _map.TreasureDungeon {
		m := message.NewBuilder().
			AddText("Hi there, ").
			ShowCharacterName().
			AddText(". This is the MV's treasure room. Use the time you have here to do whatever you want, there are a lot of things to uncover here, actually. Or else you can use the portal here to ").
			RedText().AddText("go back").
			BlackText().AddText(" to the entrance.")
		return script.SendOk(l, c, m.String())
	}
	m := message.NewBuilder().AddText("Are you sure you want to return? By returning now you are leaving your partners behind, do you really want to do it?")
	return script.SendYesNo(l, c, m.String(), script.Warp(_map.EntranceToMVsLair), script.Exit())
}
