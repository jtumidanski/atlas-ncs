package producers

import (
	"github.com/sirupsen/logrus"
	"os"
)

const topicTokenChangeMap = "TOPIC_CHANGE_MAP_COMMAND"

type changeMapEvent struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	CharacterId uint32 `json:"characterId"`
	MapId       uint32 `json:"mapId"`
	PortalId    uint32 `json:"portalId"`
}

type ChangeMapEmitter func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) error

func ChangeMap(l logrus.FieldLogger) ChangeMapEmitter {
	producer, _ := ProduceEvent(l, topicTokenChangeMap, SetBrokers([]string{os.Getenv("BOOTSTRAP_SERVERS")}))
	return produceChangeMap(producer)
}

func produceChangeMap(producer MessageProducer) ChangeMapEmitter {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) error {
		event := &changeMapEvent{WorldId: worldId, ChannelId: channelId, CharacterId: characterId, MapId: mapId, PortalId: portalId}
		return producer(CreateKey(int(characterId)), event)
	}
}
