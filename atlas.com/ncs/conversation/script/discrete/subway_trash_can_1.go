package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// SubwayTrashCan1 is located in Kerning City Subway - Line 1 <Area 1> (103000101)
type SubwayTrashCan1 struct {
}

func (r SubwayTrashCan1) NPCId() uint32 {
	return npc.SubwayTrashCan1
}

func (r SubwayTrashCan1) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Just a trash can sitting there.")
	return script.SendOk(l, c, m.String())
}
