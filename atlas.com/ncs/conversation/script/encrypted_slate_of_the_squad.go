package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// EncryptedSlateOfTheSquad is located in Cave of Life - Cave Entrance (240050000)
type EncryptedSlateOfTheSquad struct {
}

func (r EncryptedSlateOfTheSquad) NPCId() uint32 {
	return npc.EncryptedSlateOfTheSquad
}

func (r EncryptedSlateOfTheSquad) Initial(l logrus.FieldLogger, c Context) State {
	if character.HasItem(l)(c.CharacterId, item.CertificateOfTheDragonSquad) {
		m := message.NewBuilder().AddText("Do you want to access ").
			BlueText().ShowMap(_map.EntranceToHorntailsCave).
			BlackText().AddText(" right now?")
		return SendYesNo(l, c, m.String(), r.Warp, Exit())
	}
	return r.MustProveValor(l, c)
}

func (r EncryptedSlateOfTheSquad) MustProveValor(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Those who don't have the ").
		RedText().ShowItemName1(item.CertificateOfTheDragonSquad).
		BlackText().AddText(" must prove their valor before challenging ").
		BlueText().AddText("Horntail").
		BlackText().AddText(".")
	return SendOk(l, c, m.String())
}

func (r EncryptedSlateOfTheSquad) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.EntranceToHorntailsCave, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.EntranceToHorntailsCave, c.NPCId)
	}
	return Exit()(l, c)
}
