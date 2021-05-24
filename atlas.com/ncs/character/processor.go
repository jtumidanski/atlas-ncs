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
		id:           uint32(cid),
		level:        att.Level,
		meso:         att.Meso,
		jobId:        att.JobId,
		dexterity:    att.Dexterity,
		intelligence: att.Intelligence,
	}
	return &r
}

func HasItem(l logrus.FieldLogger) func(characterId uint32, itemId uint32) bool {
	return func(characterId uint32, itemId uint32) bool {
		return HasItems(l)(characterId, itemId, 1)
	}
}

func HasItems(l logrus.FieldLogger) func(characterId uint32, itemId uint32, quantity uint32) bool {
	return func(characterId uint32, itemId uint32, quantity uint32) bool {
		items, err := requestItemsForCharacter(characterId, itemId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve inventory items for character %d.", characterId)
			return false
		}

		count := uint32(0)
		for _, i := range items.Data {
			count += i.Attributes.Quantity
			if count > quantity {
				return true
			}
		}

		return false
	}
}

func CanHold(l logrus.FieldLogger) func(characterId uint32, itemId uint32) bool {
	return func(characterId uint32, itemId uint32) bool {
		return CanHoldAll(l)(characterId, itemId, 1)
	}
}

func CanHoldAll(l logrus.FieldLogger) func(characterId uint32, itemId uint32, quantity uint32) bool {
	return func(characterId uint32, itemId uint32, quantity uint32) bool {
		return true
	}
}

func ChangeJob(l logrus.FieldLogger) func(characterId uint32, jobId uint16) {
	return func(characterId uint32, jobId uint16) {
		adjustJob(l)(characterId, jobId)
	}
}

func ResetAP(l logrus.FieldLogger) func(characterId uint32) {
	return func(characterId uint32) {
		resetAP(l)(characterId)
	}
}

func IsLevel(l logrus.FieldLogger) func(characterId uint32, level byte) bool {
	return func(characterId uint32, level byte) bool {
		c, err := GetCharacterById(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character.")
			return false
		}
		return c.Level() >= level
	}
}

func HasDexterity(l logrus.FieldLogger) func(characterId uint32, amount uint16) bool {
	return func(characterId uint32, amount uint16) bool {
		c, err := GetCharacterById(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character.")
			return false
		}
		return c.Dexterity() >= amount
	}
}

func HasIntelligence(l logrus.FieldLogger) func(characterId uint32, amount uint16) bool {
	return func(characterId uint32, amount uint16) bool {
		c, err := GetCharacterById(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character.")
			return false
		}
		return c.Intelligence() >= amount
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

func GainEquipment(l logrus.FieldLogger) func(characterId uint32, itemId uint32) {
	return func(characterId uint32, itemId uint32) {
		gainEquipment(l)(characterId, itemId)
	}
}

func GainItem(l logrus.FieldLogger) func(characterId uint32, itemId uint32, amount int32) {
	return func(characterId uint32, itemId uint32, amount int32) {
		gainItem(l)(characterId, itemId, amount)
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

func IsBeginnerTree(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		c, err := GetCharacterById(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character for job check.")
			return false
		}
		return job.IsA(c.JobId(), job.Beginner, job.Noblesse, job.Legend)
	}
}

func IsJob(l logrus.FieldLogger) func(characterId uint32, option uint16) bool {
	return func(characterId uint32, option uint16) bool {
		c, err := GetCharacterById(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character for job check.")
			return false
		}
		return job.IsA(c.JobId(), option)
	}
}

func CompleteQuest(l logrus.FieldLogger) func(characterId uint32, questId uint32) {
	return func(characterId uint32, questId uint32) {

	}
}

func QuestStarted(l logrus.FieldLogger) func(characterId uint32, questId uint32) bool {
	return func(characterId uint32, questId uint32) bool {
		return false
	}
}

func QuestCompleted(l logrus.FieldLogger) func(characterId uint32, questId uint32) bool {
	return func(characterId uint32, questId uint32) bool {
		return false
	}
}

func StartQuest(l logrus.FieldLogger) func(characterId uint32, questId uint32) {
	return func(characterId uint32, questId uint32) {

	}
}

func SaveLocation(l logrus.FieldLogger) func(characterId uint32, location string) {
	return func(characterId uint32, location string) {

	}
}

func BuddyCapacity(l logrus.FieldLogger) func(characterId uint32) uint8 {
	return func(characterId uint32) uint8 {
		return 0
	}
}

func IncreaseBuddyCapacity(l logrus.FieldLogger) func(characterId uint32, amount int8) error {
	return func(characterId uint32, amount int8) error {
		return nil
	}
}

func HasPets(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return false
	}
}

func GainCloseness(l logrus.FieldLogger) func(characterId uint32, amount int8) {
	return func(characterId uint32, amount int8) {
		
	}
}