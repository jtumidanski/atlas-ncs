package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DocumentRoll is located in Ludibrium - Chloe's House (220000304)
type DocumentRoll struct {
}

func (r DocumentRoll) NPCId() uint32 {
	return npc.DocumentRoll
}

func (r DocumentRoll) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("A document roll on the ground.")
	return script.SendOk(l, c, m.String())
}
