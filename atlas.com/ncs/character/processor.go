package character

import (
	"atlas-ncs/job"
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetCharacterById(characterId uint32) (*Model, error) {
	cs, err := requestCharacter(characterId)
	if err != nil {
		return nil, err
	}
	ca := makeCharacterAttributes(cs.Data())
	if ca == nil {
		return nil, errors.New("unable to make character attributes")
	}
	return ca, nil
}

func makeCharacterAttributes(ca *dataBody) *Model {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return nil
	}
	att := ca.Attributes
	r := Model{
		id:    uint32(cid),
		level: att.Level,
		meso:  att.Meso,
		jobId: att.JobId,
	}
	return &r
}

func HasItem(l logrus.FieldLogger) func(characterId uint32, itemId uint32) bool {
	return func(characterId uint32, itemId uint32) bool {
		return false
	}
}

func IsAboveLevel(l logrus.FieldLogger) func(characterId uint32, level byte) bool {
	return func(characterId uint32, level byte) bool {
		c, err := GetCharacterById(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character.")
			return false
		}
		return c.Level() > level
	}
}

func HasMeso(l logrus.FieldLogger) func(characterId uint32, amount uint32) bool {
	return func(characterId uint32, amount uint32) bool {
		c, err := GetCharacterById(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character.")
			return false
		}
		return c.Meso() > amount
	}
}

func GainItem(l logrus.FieldLogger) func(characterId uint32, itemId uint32, amount int16) {
	return func(characterId uint32, itemId uint32, amount int16) {

	}
}

func GainMeso(l logrus.FieldLogger) func(characterId uint32, amount int32) error {
	adjuster, _ := AdjustMeso(l)
	return func(characterId uint32, amount int32) error {
		err := adjuster(characterId, amount)
		if err != nil {
			l.WithError(err).Errorf("Unable to adjust %d meso by %d.", characterId, amount)
		}
		return err
	}
}

func IsBeginner(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		c, err := GetCharacterById(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character for job check.")
			return false
		}
		return job.IsA(c.JobId(), job.Beginner, job.Noblesse, job.Legend)
	}
}
