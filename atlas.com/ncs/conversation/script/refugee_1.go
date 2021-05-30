package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Refugee1 is located in Black Road - Ready to Leave (914000100)
type Refugee1 struct {
}

func (r Refugee1) NPCId() uint32 {
	return npc.Refugee1
}

func (r Refugee1) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The ").
		RedText().AddText("Black Magician").
		BlackText().AddText("'s forces approaches here in an unstoppable pace... We have no other way than to flee this area now, leaving our home behind. Oh, the tragedy!")
	return SendOk(l, c, m.String())
}
