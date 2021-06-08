package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/job"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// AthenaPierce is located in Victoria Road - Bowman Instructional School (100000201)
type AthenaPierce struct {
}

func (r AthenaPierce) NPCId() uint32 {
	return npc.AthenaPierce
}

func (r AthenaPierce) Initial(l logrus.FieldLogger, c Context) State {
	if character.IsJob(l)(c.CharacterId, job.Beginner) {
		return r.FirstJobInitial(l, c)
	}
	return nil
}

func (r AthenaPierce) FirstJobInitial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("So you decided to become a ").
		RedText().AddText("bowman").
		BlackText().AddText("? There are some standards to meet, y'know... ").
		BlueText().AddText("Your level should be at least 10, with at least 25 DEX").
		BlackText().AddText(". Let's see.")
	return SendNext(l, c, m.String(), r.FirstJobRequirementCheck)
}

func (r AthenaPierce) FirstJobRequirementCheck(l logrus.FieldLogger, c Context) State {
	if character.MeetsCriteria(l)(c.CharacterId, character.IsLevelCriteria(10), character.HasDexterityCriteria(25)) {
		m := message.NewBuilder().
			AddText("It is an important and final choice. You will not be able to turn back.")
		return SendNextPrevious(l, c, m.String(), r.AwardFirstJob, r.FirstJobInitial)
	}

	m := message.NewBuilder().
		AddText("Train a bit more until you reach the base requirements and I can show you the way of the ").
		RedText().AddText("Bowman").
		BlackText().AddText(".")
	return SendOk(l, c, m.String())
}

func (r AthenaPierce) AwardFirstJob(l logrus.FieldLogger, c Context) State {
	if !character.CanHold(l)(c.CharacterId, item.BeginnerBowmanBow) || !character.CanHold(l)(c.CharacterId, item.ArrowForBow) {
		m := message.NewBuilder().
			AddText("Make some room in your inventory and talk back to me.")
		return SendNext(l, c, m.String(), Exit())
	}

	character.ChangeJob(l)(c.CharacterId, job.Bowman)
	character.GainEquipment(l)(c.CharacterId, item.BeginnerBowmanBow)
	character.GainItem(l)(c.CharacterId, item.ArrowForBow, 1000)
	character.ResetAP(l)(c.CharacterId)

	m := message.NewBuilder().
		AddText("Alright, from here out, you are a part of us! You'll be living the life of a wanderer at ..., but just be patient as soon, you'll be living the high life. Alright, it ain't much, but I'll give you some of my abilities... HAAAHHH!!!")
	return SendNext(l, c, m.String(), Exit())
}
