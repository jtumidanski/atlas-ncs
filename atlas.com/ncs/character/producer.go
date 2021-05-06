package character

import (
	"atlas-ncs/kafka/producers"
	"github.com/sirupsen/logrus"
)

type adjustMesoEvent struct {
	CharacterId uint32 `json:"characterId"`
	Amount      int32  `json:"amount"`
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

type gainEquipmentCommand struct {
	CharacterId uint32 `json:"characterId"`
	ItemId      uint32 `json:"itemId"`
}

func gainEquipment(l logrus.FieldLogger) func(characterId uint32, itemId uint32) {
	producer, _ := producers.ProduceEvent(l, "TOPIC_CHARACTER_GAIN_EQUIPMENT")
	return func(characterId uint32, itemId uint32) {
		e := &gainEquipmentCommand{CharacterId: characterId, ItemId: itemId}
		_ = producer(producers.CreateKey(int(characterId)), e)
	}
}

type gainItemCommand struct {
	CharacterId uint32 `json:"characterId"`
	ItemId      uint32 `json:"itemId"`
	Quantity    int32 `json:"quantity"`
}

func gainItem(l logrus.FieldLogger) func(characterId uint32, itemId uint32, quantity int32) {
	producer, _ := producers.ProduceEvent(l, "TOPIC_CHARACTER_GAIN_ITEM")
	return func(characterId uint32, itemId uint32, quantity int32) {
		e := &gainItemCommand{CharacterId: characterId, ItemId: itemId, Quantity: quantity}
		_ = producer(producers.CreateKey(int(characterId)), e)
	}
}

type adjustJobCommand struct {
	CharacterId uint32 `json:"characterId"`
	JobId       uint16 `json:"jobId"`
}

func adjustJob(l logrus.FieldLogger) func(characterId uint32, jobId uint16) {
	producer, _ := producers.ProduceEvent(l, "TOPIC_CHARACTER_ADJUST_JOB")
	return func(characterId uint32, jobId uint16) {
		e := &adjustJobCommand{CharacterId: characterId, JobId: jobId}
		_ = producer(producers.CreateKey(int(characterId)), e)
	}
}

type resetAPCommand struct {
	CharacterId uint32 `json:"characterId"`
}

func resetAP(l logrus.FieldLogger) func(characterId uint32) {
	producer, _ := producers.ProduceEvent(l, "TOPIC_CHARACTER_RESET_AP")
	return func(characterId uint32) {
		e := &resetAPCommand{CharacterId: characterId}
		_ = producer(producers.CreateKey(int(characterId)), e)
	}

}