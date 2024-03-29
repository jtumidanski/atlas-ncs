package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// ShamanRock2 is located in Dungeon - The Tunnel That Lost Light II (105090100)
type ShamanRock2 struct {
}

func (r ShamanRock2) NPCId() uint32 {
	return npc.ShamanRock2
}

func (r ShamanRock2) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 2236) || !character.HasItems(l, span)(c.CharacterId, item.ShamanCharm, 1) {
		return script.Exit()(l, span, c)
	}

	progress := quest.Progress(l)(c.CharacterId, 100300)
	if c.MapId == _map.AntTunnelIII {
		return r.ActiveShamanRock(0, progress)(l, span, c)
	} else if c.MapId == _map.DangerousSteam {
		return r.ActiveShamanRock(1, progress)(l, span, c)
	} else if c.MapId == _map.DeepAntTunnelII {
		return r.ActiveShamanRock(2, progress)(l, span, c)
	} else if c.MapId == _map.TheTunnelThatLostLightI {
		id := quest.ProgressInt(l)(c.CharacterId, 2236, 1)
		if id == 0 {
			return r.ProgressQuest(progress)(l, span, c)
		} else if c.NPCObjectId != uint32(id) {
			return r.ActiveShamanRock(4, progress)(l, span, c)
		} else {
			return script.Exit()(l, span, c)
		}
	} else if c.MapId == _map.TheTunnelThatLostLightII {
		return r.ActiveShamanRock(5, progress)(l, span, c)
	}
	return script.Exit()(l, span, c)
}

func (r ShamanRock2) ActiveShamanRock(i int, progress string) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		ch := progress[i]
		if ch != '0' {
			return script.Exit()(l, span, c)
		}

		next := progress[0:i] + string('1') + progress[i+1:]
		quest.SetProgressString(l)(c.CharacterId, 2236, next)
		character.GainItem(l, span)(c.CharacterId, item.ShamanCharm, -1)
		m := message.NewBuilder().
			AddText("The seal took it's place, repelling the evil in the area.")
		return script.SendOk(l, span, c, m.String())
	}
}

func (r ShamanRock2) ProgressQuest(progress string) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		quest.SetProgress(l)(c.CharacterId, 100300, 1, c.NPCObjectId)
		return r.ActiveShamanRock(3, progress)(l, span, c)
	}
}