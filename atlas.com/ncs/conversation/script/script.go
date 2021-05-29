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
	NPCObjectId uint32
}

type Script interface {
	NPCId() uint32

	Initial(l logrus.FieldLogger, c Context) State
}

type StateProducer func(l logrus.FieldLogger, c Context) State

type ProcessNumber func(selection int32) StateProducer

type ProcessText func(text string) StateProducer

type State func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State

type ProcessSelection func(selection int32) StateProducer

func Exit() StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := npc.Dispose(l)(c.CharacterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to dispose conversation.")
		}
		return nil
	}
}

func SendListSelection(l logrus.FieldLogger, c Context, message string, s ProcessSelection) State {
	err := npc.SendSimple(l, c)(message)
	if err != nil {
		l.WithError(err).Errorf("Sending list selection for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doListSelectionExit(Exit(), s)
}

func SendListSelectionExit(l logrus.FieldLogger, c Context, message string, s ProcessSelection, exit StateProducer) State {
	err := npc.SendSimple(l, c)(message)
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
	err := npc.SendNext(l, c)(message)
	if err != nil {
		l.WithError(err).Errorf("Sending next message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doNextExit(Exit(), next)
}

func SendNextExit(l logrus.FieldLogger, c Context, message string, next StateProducer, exit StateProducer) State {
	err := npc.SendNext(l, c)(message)
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
	err := npc.SendNextPrevious(l, c)(message)
	if err != nil {
		l.WithError(err).Errorf("Sending next / previous message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doNextPreviousExit(Exit(), next, previous)
}

func SendNextPreviousExit(l logrus.FieldLogger, c Context, message string, next StateProducer, previous StateProducer, exit StateProducer) State {
	err := npc.SendNextPrevious(l, c)(message)
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
	err := npc.SendPrevious(l, c)(message)
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
	err := npc.SendYesNo(l, c)(message)
	if err != nil {
		l.WithError(err).Errorf("Sending yes / no message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doYesNoExit(Exit(), yes, no)
}

func SendYesNoExit(l logrus.FieldLogger, c Context, message string, yes StateProducer, no StateProducer, exit StateProducer) State {
	err := npc.SendYesNo(l, c)(message)
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
	err := npc.SendOk(l, c)(message)
	if err != nil {
		l.WithError(err).Errorf("Sending ok message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return nil
}

func SendOkTrigger(l logrus.FieldLogger, c Context, message string, next StateProducer) State {
	err := npc.SendOk(l, c)(message)
	if err != nil {
		l.WithError(err).Errorf("Sending ok message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return next(l, c)
}

func SendGetNumber(l logrus.FieldLogger, c Context, message string, s ProcessNumber, defaultValue int32, minimumValue int32, maximumValue int32) State {
	err := npc.SendGetNumber(l, c)(message, defaultValue, minimumValue, maximumValue)
	if err != nil {
		l.WithError(err).Errorf("Sending get number for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doGetNumberExit(Exit(), s)
}

func SendGetNumberExit(l logrus.FieldLogger, c Context, message string, s ProcessNumber, e StateProducer, defaultValue int32, minimumValue int32, maximumValue int32) State {
	err := npc.SendGetNumber(l, c)(message, defaultValue, minimumValue, maximumValue)
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

func SendGetText(l logrus.FieldLogger, c Context, message string, s ProcessText) State {
	return SendGetTextExit(l, c, message, s, Exit())
}

func SendGetTextExit(l logrus.FieldLogger, c Context, message string, s ProcessText, e StateProducer) State {
	err := npc.SendGetText(l, c)(message)
	if err != nil {
		l.WithError(err).Errorf("Sending get number for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doGetTextExit(e, s)
}

func doGetTextExit(e StateProducer, s ProcessText) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 0 && theType == 4 {
			return e(l, c)
		}

		//TODO get text somehow
		text := ""
		f := s(text)
		if f == nil {
			l.Errorf("unhandled selection %d for npc %d.", selection, c.NPCId)
			return nil
		}
		return f(l, c)
	}
}

func SendStyle(l logrus.FieldLogger, c Context, message string, next ProcessSelection, options []uint32) State {
	err := npc.SendStyle(l, c)(message, options)
	if err != nil {
		l.WithError(err).Errorf("Sending style for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doSendStyleExit(Exit(), next)
}

func doSendStyleExit(e StateProducer, next ProcessSelection) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 0 && theType == 7 {
			return e(l, c)
		}
		return next(selection)(l, c)
	}
}

func SendAcceptDecline(l logrus.FieldLogger, c Context, message string, accept StateProducer, decline StateProducer) State {
	err := npc.SendAcceptDecline(l, c)(message)
	if err != nil {
		l.WithError(err).Errorf("Sending yes / no message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doAcceptDeclineExit(Exit(), accept, decline)
}

func SendAcceptDeclineExit(l logrus.FieldLogger, c Context, message string, accept StateProducer, decline StateProducer, exit StateProducer) State {
	err := npc.SendAcceptDecline(l, c)(message)
	if err != nil {
		l.WithError(err).Errorf("Sending yes / no message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doAcceptDeclineExit(exit, accept, decline)
}

func doAcceptDeclineExit(e StateProducer, accept StateProducer, decline StateProducer) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 255 && theType == 0 {
			return e(l, c)
		}
		if mode == 0 && decline != nil {
			return decline(l, c)
		} else if mode == 1 && accept != nil {
			return accept(l, c)
		}
		return nil
	}
}
