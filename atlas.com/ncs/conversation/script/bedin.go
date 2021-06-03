package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Bedin is located in Zenumist Research Institute - Lab - 1st Floor Hallway (261010000)
type Bedin struct {
}

func (r Bedin) NPCId() uint32 {
	return npc.Bedin
}

func (r Bedin) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Make it clear of your position! Are you Zenumist or Alcadno?")
	return SendOk(l, c, m.String())
}
