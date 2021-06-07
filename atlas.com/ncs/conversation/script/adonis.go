package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Adonis is located in El Nath - El Nath (211000000)
type Adonis struct {
}

func (r Adonis) NPCId() uint32 {
	return npc.Adonis
}

func (r Adonis) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I came from far-away places looking for people powerful enough to join my expedition against the evil that lays waste on this land. Are you, by any chance, one of those people?")
	return SendOk(l, c, m.String())
}
