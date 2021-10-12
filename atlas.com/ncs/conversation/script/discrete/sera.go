package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

// Sera is located in Maple Road : Entrance - Mushroom Town Training Camp (0), Maple Road: Upper level of the Training Camp (1), Maple Road : Entrance - Mushroom Town Training Camp (3)
type Sera struct {
}

func (r Sera) NPCId() uint32 {
	return npc.Sera
}

func (r Sera) Initial(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	if c.MapId == _map.MushroomTownTrainingCampEntrance || c.MapId == _map.MushroomTownTrainingCampEntrance2 {
		return r.Welcome(l, span, c)
	} else {
		return r.FirstTraining(l, span, c)
	}
}

func (r Sera) Welcome(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Welcome to the world of MapleStory. The purpose of this training camp is to help beginners. Would you like to enter this training camp? Some people start their journey without taking the training program. But I strongly recommend you take the training program first.")
	return script.SendYesNo(l, span, c, m.String(), r.OkThen, r.ConfirmStartJourney)
}

func (r Sera) FirstTraining(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("This is the image room where your first training program begins. In this room, you will have an advance look into the job of your choice.")
	return script.SendNext(l, span, c, m.String(), r.EntitledToAJob)
}

func (r Sera) ConfirmStartJourney(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Do you really want to start your journey right away?")
	return script.SendYesNo(l, span, c, m.String(), r.Skip, r.CancelSkip)
}

func (r Sera) Skip(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("It seems like you want to start your journey without taking the training program. Then, I will let you move on to the training ground. Be careful~")
	return script.SendNext(l, span, c, m.String(), r.WarpSkip)
}

func (r Sera) WarpTraining(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.UpperLevelOfTheTrainingCamp, 0)(l, span, c)
}

func (r Sera) WarpSkip(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	return script.WarpById(_map.InASmallForest, 0)(l, span, c)
}

func (r Sera) CancelSkip(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Please talk to me again when you finally made your decision.")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r Sera) EntitledToAJob(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Once you train hard enough, you will be entitled to occupy a job. You can become a Bowman in Henesys, a Magician in Ellinia, a Warrior in Perion, and a Thief in Kerning City...")
	return script.SendNext(l, span, c, m.String(), script.Exit())
}

func (r Sera) OkThen(l logrus.FieldLogger, span opentracing.Span, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Ok then, I will let you enter the training camp. Please follow your instructor's lead.")
	return script.SendNext(l, span, c, m.String(), r.WarpTraining)
}
