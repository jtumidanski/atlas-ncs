package character

import (
	"atlas-ncs/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type adjustMesoEvent struct {
	CharacterId uint32 `json:"characterId"`
	Amount      int32  `json:"amount"`
	Show        bool   `json:"show"`
}

func emitMesoAdjustment(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, amount int32) {
	producer := kafka.ProduceEvent(l, span, "TOPIC_ADJUST_MESO")
	return func(characterId uint32, amount int32) {
		event := &adjustMesoEvent{characterId, amount, true}
		producer(kafka.CreateKey(int(characterId)), event)
	}
}

type gainItemCommand struct {
	CharacterId uint32 `json:"characterId"`
	ItemId      uint32 `json:"itemId"`
	Quantity    int32  `json:"quantity"`
}

func emitItemGain(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, itemId uint32, quantity int32) {
	producer := kafka.ProduceEvent(l, span, "TOPIC_CHARACTER_GAIN_ITEM")
	return func(characterId uint32, itemId uint32, quantity int32) {
		e := &gainItemCommand{CharacterId: characterId, ItemId: itemId, Quantity: quantity}
		producer(kafka.CreateKey(int(characterId)), e)
	}
}

type adjustJobCommand struct {
	CharacterId uint32 `json:"characterId"`
	JobId       uint16 `json:"jobId"`
}

func emitJobAdjustment(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, jobId uint16) {
	producer := kafka.ProduceEvent(l, span, "TOPIC_CHARACTER_ADJUST_JOB")
	return func(characterId uint32, jobId uint16) {
		e := &adjustJobCommand{CharacterId: characterId, JobId: jobId}
		producer(kafka.CreateKey(int(characterId)), e)
	}
}

type resetAPCommand struct {
	CharacterId uint32 `json:"characterId"`
}

func emitAPReset(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) {
	producer := kafka.ProduceEvent(l, span, "TOPIC_CHARACTER_RESET_AP")
	return func(characterId uint32) {
		e := &resetAPCommand{CharacterId: characterId}
		producer(kafka.CreateKey(int(characterId)), e)
	}
}
