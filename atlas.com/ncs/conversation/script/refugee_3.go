package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Refugee3 is located in Black Road - Ready to Leave (914000100)
type Refugee3 struct {
}

func (r Refugee3) NPCId() uint32 {
	return npc.Refugee3
}

func (r Refugee3) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I hope for this travel to be a safe one, and that we get to live on a more peaceful place there... Hey, darling, let's go.")
	return SendOk(l, c, m.String())
}