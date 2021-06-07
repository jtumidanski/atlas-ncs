package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// LittleSuzy is located in New Leaf City Town Street - New Leaf City - Town Center (600000000)
type LittleSuzy struct {
}

func (r LittleSuzy) NPCId() uint32 {
	return npc.LittleSuzy
}

func (r LittleSuzy) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Have you heard the fantastic Jack Masque appeared around the city these days? That is sooooo nice!")
	return SendOk(l, c, m.String())
}
