package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Ridley is located in Crimsonwood Mountain - Crimsonwood Keep (610020006)
type Ridley struct {
}

func (r Ridley) NPCId() uint32 {
	return npc.Ridley
}

func (r Ridley) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if character.IsLevel(l)(c.CharacterId, 100) {
		m := message.NewBuilder().
			AddText("Expeditions are frequently being held inside the Crimsonwood Keep by adventurers like you, where many people from many parties cooperate together, solving puzzles therein and taking down strong enemies, being able to get many prizes in the process. To find more info about this, go ahead inside the keep at the top-right room there.")
		return script.SendOk(l, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Inside the Keep, expeditions can be formed to attempt the Crimsonwood Keep PQ, which requires maplers from level 100 or more. It seems you are not suitable for attempting it yet, train some more if you want to attempt it.")
	return script.SendOk(l, c, m.String())
}
