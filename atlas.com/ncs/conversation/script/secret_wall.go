package script

import (
	"atlas-ncs/character"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// SecretWall is located in Ariant - The Town of Ariant (260000200)
type SecretWall struct {
}

func (r SecretWall) NPCId() uint32 {
	return npc.SecretWall
}

func (r SecretWall) Initial(l logrus.FieldLogger, c Context) State {
	if character.QuestStarted(l)(c.CharacterId, 3927) {
		character.SetQuestProgress(l)(c.CharacterId, 3927, 0, 1)
		m := message.NewBuilder().AddText("If I had an iron hammer and a dagger, a bow and an arrow...")
		return SendOk(l, c, m.String())
	}
	return Exit()(l, c)
}
