package producers

import (
	"context"
	"github.com/sirupsen/logrus"
)

type changeMapEvent struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	CharacterId uint32 `json:"characterId"`
	MapId       uint32 `json:"mapId"`
	PortalId    uint32 `json:"portalId"`
}

var ChangeMap = func(l logrus.FieldLogger, ctx context.Context) *changeMap {
	return &changeMap{
		l:   l,
		ctx: ctx,
	}
}

type changeMap struct {
	l   logrus.FieldLogger
	ctx context.Context
}

func (e *changeMap) Emit(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
	event := &changeMapEvent{WorldId: worldId, ChannelId: channelId, CharacterId: characterId, MapId: mapId, PortalId: portalId}
	produceEvent(e.l, "TOPIC_CHANGE_MAP_COMMAND", createKey(int(characterId)), event)
}
