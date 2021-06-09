package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
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

func (r EncryptedSlateOfTheSquad) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.HasItem(l)(c.CharacterId, item.CertificateOfTheDragonSquad) {
		m := message.NewBuilder().AddText("Do you want to access ").
			BlueText().ShowMap(_map.EntranceToHorntailsCave).
			BlackText().AddText(" right now?")
		return script.SendYesNo(l, c, m.String(), script.WarpById(_map.EntranceToHorntailsCave, 0), script.Exit())
	}
	return r.MustProveValor(l, c)
}

func (r EncryptedSlateOfTheSquad) MustProveValor(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Those who don't have the ").
		RedText().ShowItemName1(item.CertificateOfTheDragonSquad).
		BlackText().AddText(" must prove their valor before challenging ").
		BlueText().AddText("Horntail").
		BlackText().AddText(".")
	return script.SendOk(l, c, m.String())
}
