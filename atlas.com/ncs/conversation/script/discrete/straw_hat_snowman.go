package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// StrawHatSnowman is located in Hidden Street - Happyville (209000000)
type StrawHatSnowman struct {
}

func (r StrawHatSnowman) NPCId() uint32 {
	return npc.StrawHatSnowman
}

func (r StrawHatSnowman) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("We have a beautiful christmas tree.").NewLine().
		AddText("Do you want to see/decorate it?")
	return script.SendYesNo(l, span, c, m.String(), r.Warp, script.Exit())
}

func (r StrawHatSnowman) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.TheHillOfChristmas3, 0)(l, span, c)
}
