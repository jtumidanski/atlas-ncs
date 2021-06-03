package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Jiyur is located in Ariant - The Town of Ariant (260000200)
type Jiyur struct {
}

func (r Jiyur) NPCId() uint32 {
	return npc.Jiyur
}

func (r Jiyur) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("I miss my sister... She's always working at the palace as the servant and I only get to see her on Sundays. The King and Queen are so selfish.")
	return SendOk(l, c, m.String())
}
