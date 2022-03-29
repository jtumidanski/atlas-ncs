package conversation

import (
	"atlas-ncs/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	consumerNameContinue      = "continue_npc_conversation_command"
	consumerNameStart         = "start_npc_conversation_command"
	consumerNameSetReturnText = "set_return_text_command"
	topicTokenContinue        = "TOPIC_CONTINUE_NPC_CONVERSATION"
	topicTokenStart           = "TOPIC_START_NPC_CONVERSATION"
	topicTokenSetReturnText   = "TOPIC_SET_RETURN_TEXT"
)

func ContinueConsumer(groupId string) kafka.ConsumerConfig {
	return kafka.NewConsumerConfig[continueCommand](consumerNameContinue, topicTokenContinue, groupId, HandleContinue())
}

type continueCommand struct {
	CharacterId uint32 `json:"characterId"`
	Mode        byte   `json:"mode"`
	Type        byte   `json:"type"`
	Selection   int32  `json:"selection"`
}

func HandleContinue() kafka.HandlerFunc[continueCommand] {
	return func(l logrus.FieldLogger, span opentracing.Span, command continueCommand) {
		Continue(l, span)(command.CharacterId, command.Mode, command.Type, command.Selection)
	}
}

func StartConsumer(groupId string) kafka.ConsumerConfig {
	return kafka.NewConsumerConfig[startCommand](consumerNameStart, topicTokenStart, groupId, HandleStart())
}

type startCommand struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	MapId       uint32 `json:"mapId"`
	CharacterId uint32 `json:"characterId"`
	NPCId       uint32 `json:"npcId"`
	NPCObjectId uint32 `json:"npcObjectId"`
}

func HandleStart() kafka.HandlerFunc[startCommand] {
	return func(l logrus.FieldLogger, span opentracing.Span, event startCommand) {
		Start(l, span)(event.WorldId, event.ChannelId, event.MapId, event.NPCId, event.NPCObjectId, event.CharacterId)
	}
}

func SetReturnTextConsumer(groupId string) kafka.ConsumerConfig {
	return kafka.NewConsumerConfig[setReturnTextCommand](consumerNameSetReturnText, topicTokenSetReturnText, groupId, HandleSetReturnText())
}

type setReturnTextCommand struct {
	CharacterId uint32 `json:"characterId"`
	Text        string `json:"text"`
}

func HandleSetReturnText() kafka.HandlerFunc[setReturnTextCommand] {
	return func(l logrus.FieldLogger, span opentracing.Span, event setReturnTextCommand) {
		l.Debugf("Handling SetReturnText command for character %d, with value %s.", event.CharacterId, event.Text)
	}
}
