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
	MessageTypeNum          = "NUM"
	MessageTypeStyle        = "STYLE"

	SpeakerNPCLeft = "NPC_LEFT"
)

type processor struct {
	l                    logrus.FieldLogger
	enableActionsEmitter producers.EnableActionsEmitter
	changeMapEmitter     producers.ChangeMapEmitter
}

type conversation struct {
	l                   logrus.FieldLogger
	npcTalkEmitter      producers.NPCTalkEmitter
	npcTalkNumEmitter   producers.NPCTalkNumEmitter
	npcTalkStyleEmitter producers.NPCTalkStyleEmitter
	characterId         uint32
	npcId               uint32
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
	npcTalkNumEmitter, _ := producers.NPCTalkNum(p.l)
	npcTalkStyleEmitter, _ := producers.NPCTalkStyle(p.l)

	return &conversation{
		l:                   p.l,
		npcTalkEmitter:      npcTalkEmitter,
		npcTalkNumEmitter:   npcTalkNumEmitter,
		npcTalkStyleEmitter: npcTalkStyleEmitter,
		characterId:         characterId,
		npcId:               npcId,
	}
}

func (p *processor) LockUI() {

}

func (p *processor) WarpById(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) error {
	return p.changeMapEmitter(worldId, channelId, characterId, mapId, portalId)
}

func (p *processor) WarpByName(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) error {
	return nil
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

func (c *conversation) SendGetNumber(message string, defaultValue int32, minimumValue int32, maximumValue int32) error {
	return c.npcTalkNumEmitter(c.characterId, c.npcId, message, defaultValue, minimumValue, maximumValue, MessageTypeNum, SpeakerNPCLeft)
}

func (c *conversation) SendStyle(message string, options []uint32) error {
	return c.npcTalkStyleEmitter(c.characterId, c.npcId, message, options, MessageTypeStyle, SpeakerNPCLeft)
}
