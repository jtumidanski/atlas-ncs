package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Areda is located in Ariant Castle - King's Room (260000303)
type Areda struct {
}

func (r Areda) NPCId() uint32 {
	return npc.Areda
}

func (r Areda) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("NO! Abdullah, I said 17 bedrooms, and 23 bathrooms! CALL THE CONSTRUCTION COMPANY AND CHANGE IT!")
	return script.SendOk(l, c, m.String())
}
