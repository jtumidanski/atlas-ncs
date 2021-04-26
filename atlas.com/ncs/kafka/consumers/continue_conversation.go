package consumers

import (
	"atlas-ncs/conversation"
	"github.com/sirupsen/logrus"
)

type continueNPCConversationCommand struct {
	CharacterId uint32 `json:"characterId"`
	Mode        byte   `json:"mode"`
	Type        byte   `json:"type"`
	Selection   int32  `json:"selection"`
}

func ContinueNPCConversationCommandCreator() EmptyEventCreator {
	return func() interface{} {
		return &continueNPCConversationCommand{}
	}
}

func HandleContinueNPCConversationCommand() EventProcessor {
	return func(l logrus.FieldLogger, e interface{}) {
		if event, ok := e.(*continueNPCConversationCommand); ok {
			conversation.Processor(l).Continue(event.CharacterId, event.Mode, event.Type, event.Selection)
		} else {
			l.Errorf("Unable to cast event provided to handler")
		}
	}
}
