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
		strength:     att.Strength,
		dexterity:    att.Dexterity,
		intelligence: att.Intelligence,
		mapId:        att.MapId,
		gender:       att.Gender,
		hair:         att.Hair,
		face:         att.Face,
	}
	return &r
}

type AttributeCriteria func(*Model) bool

func MeetsCriteria(l logrus.FieldLogger) func(characterId uint32, criteria ...AttributeCriteria) bool {
	return func(characterId uint32, criteria ...AttributeCriteria) bool {
		c, err := GetCharacterById(characterId)
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

func HasAnyItem(l logrus.FieldLogger) func(characterId uint32, items ...uint32) bool {
	return func(characterId uint32, items ...uint32) bool {
		allItems, err := requestAllItemsForCharacter(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve inventory items for character %d.", characterId)
			return false
		}

		for _, i := range allItems.Data {
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
		return MeetsCriteria(l)(characterId, IsLevelCriteria(level))
	}
}

func IsLevelCriteria(level byte) AttributeCriteria {
	return func(c *Model) bool {
		return c.Level() >= level
	}
}

func LevelBetweenCriteria(lower byte, upper byte) AttributeCriteria {
	return func(c *Model) bool {
		return c.Level() > lower && c.Level() < upper
	}
}

func HasStrength(l logrus.FieldLogger) func(characterId uint32, amount uint16) bool {
	return func(characterId uint32, amount uint16) bool {
		return MeetsCriteria(l)(characterId, HasStrengthCriteria(amount))
	}
}

func HasStrengthCriteria(amount uint16) AttributeCriteria {
	return func(c *Model) bool {
		return c.Strength() >= amount
	}
}

func HasDexterity(l logrus.FieldLogger) func(characterId uint32, amount uint16) bool {
	return func(characterId uint32, amount uint16) bool {
		return MeetsCriteria(l)(characterId, HasDexterityCriteria(amount))
	}
}

func HasDexterityCriteria(amount uint16) AttributeCriteria {
	return func(c *Model) bool {
		return c.Dexterity() >= amount
	}
}

func HasIntelligence(l logrus.FieldLogger) func(characterId uint32, amount uint16) bool {
	return func(characterId uint32, amount uint16) bool {
		return MeetsCriteria(l)(characterId, HasIntelligenceCriteria(amount))
	}
}

func HasIntelligenceCriteria(amount uint16) AttributeCriteria {
	return func(c *Model) bool {
		return c.Intelligence() >= amount
	}
}

func AboveLevel(l logrus.FieldLogger) func(characterId uint32, level byte) bool {
	return func(characterId uint32, level byte) bool {
		return MeetsCriteria(l)(characterId, AboveLevelCriteria(level))
	}
}

func AboveLevelCriteria(level byte) AttributeCriteria {
	return func(c *Model) bool {
		return c.Level() > level
	}
}

func HasMeso(l logrus.FieldLogger) func(characterId uint32, amount uint32) bool {
	return func(characterId uint32, amount uint32) bool {
		return MeetsCriteria(l)(characterId, HasMesoCriteria(amount))
	}
}

func HasMesoCriteria(amount uint32) AttributeCriteria {
	return func(c *Model) bool {
		return c.Meso() >= amount
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

func GainFame(l logrus.FieldLogger) func(characterId uint32, amount int32) {
	return func(characterId uint32, amount int32) {

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
		return MeetsCriteria(l)(characterId, IsBeginnerTreeCriteria())
	}
}

func IsBeginnerTreeCriteria() AttributeCriteria {
	return func(c *Model) bool {
		return job.IsA(c.JobId(), job.Beginner, job.Noblesse, job.Legend)
	}
}

func IsJob(l logrus.FieldLogger) func(characterId uint32, option uint16) bool {
	return func(characterId uint32, option uint16) bool {
		return MeetsCriteria(l)(characterId, IsJobCriteria(option))
	}
}

func IsJobCriteria(option uint16) AttributeCriteria {
	return func(c *Model) bool {
		return job.IsA(c.JobId(), option)
	}
}

func IsAJobCriteria(options ...uint16) AttributeCriteria {
	return func(c *Model) bool {
		return job.IsA(c.JobId(), options...)
	}
}

func CompleteQuest(l logrus.FieldLogger) func(characterId uint32, questId uint32) {
	return func(characterId uint32, questId uint32) {

	}
}

func CompleteQuestViaNPC(l logrus.FieldLogger) func(characterId uint32, questId uint32, npcId uint32) {
	return func(characterId uint32, questId uint32, npcId uint32) {

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

func HasPet(l logrus.FieldLogger) func(characterId uint32, slot int16) bool {
	return func(characterId uint32, slot int16) bool {
		return false
	}
}

func PetIs(l logrus.FieldLogger) func(characterId uint32, slot int16, petId ...uint32) bool {
	return func(characterId uint32, slot int16, petId ...uint32) bool {
		return false
	}
}

func PetIsLevel(l logrus.FieldLogger) func(characterId uint32, slot int16, level byte) bool {
	return func(characterId uint32, slot int16, level byte) bool {
		return false
	}
}

func EvolvePet(l logrus.FieldLogger) func(characterId uint32, slot int16, itemId uint32) {
	return func(characterId uint32, slot int16, itemId uint32) {

	}
}

func GainCloseness(l logrus.FieldLogger) func(characterId uint32, amount int8) {
	return func(characterId uint32, amount int8) {

	}
}

func GetGender(l logrus.FieldLogger) func(characterId uint32) byte {
	return func(characterId uint32) byte {
		c, err := GetCharacterById(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character %d.", characterId)
			return 0
		}
		return c.Gender()
	}
}

func GetHair(l logrus.FieldLogger) func(characterId uint32) uint32 {
	return func(characterId uint32) uint32 {
		c, err := GetCharacterById(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character %d.", characterId)
			return 0
		}
		return c.Hair()
	}
}

func GetFace(l logrus.FieldLogger) func(characterId uint32) uint32 {
	return func(characterId uint32) uint32 {
		c, err := GetCharacterById(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve character %d.", characterId)
			return 0
		}
		return c.Face()
	}
}

func SetHair(l logrus.FieldLogger) func(characterId uint32, hair uint32) {
	return func(characterId uint32, hair uint32) {

	}
}

func SetSkin(l logrus.FieldLogger) func(characterId uint32, skin byte) {
	return func(characterId uint32, skin byte) {

	}
}

func SetFace(l logrus.FieldLogger) func(characterId uint32, face uint32) {
	return func(characterId uint32, face uint32) {

	}
}

func QuestNotStarted(l logrus.FieldLogger) func(characterId uint32, questId uint32) bool {
	return func(characterId uint32, questId uint32) bool {
		return true
	}
}

func ForceCompleteQuest(l logrus.FieldLogger) func(characterId uint32, questId uint32) {
	return func(characterId uint32, questId uint32) {

	}
}

func AnyQuestActive(l logrus.FieldLogger) func(characterId uint32, questId ...uint32) bool {
	return func(characterId uint32, questId ...uint32) bool {
		for _, q := range questId {
			active := QuestActive(l)(characterId, q)
			if active {
				return true
			}
		}
		return false
	}
}

func QuestActive(l logrus.FieldLogger) func(characterId uint32, questId uint32) bool {
	return func(characterId uint32, questId uint32) bool {
		return false
	}
}

func TransportBoarding(l logrus.FieldLogger) func(characterId uint32, departureMapId uint32, destinationMapId uint32) bool {
	return func(characterId uint32, departureMapId uint32, destinationMapId uint32) bool {
		return false
	}
}

func RemoveFromSlot(l logrus.FieldLogger) func(characterId uint32, inventoryType string, slot int16, amount int32) {
	return func(characterId uint32, inventoryType string, slot int16, amount int32) {

	}
}

func GainExperience(l logrus.FieldLogger) func(characterId uint32, amount int32) {
	return func(characterId uint32, amount int32) {

	}
}

func HasParty(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return false
	}
}

func QuestProgress(l logrus.FieldLogger) func(characterId uint32, questId uint32) string {
	return func(characterId uint32, questId uint32) string {
		return ""
	}
}

func QuestProgressInt(l logrus.FieldLogger) func(characterId uint32, questId uint32, infoNumber int) int {
	return func(characterId uint32, questId uint32, infoNumber int) int {
		return 0
	}
}

func SetQuestProgress(l logrus.FieldLogger) func(characterId uint32, questId uint32, infoNumber int, progress uint32) {
	return func(characterId uint32, questId uint32, infoNumber int, progress uint32) {

	}
}

func SetQuestProgressString(l logrus.FieldLogger) func(characterId uint32, questId uint32, progress string) {
	return func(characterId uint32, questId uint32, progress string) {

	}
}

func RemoveAll(l logrus.FieldLogger) func(characterId uint32, itemId uint32) {
	return func(characterId uint32, itemId uint32) {

	}
}

func SavedLocation(l logrus.FieldLogger) func(characterId uint32, location string) uint32 {
	return func(characterId uint32, location string) uint32 {
		return 0
	}
}

func ClearSavedLocation(l logrus.FieldLogger) func(characterId uint32, location string) {
	return func(characterId uint32, location string) {
	}
}

func IsCygnus(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return MeetsCriteria(l)(characterId, IsCygnusCriteria())
	}
}

func IsCygnusCriteria() AttributeCriteria {
	return func(c *Model) bool {
		return false
	}
}

func IsJobBranch(l logrus.FieldLogger) func(characterId uint32, branch uint32) bool {
	return func(characterId uint32, branch uint32) bool {
		return MeetsCriteria(l)(characterId, IsJobBranchCriteria(branch))
	}
}

func IsJobBranchCriteria(branch uint32) AttributeCriteria {
	return func(c *Model) bool {
		return false
	}
}

func UseItem(l logrus.FieldLogger) func(characterId uint32, itemId uint32) {
	return func(characterId uint32, itemId uint32) {

	}
}

func GuideHint(l logrus.FieldLogger) func(characterId uint32, hint uint32) {
	return func(characterId uint32, hint uint32) {

	}
}

func SendNotice(l logrus.FieldLogger) func(characterId uint32, noticeType string, message string) {
	return func(characterId uint32, noticeType string, message string) {

	}
}

func AreaInfo(l logrus.FieldLogger) func(characterId uint32, areaId uint16, property string) bool {
	return func(characterId uint32, areaId uint16, property string) bool {
		return false
	}
}

func ShowIntro(l logrus.FieldLogger) func(characterId uint32, path string) {
	return func(characterId uint32, path string) {

	}
}

func SetAreaInfo(l logrus.FieldLogger) func(characterId uint32, areaId uint16, property string) {
	return func(characterId uint32, areaId uint16, property string) {

	}
}

func SpawnGuide(l logrus.FieldLogger) func(characterId uint32) {
	return func(characterId uint32) {

	}
}

func HasGuild(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return false
	}
}

func IsGuildLeader(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return false
	}
}

func GuildHasAlliance(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return false
	}
}

func AllianceLeader(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return false
	}
}

func ValidAllianceName(l logrus.FieldLogger) func(text string) bool {
	return func(text string) bool {
		return false
	}
}

func CreateAlliance(l logrus.FieldLogger) func(characterId uint32, name string) error {
	return func(characterId uint32, name string) error {
		return nil
	}
}

func AllianceAtCapacity(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		return false
	}
}

func ExpandAlliance(l logrus.FieldLogger) func(characterId uint32) error {
	return func(characterId uint32) error {
		return nil
	}
}

func ShowEffect(l logrus.FieldLogger) func(characterId uint32, path string) {
	return func(characterId uint32, path string) {

	}
}

func PlaySound(l logrus.FieldLogger) func(characterId uint32, path string) {
	return func(characterId uint32, path string) {

	}
}

func HasEquipped(l logrus.FieldLogger) func(characterId uint32, itemId uint32) bool {
	return func(characterId uint32, itemId uint32) bool {
		return false
	}
}

func TeachSkill(l logrus.FieldLogger) func(characterId uint32, skillId uint32, level byte, masterLevel byte, expiration int64) {
	return func(characterId uint32, skillId uint32, level byte, masterLevel byte, expiration int64) {

	}
}

func ChangeMusic(l logrus.FieldLogger) func(characterId uint32, path string) {
	return func(characterId uint32, path string) {

	}
}

func IsLevelBetweenCriteria(lower byte, upper byte) AttributeCriteria {
	return func(c *Model) bool {
		return c.Level() >= lower && c.Level() <= upper
	}
}

func BuffSource(l logrus.FieldLogger) func(characterId uint32, source int32) uint32 {
	return func(characterId uint32, source int32) uint32 {
		return 0
	}
}
