package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Kiridu is located in Empress' Road - Kiridu's Hatchery (130010220)
type Kiridu struct {
}

func (r Kiridu) NPCId() uint32 {
	return npc.Kiridu
}

func (r Kiridu) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Yo. I am ").
		ShowNPC(npc.Kiridu).
		AddText(", in charge of mount raising and training for the Cygnus Knights' of Ereve!")
	return SendOk(l, c, m.String())
}
