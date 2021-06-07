package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// YokoYoko is located in Zipangu - Showa Street Market (801000300)
type YokoYoko struct {
}

func (r YokoYoko) NPCId() uint32 {
	return npc.YokoYoko
}

func (r YokoYoko) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The quality of the movies they are launching these days are impressive!")
	return SendOk(l, c, m.String())
}
