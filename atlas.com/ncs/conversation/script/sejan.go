package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Sejan is located in Ariant - The Town of Ariant (260000200)
type Sejan struct {
}

func (r Sejan) NPCId() uint32 {
	return npc.Sejan
}

func (r Sejan) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The light and dark always coexist...")
	return SendOk(l, c, m.String())
}
