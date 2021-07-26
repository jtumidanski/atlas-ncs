package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// ShamanRock1 is located in Dungeon - The Tunnel That Lost Light II (105090100)
type ShamanRock1 struct {
}

func (r ShamanRock1) NPCId() uint32 {
	return npc.ShamanRock1
}

func (r ShamanRock1) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 2236) || !character.HasItems(l)(c.CharacterId, item.ShamanCharm, 1) {
		return script.Exit()(l, c)
	}

	progress := quest.Progress(l)(c.CharacterId, 100300)
	if c.MapId == _map.AntTunnelIII {
		return r.ActiveShamanRock(0, progress)(l, c)
	} else if c.MapId == _map.DangerousSteam {
		return r.ActiveShamanRock(1, progress)(l, c)
	} else if c.MapId == _map.DeepAntTunnelII {
		return r.ActiveShamanRock(2, progress)(l, c)
	} else if c.MapId == _map.TheTunnelThatLostLightI {
		id := quest.ProgressInt(l)(c.CharacterId, 2236, 1)
		if id == 0 {
			return r.ProgressQuest(progress)(l, c)
		} else if c.NPCObjectId != uint32(id) {
			return r.ActiveShamanRock(4, progress)(l, c)
		} else {
			return script.Exit()(l, c)
		}
	} else if c.MapId == _map.TheTunnelThatLostLightII {
		return r.ActiveShamanRock(5, progress)(l, c)
	}
	return script.Exit()(l, c)
}

func (r ShamanRock1) ActiveShamanRock(i int, progress string) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		ch := progress[i]
		if ch != '0' {
			return script.Exit()(l, c)
		}

		next := progress[0:i] + string('1') + progress[i+1:]
		quest.SetProgressString(l)(c.CharacterId, 2236, next)
		character.GainItem(l)(c.CharacterId, item.ShamanCharm, -1)
		m := message.NewBuilder().
			AddText("The seal took it's place, repelling the evil in the area.")
		return script.SendOk(l, c, m.String())
	}
}

func (r ShamanRock1) ProgressQuest(progress string) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		quest.SetProgress(l)(c.CharacterId, 100300, 1, c.NPCObjectId)
		return r.ActiveShamanRock(3, progress)(l, c)
	}
}
