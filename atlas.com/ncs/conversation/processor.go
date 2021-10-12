package conversation

import (
	"atlas-ncs/conversation/script"
	registry2 "atlas-ncs/conversation/script/registry"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func Start(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, npcId uint32, npcObjectId uint32, characterId uint32) {
	return func(worldId byte, channelId byte, mapId uint32, npcId uint32, npcObjectId uint32, characterId uint32) {
		l.Debugf("Start conversation with NPC %d with character %d in map %d.", npcId, characterId, mapId)
		s, err := GetRegistry().GetPreviousContext(characterId)
		if err == nil {
			l.Debugf("Previous conversation between character %d and npc %d exists, avoiding starting new conversation with %d.", characterId, s.ctx.NPCId, npcId)
			return
		}

		c, err := registry2.GetRegistry().GetScript(npcId)
		if err != nil {
			l.Errorf("Script for npc %d is not implemented.", npcId)
			return
		}

		ctx := script.Context{WorldId: worldId, ChannelId: channelId, MapId: mapId, CharacterId: characterId, NPCId: npcId, NPCObjectId: npcObjectId}
		ns := (*c).Initial(l, span, ctx)
		if ns != nil {
			GetRegistry().SetContext(characterId, ctx, ns)
		} else {
			GetRegistry().ClearContext(characterId)
		}
	}
}

func Continue(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, mode byte, theType byte, selection int32) {
	return func(characterId uint32, mode byte, theType byte, selection int32) {
		s, err := GetRegistry().GetPreviousContext(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve conversation context for %d.", characterId)
			return
		}
		ctx := s.ctx
		state := s.ns

		l.Debugf("Continuing conversation with NPC %d with character %d in map %d.", ctx.NPCId, characterId, ctx.MapId)
		l.Debugf("Calling continue for NPC %d conversation with: mode %d, type %d, selection %d.", ctx.NPCId, mode, theType, selection)
		ns := state(l, span, ctx, mode, theType, selection)
		if ns != nil {
			GetRegistry().SetContext(characterId, ctx, ns)
		} else {
			GetRegistry().ClearContext(characterId)
		}
	}
}
