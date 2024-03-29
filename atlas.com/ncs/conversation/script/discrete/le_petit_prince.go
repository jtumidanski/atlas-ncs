package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// LePetitPrince is located in The Burning Sands - Dry Desert (260010500)
type LePetitPrince struct {
}

func (r LePetitPrince) NPCId() uint32 {
	return npc.LePetitPrince
}

func (r LePetitPrince) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Home is so boring... my parents ignore me so much it's unbearable. And ever since we moved from ").
		RedText().AddText("Ariant").
		BlackText().AddText(", they've been trying to get a new palace built so they don't have to live outdoors. But I love the outdoors...")
	return script.SendOk(l, span, c, m.String())
}
