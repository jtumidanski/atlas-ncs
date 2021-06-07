package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Roodolph is located in Hidden Street - Extra Frosty Snow Zone (209080000) and Hidden Street - Happyville (209000000)
type Roodolph struct {
}

func (r Roodolph) NPCId() uint32 {
	return npc.Roodolph
}

func (r Roodolph) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.Happyville {
		m := message.NewBuilder().
			AddText("Do you wish to head to where the ").
			BlueText().AddText("Snow Sprinkler").
			BlackText().AddText(" is?")
		return SendYesNo(l, c, m.String(), r.Warp(_map.ExtraFrostySnowZone), r.WhenYouWantTo)
	} else if c.MapId == _map.ExtraFrostySnowZone {
		m := message.NewBuilder().
			AddText("Do you wish to return to Happyville?")
		return SendYesNo(l, c, m.String(), r.Warp(_map.Happyville), r.WhenYouWantTo)
	}
	m := message.NewBuilder().AddText("You Alright?")
	return SendOk(l, c, m.String())
}

func (r Roodolph) WhenYouWantTo(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Talk to me again when you want to.")
	return SendOk(l, c, m.String())
}

func (r Roodolph) Warp(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return Exit()(l, c)
	}
}
