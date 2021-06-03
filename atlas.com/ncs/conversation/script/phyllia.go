package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Phyllia is located in Sunset Road - Magatia (261000000)
type Phyllia struct {
}

func (r Phyllia) NPCId() uint32 {
	return npc.Phyllia
}

func (r Phyllia) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Even though Alchemists and Fairies have antipathy for each other, I can still be in this town because the Zenumist President has been protecting us.")
	return SendOk(l, c, m.String())
}
