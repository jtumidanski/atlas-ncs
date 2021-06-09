package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Rius is located in Ellin Forest - Mossy Tree Forest Trail (300010300)
type Rius struct {
}

func (r Rius) NPCId() uint32 {
	return npc.Rius
}

func (r Rius) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("The magic of this forest is amazing...")
	return script.SendOk(l, c, m.String())
}
