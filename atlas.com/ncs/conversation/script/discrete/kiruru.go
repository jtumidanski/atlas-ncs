package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Kiruru is located in Empress' Road - To Ellinia (200090055)
type Kiruru struct {
}

func (r Kiruru) NPCId() uint32 {
	return npc.Kiruru
}

func (r Kiruru) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("The weather is so nice. At this rate, we should arrive in no time....")
	return script.SendOk(l, c, m.String())
}
