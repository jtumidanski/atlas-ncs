package script

import (
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// KnightArmor is located in Sharenian - Hall of the Knight (990000400)
type KnightArmor struct {
}

func (r KnightArmor) NPCId() uint32 {
	return npc.KnightArmor
}

func (r KnightArmor) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The plaque translates as follows: ").NewLine().
		AddText("The knights of Sharenian are proud warriors. Their Longinus Spears are both formidable weapons and the key to the castle's defense: By removing them from their platforms at the highest points of this hall, they block off the entrance from invaders.").NewLine().NewLine().
		AddText("Something seems to be etched in English on the side, barely readable: ").NewLine().
		AddText("evil stole spears, chained up behind obstacles. return to top of towers. large spear, grab from higher up.").NewLine().
		AddText("...Obviously whoever figured it out didn't have much time to live. Poor guy.")
	return SendOk(l, c, m.String())
}
