package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// FirstPipeHandle is located in Magatia - Home of the Missing Alchemist (261000001)
type FirstPipeHandle struct {
}

func (r FirstPipeHandle) NPCId() uint32 {
	return npc.FirstPipeHandle
}

func (r FirstPipeHandle) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 3339) {
		return r.Progress(l, span, c)
	}
	if quest.IsCompleted(l)(c.CharacterId, 3339) {
		return r.WarpInMap(l, span, c)
	}
	return script.Exit()(l, span, c)
}

func (r FirstPipeHandle) GetPassword(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The pipe reacts as the water starts flowing. A secret compartment with a keypad shows up. ").
		BlueText().AddText("Password").
		BlackText().AddText("!")
	return script.SendGetText(l, span, c, m.String(), r.Validate)
}

func (r FirstPipeHandle) WarpInMap(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.HomeOfTheMissingAlchemist, 1)(l, span, c)
}

func (r FirstPipeHandle) Progress(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	progress := quest.ProgressInt(l)(c.CharacterId, 23339, 1)
	if progress == 3 {
		return r.GetPassword(l, span, c)
	} else if progress == 0 {
		quest.SetProgress(l)(c.CharacterId, 23339, 1, 1)
		return script.Exit()(l, span, c)
	} else if progress < 3 {
		quest.SetProgress(l)(c.CharacterId, 23339, 1, 0)
		return script.Exit()(l, span, c)
	} else {
		return r.WarpInMap(l, span, c)
	}
}

func (r FirstPipeHandle) Validate(text string) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if text != "my love Phyllia" {
			m := message.NewBuilder().RedText().AddText("Wrong!")
			return script.SendOk(l, span, c, m.String())
		}
		quest.SetProgress(l)(c.CharacterId, 23339, 1, 4)
		return r.WarpInMap(l, span, c)
	}
}
