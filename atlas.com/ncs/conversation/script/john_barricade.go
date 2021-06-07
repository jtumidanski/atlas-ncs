package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// JohnBarricade is located in Bigger Ben - Lobby (600020000)
type JohnBarricade struct {
}

func (r JohnBarricade) NPCId() uint32 {
	return npc.JohnBarricade
}

func (r JohnBarricade) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The patrol in New Leaf City is always ready. No creatures are able to break through to the city.")
	return SendOk(l, c, m.String())
}
