package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/job"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// GrendelTheReallyOld is located in Victoria Road - Magic Library (101000003)
type GrendelTheReallyOld struct {
}

func (r GrendelTheReallyOld) NPCId() uint32 {
	return npc.GrendelTheReallyOld
}

func (r GrendelTheReallyOld) Initial(l logrus.FieldLogger, c Context) State {
	if character.IsJob(l)(c.CharacterId, job.Beginner) {
		return r.FirstJobInitial(l, c)
	}
	return nil
}

func (r GrendelTheReallyOld) FirstJobInitial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Want to be a ").
		RedText().AddText("magician").
		BlackText().AddText("? There are some standards to meet. because we can't just accept EVERYONE in... ").
		BlueText().AddText("Your level should be at least 8").
		BlackText().AddText(", with getting 20 INT as your top priority. Let's see.")
	return SendNext(l, c, m.String(), r.FirstJobRequirementCheck)
}

func (r GrendelTheReallyOld) FirstJobRequirementCheck(l logrus.FieldLogger, c Context) State {
	if character.IsLevel(l)(c.CharacterId, 8) && character.HasIntelligence(l)(c.CharacterId, 20) {
		m := message.NewBuilder().
			AddText("Oh...! You look like someone that can definitely be a part of us... all you need is a little sinister mind, and... yeah... so, what do you think? Wanna be the Magician?")
		return SendYesNo(l, c, m.String(), r.AwardFirstJob, r.FirstJobInitial)
	}

	m := message.NewBuilder().
		AddText("Train a bit more until you reach the base requirements and I can show you the way of the ").
		RedText().AddText("Magician").
		BlackText().AddText(".")
	return SendOk(l, c, m.String())
}

func (r GrendelTheReallyOld) AwardFirstJob(l logrus.FieldLogger, c Context) State {
	if !character.CanHold(l)(c.CharacterId, item.BeginnerMagicWand) {
		m := message.NewBuilder().
			AddText("Make some room in your inventory and talk back to me.")
		return SendNext(l, c, m.String(), Exit())
	}

	character.ChangeJob(l)(c.CharacterId, job.Magician)
	character.GainEquipment(l)(c.CharacterId, item.BeginnerMagicWand)
	character.ResetAP(l)(c.CharacterId)

	m := message.NewBuilder().
		AddText("Alright, from here out, you are a part of us! You'll be living the life of a wanderer at ..., but just be patient as soon, you'll be living the high life. Alright, it ain't much, but I'll give you some of my abilities... HAAAHHH!!!")
	return SendNext(l, c, m.String(), Exit())
}
