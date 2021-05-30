package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// KingPepe is located in Mushroom Castle - Wedding Hall (106021600)
type KingPepe struct {
}

func (r KingPepe) NPCId() uint32 {
	return npc.KingPepe
}

func (r KingPepe) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Let the ceremony begins, we cannot let the masses waiting! Hem~hem~heeh~~")
	return SendOk(l, c, m.String())
}
