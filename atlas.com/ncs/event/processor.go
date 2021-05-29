package event

import "github.com/sirupsen/logrus"

func StartEvent(l logrus.FieldLogger) func(characterId uint32, eventName string) bool {
	return func(characterId uint32, eventName string) bool {
		return false
	}
}

func SetProperty(l logrus.FieldLogger) func(characterId uint32, property string, value string) {
	return func(characterId uint32, property string, value string) {

	}
}