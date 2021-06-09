package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Tombstone is located in Ludibrium - Deep Inside the Clocktower (220080000)
type Tombstone struct {
}

func (r Tombstone) NPCId() uint32 {
	return npc.Tombstone
}

func (r Tombstone) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("For those capable of great feats and bearers of an unwavering resolve, the ").
		BlueText().AddText("final destination").
		BlackText().AddText(" lies ahead past the gate. The Machine Room accepts only ").
		RedText().AddText("one party at a time").
		BlackText().AddText(", so make sure your party is ready when crossing the gate.")
	return script.SendOk(l, c, m.String())
}
