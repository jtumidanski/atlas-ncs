package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Dunamis is located in Hidden Street - Cave of Black Witches (924010000)
type Dunamis struct {
}

func (r Dunamis) NPCId() uint32 {
	return npc.Dunamis
}

func (r Dunamis) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I am Advanced Knight ").
		BlueText().ShowNPC(npc.Dunamis).
		BlackText().AddText(". Thanks to your bravery I and all of Ereve have been rescued from the grasps of Eleanor. By the kindness of our Empress, well battled!")
	return script.SendOk(l, c, m.String())
}
