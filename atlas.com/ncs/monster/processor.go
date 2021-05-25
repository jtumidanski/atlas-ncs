package monster

import "github.com/sirupsen/logrus"

func SpawnMonster(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, monsterId uint32, x int16, y int16) {
	return func(worldId byte, channelId byte, mapId uint32, monsterId uint32, x int16, y int16) {

	}
}
