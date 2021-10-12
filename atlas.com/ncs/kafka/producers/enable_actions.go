package producers

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const topicTokenEnableActions = "TOPIC_ENABLE_ACTIONS"

type enableActionsEvent struct {
	CharacterId uint32 `json:"characterId"`
}

type EnableActionsEmitter func(characterId uint32)

func EnableActions(l logrus.FieldLogger, span opentracing.Span) EnableActionsEmitter {
	producer := ProduceEvent(l, span, topicTokenEnableActions)
	return func(characterId uint32) {
		event := &enableActionsEvent{characterId}
		producer(CreateKey(int(characterId)), event)
	}
}
