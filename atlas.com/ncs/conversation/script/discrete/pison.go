package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Pison is located in Florina Road - Florina Beach (110000000)
type Pison struct {
}

func (r Pison) NPCId() uint32 {
	return npc.Pison
}

func (r Pison) Initial(l logrus.FieldLogger, c script.Context) script.State {
	mapId := character.SavedLocation(l)(c.CharacterId, "FLORINA")
	m := message.NewBuilder().
		AddText("So you want to leave ").
		BlueText().ShowMap(_map.FlorinaBeach).
		BlackText().AddText("? If you want, I can take you back to ").
		BlueText().ShowMap(mapId).
		BlackText().AddText(".")
	return script.SendNextExit(l, c, m.String(), r.Confirm(mapId), r.OtherBusiness)
}

func (r Pison) Confirm(mapId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Are you sure you want to return to ").
			BlueText().ShowMap(mapId).
			BlackText().AddText("? Alright, we'll have to get going fast. Do you want to head back to ").
			ShowMap(mapId).AddText(" now?")
		return script.SendYesNoExit(l, c, m.String(), r.Warp(mapId), r.OtherBusiness, r.OtherBusiness)
	}
}

func (r Pison) OtherBusiness(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You must have some business to take care of here. It's not a bad idea to take some rest at ").
		ShowMap(_map.FlorinaBeach).
		AddText(" Look at me; I love it here so much that I wound up living here. Hahaha anyway, talk to me when you feel like going back.")
	return script.SendOk(l, c, m.String())
}

func (r Pison) Warp(mapId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		character.ClearSavedLocation(l)(c.CharacterId, "FLORINA")
		return script.WarpById(mapId, 0)(l, c)
	}
}
