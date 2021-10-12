package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// MooseExit is located in Hidden Street - On the Way to the Practice Field (924000000), Hidden Street - Moose's Practice Field (924000001), and Hidden Street - Exiting the Practice Field (924000002)
type MooseExit struct {
}

func (r MooseExit) NPCId() uint32 {
	return npc.MooseExit
}

func (r MooseExit) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you want to exit the area? If you quit, you will need to start this task from the scratch.")
	return script.SendYesNo(l, span, c, m.String(), script.WarpByName(_map.ForestCrossroad, "st00"), script.Exit())
}
