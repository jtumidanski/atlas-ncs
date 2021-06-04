package event

import (
	"github.com/sirupsen/logrus"
)

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

func GetProperty(l logrus.FieldLogger) func(eventName string, property string) string {
	return func(eventName string, property string) string {
		return ""
	}
}

func Cleared(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return false
	}
}

func Leader(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return false
	}
}

func GiveParticipantsExperience(l logrus.FieldLogger) func(amount uint32) {
	return func(amount uint32) {

	}
}

func Clear(l logrus.FieldLogger) {

}

func ReceivedReward(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return false
	}
}

func SetRewardReceived(l logrus.FieldLogger) func(characterId uint32) {
	return func(characterId uint32) {

	}
}

func GiveEventReward(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return true
	}
}

func JoinEvent(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32) bool {
	return func(worldId byte, channelId byte, characterId uint32) bool {
		//TODO implement this in event service
		e, ok := GetEvent(l)(worldId, channelId)
		if !ok {
			return false
		}

		if ok && e.Limit() > 0 {
			// if coconut harvest or event map entrance assign to team
			// reduce limit
			// warp character
			return true
		}
		return false
	}
}

func GetEvent(l logrus.FieldLogger) func(worldId byte, channelId byte) (*Model, bool) {
	return func(worldId byte, channelId byte) (*Model, bool) {
		return nil, false
	}
}

func AssignToTeam(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32) {
	return func(worldId byte, channelId byte, characterId uint32) {

	}
}

func LeaveEvent(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32) {
	return func(worldId byte, channelId byte, characterId uint32) {
		//TODO implement this in event service
		// increase limit
		// warp character
	}
}