package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// YoungAthenaPierce2 is located in 
type YoungAthenaPierce2 struct {
}

func (r YoungAthenaPierce2) NPCId() uint32 {
	return npc.YoungAthenaPierce2
}

func (r YoungAthenaPierce2) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("It is been a while since we left Ossyria to avoid the Black Magician. If not for the world tree, I do not know where we would have been. I have been trying to establish myself here, but that is not easy. I wonder how things are like back home.")
	return SendOk(l, c, m.String())
}
