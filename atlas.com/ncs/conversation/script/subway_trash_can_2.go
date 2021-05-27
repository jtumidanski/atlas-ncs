package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// SubwayTrashCan2 is located in Kerning City Subway - Line 1 <Area 1> (103000101)
type SubwayTrashCan2 struct {
}

func (r SubwayTrashCan2) NPCId() uint32 {
	return npc.SubwayTrashCan2
}

func (r SubwayTrashCan2) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Just a trash can sitting there.")
	return SendOk(l, c, m.String())
}
