package consumers

import (
	"atlas-ncs/retry"
	"atlas-ncs/topic"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"time"
)

type EmptyEventCreator func() interface{}

type EventProcessor func(logrus.FieldLogger, interface{})

type Config func(c *config)

func SetGroupId(groupId string) func(c *config) {
	return func(c *config) {
		c.groupId = groupId
	}
}

func SetTopicToken(topicToken string) func(c *config) {
	return func(c *config) {
		c.topicToken = topicToken
	}
}

func SetBrokers(brokers []string) func(c *config) {
	return func(c *config) {
		c.brokers = brokers
	}
}

type config struct {
	groupId    string
	topicToken string
	maxWait    time.Duration
	brokers    []string
}

func NewConsumer(l *logrus.Logger, processor EventProcessor, eventCreator EmptyEventCreator, options ...Config) error {
	c := &config{maxWait: 500 * time.Millisecond}

	for _, option := range options {
		option(c)
	}

	t, err := topic.GetRegistry().Get(c.topicToken)
	if err != nil {
		return err
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: c.brokers,
		Topic:   t.Name(),
		GroupID: c.groupId,
		MaxWait: c.maxWait,
	})

	return readerLoop(r, l.WithFields(logrus.Fields{"originator": t.Name(), "type": "kafka_consumer"}), eventCreator, processor)
}

func readerLoop(r *kafka.Reader, l logrus.FieldLogger, eventCreator EmptyEventCreator, processor EventProcessor) error {
	name := r.Config().Topic

	l.Infof("Creating topic consumer for %s.", name)
	for {
		var msg kafka.Message
		var err error

		readMessage := func(attempt int) (bool, error) {
			msg, err = r.ReadMessage(context.Background())
			if err != nil {
				l.WithError(err).Warnf("Could not read message on topic %s, will retry.", name)
				return true, err
			}
			return false, err
		}

		err = retry.Try(readMessage, 10)
		if err != nil {
			l.WithError(err).Errorf("Could not successfully read message on topic %s.", name)
			continue
		}

		event := eventCreator()
		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			l.WithError(err).Errorf("Could not unmarshal event from topic %s into %s.", name, msg.Value)
			continue
		}

		processor(l, event)
	}
}
