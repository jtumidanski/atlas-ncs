package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// EckhartStatue is located in Hidden Street - Quiet Ereve (913030000)
type EckhartStatue struct {
}

func (r EckhartStatue) NPCId() uint32 {
	return npc.EckhartStatue
}

func (r EckhartStatue) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("...")
	return script.SendOk(l, c, m.String())
}
