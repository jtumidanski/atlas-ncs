package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// AssistantCheng is located in Ludibrium - Toy Factory <Process 1> Zone 1 (220020000) and Hidden Street - Toy Factory<Sector 4> (922000000)
type AssistantCheng struct {
}

func (r AssistantCheng) NPCId() uint32 {
	return npc.AssistantCheng
}

func (r AssistantCheng) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if c.MapId == _map.ToyFactorySector4 {
		return r.ConfirmQuit(l, span, c)
	}
	if quest.IsStarted(l)(c.CharacterId, 3239) {
		return r.ConfirmEnter(l, span, c)
	}
	return r.AccessRestricted(l, span, c)
}

func (r AssistantCheng) ConfirmQuit(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Do you wish to quit this stage?")
	return script.SendYesNo(l, span, c, m.String(), r.ProcessExit, r.CallMe)
}

func (r AssistantCheng) CallMe(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Ok. Call me if you urge to exit, then.")
	return script.SendOk(l, span, c, m.String())
}

func (r AssistantCheng) ProcessExit(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !(quest.IsStarted(l)(c.CharacterId, 3239) && character.HasItems(l, span)(c.CharacterId, item.MachineParts, 10)) {
		character.RemoveAll(l)(c.CharacterId, item.MachineParts)
	}
	return script.WarpById(_map.SecretPassage, 0)(l, span, c)
}

func (r AssistantCheng) ConfirmEnter(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Do you want to enter ").
		BlueText().ShowMap(_map.ToyFactorySector4).
		BlackText().AddText("?")
	return script.SendYesNo(l, span, c, m.String(), r.ValidateEnter, script.Exit())
}

func (r AssistantCheng) AccessRestricted(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Access to ").
		BlueText().ShowMap(_map.ToyFactorySector4).
		BlackText().AddText(" is restricted to the public.")
	return script.SendOk(l, span, c, m.String())
}

func (r AssistantCheng) ValidateEnter(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if _map.CharacterCount(l)(c.WorldId, c.ChannelId, _map.ToyFactorySector4) != 0 {
		return r.AlreadyAttempting(l, span, c)
	}
	return r.WarpEnter(l, span, c)
}

func (r AssistantCheng) AlreadyAttempting(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().AddText("Someone else is already attempting the parts. Wait for them to finish before you enter.")
	return script.SendOk(l, span, c, m.String())
}

func (r AssistantCheng) WarpEnter(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if !(quest.IsStarted(l)(c.CharacterId, 3239) && character.HasItems(l, span)(c.CharacterId, item.MachineParts, 10)) {
		character.RemoveAll(l)(c.CharacterId, item.MachineParts)
	}
	return script.WarpById(_map.ToyFactorySector4, 0)(l, span, c)
}
