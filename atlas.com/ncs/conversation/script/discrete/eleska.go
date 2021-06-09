package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Eleska is located in Ariant - The Town of Ariant (260000200)
type Eleska struct {
}

func (r Eleska) NPCId() uint32 {
	return npc.Eleska
}

func (r Eleska) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Stay away from me, if you don't want any danger.")
	return script.SendOk(l, c, m.String())
}
