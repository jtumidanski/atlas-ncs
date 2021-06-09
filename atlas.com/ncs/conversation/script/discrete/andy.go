package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Andy is located in Tera Forest   - Tera Forest Time Gate (240070000)
type Andy struct {
}

func (r Andy) NPCId() uint32 {
	return npc.Andy
}

func (r Andy) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hi, I am Andy, the time traveler from a not so distant future. I have come to avert the creation of machines by the greedy people of this time. They went berserk on my time and consumed everything to dust. I must stop it at any cost!")
	return script.SendOk(l, c, m.String())
}
