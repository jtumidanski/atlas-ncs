package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// NeinheartDemo is located in Opening - Cygnus Knights (913040006)
type NeinheartDemo struct {
}

func (r NeinheartDemo) NPCId() uint32 {
	return npc.NeinheartDemo
}

func (r NeinheartDemo) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Becoming a Knight of Cygnus requires talent, faith, courage, and will power... and it looks like you are more than qualified to become a Knight of Cygnus. What do you think? If you wish to become one right this minute, I'll take you straight to Erev. Would you like to head over to Erev right now?")
	return script.SendAcceptDecline(l, c, m.String(), r.WarpToEreve, r.WarpBack)
}

func (r NeinheartDemo) WarpToEreve(l logrus.FieldLogger, c script.Context) script.State {
	return script.Warp(_map.Ereve)(l, c)
}

func (r NeinheartDemo) WarpBack(l logrus.FieldLogger, c script.Context) script.State {
	mapId := character.SavedLocation(l)(c.CharacterId, "CYGNUSINTRO")
	return script.Warp(mapId)(l, c)
}
