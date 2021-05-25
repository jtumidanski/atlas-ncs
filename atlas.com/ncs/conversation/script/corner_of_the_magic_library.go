package script

import (
	"atlas-ncs/character"
	"atlas-ncs/monster"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// CornerOfTheMagicLibrary is located in Hidden Street - Magic Library (910110000)
type CornerOfTheMagicLibrary struct {
}

func (r CornerOfTheMagicLibrary) NPCId() uint32 {
	return npc.CornerOfTheMagicLibrary
}

func (r CornerOfTheMagicLibrary) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestStarted(l)(c.CharacterId, 20718) {
		return Exit()(l, c)
	}
	return r.AngryMonsters(l, c)
}

func (r CornerOfTheMagicLibrary) AngryMonsters(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("A mysterious black figure appeared and summoned a lot of angry monsters!")
	return SendOkTrigger(l, c, m.String(), r.Spawn)
}

func (r CornerOfTheMagicLibrary) Spawn(l logrus.FieldLogger, c Context) State {
	for i := 0; i < 10; i++ {
		monster.SpawnMonster(l)(c.WorldId, c.ChannelId, c.MapId, monster.BlueMushroom, 117, 183)
	}
	for i := 0; i < 10; i++ {
		monster.SpawnMonster(l)(c.WorldId, c.ChannelId, c.MapId, monster.BlueMushroom, 4, 183)
	}
	for i := 0; i < 10; i++ {
		monster.SpawnMonster(l)(c.WorldId, c.ChannelId, c.MapId, monster.BlueMushroom, -109, 183)
	}
	character.CompleteQuestViaNPC(l)(c.CharacterId, 20718, npc.Hersha)
	character.GainExperience(l)(c.CharacterId, 4000)
	return Exit()(l, c)
}
