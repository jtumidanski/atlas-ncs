package consumers

import (
	"atlas-ncs/conversation"
	"atlas-ncs/kafka/handler"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type startNPCConversationCommand struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	MapId       uint32 `json:"mapId"`
	CharacterId uint32 `json:"characterId"`
	NPCId       uint32 `json:"npcId"`
	NPCObjectId uint32 `json:"npcObjectId"`
}

func StartNPCConversationCommandCreator() handler.EmptyEventCreator {
	return func() interface{} {
		return &startNPCConversationCommand{}
	}
}

func HandleStartNPCConversationCommand() handler.EventHandler {
	return func(l logrus.FieldLogger, span opentracing.Span, e interface{}) {
		if event, ok := e.(*startNPCConversationCommand); ok {
			conversation.Start(l, span)(event.WorldId, event.ChannelId, event.MapId, event.NPCId, event.NPCObjectId, event.CharacterId)
		} else {
			l.Errorf("Unable to cast event provided to handler")
		}
	}
}
