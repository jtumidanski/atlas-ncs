package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// CornerOfTheMagicLibrary2 is located in Hidden Street - Magic Library (910110000)
type CornerOfTheMagicLibrary2 struct {
}

func (r CornerOfTheMagicLibrary2) NPCId() uint32 {
	return npc.CornerOfTheMagicLibrary2
}

func (r CornerOfTheMagicLibrary2) Initial(l logrus.FieldLogger, c Context) State {
	return r.NothingRemarkable(l, c)
}

func (r CornerOfTheMagicLibrary2) NothingRemarkable(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Nothing remarkable here.")
	return SendOk(l, c, m.String())
}
