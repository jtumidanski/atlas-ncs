package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Luke is located in Warning Street - Henesys Dungeon Entrance (106010100)
type Luke struct {
}

func (r Luke) NPCId() uint32 {
	return npc.Luke
}

func (r Luke) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.IsStarted(l)(c.CharacterId, 28177) && !character.HasItem(l, span)(c.CharacterId, item.TrainingInstructorsBadge) {
		if character.CanHold(l)(c.CharacterId, item.TrainingInstructorsBadge) {
			return r.GiveItem(l, span, c)
		}
		return r.MakeRoom(l, span, c)
	}
	return r.Snooze(l, span, c)
}

func (r Luke) Snooze(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Zzzzzz...")
	return script.SendOk(l, span, c, m.String())
}

func (r Luke) MakeRoom(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hey, make a slot available before talking to me.")
	return script.SendOk(l, span, c, m.String())
}

func (r Luke) GiveItem(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.TrainingInstructorsBadge, 1)
	return r.Success(l, span, c)
}

func (r Luke) Success(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Huh, are you looking for me? Chief Stan sent you here, right? But hey, I am not the suspect you seek. If I have some proof? Here, take this and return it to ").
		BlueText().ShowNPC(npc.ChiefStan).
		BlackText().AddText(".")
	return script.SendOk(l, span, c, m.String())
}
