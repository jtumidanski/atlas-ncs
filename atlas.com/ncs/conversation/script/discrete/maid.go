package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Maid is located in Haunted House - Foyer (682000100)
type Maid struct {
}

func (r Maid) NPCId() uint32 {
	return npc.Maid
}

func (r Maid) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hello and welcome, dear guest. The Master has prepared some wonderful games for you to enjoy tonight.")
	return script.SendOk(l, c, m.String())
}
