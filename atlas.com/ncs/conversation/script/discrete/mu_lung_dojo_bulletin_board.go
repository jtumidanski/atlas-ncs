package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type MuLungDojoBulletinBoard struct {
}

func (r MuLungDojoBulletinBoard) NPCId() uint32 {
	return npc.MuLungDojoBulletinBoard
}

func (r MuLungDojoBulletinBoard) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		BoldText().AddText("< Notice >").NewLine().
		NormalText().AddText("If there is anyone who has the courage to challenge the Mu Lung Dojo, come to the Mu Lung Dojo.  - Mu Gong -").NewLine().NewLine().NewLine().
		OpenItem(0).BlueText().AddText("Challenge the Mu Lung Dojo.").CloseItem().NewLine().
		OpenItem(1).AddText("Read the notice in more detail.").CloseItem()
	return script.SendListSelection(l, span, c, m.String(), r.Selection)
}

func (r MuLungDojoBulletinBoard) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Challenge
	case 1:
		return r.Detail
	}
	return nil
}

func (r MuLungDojoBulletinBoard) Challenge(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("(Once I had placed my hands on the bulletin board, a mysterious energy began to envelop me.)").NewLine().NewLine().
		AddText("Would you like to go to Mu Lung Dojo?")
	return script.SendYesNo(l, span, c, m.String(), r.Warp, r.No)
}

func (r MuLungDojoBulletinBoard) No(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().BlueText().AddText("(As I took my hand off the bulletin board, the mysterious energy that was covering my disappeared as well.)")
	return script.SendOk(l, span, c, m.String())
}

func (r MuLungDojoBulletinBoard) Warp(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	character.SaveLocation(l)(c.CharacterId, "MIRROR")
	return script.WarpById(_map.MuLungDojoEntrance, 4)(l, span, c)
}

func (r MuLungDojoBulletinBoard) Detail(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		BoldText().AddText("< Notice : Take the challenge! >").NewLine().
		NormalText().AddText("My name is Mu Gong, the owner of the My Lung Dojo. Since long ago, I have been training in Mu Lung to the point where my skills have now reached the pinnacle. Starting today, I will take on any and all applicants for Mu Lung Dojo. The rights to the Mu Lung Dojo will be given only to the strongest person.\\r\\nIf there is anyone who wishes to learn from me, come take the challenge any time! If there is anyone who wishes to challenge me, you're welcome as well. I will make you fully aware of your own weakness.")
	return script.SendNext(l, span, c, m.String(), r.CallYourFriends)
}

func (r MuLungDojoBulletinBoard) CallYourFriends(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("PS:You can challenge me on your own. But if you don't have that kind of courage, go ahead and call all your friends.")
	return script.SendPrevious(l, span, c, m.String(), r.Detail)
}
