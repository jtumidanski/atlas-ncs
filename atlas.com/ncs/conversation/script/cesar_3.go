package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

type Cesar3 struct {
}

func (r Cesar3) NPCId() uint32 {
	return npc.Cesar3
}

func (r Cesar3) Initial(l logrus.FieldLogger, c Context) State {
	if character.MeetsCriteria(l)(c.CharacterId, character.IsLevelBetweenCriteria(20, 30)) {
		return r.Sorry(l, c)
	}
	return r.HugeFestival(l, c)
}

func (r Cesar3) HugeFestival(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I have prepared a huge festival here at Ariant for the great fighters of MapleStory. It's called ").
		BlueText().AddText("The Ariant Coliseum Challenge").
		BlackText().AddText(".")
	return SendNext(l, c, m.String(), r.Explained)
}

func (r Cesar3) Sorry(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You're not between level 20 and 30. Sorry, you may not participate.")
	return SendOk(l, c, m.String())
}

func (r Cesar3) Explained(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The Ariant Coliseum Challenge is a competition that matches the skills of monster combat against others. In this competition, your object isn't to hunt the monster;  rather, you need to ").
		BlueText().AddText("eliminate a set amount of HP from the monster, followed by absorbing it with a jewel").
		BlackText().AddText(". ").
		BlueText().AddText("The fighter that ends up with the most jewels will win the competition.")
	return SendNextPrevious(l, c, m.String(), r.HugeFestival, r.AreYouInterested)
}

func (r Cesar3) AreYouInterested(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("If you are a strong and brave warrior from ").
		BlueText().AddText("Perion").
		BlackText().AddText(", training under Dances With Balrogs, then are you interested in participating in The Ariant Coliseum Challenge?!").NewLine().
		OpenItem(0).BlueText().AddText("I'd love to participate in this great competition.").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r Cesar3) Selection(selection int32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().AddText("Okay, now I'll send you to the battle arena. I'd like to see you emerge victorious!")
		return SendNext(l, c, m.String(), r.Warp)
	}
}

func (r Cesar3) Warp(l logrus.FieldLogger, c Context) State {
	character.SaveLocation(l)(c.CharacterId, "MIRROR")
	return WarpById(_map.BattleArenaLobby, 3)(l, c)
}
