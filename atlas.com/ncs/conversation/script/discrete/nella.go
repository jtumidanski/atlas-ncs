package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Nella is located in Kerning PQ
type Nella struct {
}

func (r Nella) NPCId() uint32 {
	return npc.Nella
}

func (r Nella) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if c.MapId == _map.FirstAccompanimentExit {
		return r.ExitTheExit(l, span, c)
	}
	if c.MapId == _map.FirstAccompanimentBonus {
		return r.ExitBonus(l, span, c)
	}
	return r.ExitQuestMaps(l, span, c)
}

func (r Nella) ExitQuestMaps(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Once you leave the map, you'll have to restart the whole quest if you want to try it again.  Do you still want to leave this map?")
	return script.SendYesNo(l, span, c, m.String(), r.WarpToExit, script.Exit())
}

func (r Nella) ExitBonus(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Are you ready to leave this map?")
	return script.SendYesNo(l, span, c, m.String(), r.WarpToExit, script.Exit())
}

func (r Nella) ExitTheExit(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("To return back to the city, follow this way.")
	return script.SendNext(l, span, c, m.String(), script.Warp(_map.KerningCity))
}

func (r Nella) WarpToExit(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpByName(_map.FirstAccompanimentExit, "st00")(l, span, c)
}
