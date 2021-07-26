package pet

import "github.com/sirupsen/logrus"

func HasPets(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		//TODO
		return false
	}
}

func HasPetInSlot(l logrus.FieldLogger) func(characterId uint32, slot int16) bool {
	return func(characterId uint32, slot int16) bool {
		//TODO
		return false
	}
}

func Is(l logrus.FieldLogger) func(characterId uint32, slot int16, petId ...uint32) bool {
	return func(characterId uint32, slot int16, petId ...uint32) bool {
		//TODO
		return false
	}
}

func IsLevel(l logrus.FieldLogger) func(characterId uint32, slot int16, level byte) bool {
	return func(characterId uint32, slot int16, level byte) bool {
		//TODO
		return false
	}
}

func Evolve(l logrus.FieldLogger) func(characterId uint32, slot int16, itemId uint32) {
	return func(characterId uint32, slot int16, itemId uint32) {
		//TODO
	}
}

func GainCloseness(l logrus.FieldLogger) func(characterId uint32, amount int8) {
	return func(characterId uint32, amount int8) {
		//TODO
	}
}

