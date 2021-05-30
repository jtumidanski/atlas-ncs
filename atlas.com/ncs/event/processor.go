package event

import "github.com/sirupsen/logrus"

func StartEvent(l logrus.FieldLogger) func(characterId uint32, eventName string) bool {
	return func(characterId uint32, eventName string) bool {
		return false
	}
}

func StartPartyEvent(l logrus.FieldLogger) func(eventName string, partyId uint32, mapId uint32, difficulty uint32) bool {
	return func(eventName string, partyId uint32, mapId uint32, difficulty uint32) bool {
		return false
	}
}

func SetProperty(l logrus.FieldLogger) func(eventName string, property string, value string) {
	return func(eventName string, property string, value string) {

	}
}