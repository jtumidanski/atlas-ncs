package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/monster"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// Cygnus is located in Hidden Street - Quiet Ereve (913030000)
type Cygnus struct {
}

func (r Cygnus) NPCId() uint32 {
	return npc.Cygnus
}

func (r Cygnus) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !(quest.IsCompleted(l)(c.CharacterId, 20407) || quest.IsStarted(l)(c.CharacterId, 20407) && quest.ProgressInt(l)(c.CharacterId, 20407, int(monster.BlackWitch)) != 0) &&
		_map.MonsterCount(l)(c.WorldId, c.ChannelId, c.MapId, monster.BlackWitch) == 0 &&
		!_map.HasNPC(l)(c.WorldId, c.ChannelId, c.MapId, npc.Eleanor) {
		return r.ShesAlreadyHere(l, c)
	}
	m := message.NewBuilder().AddText("...")
	return script.SendOk(l, c, m.String())
}

func (r Cygnus) ShesAlreadyHere(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("... Hnngh... ").
		BlueText().ShowCharacterName().
		BlackText().AddText(", is that you...? ").
		RedText().ShowNPC(npc.Eleanor).
		BlackText().AddText("... She's already here... ").
		BlueText().ShowCharacterName().
		BlackText().AddText(", I'm truly sorry I can't help you right now in this state, just when a bigger threat appeared I could do nothing for my people.... Please I beg you, please defeat her, ").
		BlueText().ShowCharacterName().
		BlackText().AddText("!! ....")
	return script.SendOkTrigger(l, c, m.String(), r.SpawnEleanor)
}

func (r Cygnus) SpawnEleanor(l logrus.FieldLogger, c script.Context) script.State {
	npc.Spawn(l)(c.WorldId, c.ChannelId, c.MapId, npc.Eleanor, 850, 0)
	return script.Exit()(l, c)
}
