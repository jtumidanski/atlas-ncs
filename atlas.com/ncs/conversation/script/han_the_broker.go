package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// HanTheBroker is located in Sunset Road - Magatia (261000000)
type HanTheBroker struct {
}

func (r HanTheBroker) NPCId() uint32 {
	return npc.HanTheBroker
}

func (r HanTheBroker) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hm... Don't doubt it because I'm a back street broker. Dealing with me is trust... I keep up my credit.")
	return SendOk(l, c, m.String())
}
