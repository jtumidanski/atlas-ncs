package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Shadrion is located in Ellin Forest - Western Region of Mossy Tree Forest 2 (300010200)
type Shadrion struct {
}

func (r Shadrion) NPCId() uint32 {
	return npc.Shadrion
}

func (r Shadrion) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Being young doesn't mean I'm any different from those guys. I'll show them!")
	return script.SendOk(l, c, m.String())
}