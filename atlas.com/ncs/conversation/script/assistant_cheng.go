package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// AssistantCheng is located in Ludibrium - Toy Factory <Process 1> Zone 1 (220020000) and Hidden Street - Toy Factory<Sector 4> (922000000)
type AssistantCheng struct {
}

func (r AssistantCheng) NPCId() uint32 {
	return npc.AssistantCheng
}

func (r AssistantCheng) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.ToyFactorySector4 {
		return r.ConfirmQuit(l, c)
	}
	if character.QuestStarted(l)(c.CharacterId, 3239) {
		return r.ConfirmEnter(l, c)
	}
	return r.AccessRestricted(l, c)
}

func (r AssistantCheng) ConfirmQuit(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Do you wish to quit this stage?")
	return SendYesNo(l, c, m.String(), r.ProcessExit, r.CallMe)
}

func (r AssistantCheng) CallMe(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Ok. Call me if you urge to exit, then.")
	return SendOk(l, c, m.String())
}

func (r AssistantCheng) ProcessExit(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.SecretPassage, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.SecretPassage, c.NPCId)
	}
	if !(character.QuestStarted(l)(c.CharacterId, 3239) && character.HasItems(l)(c.CharacterId, item.MachineParts, 10)) {
		character.RemoveAll(l)(c.CharacterId, item.MachineParts)
	}
	return Exit()(l, c)
}

func (r AssistantCheng) ConfirmEnter(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Do you want to enter ").
		BlueText().ShowMap(_map.ToyFactorySector4).
		BlackText().AddText("?")
	return SendYesNo(l, c, m.String(), r.ValidateEnter, Exit())
}

func (r AssistantCheng) AccessRestricted(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Access to ").
		BlueText().ShowMap(_map.ToyFactorySector4).
		BlackText().AddText(" is restricted to the public.")
	return SendOk(l, c, m.String())
}

func (r AssistantCheng) ValidateEnter(l logrus.FieldLogger, c Context) State {
	if _map.CharacterCount(l)(c.WorldId, c.ChannelId, _map.ToyFactorySector4) != 0 {
		return r.AlreadyAttempting(l, c)
	}
	return r.WarpEnter(l, c)
}

func (r AssistantCheng) AlreadyAttempting(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Someone else is already attempting the parts. Wait for them to finish before you enter.")
	return SendOk(l, c, m.String())
}

func (r AssistantCheng) WarpEnter(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.ToyFactorySector4, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.ToyFactorySector4, c.NPCId)
	}
	if !(character.QuestStarted(l)(c.CharacterId, 3239) && character.HasItems(l)(c.CharacterId, item.MachineParts, 10)) {
		character.RemoveAll(l)(c.CharacterId, item.MachineParts)
	}
	return Exit()(l, c)
}
