package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// KinoKonoko is located in Zipangu - Mushroom Shrine (800000000)
type KinoKonoko struct {
}

func (r KinoKonoko) NPCId() uint32 {
	return npc.KinoKonoko
}

func (r KinoKonoko) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Musssshhhhroooom Shrine~~~")
	return script.SendOk(l, span, c, m.String())
}
