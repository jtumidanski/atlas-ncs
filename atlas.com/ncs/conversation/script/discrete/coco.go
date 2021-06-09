package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type Coco struct {
}

func (r Coco) NPCId() uint32 {
	return npc.Coco
}

func (r Coco) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hi, I'm ").
		BlueText().ShowNPC(npc.Coco).
		BlackText().AddText(".")
	return script.SendOk(l, c, m.String())
}
