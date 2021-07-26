package alliance

import "github.com/sirupsen/logrus"

func GuildHasAlliance(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		//TODO
		return false
	}
}

func IsLeader(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		//TODO
		return false
	}
}

func ValidName(l logrus.FieldLogger) func(text string) bool {
	return func(text string) bool {
		//TODO
		return false
	}
}

func Create(l logrus.FieldLogger) func(characterId uint32, name string) error {
	return func(characterId uint32, name string) error {
		//TODO
		return nil
	}
}

func AtCapacity(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		//TODO
		return false
	}
}

func Expand(l logrus.FieldLogger) func(characterId uint32) error {
	return func(characterId uint32) error {
		//TODO
		return nil
	}
}

