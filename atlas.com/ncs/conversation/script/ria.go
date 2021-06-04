package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Ria is located in Victoria Road - Ellinia (101000000)
type Ria struct {
}

func (r Ria) NPCId() uint32 {
	return npc.Ria
}

func (r Ria) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hi, I'm ").ShowNPC(npc.Ria).AddText(".")
	return SendOk(l, c, m.String())
}
