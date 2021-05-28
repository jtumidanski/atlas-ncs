package npc

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/kafka/producers"
	"github.com/sirupsen/logrus"
)

const (
	MessageTypeSimple       = "SIMPLE"
	MessageTypeNext         = "NEXT"
	MessageTypeNextPrevious = "NEXT_PREVIOUS"
	MessageTypePrevious     = "PREVIOUS"
	MessageTypeYesNo        = "YES_NO"
	MessageTypeOk           = "OK"
	MessageTypeNum          = "NUM"
	MessageTypeStyle        = "STYLE"

	SpeakerNPCLeft = "NPC_LEFT"
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

func WarpById(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) error {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalId uint32) error {
		return producers.ChangeMap(l)(worldId, channelId, characterId, mapId, portalId)
	}
}

func WarpByName(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) error {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, portalName string) error {
		return nil
	}
}

func SendSimple(l logrus.FieldLogger, c script.Context) func(message string) error {
	return func(message string) error {
		return producers.NPCTalk(l)(c.CharacterId, c.NPCId, message, MessageTypeSimple, SpeakerNPCLeft)
	}
}

func SendNext(l logrus.FieldLogger, c script.Context) func(message string) error {
	return func(message string) error {
		return producers.NPCTalk(l)(c.CharacterId, c.NPCId, message, MessageTypeNext, SpeakerNPCLeft)
	}
}

func SendNextPrevious(l logrus.FieldLogger, c script.Context) func(message string) error {
	return func(message string) error {
		return producers.NPCTalk(l)(c.CharacterId, c.NPCId, message, MessageTypeNextPrevious, SpeakerNPCLeft)
	}
}

func SendPrevious(l logrus.FieldLogger, c script.Context) func(message string) error {
	return func(message string) error {
		return producers.NPCTalk(l)(c.CharacterId, c.NPCId, message, MessageTypePrevious, SpeakerNPCLeft)
	}
}

func SendYesNo(l logrus.FieldLogger, c script.Context) func(message string) error {
	return func(message string) error {
		return producers.NPCTalk(l)(c.CharacterId, c.NPCId, message, MessageTypeYesNo, SpeakerNPCLeft)
	}
}

func SendOk(l logrus.FieldLogger, c script.Context) func(message string) error {
	return func(message string) error {
		return producers.NPCTalk(l)(c.CharacterId, c.NPCId, message, MessageTypeOk, SpeakerNPCLeft)
	}
}

func SendGetNumber(l logrus.FieldLogger, c script.Context) func(message string, defaultValue int32, minimumValue int32, maximumValue int32) error {
	return func(message string, defaultValue int32, minimumValue int32, maximumValue int32) error {
		return producers.NPCTalkNum(l)(c.CharacterId, c.NPCId, message, defaultValue, minimumValue, maximumValue, MessageTypeNum, SpeakerNPCLeft)
	}
}

func SendStyle(l logrus.FieldLogger, c script.Context) func(message string, options []uint32) error {
	return func(message string, options []uint32) error {
		return producers.NPCTalkStyle(l)(c.CharacterId, c.NPCId, message, options, MessageTypeStyle, SpeakerNPCLeft)
	}
}
