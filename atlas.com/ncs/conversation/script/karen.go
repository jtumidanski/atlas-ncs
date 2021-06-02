package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Karen is located in Hidden Street - Time Control Room (222020400)
type Karen struct {
}

func (r Karen) NPCId() uint32 {
	return npc.Karen
}

func (r Karen) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Up ahead is the mysterious forest of ").
		BlueText().AddText("Ellin").
		BlackText().AddText(". Go through ").
		BlueText().AddText("the portal of time").
		BlackText().AddText(" if you are ready to unveil some of the mysteries of the past of Victoria Island, as how it used to be at it's dawn.")
	return SendOk(l, c, m.String())
}
