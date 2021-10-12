package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// DarkLordDemo is located in Maple Road : Split Road of Destiny (1020000)
type DarkLordDemo struct {
}

func (r DarkLordDemo) NPCId() uint32 {
	return npc.DarkLordDemo
}

func (r DarkLordDemo) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.ThiefIntroduction(l, span, c)
}

func (r DarkLordDemo) ThiefIntroduction(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Thieves are a perfect blend of luck, dexterity, and power that are adept at the surprise attacks against helpless enemies. A high level of avoidability and speed allows Thieves to attack enemies from various angles.")
	return script.SendNext(l, span, c, m.String(), r.Demo)
}

func (r DarkLordDemo) Demo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Thief?")
	return script.SendYesNo(l, span, c, m.String(), r.DoDemo, r.SeeMeAgain)
}

func (r DarkLordDemo) DoDemo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	npc.LockUI(l)(c.CharacterId)
	return script.WarpById(_map.ThiefDemo, 0)(l, span, c)
}

func (r DarkLordDemo) SeeMeAgain(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Thief, come see me again.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}
