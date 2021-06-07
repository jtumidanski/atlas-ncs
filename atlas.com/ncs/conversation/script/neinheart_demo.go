package script

import (
	"atlas-ncs/character"
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

func (r NeinheartDemo) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Becoming a Knight of Cygnus requires talent, faith, courage, and will power... and it looks like you are more than qualified to become a Knight of Cygnus. What do you think? If you wish to become one right this minute, I'll take you straight to Erev. Would you like to head over to Erev right now?")
	return SendAcceptDecline(l, c, m.String(), r.WarpToEreve, r.WarpBack)
}

func (r NeinheartDemo) WarpToEreve(l logrus.FieldLogger, c Context) State {
	return Warp(_map.Ereve)(l, c)
}

func (r NeinheartDemo) WarpBack(l logrus.FieldLogger, c Context) State {
	mapId := character.SavedLocation(l)(c.CharacterId, "CYGNUSINTRO")
	return Warp(mapId)(l, c)
}
