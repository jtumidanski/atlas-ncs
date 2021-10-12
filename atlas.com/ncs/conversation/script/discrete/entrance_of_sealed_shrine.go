package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// EntranceOfSealedShrine is located in Mu Lung - Practice Field : Advanced Level (250020300)
type EntranceOfSealedShrine struct {
}

func (r EntranceOfSealedShrine) NPCId() uint32 {
	return npc.EntranceOfSealedShrine
}

func (r EntranceOfSealedShrine) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The entrance of the Sealed Shrine... ").
		BlueText().AddText("Password").
		BlackText().AddText("!")
	return script.SendGetText(l, span, c, m.String(), r.ProcessPassword)
}

func (r EntranceOfSealedShrine) ProcessPassword(text string) script.StateProducer {
	if text != "Actions speak louder than words" {
		return r.Wrong
	}
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		if _map.CharacterCount(l)(c.WorldId, c.ChannelId, _map.SealedTemple) > 0 {
			return r.AlreadyAttending(l, span, c)
		}
		if quest.IsStarted(l)(c.CharacterId, 21747) && quest.ProgressInt(l)(c.CharacterId, 21747, 9300351) == 0 {
			return script.WarpById(_map.SealedTemple, 0)(l, span, c)
		} else {
			return r.CorrectButMissingPrerequisites(l, span, c)
		}
	}
}

func (r EntranceOfSealedShrine) CorrectButMissingPrerequisites(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.SendNotice(l)(c.CharacterId, "PINK_TEXT", "Although you said the right answer, some mysterious forces are blocking the way in.")
	return script.Exit()(l, span, c)
}

func (r EntranceOfSealedShrine) AlreadyAttending(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Someone is already attending the Sealed Shrine.")
	return script.SendOk(l, span, c, m.String())
}

func (r EntranceOfSealedShrine) Wrong(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().RedText().AddText("Wrong!")
	return script.SendOk(l, span, c, m.String())
}
