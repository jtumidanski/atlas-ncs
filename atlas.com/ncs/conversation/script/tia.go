package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Tia is located in Victoria Road - Perion (102000000) and  Singapore - CBD (540000000)
type Tia struct {
}

func (r Tia) NPCId() uint32 {
	return npc.Tia
}

func (r Tia) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hi, I'm ").ShowNPC(npc.Tia).AddText(".")
	return SendOk(l, c, m.String())
}
