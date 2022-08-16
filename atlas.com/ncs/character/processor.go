package character

import (
	"atlas-ncs/job"
	"atlas-ncs/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetCharacterById(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) (Model, error) {
	return func(characterId uint32) (Model, error) {
		return requests.Provider[attributes, Model](l, span)(requestCharacter(characterId), makeModel)()
	}
}

func makeModel(ca requests.DataBody[attributes]) (Model, error) {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return Model{}, err
	}
	att := ca.Attributes
	r := Model{
		id:           uint32(cid),
		level:        att.Level,
		meso:         att.Meso,
		jobId:        att.JobId,
		strength:     att.Strength,
		dexterity:    att.Dexterity,
		intelligence: att.Intelligence,
		mapId:        att.MapId,
		gender:       att.Gender,
		hair:         att.Hair,
		face:         att.Face,
	}
	return r, nil
}

type AttributeCriteria func(Model) bool

func MeetsCriteria(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, criteria ...AttributeCriteria) bool {
	return func(characterId uint32, criteria ...AttributeCriteria) bool {
		c, err := GetCharacterById(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character %d for criteria check.", characterId)
			return false
		}
		for _, check := range criteria {
			if ok := check(c); !ok {
				return false
			}
		}
		return true
	}
}

func HasItem(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, itemId uint32) bool {
	return func(characterId uint32, itemId uint32) bool {
		return HasItems(l, span)(characterId, itemId, 1)
	}
}

func HasItems(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, itemId uint32, quantity uint32) bool {
	return func(characterId uint32, itemId uint32, quantity uint32) bool {
		items, err := requestItemsForCharacter(characterId, itemId)(l, span)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve inventory items for character %d.", characterId)
			return false
		}

		count := uint32(0)
		for _, i := range items.DataList() {
			attr := i.Attributes
			count += attr.Quantity
			if count >= quantity {
				return true
			}
		}

		return false
	}
}

func HasAnyItem(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, items ...uint32) bool {
	return func(characterId uint32, items ...uint32) bool {
		allItems, err := requestAllItemsForCharacter(characterId)(l, span)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve inventory items for character %d.", characterId)
			return false
		}

		for _, i := range allItems.DataList() {
			possibleId, err := strconv.Atoi(i.Id)
			if err != nil {
				l.WithError(err).Errorf("Error parsing item id %s.", i.Id)
				continue
			}
			for _, id := range items {
				if uint32(possibleId) == id {
					return true
				}
			}
		}

		return false
	}
}

type Item struct {
	ItemId   uint32
	Quantity uint32
}

func CanHold(l logrus.FieldLogger) func(characterId uint32, itemId uint32) bool {
	return func(characterId uint32, itemId uint32) bool {
		return CanHoldAll(l)(characterId, itemId, 1)
	}
}

func CanHoldAll(l logrus.FieldLogger) func(characterId uint32, itemId uint32, quantity uint32) bool {
	return func(characterId uint32, itemId uint32, quantity uint32) bool {
		return CanHoldThese(l)(characterId, Item{ItemId: itemId, Quantity: quantity})
	}
}

func CanHoldThese(l logrus.FieldLogger) func(characterId uint32, items ...Item) bool {
	return func(characterId uint32, items ...Item) bool {
		return true
	}
}

func ChangeJob(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, jobId uint16) {
	return func(characterId uint32, jobId uint16) {
		emitJobAdjustment(l, span)(characterId, jobId)
	}
}

func ResetAP(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) {
	return func(characterId uint32) {
		emitAPReset(l, span)(characterId)
	}
}

func IsLevel(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, level byte) bool {
	return func(characterId uint32, level byte) bool {
		return MeetsCriteria(l, span)(characterId, IsLevelCriteria(level))
	}
}

func IsLevelCriteria(level byte) AttributeCriteria {
	return func(c Model) bool {
		return c.Level() >= level
	}
}

func LevelBetweenCriteria(lower byte, upper byte) AttributeCriteria {
	return func(c Model) bool {
		return c.Level() > lower && c.Level() < upper
	}
}

func HasStrengthCriteria(amount uint16) AttributeCriteria {
	return func(c Model) bool {
		return c.Strength() >= amount
	}
}

func HasDexterityCriteria(amount uint16) AttributeCriteria {
	return func(c Model) bool {
		return c.Dexterity() >= amount
	}
}

func HasIntelligenceCriteria(amount uint16) AttributeCriteria {
	return func(c Model) bool {
		return c.Intelligence() >= amount
	}
}

func AboveLevel(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, level byte) bool {
	return func(characterId uint32, level byte) bool {
		return MeetsCriteria(l, span)(characterId, AboveLevelCriteria(level))
	}
}

func AboveLevelCriteria(level byte) AttributeCriteria {
	return func(c Model) bool {
		return c.Level() > level
	}
}

func HasMeso(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, amount uint32) bool {
	return func(characterId uint32, amount uint32) bool {
		return MeetsCriteria(l, span)(characterId, HasMesoCriteria(amount))
	}
}

func HasMesoCriteria(amount uint32) AttributeCriteria {
	return func(c Model) bool {
		return c.Meso() >= amount
	}
}

func GainEquipment(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, itemId uint32) {
	return func(characterId uint32, itemId uint32) {
		emitEquipmentGain(l, span)(characterId, itemId)
	}
}

func GainItem(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, itemId uint32, amount int32) {
	return func(characterId uint32, itemId uint32, amount int32) {
		emitItemGain(l, span)(characterId, itemId, amount)
	}
}

func GainFame(l logrus.FieldLogger) func(characterId uint32, amount int32) {
	return func(characterId uint32, amount int32) {
		//TODO implement fame gain.
	}
}

func GainMeso(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, amount int32) {
	return func(characterId uint32, amount int32) {
		emitMesoAdjustment(l, span)(characterId, amount)
	}
}

func IsBeginnerTree(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return MeetsCriteria(l, span)(characterId, IsBeginnerTreeCriteria())
	}
}

func IsBeginnerTreeCriteria() AttributeCriteria {
	return func(c Model) bool {
		return job.IsA(c.JobId(), job.Beginner, job.Noblesse, job.Legend)
	}
}

func IsJob(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, option uint16) bool {
	return func(characterId uint32, option uint16) bool {
		return MeetsCriteria(l, span)(characterId, IsJobCriteria(option))
	}
}

func IsJobCriteria(option uint16) AttributeCriteria {
	return func(c Model) bool {
		return job.IsA(c.JobId(), option)
	}
}

func IsAJobCriteria(options ...uint16) AttributeCriteria {
	return func(c Model) bool {
		return job.IsA(c.JobId(), options...)
	}
}

func BuddyCapacity(l logrus.FieldLogger) func(characterId uint32) uint8 {
	return func(characterId uint32) uint8 {
		//TODO
		return 0
	}
}

func IncreaseBuddyCapacity(l logrus.FieldLogger) func(characterId uint32, amount int8) error {
	return func(characterId uint32, amount int8) error {
		//TODO
		return nil
	}
}

func GetGender(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) byte {
	return func(characterId uint32) byte {
		c, err := GetCharacterById(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character %d.", characterId)
			return 0
		}
		return c.Gender()
	}
}

func GetHair(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) uint32 {
	return func(characterId uint32) uint32 {
		c, err := GetCharacterById(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character %d.", characterId)
			return 0
		}
		return c.Hair()
	}
}

func GetFace(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) uint32 {
	return func(characterId uint32) uint32 {
		c, err := GetCharacterById(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character %d.", characterId)
			return 0
		}
		return c.Face()
	}
}

func SetHair(l logrus.FieldLogger) func(characterId uint32, hair uint32) {
	return func(characterId uint32, hair uint32) {
		//TODO
	}
}

func SetSkin(l logrus.FieldLogger) func(characterId uint32, skin byte) {
	return func(characterId uint32, skin byte) {
		//TODO
	}
}

func SetFace(l logrus.FieldLogger) func(characterId uint32, face uint32) {
	return func(characterId uint32, face uint32) {
		//TODO
	}
}

func TransportBoarding(l logrus.FieldLogger) func(characterId uint32, departureMapId uint32, destinationMapId uint32) bool {
	return func(characterId uint32, departureMapId uint32, destinationMapId uint32) bool {
		//TODO
		return false
	}
}

func RemoveFromSlot(l logrus.FieldLogger) func(characterId uint32, inventoryType string, slot int16, amount int32) {
	return func(characterId uint32, inventoryType string, slot int16, amount int32) {
		//TODO
	}
}

func GainExperience(l logrus.FieldLogger) func(characterId uint32, amount int32) {
	return func(characterId uint32, amount int32) {
		//TODO
	}
}

func RemoveAll(l logrus.FieldLogger) func(characterId uint32, itemId uint32) {
	return func(characterId uint32, itemId uint32) {
		//TODO
	}
}

func SavedLocation(l logrus.FieldLogger) func(characterId uint32, location string) uint32 {
	return func(characterId uint32, location string) uint32 {
		//TODO
		return 0
	}
}

func ClearSavedLocation(l logrus.FieldLogger) func(characterId uint32, location string) {
	return func(characterId uint32, location string) {
		//TODO
	}
}

func IsCygnus(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return MeetsCriteria(l, span)(characterId, IsCygnusCriteria())
	}
}

func IsCygnusCriteria() AttributeCriteria {
	return func(c Model) bool {
		//TODO
		return false
	}
}

func IsJobBranch(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, branch uint32) bool {
	return func(characterId uint32, branch uint32) bool {
		return MeetsCriteria(l, span)(characterId, IsJobBranchCriteria(branch))
	}
}

func IsJobBranchCriteria(branch uint32) AttributeCriteria {
	return func(c Model) bool {
		//TODO
		return false
	}
}

func UseItem(l logrus.FieldLogger) func(characterId uint32, itemId uint32) {
	return func(characterId uint32, itemId uint32) {
		//TODO
	}
}

func GuideHint(l logrus.FieldLogger) func(characterId uint32, hint uint32) {
	return func(characterId uint32, hint uint32) {
		//TODO
	}
}

func SendNotice(l logrus.FieldLogger) func(characterId uint32, noticeType string, message string) {
	return func(characterId uint32, noticeType string, message string) {
		//TODO
	}
}

func AreaInfo(l logrus.FieldLogger) func(characterId uint32, areaId uint16, property string) bool {
	return func(characterId uint32, areaId uint16, property string) bool {
		//TODO
		return false
	}
}

func ShowIntro(l logrus.FieldLogger) func(characterId uint32, path string) {
	return func(characterId uint32, path string) {
		//TODO
	}
}

func SetAreaInfo(l logrus.FieldLogger) func(characterId uint32, areaId uint16, property string) {
	return func(characterId uint32, areaId uint16, property string) {
		//TODO
	}
}

func SpawnGuide(l logrus.FieldLogger) func(characterId uint32) {
	return func(characterId uint32) {
		//TODO
	}
}

func ShowEffect(l logrus.FieldLogger) func(characterId uint32, path string) {
	return func(characterId uint32, path string) {
		//TODO
	}
}

func PlaySound(l logrus.FieldLogger) func(characterId uint32, path string) {
	return func(characterId uint32, path string) {
		//TODO
	}
}

func HasEquipped(l logrus.FieldLogger) func(characterId uint32, itemId uint32) bool {
	return func(characterId uint32, itemId uint32) bool {
		//TODO
		return false
	}
}

func TeachSkill(l logrus.FieldLogger) func(characterId uint32, skillId uint32, level byte, masterLevel byte, expiration int64) {
	return func(characterId uint32, skillId uint32, level byte, masterLevel byte, expiration int64) {
		//TODO
	}
}

func ChangeMusic(l logrus.FieldLogger) func(characterId uint32, path string) {
	return func(characterId uint32, path string) {
		//TODO
	}
}

func IsLevelBetweenCriteria(lower byte, upper byte) AttributeCriteria {
	return func(c Model) bool {
		return c.Level() >= lower && c.Level() <= upper
	}
}

func BuffSource(l logrus.FieldLogger) func(characterId uint32, source int32) uint32 {
	return func(characterId uint32, source int32) uint32 {
		//TODO
		return 0
	}
}

func ItemQuantity(l logrus.FieldLogger) func(characterId uint32, itemId uint32) uint32 {
	return func(characterId uint32, itemId uint32) uint32 {
		//TODO
		return 0
	}
}
