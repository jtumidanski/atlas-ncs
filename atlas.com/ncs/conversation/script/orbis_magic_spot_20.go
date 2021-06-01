package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// OrbisMagicSpot20 is located in Orbis - Orbis Tower <20th Floor> (200080200)
type OrbisMagicSpot20 struct {
}

func (r OrbisMagicSpot20) NPCId() uint32 {
	return npc.OrbisMagicSpot20
}

func (r OrbisMagicSpot20) Initial(l logrus.FieldLogger, c Context) State {
	if !character.HasItem(l)(c.CharacterId, item.OrbisRockScroll) {
		return r.NeedScroll(l, c)
	}
	return r.Confirm(l, c)
}

func (r OrbisMagicSpot20) NeedScroll(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("There's a ").
		BlueText().ShowNPC(npc.OrbisMagicSpot20).
		BlackText().AddText(" that'll enable you to teleport to where ").
		BlueText().ShowNPC(npc.OrbisMagicSpot1).
		BlackText().AddText(" is, but you can't activate it without the scroll.")
	return SendOk(l, c, m.String())
}

func (r OrbisMagicSpot20) Confirm(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You can use ").
		BlueText().ShowItemName1(item.OrbisRockScroll).
		BlackText().AddText(" to activate ").
		BlueText().ShowNPC(npc.OrbisMagicSpot20).
		BlackText().AddText(". Will you teleport to where ").
		BlueText().ShowNPC(npc.OrbisMagicSpot1).
		BlackText().AddText(" is?")
	return SendYesNo(l, c, m.String(), r.Warp, Exit())
}

func (r OrbisMagicSpot20) Warp(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.OrbisRockScroll, -1)
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.OrbisTower1stFloor, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.OrbisTower1stFloor, c.NPCId)
	}
	return Exit()(l, c)
}
