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

	Initial(l logrus.FieldLogger, c Context) State
}

type StateProducer func(l logrus.FieldLogger, c Context) State

type State func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State

type ProcessSelection func(selection int32) StateProducer

type ExitFunction func(l logrus.FieldLogger, c Context) error

func GenericExit(l logrus.FieldLogger, c Context) error {
	return npc.Processor(l).Dispose(c.CharacterId)
}

func Exit() StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := GenericExit(l, c)
		if err != nil {
			l.WithError(err).Errorf("Unable to dispose conversation.")
		}
		return nil
	}
}

func SendListSelection(l logrus.FieldLogger, c Context, message string, s ProcessSelection) State {
	err := npc.Processor(l).Conversation(c.CharacterId, c.NPCId).SendSimple(message)
	if err != nil {
		l.WithError(err).Errorf("Sending list selection for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doListSelection(GenericExit, s)
}

func doListSelection(e ExitFunction, s ProcessSelection) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 0 && theType == 4 {
			err := e(l, c)
			if err != nil {
				l.WithError(err).Errorf("Error exiting conversation.")
			}
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

func SendNext(l logrus.FieldLogger, c Context, message string, next StateProducer) State {
	err := npc.Processor(l).Conversation(c.CharacterId, c.NPCId).SendNext(message)
	if err != nil {
		l.WithError(err).Errorf("Sending next message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doNext(GenericExit, next)
}

func doNext(e ExitFunction, next StateProducer) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 255 && theType == 0 {
			err := e(l, c)
			if err != nil {
				l.WithError(err).Errorf("Error exiting conversation.")
			}
			return nil
		}
		return next(l, c)
	}
}

func SendNextPrevious(l logrus.FieldLogger, c Context, message string, next StateProducer, previous StateProducer) State {
	err := npc.Processor(l).Conversation(c.CharacterId, c.NPCId).SendNextPrevious(message)
	if err != nil {
		l.WithError(err).Errorf("Sending next / previous message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doNextPrevious(GenericExit, next, previous)
}

func doNextPrevious(e ExitFunction, next StateProducer, previous StateProducer) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 255 && theType == 0 {
			err := e(l, c)
			if err != nil {
				l.WithError(err).Errorf("Error exiting conversation.")
			}
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

func SendPrevious(l logrus.FieldLogger, c Context, message string, previous StateProducer) State {
	err := npc.Processor(l).Conversation(c.CharacterId, c.NPCId).SendPrevious(message)
	if err != nil {
		l.WithError(err).Errorf("Sending previous message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doPrevious(GenericExit, previous)
}

func doPrevious(e ExitFunction, previous StateProducer) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 255 && theType == 0 {
			err := e(l, c)
			if err != nil {
				l.WithError(err).Errorf("Error exiting conversation.")
			}
			return nil
		}
		if mode == 0 && previous != nil {
			return previous(l, c)
		}
		return nil
	}
}

func SendYesNo(l logrus.FieldLogger, c Context, message string, yes StateProducer, no StateProducer) State {
	err := npc.Processor(l).Conversation(c.CharacterId, c.NPCId).SendYesNo(message)
	if err != nil {
		l.WithError(err).Errorf("Sending yes / no message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doYesNo(GenericExit, yes, no)
}

func doYesNo(e ExitFunction, yes StateProducer, no StateProducer) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 255 && theType == 0 {
			err := e(l, c)
			if err != nil {
				l.WithError(err).Errorf("Error exiting conversation.")
			}
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

func SendOk(l logrus.FieldLogger, c Context, message string) State {
	err := npc.Processor(l).Conversation(c.CharacterId, c.NPCId).SendOk(message)
	if err != nil {
		l.WithError(err).Errorf("Sending ok message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return nil
}
