package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// LitaLawless is located in New Leaf City Town Street - New Leaf City - Town Center (600000000)
type LitaLawless struct {
}

func (r LitaLawless) NPCId() uint32 {
	return npc.LitaLawless
}

func (r LitaLawless) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("The patrol in New Leaf City is always ready. No creatures are able to break through to the city.")
	return script.SendOk(l, c, m.String())
}
