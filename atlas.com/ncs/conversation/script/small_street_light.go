package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// SmallStreetLight is located in Kerning City Subway - Line 1 <Area 4> (103000105)
type SmallStreetLight struct {
}

func (r SmallStreetLight) NPCId() uint32 {
	return npc.SmallStreetLight
}

func (r SmallStreetLight) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("A small focus of light lighting in the immersive darkness.")
	return SendOk(l, c, m.String())
}
