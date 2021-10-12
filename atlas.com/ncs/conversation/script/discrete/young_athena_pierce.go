package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// YoungAthenaPierce is located in Black Road - Ready to Leave (914000100)
type YoungAthenaPierce struct {
}

func (r YoungAthenaPierce) NPCId() uint32 {
	return npc.YoungAthenaPierce
}

func (r YoungAthenaPierce) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Aran, you're awake! How are you feeling? Hm? You want to know what's been going on?")
	return script.SendNext(l, span, c, m.String(), r.AlmostDone)
}

func (r YoungAthenaPierce) AlmostDone(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("We're almost done preparing for the escape. You don't have to worry. Everyone I could possibly find has boarded the ark, and Shinsoo has agreed to guide the way. We'll head to Victoria Island as soon as we finish the remaining preparations.")
	return script.SendNext(l, span, c, m.String(), r.OtherHeroes)
}

func (r YoungAthenaPierce) OtherHeroes(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The other heroes? They've left to fight the Black Magician. They're buying us time to escape. What? You want to fight with them? No! You can't! You're hurt. You must leave with us!")
	return script.SendNext(l, span, c, m.String(), r.ShowTrio)
}

func (r YoungAthenaPierce) ShowTrio(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.ShowIntro(l)(c.CharacterId, "Effect/Direction1.img/aranTutorial/Trio")
	return script.Exit()(l, span, c)
}
