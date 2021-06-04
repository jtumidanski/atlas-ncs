package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Vikin is located in Victoria Road - Lith Harbor (104000000)
type Vikin struct {
}

func (r Vikin) NPCId() uint32 {
	return npc.Vikin
}

func (r Vikin) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hey hey!!! Find the Treasure Scroll! I lost the map").NewLine().
		AddText("somewhere and I can't leave without it.")
	return SendOk(l, c, m.String())
}
