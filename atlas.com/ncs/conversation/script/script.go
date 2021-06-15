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
	err := npc.SendSimple(l, c.CharacterId, c.NPCId)(message)
	if err != nil {
		l.WithError(err).Errorf("Sending list selection for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doListSelectionExit(Exit(), s)
}

func SendListSelectionExit(l logrus.FieldLogger, c Context, message string, s ProcessSelection, exit StateProducer) State {
	err := npc.SendSimple(l, c.CharacterId, c.NPCId)(message)
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

type SendTalkConfig struct {
	configurators []npc.TalkConfigurator
	exit          StateProducer
}

func (c SendTalkConfig) Exit() StateProducer {
	return c.exit
}

func (c SendTalkConfig) Configurators() []npc.TalkConfigurator {
	return c.configurators
}

type SendTalkConfigurator func(config *SendTalkConfig)

func AddSendTalkConfigurator(configurator npc.TalkConfigurator) SendTalkConfigurator {
	return func(config *SendTalkConfig) {
		config.configurators = append(config.configurators, configurator)
	}
}

func SetSendTalkExit(exit StateProducer) SendTalkConfigurator {
	return func(config *SendTalkConfig) {
		config.exit = exit
	}
}

type ProcessStateFunc func(exit StateProducer) State

func sendTalk(l logrus.FieldLogger, c Context, message string, configurations []SendTalkConfigurator, talkFunc npc.TalkFunc, do ProcessStateFunc) State {
	baseConfig := &SendTalkConfig{configurators: make([]npc.TalkConfigurator, 0), exit: Exit()}
	for _, configuration := range configurations {
		configuration(baseConfig)
	}

	err := talkFunc(message, baseConfig.Configurators()...)
	if err != nil {
		l.WithError(err).Errorf("Sending next message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return do(baseConfig.Exit())
}

func SendNext(l logrus.FieldLogger, c Context, message string, next StateProducer, configurations ...SendTalkConfigurator) State {
	return sendTalk(l, c, message, configurations, npc.SendNext(l, c.CharacterId, c.NPCId), doNext(next))
}

func SendNextSpeaker(l logrus.FieldLogger, c Context, message string, speaker string, next StateProducer) State {
	return SendNext(l, c, message, next, AddSendTalkConfigurator(npc.SetSpeaker(speaker)))
}

func SendNextExit(l logrus.FieldLogger, c Context, message string, next StateProducer, exit StateProducer) State {
	return SendNext(l, c, message, next, SetSendTalkExit(exit))
}

func doNext(next StateProducer) ProcessStateFunc {
	return func(exit StateProducer) State {
		return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
			if mode == 255 && theType == 0 {
				return exit(l, c)
			}
			return next(l, c)
		}
	}
}

func SendNextPrevious(l logrus.FieldLogger, c Context, message string, next StateProducer, previous StateProducer, configurations ...SendTalkConfigurator) State {
	return sendTalk(l, c, message, configurations, npc.SendNextPrevious(l, c.CharacterId, c.NPCId), doNextPrevious(next, previous))
}

func SendNextPreviousSpeaker(l logrus.FieldLogger, c Context, message string, speaker string, next StateProducer, previous StateProducer) State {
	return SendNextPrevious(l, c, message, next, previous, AddSendTalkConfigurator(npc.SetSpeaker(speaker)))
}

func SendNextPreviousExit(l logrus.FieldLogger, c Context, message string, next StateProducer, previous StateProducer, exit StateProducer) State {
	return SendNextPrevious(l, c, message, next, previous, SetSendTalkExit(exit))
}

func doNextPrevious(next StateProducer, previous StateProducer) ProcessStateFunc {
	return func(exit StateProducer) State {
		return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
			if mode == 255 && theType == 0 {
				return exit(l, c)
			}
			if mode == 0 && previous != nil {
				return previous(l, c)
			} else if mode == 1 && next != nil {
				return next(l, c)
			}
			return nil
		}
	}
}

func SendPrevious(l logrus.FieldLogger, c Context, message string, previous StateProducer, configurations ...SendTalkConfigurator) State {
	return sendTalk(l, c, message, configurations, npc.SendPrevious(l, c.CharacterId, c.NPCId), doPrevious(previous))
}

func doPrevious(previous StateProducer) ProcessStateFunc {
	return func(exit StateProducer) State {
		return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
			if mode == 255 && theType == 0 {
				return exit(l, c)
			}
			if mode == 0 && previous != nil {
				return previous(l, c)
			}
			return nil
		}
	}
}

func SendYesNo(l logrus.FieldLogger, c Context, message string, yes StateProducer, no StateProducer, configurations ...SendTalkConfigurator) State {
	return sendTalk(l, c, message, configurations, npc.SendYesNo(l, c.CharacterId, c.NPCId), doYesNo(yes, no))
}

func SendYesNoExit(l logrus.FieldLogger, c Context, message string, yes StateProducer, no StateProducer, exit StateProducer) State {
	return SendYesNo(l, c, message, yes, no, SetSendTalkExit(exit))
}

func doYesNo(yes StateProducer, no StateProducer) ProcessStateFunc {
	return func(exit StateProducer) State {
		return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
			if mode == 255 && theType == 0 {
				return exit(l, c)
			}
			if mode == 0 && no != nil {
				return no(l, c)
			} else if mode == 1 && yes != nil {
				return yes(l, c)
			}
			return nil
		}
	}
}

func SendOk(l logrus.FieldLogger, c Context, message string, configurations ...SendTalkConfigurator) State {
	return sendTalk(l, c, message, configurations, npc.SendOk(l, c.CharacterId, c.NPCId), func(exit StateProducer) State { return exit(l, c) })
}

func SendOkTrigger(l logrus.FieldLogger, c Context, message string, next StateProducer) State {
	return SendOk(l, c, message, SetSendTalkExit(next))
}

func SendGetNumber(l logrus.FieldLogger, c Context, message string, s ProcessNumber, defaultValue int32, minimumValue int32, maximumValue int32) State {
	err := npc.SendGetNumber(l, c.CharacterId, c.NPCId)(message, defaultValue, minimumValue, maximumValue)
	if err != nil {
		l.WithError(err).Errorf("Sending get number for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doGetNumberExit(Exit(), s)
}

func SendGetNumberExit(l logrus.FieldLogger, c Context, message string, s ProcessNumber, e StateProducer, defaultValue int32, minimumValue int32, maximumValue int32) State {
	err := npc.SendGetNumber(l, c.CharacterId, c.NPCId)(message, defaultValue, minimumValue, maximumValue)
	if err != nil {
		l.WithError(err).Errorf("Sending get number for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doGetNumberExit(e, s)
}

func doGetNumberExit(e StateProducer, s ProcessNumber) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 0 && theType == 3 {
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
	err := npc.SendGetText(l, c.CharacterId, c.NPCId)(message)
	if err != nil {
		l.WithError(err).Errorf("Sending get number for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doGetTextExit(e, s)
}

func doGetTextExit(e StateProducer, s ProcessText) State {
	return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
		if mode == 0 && theType == 3 {
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
	err := npc.SendStyle(l, c.CharacterId, c.NPCId)(message, options)
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

func SendAcceptDecline(l logrus.FieldLogger, c Context, message string, accept StateProducer, decline StateProducer, configurations ...SendTalkConfigurator) State {
	return sendTalk(l, c, message, configurations, npc.SendAcceptDecline(l, c.CharacterId, c.NPCId), doAcceptDecline(accept, decline))
}

func SendAcceptDeclineExit(l logrus.FieldLogger, c Context, message string, accept StateProducer, decline StateProducer, exit StateProducer) State {
	return SendAcceptDecline(l, c, message, accept, decline, SetSendTalkExit(exit))
}

func doAcceptDecline(accept StateProducer, decline StateProducer) ProcessStateFunc {
	return func(exit StateProducer) State {
		return func(l logrus.FieldLogger, c Context, mode byte, theType byte, selection int32) State {
			if mode == 255 && theType == 0 {
				return exit(l, c)
			}
			if mode == 0 && decline != nil {
				return decline(l, c)
			} else if mode == 1 && accept != nil {
				return accept(l, c)
			}
			return nil
		}
	}
}

func SendDimensionalMirror(l logrus.FieldLogger, c Context, message string, selection ProcessSelection) State {
	err := npc.SendDimensionalMirror(l, c.CharacterId, c.NPCId)(message)
	if err != nil {
		l.WithError(err).Errorf("Sending next message for npc %d to character %d.", c.NPCId, c.CharacterId)
	}
	return doDimensionalMirror(Exit(), selection)
}

func doDimensionalMirror(e StateProducer, s ProcessSelection) State {
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

func Warp(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := npc.WarpRandom(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return Exit()(l, c)
	}
}

func WarpById(mapId uint32, portalId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, portalId)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return Exit()(l, c)
	}
}

func WarpByName(mapId uint32, portalName string) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := npc.WarpByName(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, portalName)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return Exit()(l, c)
	}
}
