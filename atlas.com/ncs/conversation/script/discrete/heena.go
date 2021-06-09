package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Heena is located in Maple Road - Mushroom Town (10000)
type Heena struct {
}

func (r Heena) NPCId() uint32 {
	return npc.Heena
}

func (r Heena) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return r.DoneTraining(l, c)
}

func (r Heena) DoneTraining(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Are you done with your training? If you wish, I will send you out from this training camp.")
	return script.SendYesNo(l, c, m.String(), r.Yes, r.No)
}

func (r Heena) No(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Haven't you finished the training program yet? If you want to leave this place, please do not hesitate to tell me.")
	return script.SendNext(l, c, m.String(), script.Exit())
}

func (r Heena) Yes(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Then, I will send you out from here. Good job.")
	return script.SendNext(l, c, m.String(), script.WarpById(_map.InASmallForest, 0))
}
