package consumers

import (
	"atlas-ncs/kafka/handler"
	"context"
	"github.com/sirupsen/logrus"
	"sync"
)

func CreateEventConsumers(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup) {
	cec := func(topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
		createEventConsumer(l, ctx, wg, topicToken, emptyEventCreator, processor)
	}
	cec("TOPIC_START_NPC_CONVERSATION", StartNPCConversationCommandCreator(), HandleStartNPCConversationCommand())
	cec("TOPIC_CONTINUE_NPC_CONVERSATION", ContinueNPCConversationCommandCreator(), HandleContinueNPCConversationCommand())
	cec("TOPIC_SET_RETURN_TEXT", SetReturnTextCommandCreator(), HandleSetReturnTextCommand())

}

func createEventConsumer(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup, topicToken string, emptyEventCreator handler.EmptyEventCreator, processor handler.EventHandler) {
	wg.Add(1)
	go NewConsumer(l, ctx, wg, topicToken, "World Registry Service", emptyEventCreator, processor)
}
