package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// MrPickall is located in Victoria Road - Kerning City (103000000)
type MrPickall struct {
}

func (r MrPickall) NPCId() uint32 {
	return npc.MrPickall
}

func (r MrPickall) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Should you have a locked box you want to open, bring it to me.")
	return script.SendOk(l, span, c, m.String())
}
