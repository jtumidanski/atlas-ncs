package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Shanks is located in Maple Road : Southperry (60000)
type Shanks struct {
}

func (r Shanks) NPCId() uint32 {
	return npc.Shanks
}

func (r Shanks) Initial(l logrus.FieldLogger, c Context) State {
	return r.TakeTheShip(l, c)
}

func (r Shanks) TakeTheShip(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Take this ship and you'll head off to a bigger continent. For ").
		BoldText().AddText("150 mesos").
		NormalText().AddText(", I'll take you to ").
		BlueText().AddText("Victoria Island").
		BlackText().AddText(". The thing is, once you leave this place, you can't ever come back. What do you think? Do you want to go to Victoria Island?")
	return SendYesNo(l, c, m.String(), r.Yes, r.No)
}

func (r Shanks) No(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hmm... I guess you still have things to do here?")
	return SendNext(l, c, m.String(), Exit())
}

func (r Shanks) Yes(l logrus.FieldLogger, c Context) State {
	if character.HasItem(l)(c.CharacterId, item.LucasRecommendationLetter) {
		m := message.NewBuilder().
			AddText("Okay, now give me 150 mesos... Hey, what's that? Is that the recommendation letter from Lucas, the chief of Amherst? Hey, you should have told me you had this. I, Shanks, recognize greatness when I see one, and since you have been recommended by Lucas, I see that you have a great, great potential as an adventurer. No way would I charge you for this trip!")
		return SendNext(l, c, m.String(), r.ConfirmUse)
	} else {
		m := message.NewBuilder().
			AddText("Bored of this place? Here... Give me ").
			BoldText().AddText("150 mesos").
			NormalText().AddText(" first...")
		return SendNext(l, c, m.String(), r.StrongEnough)
	}
}

func (r Shanks) ConfirmUse(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Since you have the recommendation letter, I won't charge you for this. Alright, buckle up, because we're going to head to Victoria Island right now, and it might get a bit turbulent!!")
	return SendNextPrevious(l, c, m.String(), r.WarpWithItem, r.Yes)
}

func (r Shanks) WarpWithItem(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.LucasRecommendationLetter, -1)

	err := npc.Processor(l).WarpById(c.WorldId, c.ChannelId, c.CharacterId, _map.LithHarbor, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.LithHarbor, c.NPCId)
	}
	return nil
}

func (r Shanks) StrongEnough(l logrus.FieldLogger, c Context) State {
	if !character.IsAboveLevel(l)(c.CharacterId, 6) {
		m := message.NewBuilder().
			AddText("Let's see... I don't think you are strong enough. You'll have to be at least Level 7 to go to Victoria Island.")
		return SendOk(l, c, m.String())
	}

	if !character.HasMeso(l)(c.CharacterId, 150) {
		m := message.NewBuilder().
			AddText("What? You're telling me you wanted to go without any money? You're one weirdo...")
		return SendOk(l, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Awesome! ").
		BoldText().AddText("150").
		NormalText().AddText(" mesos accepted! Alright, off to Victoria Island!")
	return SendNext(l, c, m.String(), r.WarpWithMeso)
}

func (r Shanks) WarpWithMeso(l logrus.FieldLogger, c Context) State {
	err := character.GainMeso(l)(c.CharacterId, -150)
	if err != nil {
		l.WithError(err).Errorf("Unable to complete meso transaction with shanks for %d.", c.CharacterId)
		return nil
	}

	err = npc.Processor(l).WarpById(c.WorldId, c.ChannelId, c.CharacterId, _map.LithHarbor, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.LithHarbor, c.NPCId)
	}
	return nil
}
