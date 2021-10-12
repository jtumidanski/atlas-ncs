package npc

import (
	"atlas-ncs/kafka/producers"
	"atlas-ncs/map/portal"
	"github.com/opentracing/opentracing-go"
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

func Dispose(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) {
	return func(characterId uint32) {
		producers.EnableActions(l, span)(characterId)
	}
}

func LockUI(l logrus.FieldLogger) func(characterId uint32) {
	return func(characterId uint32) {

	}
}

func WarpToPortal(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32, p portal.IdProvider) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, p portal.IdProvider) {
		producers.ChangeMap(l, span)(worldId, channelId, characterId, mapId, p())
	}
}

func WarpRandom(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32) {
		WarpToPortal(l, span)(worldId, channelId, characterId, mapId, portal.RandomPortalIdProvider(l)(mapId))
	}
}

func WarpById(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) {
		WarpToPortal(l, span)(worldId, channelId, characterId, mapId, portal.FixedPortalIdProvider(portalId))
	}
}

func WarpByName(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) {
		WarpToPortal(l, span)(worldId, channelId, characterId, mapId, portal.ByNamePortalIdProvider(l)(mapId, portalName))
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

func SendNPCTalk(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32, config *TalkConfig) func(message string, configurations ...TalkConfigurator) {
	return func(characterId uint32, npcId uint32, config *TalkConfig) func(message string, configurations ...TalkConfigurator) {
		return func(message string, configurations ...TalkConfigurator) {
			for _, configuration := range configurations {
				configuration(config)
			}
			producers.NPCTalk(l, span)(characterId, npcId, message, config.MessageType(), config.Speaker())
		}
	}
}

func SetSpeaker(speaker string) TalkConfigurator {
	return func(config *TalkConfig) {
		config.speaker = speaker
	}
}

type TalkFunc func(message string, configurations ...TalkConfigurator)

func SendSimple(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32) TalkFunc {
	return func(characterId uint32, npcId uint32) TalkFunc {
		return SendNPCTalk(l, span)(characterId, npcId, &TalkConfig{messageType: MessageTypeSimple, speaker: SpeakerNPCLeft})
	}
}

func SendNext(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32) TalkFunc {
	return func(characterId uint32, npcId uint32) TalkFunc {
		return SendNPCTalk(l, span)(characterId, npcId, &TalkConfig{messageType: MessageTypeNext, speaker: SpeakerNPCLeft})
	}
}

func SendNextPrevious(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32) TalkFunc {
	return func(characterId uint32, npcId uint32) TalkFunc {
		return SendNPCTalk(l, span)(characterId, npcId, &TalkConfig{messageType: MessageTypeNextPrevious, speaker: SpeakerNPCLeft})
	}
}

func SendPrevious(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32) TalkFunc {
	return func(characterId uint32, npcId uint32) TalkFunc {
		return SendNPCTalk(l, span)(characterId, npcId, &TalkConfig{messageType: MessageTypePrevious, speaker: SpeakerNPCLeft})
	}
}

func SendYesNo(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32) TalkFunc {
	return func(characterId uint32, npcId uint32) TalkFunc {
		return SendNPCTalk(l, span)(characterId, npcId, &TalkConfig{messageType: MessageTypeYesNo, speaker: SpeakerNPCLeft})
	}
}

func SendOk(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32) TalkFunc {
	return func(characterId uint32, npcId uint32) TalkFunc {
		return SendNPCTalk(l, span)(characterId, npcId, &TalkConfig{messageType: MessageTypeOk, speaker: SpeakerNPCLeft})
	}
}

func SendAcceptDecline(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32) TalkFunc {
	return func(characterId uint32, npcId uint32) TalkFunc {
		return SendNPCTalk(l, span)(characterId, npcId, &TalkConfig{messageType: MessageTypeAcceptDecline, speaker: SpeakerNPCLeft})
	}
}

func SendGetNumber(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32) func(message string, defaultValue int32, minimumValue int32, maximumValue int32) {
	return func(characterId uint32, npcId uint32) func(message string, defaultValue int32, minimumValue int32, maximumValue int32) {
		return func(message string, defaultValue int32, minimumValue int32, maximumValue int32) {
			producers.NPCTalkNum(l, span)(characterId, npcId, message, defaultValue, minimumValue, maximumValue, MessageTypeNum, SpeakerNPCLeft)
		}
	}
}

func SendGetText(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32) func(message string) {
	return func(characterId uint32, npcId uint32) func(message string) {
		return func(message string) {
			producers.NPCTalkText(l, span)(characterId, npcId, message, MessageTypeText, SpeakerNPCLeft)
		}
	}
}

func SendStyle(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, npcId uint32) func(message string, options []uint32) {
	return func(characterId uint32, npcId uint32) func(message string, options []uint32) {
		return func(message string, options []uint32) {
			producers.NPCTalkStyle(l, span)(characterId, npcId, message, options, MessageTypeStyle, SpeakerNPCLeft)
		}
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
