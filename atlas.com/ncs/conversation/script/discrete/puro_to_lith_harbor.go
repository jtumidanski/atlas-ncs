package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// PuroToLithHarbor is located in Snow Island - To Lith Harbor (200090070)
type PuroToLithHarbor struct {
}

func (r PuroToLithHarbor) NPCId() uint32 {
	return npc.PuroToLithHarbor
}

func (r PuroToLithHarbor) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("The current is serene, which means we may arrive in lith harbor earlier than expected.")
	return script.SendOk(l, c, m.String())
}
