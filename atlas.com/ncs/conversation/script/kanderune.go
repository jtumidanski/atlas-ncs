package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Kanderune is located in Ellin Forest - Boulder Mountain Entrance (300010400)
type Kanderune struct {
}

func (r Kanderune) NPCId() uint32 {
	return npc.Kanderune
}

func (r Kanderune) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Hmmm! For you to make your way here, far away from the Camp, you must be one strong individual. Let's explore new areas and find a place to establish our own town!!")
	return SendOk(l, c, m.String())
}
