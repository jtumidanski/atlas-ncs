package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/character/location"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type Cesar3 struct {
}

func (r Cesar3) NPCId() uint32 {
	return npc.Cesar3
}

func (r Cesar3) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.MeetsCriteria(l, span)(c.CharacterId, character.IsLevelBetweenCriteria(20, 30)) {
		return r.Sorry(l, span, c)
	}
	return r.HugeFestival(l, span, c)
}

func (r Cesar3) HugeFestival(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I have prepared a huge festival here at Ariant for the great fighters of MapleStory. It's called ").
		BlueText().AddText("The Ariant Coliseum Challenge").
		BlackText().AddText(".")
	return script.SendNext(l, span, c, m.String(), r.Explained)
}

func (r Cesar3) Sorry(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("You're not between level 20 and 30. Sorry, you may not participate.")
	return script.SendOk(l, span, c, m.String())
}

func (r Cesar3) Explained(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("The Ariant Coliseum Challenge is a competition that matches the skills of monster combat against others. In this competition, your object isn't to hunt the monster;  rather, you need to ").
		BlueText().AddText("eliminate a set amount of HP from the monster, followed by absorbing it with a jewel").
		BlackText().AddText(". ").
		BlueText().AddText("The fighter that ends up with the most jewels will win the competition.")
	return script.SendNextPrevious(l, span, c, m.String(), r.HugeFestival, r.AreYouInterested)
}

func (r Cesar3) AreYouInterested(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("If you are a strong and brave warrior from ").
		BlueText().AddText("Perion").
		BlackText().AddText(", training under Dances With Balrogs, then are you interested in participating in The Ariant Coliseum Challenge?!").NewLine().
		OpenItem(0).BlueText().AddText("I'd love to participate in this great competition.").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r Cesar3) Selection(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
		m := message.NewBuilder().AddText("Okay, now I'll send you to the battle arena. I'd like to see you emerge victorious!")
		return script.SendNext(l, span, c, m.String(), r.Warp)
	}
}

func (r Cesar3) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	location.SaveLocation(l, span)(c.CharacterId, "MIRROR")
	return script.WarpById(_map.BattleArenaLobby, 3)(l, span, c)
}
