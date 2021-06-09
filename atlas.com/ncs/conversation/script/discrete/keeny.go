package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Keeny is located in Sunset Road - Magatia (261000000)
type Keeny struct {
}

func (r Keeny) NPCId() uint32 {
	return npc.Keeny
}

func (r Keeny) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Almost all Magatian people are Alchemists. Everyone concentrates on Alchemy. So...Magatia is always quiet.")
	return script.SendOk(l, c, m.String())
}
