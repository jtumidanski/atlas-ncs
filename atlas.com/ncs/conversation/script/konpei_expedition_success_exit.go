package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// KonpeiExpeditionSuccessExit is located in Zipangu - Near the Hideout (Beautiful Sky) (801040101)
type KonpeiExpeditionSuccessExit struct {
}

func (r KonpeiExpeditionSuccessExit) NPCId() uint32 {
	return npc.KonpeiExpeditionSuccessExit
}

func (r KonpeiExpeditionSuccessExit) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Ah, The Boss has been defeated. What a happy day this turns out to be! Congratulations, everyone. Follow this way back to town.")
	return SendNext(l, c, m.String(), Warp(_map.ShowaTown))
}
