package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Byron is located in Ariant - The Town of Ariant (260000200)
type Byron struct {
}

func (r Byron) NPCId() uint32 {
	return npc.Byron
}

func (r Byron) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I thought ").
		RedText().AddText("Ariant").
		BlackText().AddText(" was amazing, but Valvendale is just SUPERB! There are so many different landscapes and forms of life. Like ").
		BlueText().BoldText().AddText("those weird slimes outside of town").
		NormalText().BlackText().AddText("... those are just WEIRD!")
	return script.SendOk(l, c, m.String())
}
