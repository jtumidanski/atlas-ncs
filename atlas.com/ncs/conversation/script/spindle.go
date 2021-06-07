package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Spindle is located in Omega Sector - Omega Sector (221000000)
type Spindle struct {
}

func (r Spindle) NPCId() uint32 {
	return npc.Spindle
}

func (r Spindle) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Have you ever come to know about the card game based on MapleStory, the MapleStory iTCG?")
	return SendOk(l, c, m.String())
}
