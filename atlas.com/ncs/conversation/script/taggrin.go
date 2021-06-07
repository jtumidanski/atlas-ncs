package script

import (
	"atlas-ncs/character"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Taggrin is located in Phantom Forest - Dead Man's Gorge (610010004)
type Taggrin struct {
}

func (r Taggrin) NPCId() uint32 {
	return npc.Taggrin
}

func (r Taggrin) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestProgressInt(l)(c.CharacterId, 8224, 0) == 2 {
		m := message.NewBuilder().AddText("Well met, fellow clan member. If you need anything we can be of help, try talking to one of our members.")
		return SendOk(l, c, m.String())
	}

	m := message.NewBuilder().AddText("Hello there, stranger. We are the renowned Raven Claw clan of mercenaries, and I'm their leader.")
	return SendOk(l, c, m.String())
}
