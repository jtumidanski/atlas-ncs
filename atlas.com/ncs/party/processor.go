package party

import "github.com/sirupsen/logrus"

func GetForCharacter(l logrus.FieldLogger) func(characterId uint32) (*Model, error) {
	return func(characterId uint32) (*Model, error) {
		return nil, nil
	}
}

func IsLeader(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return false
	}
}

func Warp(l logrus.FieldLogger) func(characterId uint32, mapId uint32) {
	return func(characterId uint32, mapId uint32) {

	}
}

func WarpById(l logrus.FieldLogger) func(characterId uint32, mapId uint32, portalId uint32) {
	return func(characterId uint32, mapId uint32, portalId uint32) {

	}
}

func HasParty(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		//TODO
		return false
	}
}

