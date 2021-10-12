package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Shulynch2 is located in Hidden Street - Looking for Delli 1 (925010000)
type Shulynch2 struct {
}

func (r Shulynch2) NPCId() uint32 {
	return npc.Shulynch2
}

func (r Shulynch2) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		BlueText().ShowNPC(npc.Delli).
		BlackText().AddText(" must be some way up this cliff, according to our latest reports... Or are you saying you want to ").
		RedText().AddText("leave here").
		BlackText().AddText("?")
	return script.SendNext(l, span, c, m.String(), r.Warp)
}

func (r Shulynch2) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.TrainingRoom, 0)(l, span, c)
}
