package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Refugee4 is located in Black Road - Ready to Leave (914000100)
type Refugee4 struct {
}

func (r Refugee4) NPCId() uint32 {
	return npc.Refugee4
}

func (r Refugee4) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Alright, embarking in...")
	return SendOk(l, c, m.String())
}
