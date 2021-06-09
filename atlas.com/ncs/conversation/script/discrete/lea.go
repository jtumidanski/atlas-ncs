package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Lea is located in Orbis - Guild Headquarters <Hall of Fame> (200000301)
type Lea struct {
}

func (r Lea) NPCId() uint32 {
	return npc.Lea
}

func (r Lea) Initial(l logrus.FieldLogger, c script.Context) script.State {
	return r.Hello(l, c)
}

func (r Lea) Hello(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("What would you like to do?").NewLine().
		OpenItem(0).BlueText().AddText("Create/Change your Guild Emblem").CloseItem()
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r Lea) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.ChangeEmblem
	}
	return nil
}

func (r Lea) ChangeEmblem(l logrus.FieldLogger, c script.Context) script.State {
	if !character.IsGuildLeader(l)(c.CharacterId) {
		return r.MustBeLeader(l, c)
	}
	return r.Confirmation(l, c)
}

func (r Lea) MustBeLeader(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("You must be the Guild Leader to change the Emblem. Please tell your leader to speak with me.")
	return script.SendOk(l, c, m.String())
}

func (r Lea) Confirmation(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Creating or changing Guild Emblem costs ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 5000000)).
		BlackText().AddText(", are you sure you want to continue?")
	return script.SendYesNo(l, c, m.String(), r.ValidateChange, script.Exit())
}

func (r Lea) ValidateChange(l logrus.FieldLogger, c script.Context) script.State {
	//TODO implement
	return script.Exit()(l, c)
}
