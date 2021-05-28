package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Sera is located in Maple Road : Entrance - Mushroom Town Training Camp (0), Maple Road: Upper level of the Training Camp (1), Maple Road : Entrance - Mushroom Town Training Camp (3)
type Sera struct {
}

func (r Sera) NPCId() uint32 {
	return npc.Sera
}

func (r Sera) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.MushroomTownTrainingCampEntrance || c.MapId == _map.MushroomTownTrainingCampEntrance2 {
		return r.Welcome(l, c)
	} else {
		return r.FirstTraining(l, c)
	}
}

func (r Sera) Welcome(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Welcome to the world of MapleStory. The purpose of this training camp is to help beginners. Would you like to enter this training camp? Some people start their journey without taking the training program. But I strongly recommend you take the training program first.")
	return SendYesNo(l, c, m.String(), r.OkThen, r.ConfirmStartJourney)
}

func (r Sera) FirstTraining(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("This is the image room where your first training program begins. In this room, you will have an advance look into the job of your choice.")
	return SendNext(l, c, m.String(), r.EntitledToAJob)
}

func (r Sera) ConfirmStartJourney(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Do you really want to start your journey right away?")
	return SendYesNo(l, c, m.String(), r.Skip, r.CancelSkip)
}

func (r Sera) Skip(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("It seems like you want to start your journey without taking the training program. Then, I will let you move on to the training ground. Be careful~")
	return SendNext(l, c, m.String(), r.WarpSkip)
}

func (r Sera) WarpTraining(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.UpperLevelOfTheTrainingCamp, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.UpperLevelOfTheTrainingCamp, c.NPCId)
	}
	return Exit()(l, c)
}

func (r Sera) WarpSkip(l logrus.FieldLogger, c Context) State {
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, _map.InASmallForest, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, _map.InASmallForest, c.NPCId)
	}
	return Exit()(l, c)
}

func (r Sera) CancelSkip(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Please talk to me again when you finally made your decision.")
	return SendNext(l, c, m.String(), Exit())
}

func (r Sera) EntitledToAJob(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Once you train hard enough, you will be entitled to occupy a job. You can become a Bowman in Henesys, a Magician in Ellinia, a Warrior in Perion, and a Thief in Kerning City...")
	return SendNext(l, c, m.String(), Exit())
}

func (r Sera) OkThen(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Ok then, I will let you enter the training camp. Please follow your instructor's lead.")
	return SendNext(l, c, m.String(), r.WarpTraining)
}
