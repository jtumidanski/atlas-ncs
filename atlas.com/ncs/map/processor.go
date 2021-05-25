package _map

import "github.com/sirupsen/logrus"

func CharacterCount(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) uint32 {
	return func(worldId byte, channelId byte, mapId uint32) uint32 {
		return 0
	}
}
