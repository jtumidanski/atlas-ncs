package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Francis2 is located in Hidden Street - Puppeteer's Cave (910510200)
type Francis2 struct {
}

func (r Francis2) NPCId() uint32 {
	return npc.Francis2
}

func (r Francis2) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I'm Francis, the Puppeteer of the Black Wings. How dare you disturb my puppets... It really upsets me, but I'll let it slide this time. If I catch you doing it again though, I swear in the name of the Black Magician, I will make you pay for it.")
	return script.SendNextSpeaker(l, c, m.String(), npc.SpeakerUnknown2, r.ReportToTru)
}

func (r Francis2) ReportToTru(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		BlueText().AddText("(The Black Wings? Huh? Who are they? And how is all this related to the Black Magician? Hm, maybe you should report this info to Tru.)")
	return script.SendNextPreviousSpeaker(l, c, m.String(), npc.SpeakerCharacterRight, r.Process, r.Initial)
}

func (r Francis2) Process(l logrus.FieldLogger, c script.Context) script.State {
	character.CompleteQuest(l)(c.CharacterId, 21719)
	return script.WarpById(_map.HuntingGroundInTheDeepForestII, 0)(l, c)
}
