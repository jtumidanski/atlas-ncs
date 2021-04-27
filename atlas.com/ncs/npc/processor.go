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

func (c *conversation) SendSimple(message string) {
	producers.NPCTalk(c.l, context.Background()).Emit(c.characterId, c.npcId, message, MessageTypeSimple, SpeakerNPCLeft)
}

func (c *conversation) SendNext(message string) {
	producers.NPCTalk(c.l, context.Background()).Emit(c.characterId, c.npcId, message, MessageTypeNext, SpeakerNPCLeft)
}

func (c *conversation) send() {

}
