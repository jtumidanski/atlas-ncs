package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// PracticeChart is located in Mu Lung - Practice Field : Beginner (250020000)
type PracticeChart struct {
}

func (r PracticeChart) NPCId() uint32 {
	return npc.PracticeChart
}

func (r PracticeChart) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Amateurs train on this map. Adepts train on the next. Professionals train on the last, where the boss will be awaiting.")
	return SendOk(l, c, m.String())
}
