package npc

import (
	"atlas-ncs/kafka/producers"
	"atlas-ncs/map/portal"
	"github.com/sirupsen/logrus"
)

const (
	MessageTypeSimple        = "SIMPLE"
	MessageTypeNext          = "NEXT"
	MessageTypeNextPrevious  = "NEXT_PREVIOUS"
	MessageTypePrevious      = "PREVIOUS"
	MessageTypeYesNo         = "YES_NO"
	MessageTypeOk            = "OK"
	MessageTypeNum           = "NUM"
	MessageTypeText          = "TEXT"
	MessageTypeStyle         = "STYLE"
	MessageTypeAcceptDecline = "ACCEPT_DECLINE"

	SpeakerNPCLeft        = "NPC_LEFT"
	SpeakerNPCRight       = "NPC_RIGHT"
	SpeakerCharacterLeft  = "CHARACTER_LEFT"
	SpeakerCharacterRight = "CHARACTER_RIGHT"
	SpeakerUnknown        = "UNKNOWN"
	SpeakerUnknown2       = "UNKNOWN2"
)

func Dispose(l logrus.FieldLogger) func(characterId uint32) error {
	return func(characterId uint32) error {
		return producers.EnableActions(l)(characterId)
	}
}

func LockUI(l logrus.FieldLogger) func(characterId uint32) {
	return func(characterId uint32) {

	}
}

func WarpToPortal(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, p portal.IdProvider) error {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, p portal.IdProvider) error {
		return producers.ChangeMap(l)(worldId, channelId, characterId, mapId, p())
	}
}

func WarpRandom(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32) error {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32) error {
		return WarpToPortal(l)(worldId, channelId, characterId, mapId, portal.RandomPortalIdProvider(l)(mapId))
	}
}

func WarpById(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) error {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) error {
		return WarpToPortal(l)(worldId, channelId, characterId, mapId, portal.FixedPortalIdProvider(portalId))
	}
}

func WarpByName(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) error {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) error {
		return WarpToPortal(l)(worldId, channelId, characterId, mapId, portal.ByNamePortalIdProvider(l)(mapId, portalName))
	}
}

type TalkConfig struct {
	messageType string
	speaker     string
}

func (c TalkConfig) MessageType() string {
	return c.messageType
}

func (c TalkConfig) Speaker() string {
	return c.speaker
}

type TalkConfigurator func(config *TalkConfig)

func SendNPCTalk(l logrus.FieldLogger, characterId uint32, npcId uint32, config *TalkConfig) func(message string, configurations ...TalkConfigurator) error {
	return func(message string, configurations ...TalkConfigurator) error {
		for _, configuration := range configurations {
			configuration(config)
		}
		return producers.NPCTalk(l)(characterId, npcId, message, config.MessageType(), config.Speaker())
	}
}

func SetSpeaker(speaker string) TalkConfigurator {
	return func(config *TalkConfig) {
		config.speaker = speaker
	}
}

type TalkFunc func(message string, configurations ...TalkConfigurator) error

func SendSimple(l logrus.FieldLogger, characterId uint32, npcId uint32) TalkFunc {
	return SendNPCTalk(l, characterId, npcId, &TalkConfig{messageType: MessageTypeSimple, speaker: SpeakerNPCLeft})
}

func SendNext(l logrus.FieldLogger, characterId uint32, npcId uint32) TalkFunc {
	return SendNPCTalk(l, characterId, npcId, &TalkConfig{messageType: MessageTypeNext, speaker: SpeakerNPCLeft})
}

func SendNextPrevious(l logrus.FieldLogger, characterId uint32, npcId uint32) TalkFunc {
	return SendNPCTalk(l, characterId, npcId, &TalkConfig{messageType: MessageTypeNextPrevious, speaker: SpeakerNPCLeft})
}

func SendPrevious(l logrus.FieldLogger, characterId uint32, npcId uint32) TalkFunc {
	return SendNPCTalk(l, characterId, npcId, &TalkConfig{messageType: MessageTypePrevious, speaker: SpeakerNPCLeft})
}

func SendYesNo(l logrus.FieldLogger, characterId uint32, npcId uint32) TalkFunc {
	return SendNPCTalk(l, characterId, npcId, &TalkConfig{messageType: MessageTypeYesNo, speaker: SpeakerNPCLeft})
}

func SendOk(l logrus.FieldLogger, characterId uint32, npcId uint32) TalkFunc {
	return SendNPCTalk(l, characterId, npcId, &TalkConfig{messageType: MessageTypeOk, speaker: SpeakerNPCLeft})
}

func SendAcceptDecline(l logrus.FieldLogger, characterId uint32, npcId uint32) TalkFunc {
	return SendNPCTalk(l, characterId, npcId, &TalkConfig{messageType: MessageTypeAcceptDecline, speaker: SpeakerNPCLeft})
}

func SendGetNumber(l logrus.FieldLogger, characterId uint32, npcId uint32) func(message string, defaultValue int32, minimumValue int32, maximumValue int32) error {
	return func(message string, defaultValue int32, minimumValue int32, maximumValue int32) error {
		return producers.NPCTalkNum(l)(characterId, npcId, message, defaultValue, minimumValue, maximumValue, MessageTypeNum, SpeakerNPCLeft)
	}
}

func SendGetText(l logrus.FieldLogger, characterId uint32, npcId uint32) func(message string) error {
	return func(message string) error {
		return producers.NPCTalkText(l)(characterId, npcId, message, MessageTypeText, SpeakerNPCLeft)
	}
}

func SendStyle(l logrus.FieldLogger, characterId uint32, npcId uint32) func(message string, options []uint32) error {
	return func(message string, options []uint32) error {
		return producers.NPCTalkStyle(l)(characterId, npcId, message, options, MessageTypeStyle, SpeakerNPCLeft)
	}
}

func Spawn(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, npcId uint32, x int16, y int16) {
	return func(worldId byte, channelId byte, mapId uint32, npcId uint32, x int16, y int16) {

	}
}

func Destroy(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, npcId uint32) {
	return func(worldId byte, channelId byte, mapId uint32, npcId uint32) {

	}
}

func SendDimensionalMirror(l logrus.FieldLogger, characterId uint32, npcId uint32) func(message string) error {
	return func(message string) error {
		return nil
	}
}

func OpenShop(l logrus.FieldLogger) func(characterId uint32, shopId uint32) {
	return func(characterId uint32, shopId uint32) {

	}
}
