package character

import (
	"atlas-ncs/kafka/producers"
	"github.com/sirupsen/logrus"
)

type adjustMesoEvent struct {
	CharacterId uint32 `json:"characterId"`
	Amount      int32 `json:"amount"`
	Show        bool   `json:"show"`
}

type MesoAdjuster func(characterId uint32, amount int32) error

func AdjustMeso(l logrus.FieldLogger) (MesoAdjuster, error) {
	producer, err := producers.ProduceEvent(l, "TOPIC_ADJUST_MESO")
	if err != nil {
		l.WithError(err).Errorf("Unable to create meso adjustment producer.")
		return nil, err
	}

	return func(characterId uint32, amount int32) error {
		event := &adjustMesoEvent{characterId, amount, true}
		return producer(producers.CreateKey(int(characterId)), event)
	}, nil
}
