package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type Steward struct {
}

func (r Steward) NPCId() uint32 {
	return npc.Steward
}

func (r Steward) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("At your service, my friend.")
	return script.SendOk(l, c, m.String())
}
