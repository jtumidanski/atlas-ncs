package npc

import (
	"atlas-ncs/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	topicTokenChangeMap     = "TOPIC_CHANGE_MAP_COMMAND"
	topicTokenEnableActions = "TOPIC_ENABLE_ACTIONS"
	topicTokenNPCTalkText   = "TOPIC_NPC_TALK_TEXT_COMMAND"
	topicTokenNPCTalkStyle  = "TOPIC_NPC_TALK_STYLE_COMMAND"
	topicTokenNPCTalkNum    = "TOPIC_NPC_TALK_NUM_COMMAND"
	topicTokenNPCTalk       = "TOPIC_NPC_TALK_COMMAND"
)

type changeMapEvent struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	CharacterId uint32 `json:"characterId"`
	MapId       uint32 `json:"mapId"`
	PortalId    uint32 `json:"portalId"`
}

func emitChangeMap(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
	producer := kafka.ProduceEvent(l, span, topicTokenChangeMap)
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
		event := &changeMapEvent{WorldId: worldId, ChannelId: channelId, CharacterId: characterId, MapId: mapId, PortalId: portalId}
		producer(kafka.CreateKey(int(characterId)), event)
	}
}

type enableActionsEvent struct {
	CharacterId uint32 `json:"characterId"`
}

func emitEnableActions(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) {
	producer := kafka.ProduceEvent(l, span, topicTokenEnableActions)
	return func(characterId uint32) {
		event := &enableActionsEvent{characterId}
		producer(kafka.CreateKey(int(characterId)), event)
	}
}

type npcTalkCommand struct {
	CharacterId uint32 `json:"characterId"`
	NPCId       uint32 `json:"npcId"`
	Message     string `json:"message"`
	Type        string `json:"type"`
	Speaker     string `json:"speaker"`
}

func emitTalk(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32, message string, messageType string, speaker string) {
	producer := kafka.ProduceEvent(l, span, topicTokenNPCTalk)
	return func(characterId uint32, npcId uint32, message string, messageType string, speaker string) {
		key := kafka.CreateKey(int(characterId))
		event := &npcTalkCommand{characterId, npcId, message, messageType, speaker}
		producer(key, event)
	}
}

type npcTalkCommandNum struct {
	CharacterId  uint32 `json:"characterId"`
	NPCId        uint32 `json:"npcId"`
	Message      string `json:"message"`
	Type         string `json:"type"`
	Speaker      string `json:"speaker"`
	DefaultValue int32  `json:"defaultValue"`
	MinimumValue int32  `json:"minimumValue"`
	MaximumValue int32  `json:"maximumValue"`
}

func emitTalkNumber(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32, message string, defaultValue int32, minimumValue int32, maximumValue int32, messageType string, speaker string) {
	producer := kafka.ProduceEvent(l, span, topicTokenNPCTalkNum)
	return func(characterId uint32, npcId uint32, message string, defaultValue int32, minimumValue int32, maximumValue int32, messageType string, speaker string) {
		key := kafka.CreateKey(int(characterId))
		event := &npcTalkCommandNum{characterId, npcId, message, messageType, speaker, defaultValue, minimumValue, maximumValue}
		producer(key, event)
	}
}

type npcTalkCommandStyle struct {
	CharacterId uint32   `json:"characterId"`
	NPCId       uint32   `json:"npcId"`
	Message     string   `json:"message"`
	Type        string   `json:"type"`
	Speaker     string   `json:"speaker"`
	Styles      []uint32 `json:"styles"`
}

func emitTalkStyle(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32, message string, options []uint32, messageType string, speaker string) {
	producer := kafka.ProduceEvent(l, span, topicTokenNPCTalkStyle)
	return func(characterId uint32, npcId uint32, message string, options []uint32, messageType string, speaker string) {
		key := kafka.CreateKey(int(characterId))
		event := &npcTalkCommandStyle{characterId, npcId, message, messageType, speaker, options}
		producer(key, event)
	}
}

type npcTalkCommandText struct {
	CharacterId uint32 `json:"characterId"`
	NPCId       uint32 `json:"npcId"`
	Message     string `json:"message"`
	Type        string `json:"type"`
	Speaker     string `json:"speaker"`
}

func emitTalkText(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32, message string, messageType string, speaker string) {
	producer := kafka.ProduceEvent(l, span, topicTokenNPCTalkText)
	return func(characterId uint32, npcId uint32, message string, messageType string, speaker string) {
		key := kafka.CreateKey(int(characterId))
		event := &npcTalkCommandText{characterId, npcId, message, messageType, speaker}
		producer(key, event)
	}
}
