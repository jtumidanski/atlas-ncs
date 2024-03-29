package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/monster"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// CornerOfTheMagicLibrary is located in Hidden Street - Magic Library (910110000)
type CornerOfTheMagicLibrary struct {
}

func (r CornerOfTheMagicLibrary) NPCId() uint32 {
	return npc.CornerOfTheMagicLibrary
}

func (r CornerOfTheMagicLibrary) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 20718) {
		return script.Exit()(l, span, c)
	}
	return r.AngryMonsters(l, span, c)
}

func (r CornerOfTheMagicLibrary) AngryMonsters(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("A mysterious black figure appeared and summoned a lot of angry monsters!")
	return script.SendOkTrigger(l, span, c, m.String(), r.Spawn)
}

func (r CornerOfTheMagicLibrary) Spawn(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	for i := 0; i < 10; i++ {
		monster.SpawnMonster(l)(c.WorldId, c.ChannelId, c.MapId, monster.BlueMushroom, 117, 183)
	}
	for i := 0; i < 10; i++ {
		monster.SpawnMonster(l)(c.WorldId, c.ChannelId, c.MapId, monster.BlueMushroom, 4, 183)
	}
	for i := 0; i < 10; i++ {
		monster.SpawnMonster(l)(c.WorldId, c.ChannelId, c.MapId, monster.BlueMushroom, -109, 183)
	}
	quest.CompleteViaNPC(l)(c.CharacterId, 20718, npc.Hersha)
	character.GainExperience(l)(c.CharacterId, 4000)
	return script.Exit()(l, span, c)
}
