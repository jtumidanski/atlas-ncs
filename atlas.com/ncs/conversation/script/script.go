package script

import (
	"atlas-ncs/npc"
	"github.com/opentracing/opentracing-go"
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

	Initial(l logrus.FieldLogger, span opentracing.Span, c Context) State
}

type StateProducer func(l logrus.FieldLogger, span opentracing.Span, c Context) State

type ProcessNumber func(selection int32) StateProducer

type ProcessText func(text string) StateProducer

type State func(l logrus.FieldLogger, span opentracing.Span, c Context, mode byte, theType byte, selection int32) State

type ProcessSelection func(selection int32) StateProducer

func Exit() StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c Context) State {
		npc.Dispose(l, span)(c.CharacterId)
		return nil
	}
}

func SendListSelection(l logrus.FieldLogger, span opentracing.Span, c Context, message string, s ProcessSelection) State {
	npc.SendSimple(l, span)(c.CharacterId, c.NPCId)(message)
	return doListSelectionExit(Exit(), s)
}

func SendListSelectionExit(l logrus.FieldLogger, span opentracing.Span, c Context, message string, s ProcessSelection, exit StateProducer) State {
	npc.SendSimple(l, span)(c.CharacterId, c.NPCId)(message)
	return doListSelectionExit(exit, s)
}

func doListSelectionExit(e StateProducer, s ProcessSelection) State {
	return func(l logrus.FieldLogger, span opentracing.Span, c Context, mode byte, theType byte, selection int32) State {
		if mode == 0 && theType == 4 {
			return e(l, span, c)
		}

		f := s(selection)
		if f == nil {
			l.Errorf("unhandled selection %d for npc %d.", selection, c.NPCId)
			return nil
		}
		return f(l, span, c)
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

	talkFunc(message, baseConfig.Configurators()...)
	return do(baseConfig.Exit())
}

func SendNext(l logrus.FieldLogger, span opentracing.Span, c Context, message string, next StateProducer, configurations ...SendTalkConfigurator) State {
	return sendTalk(l, c, message, configurations, npc.SendNext(l, span)(c.CharacterId, c.NPCId), doNext(next))
}

func SendNextSpeaker(l logrus.FieldLogger, span opentracing.Span, c Context, message string, speaker string, next StateProducer) State {
	return SendNext(l, span, c, message, next, AddSendTalkConfigurator(npc.SetSpeaker(speaker)))
}

func SendNextExit(l logrus.FieldLogger, span opentracing.Span, c Context, message string, next StateProducer, exit StateProducer) State {
	return SendNext(l, span, c, message, next, SetSendTalkExit(exit))
}

func doNext(next StateProducer) ProcessStateFunc {
	return func(exit StateProducer) State {
		return func(l logrus.FieldLogger, span opentracing.Span, c Context, mode byte, theType byte, selection int32) State {
			if mode == 255 && theType == 0 {
				return exit(l, span, c)
			}
			return next(l, span, c)
		}
	}
}

func SendNextPrevious(l logrus.FieldLogger, span opentracing.Span, c Context, message string, next StateProducer, previous StateProducer, configurations ...SendTalkConfigurator) State {
	return sendTalk(l, c, message, configurations, npc.SendNextPrevious(l, span)(c.CharacterId, c.NPCId), doNextPrevious(next, previous))
}

func SendNextPreviousSpeaker(l logrus.FieldLogger, span opentracing.Span, c Context, message string, speaker string, next StateProducer, previous StateProducer) State {
	return SendNextPrevious(l, span, c, message, next, previous, AddSendTalkConfigurator(npc.SetSpeaker(speaker)))
}

func SendNextPreviousExit(l logrus.FieldLogger, span opentracing.Span, c Context, message string, next StateProducer, previous StateProducer, exit StateProducer) State {
	return SendNextPrevious(l, span, c, message, next, previous, SetSendTalkExit(exit))
}

func doNextPrevious(next StateProducer, previous StateProducer) ProcessStateFunc {
	return func(exit StateProducer) State {
		return func(l logrus.FieldLogger, span opentracing.Span, c Context, mode byte, theType byte, selection int32) State {
			if mode == 255 && theType == 0 {
				return exit(l, span, c)
			}
			if mode == 0 && previous != nil {
				return previous(l, span, c)
			} else if mode == 1 && next != nil {
				return next(l, span, c)
			}
			return nil
		}
	}
}

func SendPrevious(l logrus.FieldLogger, span opentracing.Span, c Context, message string, previous StateProducer, configurations ...SendTalkConfigurator) State {
	return sendTalk(l, c, message, configurations, npc.SendPrevious(l, span)(c.CharacterId, c.NPCId), doPrevious(previous))
}

func doPrevious(previous StateProducer) ProcessStateFunc {
	return func(exit StateProducer) State {
		return func(l logrus.FieldLogger, span opentracing.Span, c Context, mode byte, theType byte, selection int32) State {
			if mode == 255 && theType == 0 {
				return exit(l, span, c)
			}
			if mode == 0 && previous != nil {
				return previous(l, span, c)
			}
			return nil
		}
	}
}

func SendYesNo(l logrus.FieldLogger, span opentracing.Span, c Context, message string, yes StateProducer, no StateProducer, configurations ...SendTalkConfigurator) State {
	return sendTalk(l, c, message, configurations, npc.SendYesNo(l, span)(c.CharacterId, c.NPCId), doYesNo(yes, no))
}

func SendYesNoExit(l logrus.FieldLogger, span opentracing.Span, c Context, message string, yes StateProducer, no StateProducer, exit StateProducer) State {
	return SendYesNo(l, span, c, message, yes, no, SetSendTalkExit(exit))
}

func doYesNo(yes StateProducer, no StateProducer) ProcessStateFunc {
	return func(exit StateProducer) State {
		return func(l logrus.FieldLogger, span opentracing.Span, c Context, mode byte, theType byte, selection int32) State {
			if mode == 255 && theType == 0 {
				return exit(l, span, c)
			}
			if mode == 0 && no != nil {
				return no(l, span, c)
			} else if mode == 1 && yes != nil {
				return yes(l, span, c)
			}
			return nil
		}
	}
}

func SendOk(l logrus.FieldLogger, span opentracing.Span, c Context, message string, configurations ...SendTalkConfigurator) State {
	return sendTalk(l, c, message, configurations, npc.SendOk(l, span)(c.CharacterId, c.NPCId), func(exit StateProducer) State { return exit(l, span, c) })
}

func SendOkTrigger(l logrus.FieldLogger, span opentracing.Span, c Context, message string, next StateProducer) State {
	return SendOk(l, span, c, message, SetSendTalkExit(next))
}

func SendGetNumber(l logrus.FieldLogger, span opentracing.Span, c Context, message string, s ProcessNumber, defaultValue int32, minimumValue int32, maximumValue int32) State {
	npc.SendGetNumber(l, span)(c.CharacterId, c.NPCId)(message, defaultValue, minimumValue, maximumValue)
	return doGetNumberExit(Exit(), s)
}

func SendGetNumberExit(l logrus.FieldLogger, span opentracing.Span, c Context, message string, s ProcessNumber, e StateProducer, defaultValue int32, minimumValue int32, maximumValue int32) State {
	npc.SendGetNumber(l, span)(c.CharacterId, c.NPCId)(message, defaultValue, minimumValue, maximumValue)
	return doGetNumberExit(e, s)
}

func doGetNumberExit(e StateProducer, s ProcessNumber) State {
	return func(l logrus.FieldLogger, span opentracing.Span, c Context, mode byte, theType byte, selection int32) State {
		if mode == 0 && theType == 3 {
			return e(l, span, c)
		}

		f := s(selection)
		if f == nil {
			l.Errorf("unhandled selection %d for npc %d.", selection, c.NPCId)
			return nil
		}
		return f(l, span, c)
	}
}

func SendGetText(l logrus.FieldLogger, span opentracing.Span, c Context, message string, s ProcessText) State {
	return SendGetTextExit(l, span, c, message, s, Exit())
}

func SendGetTextExit(l logrus.FieldLogger, span opentracing.Span, c Context, message string, s ProcessText, e StateProducer) State {
	npc.SendGetText(l, span)(c.CharacterId, c.NPCId)(message)
	return doGetTextExit(e, s)
}

func doGetTextExit(e StateProducer, s ProcessText) State {
	return func(l logrus.FieldLogger, span opentracing.Span, c Context, mode byte, theType byte, selection int32) State {
		if mode == 0 && theType == 3 {
			return e(l, span, c)
		}

		//TODO get text somehow
		text := ""
		f := s(text)
		if f == nil {
			l.Errorf("unhandled selection %d for npc %d.", selection, c.NPCId)
			return nil
		}
		return f(l, span, c)
	}
}

func SendStyle(l logrus.FieldLogger, span opentracing.Span, c Context, message string, next ProcessSelection, options []uint32) State {
	npc.SendStyle(l, span)(c.CharacterId, c.NPCId)(message, options)
	return doSendStyleExit(Exit(), next)
}

func doSendStyleExit(e StateProducer, next ProcessSelection) State {
	return func(l logrus.FieldLogger, span opentracing.Span, c Context, mode byte, theType byte, selection int32) State {
		if mode == 0 && theType == 7 {
			return e(l, span, c)
		}
		return next(selection)(l, span, c)
	}
}

func SendAcceptDecline(l logrus.FieldLogger, span opentracing.Span, c Context, message string, accept StateProducer, decline StateProducer, configurations ...SendTalkConfigurator) State {
	return sendTalk(l, c, message, configurations, npc.SendAcceptDecline(l, span)(c.CharacterId, c.NPCId), doAcceptDecline(accept, decline))
}

func SendAcceptDeclineExit(l logrus.FieldLogger, span opentracing.Span, c Context, message string, accept StateProducer, decline StateProducer, exit StateProducer) State {
	return SendAcceptDecline(l, span, c, message, accept, decline, SetSendTalkExit(exit))
}

func doAcceptDecline(accept StateProducer, decline StateProducer) ProcessStateFunc {
	return func(exit StateProducer) State {
		return func(l logrus.FieldLogger, span opentracing.Span, c Context, mode byte, theType byte, selection int32) State {
			if mode == 255 && theType == 0 {
				return exit(l, span, c)
			}
			if mode == 0 && decline != nil {
				return decline(l, span, c)
			} else if mode == 1 && accept != nil {
				return accept(l, span, c)
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
	return func(l logrus.FieldLogger, span opentracing.Span, c Context, mode byte, theType byte, selection int32) State {
		if mode == 0 && theType == 4 {
			return e(l, span, c)
		}

		f := s(selection)
		if f == nil {
			l.Errorf("unhandled selection %d for npc %d.", selection, c.NPCId)
			return nil
		}
		return f(l, span, c)
	}
}

func Warp(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c Context) State {
		npc.WarpRandom(l, span)(c.WorldId, c.ChannelId, c.CharacterId, mapId)
		return Exit()(l, span, c)
	}
}

func WarpById(mapId uint32, portalId uint32) StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c Context) State {
		npc.WarpById(l, span)(c.WorldId, c.ChannelId, c.CharacterId, mapId, portalId)
		return Exit()(l, span, c)
	}
}

func WarpByName(mapId uint32, portalName string) StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c Context) State {
		npc.WarpByName(l, span)(c.WorldId, c.ChannelId, c.CharacterId, mapId, portalName)
		return Exit()(l, span, c)
	}
}
