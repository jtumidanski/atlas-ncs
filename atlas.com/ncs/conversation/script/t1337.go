package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// T1337 is located in New Leaf City Town Street - New Leaf City - Town Center (600000000)
type T1337 struct {
}

func (r T1337) NPCId() uint32 {
	return npc.T1337
}

func (r T1337) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The patrol in New Leaf City is always ready. No creatures are able to break through to the city.")
	return SendOk(l, c, m.String())
}
