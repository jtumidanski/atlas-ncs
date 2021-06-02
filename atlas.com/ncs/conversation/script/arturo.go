package script

import (
	"atlas-ncs/event"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Arturo is located in Hidden Street - Abandoned Tower<Determine to adventure> (922011100)
type Arturo struct {
}

func (r Arturo) NPCId() uint32 {
	return npc.Arturo
}

func (r Arturo) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Congratulations on sealing the dimensional crack! For all of your hard work, I have a gift for you! Here take this prize.")
	return SendNext(l, c, m.String(), r.Validate)
}

func (r Arturo) Validate(l logrus.FieldLogger, c Context) State {
	ok := event.GiveEventReward(l)(c.CharacterId)
	if !ok {
		return r.NeedInventorySpace(l, c)
	}
	return r.Warp(l, c)
}

func (r Arturo) NeedInventorySpace(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("It seems you don't have a free slot in either your ").
		RedText().AddText("Equip").
		BlackText().AddText(", ").
		RedText().AddText("Use").
		BlackText().AddText(" or ").
		RedText().AddText("Etc").
		BlackText().AddText(" inventories. Please make some room and try again.")
	return SendOk(l, c, m.String())
}

func (r Arturo) Warp(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.EosTower101stFloor, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.EosTower101stFloor, c.NPCId)
	}
	return Exit()(l, c)
}
