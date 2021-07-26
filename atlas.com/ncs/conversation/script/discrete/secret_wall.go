package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// SecretWall is located in Ariant - The Town of Ariant (260000200)
type SecretWall struct {
}

func (r SecretWall) NPCId() uint32 {
	return npc.SecretWall
}

func (r SecretWall) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3927) {
		quest.SetProgress(l)(c.CharacterId, 3927, 0, 1)
		m := message.NewBuilder().AddText("If I had an iron hammer and a dagger, a bow and an arrow...")
		return script.SendOk(l, c, m.String())
	}
	return script.Exit()(l, c)
}
