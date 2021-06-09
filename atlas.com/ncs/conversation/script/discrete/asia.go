package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Asia is located in Neo City - <Year 2503> Air Battleship Bow (240070600)
type Asia struct {
}

func (r Asia) NPCId() uint32 {
	return npc.Asia
}

func (r Asia) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.QuestStarted(l)(c.CharacterId, 3749) {
		return r.FilledWithDespair(l, c)
	} else {
		return r.SeeMySister(l, c)
	}
}

func (r Asia) SeeMySister(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("We've already located the enemy's ultimate weapon! Follow along the ship's bow area ahead and you will find my sister ").
		BlueText().ShowNPC(npc.Ashura).
		BlackText().AddText(". Report to her for further instructions on the mission.")
	return script.SendOk(l, c, m.String())
}

func (r Asia) FilledWithDespair(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("The future is filled with despair.")
	return script.SendOk(l, c, m.String())
}
