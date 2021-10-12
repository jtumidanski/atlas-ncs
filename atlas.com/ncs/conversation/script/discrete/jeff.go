package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Jeff is located in El Nath - Ice Valley II (211040200)
type Jeff struct {
}

func (r Jeff) NPCId() uint32 {
	return npc.Jeff
}

func (r Jeff) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.HasItem(l, span)(c.CharacterId, item.OrihalconHammer) {
		return script.WarpById(_map.IceValley, 1)(l, span, c)
	}
	return r.FurtherAndDeeper(l, span, c)
}

func (r Jeff) FurtherAndDeeper(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Hey, you look like you want to go farther and deeper past this place. Over there, though, you'll find yourself surrounded by aggressive, dangerous monsters, so even if you feel that you're ready to go, please be careful. Long ago, a few brave men from our town went in wanting to eliminate anyone threatening the town, but never came back out...")
	return script.SendNext(l, span, c, m.String(), r.VerifyLevel)
}

func (r Jeff) VerifyLevel(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if character.IsLevel(l, span)(c.CharacterId, 50) {
		return r.DoYouWantTo(l, span, c)
	} else {
		return r.ToWeak(l, span, c)
	}
}

func (r Jeff) ToWeak(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("If you are thinking of going in, I suggest you change your mind. But if you really want to go in... I'm only letting in the ones that are strong enough to stay alive in there. I do not wish to see anyone else die. Let's see... Hmmm... You haven't reached Level 50 yet. I can't let you in, then, so forget it.")
	return script.SendPrevious(l, span, c, m.String(), r.FurtherAndDeeper)
}

func (r Jeff) DoYouWantTo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("If you are thinking of going in, I suggest you change your mind. But if you really want to go in... I'm only letting in the ones that are strong enough to stay alive in there. I do not wish to see anyone else die. Let's see... Hmmm...! You look pretty strong. All right, do you want to go in?")
	return script.SendYesNo(l, span, c, m.String(), script.WarpById(_map.SharpCliffI, 5), r.ChangeYourMind)
}

func (r Jeff) ChangeYourMind(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Even if your level's high it's hard to actually go in there, but if you ever change your mind, please find me. After all, my job is to protect this place.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}
