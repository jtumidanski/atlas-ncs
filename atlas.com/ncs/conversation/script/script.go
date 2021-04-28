package script

import (
	"atlas-ncs/npc"
	"github.com/sirupsen/logrus"
)

type Context struct {
	WorldId     byte
	ChannelId   byte
	CharacterId uint32
	MapId       uint32
	NPCId       uint32
}

type Script interface {
	NPCId() uint32

	Initial() StateProducer
}

type StateProducer func(l logrus.FieldLogger, c Context) State

type State func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State

type ProcessSelection func(selection int32) StateProducer

type ExitFunction func(l logrus.FieldLogger, c Context)

func GenericExit(l logrus.FieldLogger, c Context) {
	npc.Processor(l).Dispose(c.CharacterId)
}

func ListSelection(e ExitFunction, s ProcessSelection) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 0 && theType == 4 {
			e(l, c)
			return nil
		}

		f := s(selection)
		if f == nil {
			l.Errorf("unhandled selection %d for npc %d.", selection, c.NPCId)
			return nil
		}
		return f(l, c)
	}
}

func Next(e ExitFunction, next StateProducer) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 255 && theType == 0 {
			e(l, c)
			return nil
		}
		return next(l, c)
	}
}

func NextPrevious(e ExitFunction, next StateProducer, previous StateProducer) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 255 && theType == 0 {
			e(l, c)
			return nil
		}
		if mode == 0 && previous != nil {
			return previous(l, c)
		} else if mode == 1 && next != nil {
			return next(l, c)
		}
		return nil
	}
}

func Previous(e ExitFunction, previous StateProducer) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 255 && theType == 0 {
			e(l, c)
			return nil
		}
		if mode == 0 && previous != nil {
			return previous(l, c)
		}
		return nil
	}
}

func YesNo(e ExitFunction, yes StateProducer, no StateProducer) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 255 && theType == 0 {
			e(l, c)
			return nil
		}
		if mode == 0 && no != nil {
			return no(l, c)
		} else if mode == 1 && yes != nil {
			return yes(l, c)
		}
		return nil
	}
}