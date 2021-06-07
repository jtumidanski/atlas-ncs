package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DaveAndIris is located in Singapore - Boat Quay Town (541000000)
type DaveAndIris struct {
}

func (r DaveAndIris) NPCId() uint32 {
	return npc.DaveAndIris
}

func (r DaveAndIris) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("We had our wedding at Amoria, such a beautiful place, and their people are even more amiable. Now our honeymoon on this paradisiac place... Ah, glorious, glorious!")
	return SendOk(l, c, m.String())
}
