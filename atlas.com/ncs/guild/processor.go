package guild

import "github.com/sirupsen/logrus"

func IsLeader(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		//TODO
		return false
	}
}

func HasGuild(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		//TODO
		return false
	}
}

