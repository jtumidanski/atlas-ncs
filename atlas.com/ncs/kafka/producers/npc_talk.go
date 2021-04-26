package producers

import (
	"context"
	log "github.com/sirupsen/logrus"
)

type npcTalkCommand struct {
	CharacterId uint32 `json:"characterId"`
	NPCId       uint32 `json:"npcId"`
	Message     string `json:"message"`
	Type        string `json:"type"`
	Speaker     string `json:"speaker"`
}

var NPCTalk = func(l log.FieldLogger, ctx context.Context) *npcTalk {
	return &npcTalk{
		l:   l,
		ctx: ctx,
	}
}

type npcTalk struct {
	l   log.FieldLogger
	ctx context.Context
}

func (e *npcTalk) Emit(characterId uint32, npcId uint32, message string, messageType string, speaker string) {
	event := &npcTalkCommand{characterId, npcId, message, messageType, speaker}
	produceEvent(e.l, "TOPIC_NPC_TALK_COMMAND", createKey(int(characterId)), event)
}
