package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// StrangeLookingStatue is located in Hidden Street - Puppeteer's Secret Passage (910510100)
type StrangeLookingStatue struct {
}

func (r StrangeLookingStatue) NPCId() uint32 {
	return npc.StrangeLookingStatue
}

func (r StrangeLookingStatue) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.Hello(l, span, c)
}

func (r StrangeLookingStatue) Hello(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Will you exit this trial?")
	return script.SendYesNo(l, span, c, m.String(), r.Warp, script.Exit())
}

func (r StrangeLookingStatue) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.PuppeteersHidingPlace, 2)(l, span, c)
}
