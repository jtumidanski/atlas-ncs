package producers

import (
	"github.com/sirupsen/logrus"
	"os"
)

const topicTokenNPCTalk = "TOPIC_NPC_TALK_COMMAND"

type npcTalkCommand struct {
	CharacterId uint32 `json:"characterId"`
	NPCId       uint32 `json:"npcId"`
	Message     string `json:"message"`
	Type        string `json:"type"`
	Speaker     string `json:"speaker"`
}

type NPCTalkEmitter func(characterId uint32, npcId uint32, message string, messageType string, speaker string) error

func NPCTalk(l logrus.FieldLogger) NPCTalkEmitter {
	producer, _ := ProduceEvent(l, topicTokenNPCTalk, SetBrokers([]string{os.Getenv("BOOTSTRAP_SERVERS")}))
	return func(characterId uint32, npcId uint32, message string, messageType string, speaker string) error {
		key := CreateKey(int(characterId))
		event := &npcTalkCommand{characterId, npcId, message, messageType, speaker}
		return producer(key, event)
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

type NPCTalkNumEmitter func(characterId uint32, npcId uint32, message string, defaultValue int32, minimumValue int32, maximumValue int32, messageType string, speaker string) error

func NPCTalkNum(l logrus.FieldLogger) NPCTalkNumEmitter {
	producer, _ := ProduceEvent(l, topicTokenNPCTalkNum, SetBrokers([]string{os.Getenv("BOOTSTRAP_SERVERS")}))
	return func(characterId uint32, npcId uint32, message string, defaultValue int32, minimumValue int32, maximumValue int32, messageType string, speaker string) error {
		key := CreateKey(int(characterId))
		event := &npcTalkCommandNum{characterId, npcId, message, messageType, speaker, defaultValue, minimumValue, maximumValue}
		return producer(key, event)
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

type NPCTalkStyleEmitter func(characterId uint32, npcId uint32, message string, options []uint32, messageType string, speaker string) error

func NPCTalkStyle(l logrus.FieldLogger) NPCTalkStyleEmitter {
	producer, _ := ProduceEvent(l, topicTokenNPCTalkStyle, SetBrokers([]string{os.Getenv("BOOTSTRAP_SERVERS")}))
	return func(characterId uint32, npcId uint32, message string, options []uint32, messageType string, speaker string) error {
		key := CreateKey(int(characterId))
		event := &npcTalkCommandStyle{characterId, npcId, message, messageType, speaker, options}
		return producer(key, event)
	}
}
