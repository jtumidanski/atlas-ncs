package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type Sage struct {
}

func (r Sage) NPCId() uint32 {
	return npc.Sage
}

func (r Sage) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if c.MapId == _map.ValleyOfHeroes1 {
		m := message.NewBuilder().
			AddText("O, brave adventurer. Just by reaching this spot, you are truly distinct among the masses, congratulations. However, ").
			RedText().AddText("pay heed").
			BlackText().AddText(": on the path ahead, which leads to the mighty fortress of ").
			BlueText().AddText("Crimsonwood Keep").
			BlackText().AddText(", ").
			RedText().AddText("deadly Menhirs").
			BlackText().AddText(" are deployed as traps for those unaware of the dangers ahead. ").
			RedText().AddText("One hit from it is enough to take you down").
			BlackText().AddText(", so beware. If you aim to reach the Keep, follow the trail ahead carefully.")
		return script.SendOk(l, c, m.String())
	} else if c.MapId == _map.ThePathOfStrength {
		m := message.NewBuilder().
			AddText("You seem worthy now to receive a hint for what lies ahead. Once inside the main room of the Keep, make sure you remember the layout of the statue you see there. That's it.")
		return script.SendOk(l, c, m.String())
	} else if c.MapId == _map.ThePathOfPeril {
		m := message.NewBuilder().
			AddText("You seem worthy now to receive a hint for what lies ahead. Devices known as Sigils are activated by detection when some skills of certain jobs are activated nearby, make sure your team is made whole for when the time comes. That's it.")
		return script.SendOk(l, c, m.String())
	} else {
		m := message.NewBuilder().
			AddText("So far your progress is splendid, good job. However, to make it to the Keep, you must face and accomplish this ordeal, carry on.")
		return script.SendOk(l, c, m.String())
	}
}
