package consumers

import (
	"github.com/sirupsen/logrus"
	"os"
)

func CreateEventConsumers(l *logrus.Logger) {
	brokers := []string{os.Getenv("BOOTSTRAP_SERVERS")}

	creator := func(topicToken string, creator EmptyEventCreator, processor EventProcessor) {
		go func() {
			err := NewConsumer(l, processor, creator, SetGroupId("Character Orchestration Service"), SetTopicToken(topicToken), SetBrokers(brokers))
			if err != nil {
				l.WithError(err).Errorf("Consumer exited with error.")
			}
		}()
	}

	creator("TOPIC_START_NPC_CONVERSATION", StartNPCConversationCommandCreator(), HandleStartNPCConversationCommand())
	creator("TOPIC_CONTINUE_NPC_CONVERSATION", ContinueNPCConversationCommandCreator(), HandleContinueNPCConversationCommand())
	creator("TOPIC_SET_RETURN_TEXT", SetReturnTextCommandCreator(), HandleSetReturnTextCommand())
}
