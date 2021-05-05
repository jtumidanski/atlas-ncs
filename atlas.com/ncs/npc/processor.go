package npc

import (
	"atlas-ncs/kafka/producers"
	"github.com/sirupsen/logrus"
)

const (
	MessageTypeSimple       = "SIMPLE"
	MessageTypeNext         = "NEXT"
	MessageTypeNextPrevious = "NEXT_PREVIOUS"
	MessageTypePrevious     = "PREVIOUS"
	MessageTypeYesNo        = "YES_NO"
	MessageTypeOk           = "OK"

	SpeakerNPCLeft = "NPC_LEFT"
)

type processor struct {
	l                    logrus.FieldLogger
	enableActionsEmitter producers.EnableActionsEmitter
	changeMapEmitter     producers.ChangeMapEmitter
}

type conversation struct {
	l              logrus.FieldLogger
	npcTalkEmitter producers.NPCTalkEmitter
	characterId    uint32
	npcId          uint32
}

var Processor = func(l logrus.FieldLogger) *processor {
	enableActionsEmitter, _ := producers.EnableActions(l)
	changeMapEmitter, _ := producers.ChangeMap(l)

	return &processor{l, enableActionsEmitter, changeMapEmitter}
}

func (p *processor) Dispose(characterId uint32) error {
	return p.enableActionsEmitter(characterId)
}

func (p *processor) Conversation(characterId uint32, npcId uint32) *conversation {
	npcTalkEmitter, _ := producers.NPCTalk(p.l)

	return &conversation{
		l:              p.l,
		npcTalkEmitter: npcTalkEmitter,
		characterId:    characterId,
		npcId:          npcId,
	}
}

func (p *processor) LockUI() {

}

func (p *processor) Warp(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) error {
	return p.changeMapEmitter(worldId, channelId, characterId, mapId, portalId)
}

func (c *conversation) SendSimple(message string) error {
	return c.npcTalkEmitter(c.characterId, c.npcId, message, MessageTypeSimple, SpeakerNPCLeft)
}

func (c *conversation) SendNext(message string) error {
	return c.npcTalkEmitter(c.characterId, c.npcId, message, MessageTypeNext, SpeakerNPCLeft)
}

func (c *conversation) SendNextPrevious(message string) error {
	return c.npcTalkEmitter(c.characterId, c.npcId, message, MessageTypeNextPrevious, SpeakerNPCLeft)
}

func (c *conversation) SendPrevious(message string) error {
	return c.npcTalkEmitter(c.characterId, c.npcId, message, MessageTypePrevious, SpeakerNPCLeft)
}

func (c *conversation) SendYesNo(message string) error {
	return c.npcTalkEmitter(c.characterId, c.npcId, message, MessageTypeYesNo, SpeakerNPCLeft)
}

func (c *conversation) SendOk(message string) error {
	return c.npcTalkEmitter(c.characterId, c.npcId, message, MessageTypeOk, SpeakerNPCLeft)
}
