package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Karen is located in Hidden Street - Time Control Room (222020400)
type Karen struct {
}

func (r Karen) NPCId() uint32 {
	return npc.Karen
}

func (r Karen) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Up ahead is the mysterious forest of ").
		BlueText().AddText("Ellin").
		BlackText().AddText(". Go through ").
		BlueText().AddText("the portal of time").
		BlackText().AddText(" if you are ready to unveil some of the mysteries of the past of Victoria Island, as how it used to be at it's dawn.")
	return script.SendOk(l, span, c, m.String())
}
