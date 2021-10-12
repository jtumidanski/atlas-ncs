package producers

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const topicTokenNPCTalk = "TOPIC_NPC_TALK_COMMAND"

type npcTalkCommand struct {
	CharacterId uint32 `json:"characterId"`
	NPCId       uint32 `json:"npcId"`
	Message     string `json:"message"`
	Type        string `json:"type"`
	Speaker     string `json:"speaker"`
}

type NPCTalkEmitter func(characterId uint32, npcId uint32, message string, messageType string, speaker string)

func NPCTalk(l logrus.FieldLogger, span opentracing.Span) NPCTalkEmitter {
	producer := ProduceEvent(l, span, topicTokenNPCTalk)
	return func(characterId uint32, npcId uint32, message string, messageType string, speaker string) {
		key := CreateKey(int(characterId))
		event := &npcTalkCommand{characterId, npcId, message, messageType, speaker}
		producer(key, event)
	}
}

const topicTokenNPCTalkNum = "TOPIC_NPC_TALK_NUM_COMMAND"

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

type NPCTalkNumEmitter func(characterId uint32, npcId uint32, message string, defaultValue int32, minimumValue int32, maximumValue int32, messageType string, speaker string)

func NPCTalkNum(l logrus.FieldLogger, span opentracing.Span) NPCTalkNumEmitter {
	producer := ProduceEvent(l, span, topicTokenNPCTalkNum)
	return func(characterId uint32, npcId uint32, message string, defaultValue int32, minimumValue int32, maximumValue int32, messageType string, speaker string) {
		key := CreateKey(int(characterId))
		event := &npcTalkCommandNum{characterId, npcId, message, messageType, speaker, defaultValue, minimumValue, maximumValue}
		producer(key, event)
	}
}

const topicTokenNPCTalkStyle = "TOPIC_NPC_TALK_STYLE_COMMAND"

type npcTalkCommandStyle struct {
	CharacterId uint32   `json:"characterId"`
	NPCId       uint32   `json:"npcId"`
	Message     string   `json:"message"`
	Type        string   `json:"type"`
	Speaker     string   `json:"speaker"`
	Styles      []uint32 `json:"styles"`
}

type NPCTalkStyleEmitter func(characterId uint32, npcId uint32, message string, options []uint32, messageType string, speaker string)

func NPCTalkStyle(l logrus.FieldLogger, span opentracing.Span) NPCTalkStyleEmitter {
	producer := ProduceEvent(l, span, topicTokenNPCTalkStyle)
	return func(characterId uint32, npcId uint32, message string, options []uint32, messageType string, speaker string) {
		key := CreateKey(int(characterId))
		event := &npcTalkCommandStyle{characterId, npcId, message, messageType, speaker, options}
		producer(key, event)
	}
}

const topicTokenNPCTalkText = "TOPIC_NPC_TALK_TEXT_COMMAND"

type npcTalkCommandText struct {
	CharacterId uint32 `json:"characterId"`
	NPCId       uint32 `json:"npcId"`
	Message     string `json:"message"`
	Type        string `json:"type"`
	Speaker     string `json:"speaker"`
}

type NPCTalkTextEmitter func(characterId uint32, npcId uint32, message string, messageType string, speaker string)

func NPCTalkText(l logrus.FieldLogger, span opentracing.Span) NPCTalkTextEmitter {
	producer := ProduceEvent(l, span, topicTokenNPCTalkText)
	return func(characterId uint32, npcId uint32, message string, messageType string, speaker string) {
		key := CreateKey(int(characterId))
		event := &npcTalkCommandText{characterId, npcId, message, messageType, speaker}
		producer(key, event)
	}
}
