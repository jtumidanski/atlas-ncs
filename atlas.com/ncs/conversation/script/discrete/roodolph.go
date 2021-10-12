package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Roodolph is located in Hidden Street - Extra Frosty Snow Zone (209080000) and Hidden Street - Happyville (209000000)
type Roodolph struct {
}

func (r Roodolph) NPCId() uint32 {
	return npc.Roodolph
}

func (r Roodolph) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if c.MapId == _map.Happyville {
		m := message.NewBuilder().
			AddText("Do you wish to head to where the ").
			BlueText().AddText("Snow Sprinkler").
			BlackText().AddText(" is?")
		return script.SendYesNo(l, span, c, m.String(), script.WarpById(_map.ExtraFrostySnowZone, 0), r.WhenYouWantTo)
	} else if c.MapId == _map.ExtraFrostySnowZone {
		m := message.NewBuilder().
			AddText("Do you wish to return to Happyville?")
		return script.SendYesNo(l, span, c, m.String(), script.WarpById(_map.Happyville, 0), r.WhenYouWantTo)
	}
	m := message.NewBuilder().AddText("You Alright?")
	return script.SendOk(l, span, c, m.String())
}

func (r Roodolph) WhenYouWantTo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Talk to me again when you want to.")
	return script.SendOk(l, span, c, m.String())
}
