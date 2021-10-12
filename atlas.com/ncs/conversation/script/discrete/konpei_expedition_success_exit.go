package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// KonpeiExpeditionSuccessExit is located in Zipangu - Near the Hideout (Beautiful Sky) (801040101)
type KonpeiExpeditionSuccessExit struct {
}

func (r KonpeiExpeditionSuccessExit) NPCId() uint32 {
	return npc.KonpeiExpeditionSuccessExit
}

func (r KonpeiExpeditionSuccessExit) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Ah, The Boss has been defeated. What a happy day this turns out to be! Congratulations, everyone. Follow this way back to town.")
	return script.SendNext(l, span, c, m.String(), script.Warp(_map.ShowaTown))
}
