package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// PrinceGiuseppe is located in Mushroom Castle - Wedding Hall (106021600)
type PrinceGiuseppe struct {
}

func (r PrinceGiuseppe) NPCId() uint32 {
	return npc.PrinceGiuseppe
}

func (r PrinceGiuseppe) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hey guys, what's going on? I'm already getting married at this age? But I'm only a child!!! How comes?")
	return SendOk(l, c, m.String())
}
