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

func EnableActions(l logrus.FieldLogger) (EnableActionsEmitter, error) {
	producer, err := create(l, topicTokenEnableActions, SetBrokers([]string{os.Getenv("BOOTSTRAP_SERVERS")}))
	if err != nil {
		return nil, err
	}
	return produceEnableActions(producer), nil
}

func produceEnableActions(producer MessageProducer) EnableActionsEmitter {
	return func(characterId uint32) error {
		event := &enableActionsEvent{characterId}
		return producer(createKey(int(characterId)), event)
	}
}
