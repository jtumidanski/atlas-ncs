package consumers

import (
	"context"
	"github.com/sirupsen/logrus"
)

func CreateEventConsumers(l *logrus.Logger) {
	cec := func(topicToken string, emptyEventCreator EmptyEventCreator, processor EventProcessor) {
		createEventConsumer(l, topicToken, emptyEventCreator, processor)
	}
	cec("TOPIC_START_NPC_CONVERSATION", StartNPCConversationCommandCreator(), HandleStartNPCConversationCommand())
	cec("TOPIC_CONTINUE_NPC_CONVERSATION", ContinueNPCConversationCommandCreator(), HandleContinueNPCConversationCommand())
	cec("TOPIC_SET_RETURN_TEXT", SetReturnTextCommandCreator(), HandleSetReturnTextCommand())

}

func createEventConsumer(l *logrus.Logger, topicToken string, emptyEventCreator EmptyEventCreator, processor EventProcessor) {
	h := func(logger logrus.FieldLogger, event interface{}) {
		processor(logger, event)
	}

	c := NewConsumer(l, context.Background(), h,
		SetGroupId("Character Orchestration Service"),
		SetTopicToken(topicToken),
		SetEmptyEventCreator(emptyEventCreator))
	go c.Init()
}
