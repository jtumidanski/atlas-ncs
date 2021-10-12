package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Egnet is located in Orbis - Station <To Ariant> (200000152)
type Egnet struct {
}

func (r Egnet) NPCId() uint32 {
	return npc.Egnet
}

func (r Egnet) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you wish to leave the genie?")
	return script.SendYesNo(l, span, c, m.String(), r.Alright, script.Exit())
}

func (r Egnet) Alright(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Alright, see you next time. Take care.")
	return script.SendNext(l, span, c, m.String(), script.WarpById(_map.StationToAriant, 0))
}
