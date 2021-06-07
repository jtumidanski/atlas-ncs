package script

import (
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Konpei is located in Zipangu - Showa Town (801000000)
type Konpei struct {
}

func (r Konpei) NPCId() uint32 {
	return npc.Konpei
}

func (r Konpei) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("What do you want from me?").NewLine().
		OpenItem(0).BlueText().AddText("Gather up some information on the hideout.").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Take me to the hideout").CloseItem().NewLine().
		OpenItem(2).BlackText().AddText("Nothing").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r Konpei) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Infested
	case 1:
		return r.ToTheHideout
	case 2:
		return r.IAmBusy
	}
	return nil
}

func (r Konpei) Infested(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I can take you to the hideout, but the place is infested with thugs looking for trouble. You'll need to be both incredibly strong and brave to enter the premise. At the hideaway, you'll find the Boss that controls all the other bosses around this area. It's easy to get to the hideout, but the room on the top floor of the place can only be entered ONCE a day. The Boss's Room is not a place to mess around. I suggest you don't stay there for too long; you'll need to swiftly take care of the business once inside. The boss himself is a difficult foe, but you'll run into some incredibly powerful enemies on you way to meeting the boss! It ain't going to be easy.")
	return SendOk(l, c, m.String())
}

func (r Konpei) ToTheHideout(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Oh, the brave one. I've been awaiting your arrival. If these").NewLine().
		AddText("thugs are left unchecked, there's no telling what going to").NewLine().
		AddText("happen in this neighborhood. Before that happens, I hope").NewLine().
		AddText("you take care of all them and beat the boss, who resides").NewLine().
		AddText("on the 5th floor. You'll need to be on alert at all times, since").NewLine().
		AddText("the boss is too tough for even wise men to handle.").NewLine().
		AddText("Looking at your eyes, however, I can see that eye of the").NewLine().
		AddText("tiger, the eyes that tell me you can do this. Let's go!")
	return SendNext(l, c, m.String(), WarpByName(_map.NearTheHideout, "in00"))
}

func (r Konpei) IAmBusy(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("I'm a busy person! Leave me alone if that's all you need!")
	return SendOk(l, c, m.String())
}
