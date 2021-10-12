package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/job"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// MagicianStatue is located in Victoria Road - Lith Harbor (104000000)
type MagicianStatue struct {
}

func (r MagicianStatue) NPCId() uint32 {
	return npc.MagicianStatue
}

func (r MagicianStatue) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !character.IsJob(l, span)(c.CharacterId, job.Beginner) {
		m := message.NewBuilder().AddText("You're much stronger now. Keep training!")
		return script.SendOk(l, span, c, m.String())
	}

	if !character.MeetsCriteria(l, span)(c.CharacterId, character.IsLevelCriteria(10), character.HasIntelligenceCriteria(20)) {
		m := message.NewBuilder().
			AddText("If you want to be a ").
			BlueText().AddText("Magician").
			BlackText().AddText(", train yourself further until you reach ").
			BlueText().AddText("level 10 with 20 INT").
			BlackText().AddText(".")
		return script.SendOk(l, span, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Hey ").
		ShowCharacterName().
		AddText(", I can send you to ").
		BlueText().ShowMap(_map.MagicLibrary).
		BlackText().AddText(" if you want to be a ").
		BlueText().AddText("Magician").
		BlackText().AddText(". Do you want to go now?")
	return script.SendYesNoExit(l, span, c, m.String(), script.WarpById(_map.MagicLibrary, 0), r.ComeBack, r.ComeBack)
}

func (r MagicianStatue) ComeBack(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Come back to me if you decided to be a ").
		BlueText().AddText("Magician").
		BlackText().AddText(".")
	return script.SendOk(l, span, c, m.String())
}
