package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type Mia2 struct {
}

func (r Mia2) NPCId() uint32 {
	return npc.Mia2
}

func (r Mia2) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hi, I'm ").ShowNPC(npc.Mia2).AddText(".")
	return script.SendOk(l, c, m.String())
}
