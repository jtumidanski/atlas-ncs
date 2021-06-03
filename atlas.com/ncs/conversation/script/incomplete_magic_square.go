package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// IncompleteMagicSquare is located in Zenumist Research Institute - Lab - Unit 202 (261010102)
type IncompleteMagicSquare struct {
}

func (r IncompleteMagicSquare) NPCId() uint32 {
	return npc.IncompleteMagicSquare
}

func (r IncompleteMagicSquare) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("This chalkboard has some hard-founded studies annotated on it...")
	return SendOk(l, c, m.String())
}
