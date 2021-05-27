package party

import "github.com/sirupsen/logrus"

func GetParty(l logrus.FieldLogger) func(characterId uint32) (*Model, error) {
	return func(characterId uint32) (*Model, error) {
		return nil, nil
	}
}
