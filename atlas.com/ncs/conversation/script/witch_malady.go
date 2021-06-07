package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// WitchMalady is located in Phantom Forest - Haunted House (682000000)
type WitchMalady struct {
}

func (r WitchMalady) NPCId() uint32 {
	return npc.WitchMalady
}

func (r WitchMalady) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Young one, you seem very proud of yourself, don't you? Can you face the real nightmare that is this place? If you think you can do it, then go ahead, ehehehehehe.")
	return SendOk(l, c, m.String())
}
