package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Maed is located in Magatia - Alcadno Society (261000020)
type Maed struct {
}

func (r Maed) NPCId() uint32 {
	return npc.Maed
}

func (r Maed) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Zenumist......I know what they say. They don't like combination the of life with machine. But it is about being fearful of machine only. Seeking Pure Alchemy won't achieve anything.")
	return SendOk(l, c, m.String())
}
