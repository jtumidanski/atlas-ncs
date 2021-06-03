package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// AbdullahVIII is located in Ariant Castle - King's Room (260000303)
type AbdullahVIII struct {
}

func (r AbdullahVIII) NPCId() uint32 {
	return npc.AbdullahVIII
}

func (r AbdullahVIII) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Yawnnnn~!")
	return SendOk(l, c, m.String())
}
