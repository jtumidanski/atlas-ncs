package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DemonsDoorwayHenesys is located in Victoria Road - Henesys Hunting Ground III (104040002)
type DemonsDoorwayHenesys struct {
}

func (r DemonsDoorwayHenesys) NPCId() uint32 {
	return npc.DemonsDoorwayHenesys
}

func (r DemonsDoorwayHenesys) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestStarted(l)(c.CharacterId, 28238) {
		m := message.NewBuilder().AddText("The entrance is blocked by a strange force.")
		return SendOk(l, c, m.String())
	}

	if !character.HasItem(l)(c.CharacterId, item.AmdusiasEmblem) {
		m := message.NewBuilder().AddText("he entrance is blocked by a force that can only be lifted by those holding an emblem.")
		return SendOk(l, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Would you like to move to ").
		BlueText().ShowMap(_map.AmdusiasStrollingPath).
		BlackText().AddText("?")
	return SendYesNo(l, c, m.String(), r.Process, Exit())
}

func (r DemonsDoorwayHenesys) Process(l logrus.FieldLogger, c Context) State {
	if character.HasItem(l)(c.CharacterId, item.UnseenlyOcarina) {
		character.GainItem(l)(c.CharacterId, item.UnseenlyOcarina, -1)
	}
	if character.HasItem(l)(c.CharacterId, item.DarkweensMonsterDrum) {
		character.GainItem(l)(c.CharacterId, item.DarkweensMonsterDrum, -1)
	}
	if character.HasItem(l)(c.CharacterId, item.SolomonsSealedBow) {
		character.GainItem(l)(c.CharacterId, item.SolomonsSealedBow, -1)
	}
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.AmdusiasStrollingPath, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.AmdusiasStrollingPath, c.NPCId)
	}
	return Exit()(l, c)
}
