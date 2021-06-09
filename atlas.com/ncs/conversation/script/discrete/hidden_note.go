package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// HiddenNote is located in Victoria Road - Pet-Walking Road (100000202)
type HiddenNote struct {
}

func (r HiddenNote) NPCId() uint32 {
	return npc.HiddenNote
}

func (r HiddenNote) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.QuestStarted(l)(c.CharacterId, 4646) {
		return r.CouldNotFindAnything(l, c)
	}

	if character.HasItem(l)(c.CharacterId, item.HiddenNote) {
		return r.Eww(l, c)
	}

	return r.SeeSomething(l, c)
}

func (r HiddenNote) Eww(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().BlueText().
		AddText("(What's this... eww... a pet's poop was in there!)")
	return script.SendOk(l, c, m.String())
}

func (r HiddenNote) SeeSomething(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().BlueText().
		AddText("(I can see something covered in grass. Should I pull it out?)")
	return script.SendYesNo(l, c, m.String(), r.AwardNote, r.Covered)
}

func (r HiddenNote) CouldNotFindAnything(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().BlueText().
		AddText("(I couldn't find anything.)")
	return script.SendOk(l, c, m.String())
}

func (r HiddenNote) AwardNote(l logrus.FieldLogger, c script.Context) script.State {
	character.GainItem(l)(c.CharacterId, item.HiddenNote, 1)
	m := message.NewBuilder().
		AddText("I found the item that Pet Trainer Bartos hid... this note.")
	return script.SendOk(l, c, m.String())
}

func (r HiddenNote) Covered(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().BlueText().
		AddText("(I didn't touch this hidden item covered in grass)")
	return script.SendOk(l, c, m.String())
}
