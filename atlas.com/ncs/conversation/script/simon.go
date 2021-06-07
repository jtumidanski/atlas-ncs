package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Simon is located in Hidden Street - Happyville (209000000) and Hidden Street - Shalom Temple (681000000)
type Simon struct {
}

func (r Simon) NPCId() uint32 {
	return npc.Simon
}

func (r Simon) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.Happyville {
		m := message.NewBuilder().
			AddText("The Shalom Temple is unlike any other place in Happyville, would you like to head to ").
			BlueText().AddText("Shalom Temple").
			BlackText().AddText("?")
		return SendYesNo(l, c, m.String(), r.Warp(_map.ShalomTemple), r.LetMeKnow)
	} else if c.MapId == _map.ShalomTemple {
		m := message.NewBuilder().
			AddText("Would you like to head back to Happyville?")
		return SendYesNo(l, c, m.String(), r.Warp(_map.Happyville), r.LetMeKnow)
	}
	return Exit()(l, c)
}

func (r Simon) LetMeKnow(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Let me know if you've changed your mind!")
	return SendOk(l, c, m.String())
}

func (r Simon) Warp(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return Exit()(l, c)
	}
}
