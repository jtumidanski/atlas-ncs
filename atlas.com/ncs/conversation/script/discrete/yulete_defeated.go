package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

type YuleteDefeated struct {
}

func (r YuleteDefeated) NPCId() uint32 {
	return npc.YuleteDefeated
}

func (r YuleteDefeated) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Defeated... So, that's how Yulete's legacy will reach it's end, oh how woe is this... Hope you guys are happy now, as I will pass my days rotting in a dark cellar. Everything I've done was for the sake of Magatia!! (sob)").NewLine().
		OpenItem(0).AddText("Hey man, come now, cheer up! There were not many damages that couldn't be resolved here. Magatia created these forbidding laws to protect it's people from the undoings a greater power like this would do if it reaches wrong hands. That's not the end for you, accept rehabilitation from the Societies and everything will work out!").CloseItem()
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r YuleteDefeated) Selection(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("... Are you guys forgiving me after all that I've done? Well, I guess I was blinded by the great source of power that could be discovered that way, maybe they're right saying a human can't simply fathom on the usage of those powers without corrupting themselves along the way... I am profoundly sorry, and to make myself up with everyone I'm willing to help the Societies again wherever I can on the progress of alchemy. Thank you.")
		return script.SendNext(l, c, m.String(), r.Process)
	}
}

func (r YuleteDefeated) Process(l logrus.FieldLogger, c script.Context) script.State {
	if !quest.IsCompleted(l)(c.CharacterId, 7770) {
		quest.Complete(l)(c.CharacterId, 7770)
	}
	return script.WarpById(_map.RomeoAndJuliet, 0)(l, c)
}
