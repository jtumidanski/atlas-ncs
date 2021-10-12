package consumers

import (
	"atlas-ncs/kafka/handler"
	"context"
	"github.com/sirupsen/logrus"
	"sync"
)

const (
	StartNPCConversationCommand    = "start_npc_conversation_command"
	ContinueNPCConversationCommand = "continue_npc_conversation_command"
	SetReturnTextCommand           = "set_return_text_command"
)

func CreateEventConsumers(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup) {
	cec := func(topicToken string, name string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
		createEventConsumer(l, ctx, wg, name, topicToken, emptyEventCreator, processor)
	}
	cec("TOPIC_START_NPC_CONVERSATION", StartNPCConversationCommand, StartNPCConversationCommandCreator(), HandleStartNPCConversationCommand())
	cec("TOPIC_CONTINUE_NPC_CONVERSATION", ContinueNPCConversationCommand, ContinueNPCConversationCommandCreator(), HandleContinueNPCConversationCommand())
	cec("TOPIC_SET_RETURN_TEXT", SetReturnTextCommand, SetReturnTextCommandCreator(), HandleSetReturnTextCommand())

}

func createEventConsumer(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup, name string, topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
	wg.Add(1)
	go NewConsumer(l, ctx, wg, name, topicToken, "NPC Conversation Service", emptyEventCreator, processor)
}
