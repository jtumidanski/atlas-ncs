package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// ThreeRefugees is located in Black Road - Ready to Leave (914000100)
type ThreeRefugees struct {
}

func (r ThreeRefugees) NPCId() uint32 {
	return npc.ThreeRefugees
}

func (r ThreeRefugees) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("We are departing to ").
		BlueText().AddText("Victoria Island").
		BlackText().AddText(" briefly. I've heard the ").
		RedText().AddText("Black Magician").
		BlackText().AddText(" himself cannot take that place on his grasp yet, thanks to ").
		BlueText().AddText("the seal that has been casted on that area").
		BlackText().AddText(". We pray for their safety, but if fortune does not favor the Heroes, at least we will be safe once we reach the continent.")
	return SendOk(l, c, m.String())
}
