package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// PuroToRien is located in Snow Island - To Rien (200090060)
type PuroToRien struct {
}

func (r PuroToRien) NPCId() uint32 {
	return npc.PuroToRien
}

func (r PuroToRien) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Ahhhh, this is so boring... The whale is controlling the ship so i'm left with nothing to do but look up and stare at the clouds.")
	return script.SendOk(l, span, c, m.String())
}
