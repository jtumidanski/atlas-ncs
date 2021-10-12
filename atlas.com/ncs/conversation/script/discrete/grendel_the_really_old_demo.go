package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// GrendelTheReallyOldDemo is located in Maple Road : Split Road of Destiny (1020000)
type GrendelTheReallyOldDemo struct {
}

func (r GrendelTheReallyOldDemo) NPCId() uint32 {
	return npc.GrendelTheReallyOldDemo
}

func (r GrendelTheReallyOldDemo) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return r.MagicianIntroduction(l, span, c)
}

func (r GrendelTheReallyOldDemo) MagicianIntroduction(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Magicians are armed with flashy element-based spells and secondary magic that aids party as a whole. After the 2nd job adv., the elemental-based magic will provide ample amount of damage to enemies of opposite element.")
	return script.SendNext(l, span, c, m.String(), r.Demo)
}

func (r GrendelTheReallyOldDemo) Demo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Would you like to experience what it's like to be a Magician?")
	return script.SendYesNo(l, span, c, m.String(), r.DoDemo, r.SeeMeAgain)
}

func (r GrendelTheReallyOldDemo) DoDemo(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	npc.LockUI(l)(c.CharacterId)
	return script.WarpById(_map.MagicianDemo, 0)(l, span, c)
}

func (r GrendelTheReallyOldDemo) SeeMeAgain(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("If you wish to experience what it's like to be a Magician, come see me again.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}