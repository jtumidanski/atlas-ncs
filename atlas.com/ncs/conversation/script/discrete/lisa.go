package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Lisa is located in Orbis - Orbis (200000000)
type Lisa struct {
}

func (r Lisa) NPCId() uint32 {
	return npc.Lisa
}

func (r Lisa) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Can you help me find the Ancient Book? I lost it somewhere in El Nath...")
	return script.SendOk(l, c, m.String())
}
