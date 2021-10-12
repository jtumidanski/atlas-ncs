package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// ScarfSnowman is located in Hidden Street - The Hill of Christmas (209000001-209000015)
type ScarfSnowman struct {
}

func (r ScarfSnowman) NPCId() uint32 {
	return npc.ScarfSnowman
}

func (r ScarfSnowman) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("So, are you ready to head out of here?")
	return script.SendYesNo(l, span, c, m.String(), r.Warp, script.Exit())
}

func (r ScarfSnowman) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.Happyville, 5)(l, span, c)
}
