package npc

import (
	"atlas-ncs/kafka/producers"
	"context"
	"github.com/sirupsen/logrus"
)

const (
	MessageTypeSimple       = "SIMPLE"
	MessageTypeNextPrevious = "NEXT_PREVIOUS"

	SpeakerLeft          = "NPC_LEFT"
	SpeakerCharacterLeft = "CHARACTER_LEFT"
)

type processor struct {
	l logrus.FieldLogger
}

var Processor = func(l logrus.FieldLogger) *processor {
	return &processor{l}
}

func (p *processor) SendSimple(characterId uint32, npcId uint32, message string) {
	producers.NPCTalk(p.l, context.Background()).Emit(characterId, npcId, message, MessageTypeSimple, SpeakerLeft)
}

func (p *processor) SendNext(characterId uint32, npcId uint32, message string) {
	producers.NPCTalk(p.l, context.Background()).Emit(characterId, npcId, message, MessageTypeNextPrevious, SpeakerCharacterLeft)
}

func (p *processor) Dispose(characterId uint32) {
	producers.EnableActions(p.l, context.Background()).Emit(characterId)
}
