package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// TheGlimmerManNLC is located in New Leaf City Town Street - New Leaf City - Town Center (600000000)
type TheGlimmerManNLC struct {
}

func (r TheGlimmerManNLC) NPCId() uint32 {
	return npc.TheGlimmerManNLC
}

func (r TheGlimmerManNLC) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The patrol in New Leaf City is always ready. No creatures are able to break through to the city.")
	return SendOk(l, c, m.String())
}
