package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// WaterFilter is located in The Nautilus - Bedroom (120000202)
type WaterFilter struct {
}

func (r WaterFilter) NPCId() uint32 {
	return npc.WaterFilter
}

func (r WaterFilter) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("I'm bored! Someone come play with me!")
	return SendOk(l, c, m.String())
}
