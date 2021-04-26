package consumers

import (
	"atlas-ncs/conversation"
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

func StartNPCConversationCommandCreator() EmptyEventCreator {
	return func() interface{} {
		return &startNPCConversationCommand{}
	}
}

func HandleStartNPCConversationCommand() EventProcessor {
	return func(l logrus.FieldLogger, e interface{}) {
		if event, ok := e.(*startNPCConversationCommand); ok {
			conversation.Processor(l).Start(event.WorldId, event.ChannelId, event.MapId, event.NPCId, event.CharacterId)
		} else {
			l.Errorf("Unable to cast event provided to handler")
		}
	}
}
