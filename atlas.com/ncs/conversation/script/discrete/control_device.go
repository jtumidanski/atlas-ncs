package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// ControlDevice is located in Hidden Street - Authorized Personnel Only (261020401)
type ControlDevice struct {
}

func (r ControlDevice) NPCId() uint32 {
	return npc.ControlDevice
}

func (r ControlDevice) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("This control device seems to be monitoring something...")
	return script.SendOk(l, span, c, m.String())
}
