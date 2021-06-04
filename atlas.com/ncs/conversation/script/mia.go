package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Mia is located in Victoria Road - Henesys (100000000)
type Mia struct {
}

func (r Mia) NPCId() uint32 {
	return npc.Mia
}

func (r Mia) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hi, I'm ").ShowNPC(npc.Mia).AddText(".")
	return SendOk(l, c, m.String())
}
