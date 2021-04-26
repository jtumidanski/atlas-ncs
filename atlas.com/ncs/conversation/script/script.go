package script

import "github.com/sirupsen/logrus"

type Context struct {
	WorldId     byte
	ChannelId   byte
	CharacterId uint32
	MapId       uint32
	NPCId       uint32
}

type Script interface {
	NPCId() uint32

	Start(l logrus.FieldLogger, c Context)

	Continue(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32)
}
