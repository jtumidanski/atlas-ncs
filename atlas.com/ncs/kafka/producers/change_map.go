package producers

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const topicTokenChangeMap = "TOPIC_CHANGE_MAP_COMMAND"

type changeMapEvent struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	CharacterId uint32 `json:"characterId"`
	MapId       uint32 `json:"mapId"`
	PortalId    uint32 `json:"portalId"`
}

type ChangeMapEmitter func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32)

func ChangeMap(l logrus.FieldLogger, span opentracing.Span) ChangeMapEmitter {
	producer := ProduceEvent(l, span, topicTokenChangeMap)
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
		event := &changeMapEvent{WorldId: worldId, ChannelId: channelId, CharacterId: characterId, MapId: mapId, PortalId: portalId}
		producer(CreateKey(int(characterId)), event)
	}
}
