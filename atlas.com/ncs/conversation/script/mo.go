package script

import (
	"atlas-ncs/character"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Mo is located in Phantom Forest - Dead Man's Gorge (610010004)
type Mo struct {
}

func (r Mo) NPCId() uint32 {
	return npc.Mo
}

func (r Mo) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestCompleted(l)(c.CharacterId, 8224) {
		m := message.NewBuilder().AddText("Hm, at who do you think you are looking at?")
		return SendOk(l, c, m.String())
	}
	npc.OpenShop(l)(c.CharacterId, 9201099)
	return Exit()(l, c)
}
