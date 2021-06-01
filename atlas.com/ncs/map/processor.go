package _map

import "github.com/sirupsen/logrus"

func CharacterCount(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) uint32 {
	return func(worldId byte, channelId byte, mapId uint32) uint32 {
		return 0
	}
}

func MonsterCount(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, monsterId uint32) uint32 {
	return func(worldId byte, channelId byte, mapId uint32, monsterId uint32) uint32 {
		return 0
	}
}

func HasNPC(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, npcId uint32) bool {
	return func(worldId byte, channelId byte, mapId uint32, npcId uint32) bool {
		return false
	}
}

func PlaySound(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, sound string) {
	return func(worldId byte, channelId byte, mapId uint32, sound string) {

	}
}