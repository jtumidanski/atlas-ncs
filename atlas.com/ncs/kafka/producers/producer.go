package producers

import (
	"atlas-ncs/retry"
	"atlas-ncs/topic"
	"context"
	"encoding/binary"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func CreateKey(key int) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint32(b, uint32(key))
	return b
}

type Config func(c *config)

func SetBrokers(brokers []string) func(c *config) {
	return func(c *config) {
		c.brokers = brokers
	}
}

type config struct {
	batchTimeout time.Duration
	brokers      []string
}

type MessageProducer func([]byte, interface{}) error

func ProduceEvent(l logrus.FieldLogger, topicToken string, options ...Config) (MessageProducer, error) {
	c := &config{
		batchTimeout: 50 * time.Millisecond,
		brokers: []string{os.Getenv("BOOTSTRAP_SERVERS")},
	}

	for _, option := range options {
		option(c)
	}

	t := topic.GetRegistry().Get(l, topicToken)

	w := &kafka.Writer{
		Addr:         kafka.TCP(c.brokers...),
		Topic:        t,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: c.batchTimeout,
	}

	return produce(l, w), nil
}

func produce(l logrus.FieldLogger, w *kafka.Writer) MessageProducer {
	return func(key []byte, event interface{}) error {
		r, err := json.Marshal(event)
		l.WithField("message", string(r)).Debugf("Writing message to topic %s.", w.Topic)
		if err != nil {
			l.WithError(err).Errorf("Unable to marshall event for topic %s.", w.Topic)
			return err
		}

		produceMessage := func(attempt int) (bool, error) {
			err = w.WriteMessages(context.Background(), kafka.Message{Key: key, Value: r})
			if err != nil {
				l.WithError(err).Warnf("Unable to produce event on topic %s, will retry.", w.Topic)
				return true, err
			}
			return false, err
		}

		err = retry.Try(produceMessage, 10)
		if err != nil {
			l.WithError(err).Fatalf("Unable to produce event on topic %s.", w.Topic)
			return err
		}
		return nil
	}
}
