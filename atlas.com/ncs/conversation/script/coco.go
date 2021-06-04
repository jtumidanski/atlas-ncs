package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type Coco struct {
}

func (r Coco) NPCId() uint32 {
	return npc.Coco
}

func (r Coco) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hi, I'm ").
		BlueText().ShowNPC(npc.Coco).
		BlackText().AddText(".")
	return SendOk(l, c, m.String())
}
