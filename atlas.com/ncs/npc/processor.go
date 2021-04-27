package npc

import (
	"atlas-ncs/kafka/producers"
	"context"
	"github.com/sirupsen/logrus"
)

const (
	MessageTypeSimple       = "SIMPLE"
	MessageTypeNext         = "NEXT"
	MessageTypeNextPrevious = "NEXT_PREVIOUS"
	MessageTypePrevious     = "PREVIOUS"
	MessageTypeYesNo        = "YES_NO"

	SpeakerNPCLeft       = "NPC_LEFT"
	SpeakerCharacterLeft = "CHARACTER_LEFT"
)

type processor struct {
	l logrus.FieldLogger
}

type conversation struct {
	l           logrus.FieldLogger
	characterId uint32
	npcId       uint32
}

var Processor = func(l logrus.FieldLogger) *processor {
	return &processor{l}
}

func (p *processor) Dispose(characterId uint32) {
	producers.EnableActions(p.l, context.Background()).Emit(characterId)
}

func (p *processor) Conversation(characterId uint32, npcId uint32) *conversation {
	return &conversation{l: p.l, characterId: characterId, npcId: npcId}
}

func (p *processor) LockUI() {

}

func (p *processor) Warp(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
	producers.ChangeMap(p.l, context.Background()).Emit(worldId, channelId, characterId, mapId, portalId)
}

func (c *conversation) SendSimple(message string) {
	producers.NPCTalk(c.l, context.Background()).Emit(c.characterId, c.npcId, message, MessageTypeSimple, SpeakerNPCLeft)
}

func (c *conversation) SendNext(message string) {
	producers.NPCTalk(c.l, context.Background()).Emit(c.characterId, c.npcId, message, MessageTypeNext, SpeakerNPCLeft)
}

func (c *conversation) send() {

}

func (c *conversation) SendNextPrevious(message string) {
	producers.NPCTalk(c.l, context.Background()).Emit(c.characterId, c.npcId, message, MessageTypeNextPrevious, SpeakerNPCLeft)
}

func (c *conversation) SendPrevious(message string) {
	producers.NPCTalk(c.l, context.Background()).Emit(c.characterId, c.npcId, message, MessageTypePrevious, SpeakerNPCLeft)
}

func (c *conversation) SendYesNo(message string) {
	producers.NPCTalk(c.l, context.Background()).Emit(c.characterId, c.npcId, message, MessageTypeYesNo, SpeakerNPCLeft)
}
