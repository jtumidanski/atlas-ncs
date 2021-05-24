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

type ProcessNumber func(selection int32) StateProducer

type State func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State

type ProcessSelection func(selection int32) StateProducer

func Exit() StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := npc.Processor(l).Dispose(c.CharacterId)
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
	return doListSelectionExit(Exit(), s)
}

func SendListSelectionExit(l logrus.FieldLogger, c Context, message string, s ProcessSelection, exit StateProducer) State {
	err := npc.Processor(l).Conversation(c.CharacterId, c.NPCId).SendSimple(message)
	if err != nil {
		l.WithError(err).Errorf("Sending list selection for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doListSelectionExit(exit, s)
}

func doListSelectionExit(e StateProducer, s ProcessSelection) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 0 && theType == 4 {
			return e(l, c)
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
	return doNextExit(Exit(), next)
}

func SendNextExit(l logrus.FieldLogger, c Context, message string, next StateProducer, exit StateProducer) State {
	err := npc.Processor(l).Conversation(c.CharacterId, c.NPCId).SendNext(message)
	if err != nil {
		l.WithError(err).Errorf("Sending next message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doNextExit(exit, next)
}

func doNextExit(e StateProducer, next StateProducer) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 255 && theType == 0 {
			return e(l, c)
		}
		return next(l, c)
	}
}

func SendNextPrevious(l logrus.FieldLogger, c Context, message string, next StateProducer, previous StateProducer) State {
	err := npc.Processor(l).Conversation(c.CharacterId, c.NPCId).SendNextPrevious(message)
	if err != nil {
		l.WithError(err).Errorf("Sending next / previous message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doNextPreviousExit(Exit(), next, previous)
}

func SendNextPreviousExit(l logrus.FieldLogger, c Context, message string, next StateProducer, previous StateProducer, exit StateProducer) State {
	err := npc.Processor(l).Conversation(c.CharacterId, c.NPCId).SendNextPrevious(message)
	if err != nil {
		l.WithError(err).Errorf("Sending next / previous message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doNextPreviousExit(exit, next, previous)
}

func doNextPreviousExit(e StateProducer, next StateProducer, previous StateProducer) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 255 && theType == 0 {
			return e(l, c)
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
	return doPreviousExit(Exit(), previous)
}

func doPreviousExit(e StateProducer, previous StateProducer) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 255 && theType == 0 {
			return e(l, c)
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
	return doYesNoExit(Exit(), yes, no)
}

func SendYesNoExit(l logrus.FieldLogger, c Context, message string, yes StateProducer, no StateProducer, exit StateProducer) State {
	err := npc.Processor(l).Conversation(c.CharacterId, c.NPCId).SendYesNo(message)
	if err != nil {
		l.WithError(err).Errorf("Sending yes / no message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doYesNoExit(exit, yes, no)
}

func doYesNoExit(e StateProducer, yes StateProducer, no StateProducer) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 255 && theType == 0 {
			return e(l, c)
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

func SendGetNumber(l logrus.FieldLogger, c Context, message string, s ProcessNumber, defaultValue int32, minimumValue int32, maximumValue int32) State {
	err := npc.Processor(l).Conversation(c.CharacterId, c.NPCId).SendGetNumber(message, defaultValue, minimumValue, maximumValue)
	if err != nil {
		l.WithError(err).Errorf("Sending get number for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doGetNumberExit(Exit(), s)
}

func SendGetNumberExit(l logrus.FieldLogger, c Context, message string, s ProcessNumber, e StateProducer, defaultValue int32, minimumValue int32, maximumValue int32) State {
	err := npc.Processor(l).Conversation(c.CharacterId, c.NPCId).SendGetNumber(message, defaultValue, minimumValue, maximumValue)
	if err != nil {
		l.WithError(err).Errorf("Sending get number for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doGetNumberExit(e, s)
}

func doGetNumberExit(e StateProducer, s ProcessNumber) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 0 && theType == 4 {
			return e(l, c)
		}

		f := s(selection)
		if f == nil {
			l.Errorf("unhandled selection %d for npc %d.", selection, c.NPCId)
			return nil
		}
		return f(l, c)
	}
}