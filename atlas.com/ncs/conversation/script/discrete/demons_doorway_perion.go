package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DemonsDoorwayPerion is located in Victoria Road - West Rocky Mountain IV (102020300)
type DemonsDoorwayPerion struct {
}

func (r DemonsDoorwayPerion) NPCId() uint32 {
	return npc.DemonsDoorwayPerion
}

func (r DemonsDoorwayPerion) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.QuestStarted(l)(c.CharacterId, 28179) {
		m := message.NewBuilder().AddText("The entrance is blocked by a strange force.")
		return script.SendOk(l, c, m.String())
	}

	if !character.HasItem(l)(c.CharacterId, item.AndrasEmblem) {
		m := message.NewBuilder().AddText("he entrance is blocked by a force that can only be lifted by those holding an emblem.")
		return script.SendOk(l, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Would you like to move to ").
		BlueText().ShowMap(_map.AndrasStrollingPath).
		BlackText().AddText("?")
	return script.SendYesNo(l, c, m.String(), r.Process, script.Exit())
}

func (r DemonsDoorwayPerion) Process(l logrus.FieldLogger, c script.Context) script.State {
	if character.HasItem(l)(c.CharacterId, item.BrokenIronFragment) {
		character.GainItem(l)(c.CharacterId, item.BrokenIronFragment, -1)
	}
	if character.HasItem(l)(c.CharacterId, item.OrangeMushroomWine) {
		character.GainItem(l)(c.CharacterId, item.OrangeMushroomWine, -1)
	}
	return script.WarpById(_map.AndrasStrollingPath, 0)(l, c)
}
