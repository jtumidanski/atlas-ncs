package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Peter is located in Entrance - Mushroom Town Training Camp (000000003)
type Peter struct {
}

func (r Peter) NPCId() uint32 {
	return npc.Peter
}

func (r Peter) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You have finished all your trainings. Good job. You seem to be ready to start with the journey right away! Good, I will let you move on to the next place.")
	return SendNext(l, c, m.String(), r.Remember)
}

func (r Peter) Remember(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("But remember, once you get out of here, you will enter a village full with monsters. Well them, good bye!")
	return SendNextPrevious(l, c, m.String(), r.Process, r.Initial)
}

func (r Peter) Process(l logrus.FieldLogger, c Context) State {
	character.GainExperience(l)(c.CharacterId, 3)
	return WarpById(_map.InASmallForest, 0)(l, c)
}
