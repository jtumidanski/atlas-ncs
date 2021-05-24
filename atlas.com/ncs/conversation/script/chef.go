package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Chef is located in Victoria Road - Lith Harbor (104000000)
type Chef struct {
}

func (r Chef) NPCId() uint32 {
	return npc.Chef
}

func (r Chef) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r Chef) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hi, I'm ").
		BlueText().ShowNPC(r.NPCId()).
		BlackText().AddText(". Nice to meet you.")
	return SendNext(l, c, m.String(), Exit())
}
