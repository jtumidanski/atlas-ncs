package producers

import (
	"context"
	"github.com/sirupsen/logrus"
)

type enableActionsEvent struct {
	CharacterId uint32 `json:"characterId"`
}

var EnableActions = func(l logrus.FieldLogger, ctx context.Context) *enableActions {
	return &enableActions{
		l:   l,
		ctx: ctx,
	}
}

type enableActions struct {
	l   logrus.FieldLogger
	ctx context.Context
}

func (e *enableActions) Emit(characterId uint32) {
	event := &enableActionsEvent{characterId}
	produceEvent(e.l, "TOPIC_ENABLE_ACTIONS", createKey(int(characterId)), event)
}
