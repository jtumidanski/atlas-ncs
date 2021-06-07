package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DemonsDoorwayEllinia is located in Victoria Road - The Tree That Grew III (101010102)
type DemonsDoorwayEllinia struct {
}

func (r DemonsDoorwayEllinia) NPCId() uint32 {
	return npc.DemonsDoorwayEllinia
}

func (r DemonsDoorwayEllinia) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestStarted(l)(c.CharacterId, 28198) {
		m := message.NewBuilder().AddText("The entrance is blocked by a strange force.")
		return SendOk(l, c, m.String())
	}

	if !character.HasItem(l)(c.CharacterId, item.MarbasEmblem) {
		m := message.NewBuilder().AddText("he entrance is blocked by a force that can only be lifted by those holding an emblem.")
		return SendOk(l, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Would you like to move to ").
		BlueText().ShowMap(_map.MarbasStrollingPath).
		BlackText().AddText("?")
	return SendYesNo(l, c, m.String(), r.Process, Exit())
}

func (r DemonsDoorwayEllinia) Process(l logrus.FieldLogger, c Context) State {
	return WarpById(_map.MarbasStrollingPath, 0)(l, c)
}