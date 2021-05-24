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

func NPCTalk(l logrus.FieldLogger) (NPCTalkEmitter, error) {
	producer, err := ProduceEvent(l, topicTokenNPCTalk, SetBrokers([]string{os.Getenv("BOOTSTRAP_SERVERS")}))
	if err != nil {
		return nil, err
	}
	return func(characterId uint32, npcId uint32, message string, messageType string, speaker string) error {
		key := CreateKey(int(characterId))
		event := &npcTalkCommand{characterId, npcId, message, messageType, speaker}
		return producer(key, event)
	}, nil
}

const topicTokenNPCTalkNum = "TOPIC_NPC_TALK_COMMAND"

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

func NPCTalkNum(l logrus.FieldLogger) (NPCTalkNumEmitter, error) {
	producer, err := ProduceEvent(l, topicTokenNPCTalkNum, SetBrokers([]string{os.Getenv("BOOTSTRAP_SERVERS")}))
	if err != nil {
		return nil, err
	}
	return func(characterId uint32, npcId uint32, message string, defaultValue int32, minimumValue int32, maximumValue int32, messageType string, speaker string) error {
		key := CreateKey(int(characterId))
		event := &npcTalkCommandNum{characterId, npcId, message, messageType, speaker, defaultValue, minimumValue, maximumValue}
		return producer(key, event)
	}, nil
}
