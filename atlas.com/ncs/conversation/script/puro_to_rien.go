package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// PuroToRien is located in Snow Island - To Rien (200090060)
type PuroToRien struct {
}

func (r PuroToRien) NPCId() uint32 {
	return npc.PuroToRien
}

func (r PuroToRien) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Ahhhh, this is so boring... The whale is controlling the ship so i'm left with nothing to do but look up and stare at the clouds.")
	return SendOk(l, c, m.String())
}
