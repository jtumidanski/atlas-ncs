package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Yuris is located in Ellin Forest - Altaire Camp (300000000)
type Yuris struct {
}

func (r Yuris) NPCId() uint32 {
	return npc.Yuris
}

func (r Yuris) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("My name is ").
		ShowNPC(npc.Yuris).
		AddText("... As you can see, I am a fairy. People tell me I do not act fairy-like, but... I like making things out of metal objects. Shhh, don't tell this to anyone, but I also like MMA.")
	return SendOk(l, c, m.String())
}
