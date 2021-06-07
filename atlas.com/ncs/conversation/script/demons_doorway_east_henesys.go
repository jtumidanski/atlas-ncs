package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// DemonsDoorwayEastHenesys is located in Victoria Road - The Forest East of Henesys (100030000)
type DemonsDoorwayEastHenesys struct {
}

func (r DemonsDoorwayEastHenesys) NPCId() uint32 {
	return npc.DemonsDoorwayEastHenesys
}

func (r DemonsDoorwayEastHenesys) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestStarted(l)(c.CharacterId, 28256) {
		m := message.NewBuilder().AddText("The entrance is blocked by a strange force.")
		return SendOk(l, c, m.String())
	}

	if !character.HasItem(l)(c.CharacterId, item.CrocellsEmblem) {
		m := message.NewBuilder().AddText("he entrance is blocked by a force that can only be lifted by those holding an emblem.")
		return SendOk(l, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Would you like to move to ").
		BlueText().ShowMap(_map.CrocellStrollingPath).
		BlackText().AddText("?")
	return SendYesNo(l, c, m.String(), r.Process, Exit())
}

func (r DemonsDoorwayEastHenesys) Process(l logrus.FieldLogger, c Context) State {
	if character.HasItem(l)(c.CharacterId, item.OldWornOutPaper) {
		character.GainItem(l)(c.CharacterId, item.OldWornOutPaper, -1)
	}
	if character.HasItem(l)(c.CharacterId, item.OldKey) {
		character.GainItem(l)(c.CharacterId, item.OldKey, -1)
	}
	if character.HasItem(l)(c.CharacterId, item.GreenSlimeEraser) {
		character.GainItem(l)(c.CharacterId, item.GreenSlimeEraser, -1)
	}
	if character.HasItem(l)(c.CharacterId, item.AyanMercurysMicrophone) {
		character.GainItem(l)(c.CharacterId, item.AyanMercurysMicrophone, -1)
	}
	if character.HasItem(l)(c.CharacterId, item.InkSack) {
		character.GainItem(l)(c.CharacterId, item.InkSack, -1)
	}
	if character.HasItem(l)(c.CharacterId, item.DirtyTreasureMap) {
		character.GainItem(l)(c.CharacterId, item.DirtyTreasureMap, -1)
	}
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.CrocellStrollingPath, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.CrocellStrollingPath, c.NPCId)
	}
	return Exit()(l, c)
}