package event

import "github.com/sirupsen/logrus"

func StartEvent(l logrus.FieldLogger) func(characterId uint32, eventName string) bool {
	return func(characterId uint32, eventName string) bool {
		return false
	}
}
