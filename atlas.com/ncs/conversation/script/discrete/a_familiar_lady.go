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

// AFamiliarLady is located in Hidden Street - Gloomy Forest (922220000)
type AFamiliarLady struct {
}

func (r AFamiliarLady) NPCId() uint32 {
	return npc.AFamiliarLady
}

func (r AFamiliarLady) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if quest.ProgressInt(l)(c.CharacterId, 23647, 1) != 0 {
		return script.Exit()(l, span, c)
	}

	if !character.HasItem(l, span)(c.CharacterId, item.OldFoxsTail) {
		return r.LostInTheWoods(l, span, c)
	}

	return r.Confirm(l, span, c)
}

func (r AFamiliarLady) LostInTheWoods(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Umm... Hey... Would you help me find a ").
		BlueText().AddText("soft and shiny silver fur").
		BlackText().AddText(" that I lost on the woods? I need it, I need it, I need it sooooo much!")
	return script.SendOk(l, span, c, m.String())
}

func (r AFamiliarLady) Confirm(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hey... Umm... Would you help me find a ").
		BlueText().AddText("soft and shiny silver fur").
		BlackText().AddText(" that I lost on the woods? I need it, I need it, I need it sooooo much! ... Oh you found it!!! Will you give it to me?")
	return script.SendYesNo(l, span, c, m.String(), r.Process, script.Exit())
}

func (r AFamiliarLady) Process(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.GainItem(l, span)(c.CharacterId, item.OldFoxsTail, -1)
	character.GainFame(l)(c.CharacterId, -5)
	quest.SetProgress(l)(c.CharacterId, 23647, 1, 1)
	return r.Success(l, span, c)
}

func (r AFamiliarLady) Success(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Teehehee~ That's your reward for taking it from me, serves you well.")
	return script.SendOk(l, span, c, m.String())
}
