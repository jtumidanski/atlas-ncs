package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// SubwayTrashCan3 is located in Kerning City Subway - Line 1 <Area 1> (103000101)
type SubwayTrashCan3 struct {
}

func (r SubwayTrashCan3) NPCId() uint32 {
	return npc.SubwayTrashCan3
}

func (r SubwayTrashCan3) Initial(l logrus.FieldLogger, c Context) State {
	return r.Hello(l, c)
}

func (r SubwayTrashCan3) Hello(l logrus.FieldLogger, c Context) State {
	if !character.QuestStarted(l)(c.CharacterId, 20710) {
		return r.JustATrashCan(l, c)
	}

	if character.HasItem(l)(c.CharacterId, item.BubblingDoll) {
		return r.JustATrashCan(l, c)
	}

	if !character.CanHold(l)(c.CharacterId, item.BubblingDoll) {
		return r.NotEnoughSpace(l, c)
	}

	return r.GiveItem(l, c)
}

func (r SubwayTrashCan3) JustATrashCan(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Just a trash can sitting there.")
	return SendOk(l, c, m.String())
}

func (r SubwayTrashCan3) NotEnoughSpace(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Not enough space in your ETC inventory.")
	return SendOk(l, c, m.String())
}

func (r SubwayTrashCan3) GiveItem(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.BubblingDoll, 1)
	m := message.NewBuilder().
		AddText("You have found a ").
		BlueText().ShowItemName1(item.BubblingDoll).
		BlackText().AddText(" in the trash can!")
	return SendNext(l, c, m.String(), Exit())
}
