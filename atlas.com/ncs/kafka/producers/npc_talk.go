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
	producer, err := create(l, topicTokenNPCTalk, SetBrokers([]string{os.Getenv("BOOTSTRAP_SERVERS")}))
	if err != nil {
		return nil, err
	}
	return produceNPCTalk(producer), nil
}

func produceNPCTalk(producer MessageProducer) NPCTalkEmitter {
	return func(characterId uint32, npcId uint32, message string, messageType string, speaker string) error {
		key := createKey(int(characterId))
		event := &npcTalkCommand{characterId, npcId, message, messageType, speaker}
		return producer(key, event)
	}
}
