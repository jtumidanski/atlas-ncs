package producers

import (
	"github.com/sirupsen/logrus"
	"os"
)

const topicTokenEnableActions = "TOPIC_ENABLE_ACTIONS"

type enableActionsEvent struct {
	CharacterId uint32 `json:"characterId"`
}

type EnableActionsEmitter func(characterId uint32) error

func EnableActions(l logrus.FieldLogger) EnableActionsEmitter {
	producer, _ := ProduceEvent(l, topicTokenEnableActions, SetBrokers([]string{os.Getenv("BOOTSTRAP_SERVERS")}))
	return produceEnableActions(producer)
}

func produceEnableActions(producer MessageProducer) EnableActionsEmitter {
	return func(characterId uint32) error {
		event := &enableActionsEvent{characterId}
		return producer(CreateKey(int(characterId)), event)
	}
}
