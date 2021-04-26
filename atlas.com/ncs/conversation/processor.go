package conversation

import (
	"atlas-ncs/conversation/script"
	"github.com/sirupsen/logrus"
)

type processor struct {
	l logrus.FieldLogger
}

var Processor = func(l logrus.FieldLogger) *processor {
	return &processor{l}
}

func (p processor) Start(worldId byte, channelId byte, mapId uint32, npcId uint32, characterId uint32) {
	p.l.Debugf("Start conversation with NPC %d with character %d in map %d.", npcId, characterId, mapId)

	c, err := script.GetRegistry().GetScript(npcId)
	if err != nil {
		p.l.Errorf("Script for npc %d is not implemented.", npcId)
		return
	}

	ctx := script.Context{WorldId: worldId, ChannelId: channelId, MapId: mapId, CharacterId: characterId, NPCId: npcId}
	(*c).Start(p.l, ctx)
	GetRegistry().SetContext(characterId, ctx)
}

func (p processor) Continue(characterId uint32, mode byte, theType byte, selection int32) {
	ctx, err := GetRegistry().GetPreviousContext(characterId)
	if err != nil {
		p.l.WithError(err).Errorf("Unable to retrieve conversation context for %d.", characterId)
		return
	}

	p.l.Debugf("Continuing conversation with NPC %d with character %d in map %d.", ctx.NPCId, characterId, ctx.MapId)

	c, err := script.GetRegistry().GetScript(ctx.NPCId)
	if err != nil {
		p.l.Errorf("Script for npc %d is not implemented.", ctx.NPCId)
		return
	}

	p.l.Debugf("Calling continue for NPC %d conversation with: mode %d, type %d, selection %d.", ctx.NPCId, mode, theType, selection)
	(*c).Continue(p.l, *ctx, mode, theType, selection)
}
