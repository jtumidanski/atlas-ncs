package script

import (
	"atlas-ncs/character"
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

func (r Pison) Initial(l logrus.FieldLogger, c Context) State {
	mapId := character.SavedLocation(l)(c.CharacterId, "FLORINA")
	m := message.NewBuilder().
		AddText("So you want to leave ").
		BlueText().ShowMap(_map.FlorinaBeach).
		BlackText().AddText("? If you want, I can take you back to ").
		BlueText().ShowMap(mapId).
		BlackText().AddText(".")
	return SendNextExit(l, c, m.String(), r.Confirm(mapId), r.OtherBusiness)
}

func (r Pison) Confirm(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("Are you sure you want to return to ").
			BlueText().ShowMap(mapId).
			BlackText().AddText("? Alright, we'll have to get going fast. Do you want to head back to ").
			ShowMap(mapId).AddText(" now?")
		return SendYesNoExit(l, c, m.String(), r.Warp(mapId), r.OtherBusiness, r.OtherBusiness)
	}
}

func (r Pison) OtherBusiness(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You must have some business to take care of here. It's not a bad idea to take some rest at ").
		ShowMap(_map.FlorinaBeach).
		AddText(" Look at me; I love it here so much that I wound up living here. Hahaha anyway, talk to me when you feel like going back.")
	return SendOk(l, c, m.String())
}

func (r Pison) Warp(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		character.ClearSavedLocation(l)(c.CharacterId, "FLORINA")
		err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return nil
	}
}
