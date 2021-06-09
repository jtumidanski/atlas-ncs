package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// HawkeyeStatue is located in Hidden Street - Quiet Ereve (913030000)
type HawkeyeStatue struct {
}

func (r HawkeyeStatue) NPCId() uint32 {
	return npc.HawkeyeStatue
}

func (r HawkeyeStatue) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("...")
	return script.SendOk(l, c, m.String())
}
