package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type Mia2 struct {
}

func (r Mia2) NPCId() uint32 {
	return npc.Mia2
}

func (r Mia2) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hi, I'm ").ShowNPC(npc.Mia2).AddText(".")
	return SendOk(l, c, m.String())
}
