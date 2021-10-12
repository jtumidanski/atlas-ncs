package consumers

import (
	"atlas-ncs/conversation"
	"atlas-ncs/kafka/handler"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type continueNPCConversationCommand struct {
	CharacterId uint32 `json:"characterId"`
	Mode        byte   `json:"mode"`
	Type        byte   `json:"type"`
	Selection   int32  `json:"selection"`
}

func ContinueNPCConversationCommandCreator() handler.EmptyEventCreator {
	return func() interface{} {
		return &continueNPCConversationCommand{}
	}
}

func HandleContinueNPCConversationCommand() handler.EventHandler {
	return func(l logrus.FieldLogger, span opentracing.Span, e interface{}) {
		if event, ok := e.(*continueNPCConversationCommand); ok {
			conversation.Continue(l, span)(event.CharacterId, event.Mode, event.Type, event.Selection)
		} else {
			l.Errorf("Unable to cast event provided to handler")
		}
	}
}
