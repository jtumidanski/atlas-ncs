package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// SecondEOSRock is located in Ludibrium - Eos Tower 71st Floor (221022900)
type SecondEOSRock struct {
}

func (r SecondEOSRock) NPCId() uint32 {
	return npc.SecondEOSRock
}

func (r SecondEOSRock) Initial(l logrus.FieldLogger, c Context) State {
	if character.HasItem(l)(c.CharacterId, item.EOSRockScroll) {
		return r.ToNext(l, c)
	}
	return r.NeedRock(l, c)
}

func (r SecondEOSRock) ToNext(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You can use ").
		BlueText().AddText("Eos Rock Scroll").
		BlackText().AddText(" to activate ").
		BlueText().AddText("Second Eos Rock").
		BlackText().AddText(". Which of these rocks would you like to teleport to?").NewLine().
		OpenItem(0).BlueText().AddText("First Eos Rock (100th Floor)").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Third Eos Rock (41st Floor)").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r SecondEOSRock) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Confirm100
	case 1:
		return r.Confirm41
	}
	return nil
}

func (r SecondEOSRock) NeedRock(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("There's a rock that will enable you to teleport to ").
		BlueText().AddText("First Eos Rock or Third Eos Rock").
		BlackText().AddText(", but it cannot be activated without the scroll.")
	return SendOk(l, c, m.String())

}

func (r SecondEOSRock) Confirm100(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You can use ").
		BlueText().AddText("Eos Rock Scroll").
		BlackText().AddText(" to activate ").
		BlueText().AddText("Second Eos Rock").
		BlackText().AddText(". Will you teleport to ").
		BlueText().AddText("First Eos Rock").
		BlackText().AddText(" at the 100th Floor?")
	return SendYesNo(l, c, m.String(), r.Process(_map.EosTower100thFloor), Exit())
}

func (r SecondEOSRock) Confirm41(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You can use ").
		BlueText().AddText("Eos Rock Scroll").
		BlackText().AddText(" to activate ").
		BlueText().AddText("Second Eos Rock").
		BlackText().AddText(". Will you teleport to ").
		BlueText().AddText("Third Eos Rock").
		BlackText().AddText(" at the 41st Floor?")
	return SendYesNo(l, c, m.String(), r.Process(_map.EosTower41stFloor), Exit())
}

func (r SecondEOSRock) Process(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		character.GainItem(l)(c.CharacterId, item.EOSRockScroll, -1)
		err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 3)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return Exit()(l, c)
	}
}
