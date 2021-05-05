package character

import (
	"github.com/sirupsen/logrus"
)

func HasItem(l logrus.FieldLogger) func(characterId uint32, itemId uint32) bool {
	return func(characterId uint32, itemId uint32) bool {
		return false
	}
}

func IsAboveLevel(l logrus.FieldLogger) func(characterId uint32, level int8) bool {
	return func(characterId uint32, level int8) bool {
		return false
	}
}

func HasMeso(l logrus.FieldLogger) func(characterId uint32, amount uint32) bool {
	return func(characterId uint32, amount uint32) bool {
		return false
	}
}

func GainItem(l logrus.FieldLogger) func(characterId uint32, itemId uint32, amount int16) {
	return func(characterId uint32, itemId uint32, amount int16) {

	}
}

func GainMeso(l logrus.FieldLogger) func(characterId uint32, amount int16) {
	return func(characterId uint32, amount int16) {
		
	}
}