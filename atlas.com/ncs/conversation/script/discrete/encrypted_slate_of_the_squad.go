package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// EncryptedSlateOfTheSquad is located in Cave of Life - Cave Entrance (240050000)
type EncryptedSlateOfTheSquad struct {
}

func (r EncryptedSlateOfTheSquad) NPCId() uint32 {
	return npc.EncryptedSlateOfTheSquad
}

func (r EncryptedSlateOfTheSquad) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.CertificateOfTheDragonSquad) {
		m := message.NewBuilder().AddText("Do you want to access ").
			BlueText().ShowMap(_map.EntranceToHorntailsCave).
			BlackText().AddText(" right now?")
		return script.SendYesNo(l, span, c, m.String(), script.WarpById(_map.EntranceToHorntailsCave, 0), script.Exit())
	}
	return r.MustProveValor(l, span, c)
}

func (r EncryptedSlateOfTheSquad) MustProveValor(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Those who don't have the ").
		RedText().ShowItemName1(item.CertificateOfTheDragonSquad).
		BlackText().AddText(" must prove their valor before challenging ").
		BlueText().AddText("Horntail").
		BlackText().AddText(".")
	return script.SendOk(l, span, c, m.String())
}
