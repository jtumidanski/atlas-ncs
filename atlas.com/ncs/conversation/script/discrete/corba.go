package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Corba is located in Leafre - Station (240000110)
type Corba struct {
}

func (r Corba) NPCId() uint32 {
	return npc.Corba
}

func (r Corba) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("If you had wings, I'm sure you could go there.  But, that alone won't be enough.  If you want to fly though the wind that's sharper than a blade, you'll need tough scales as well.  I'm the only Halfling left that knows the way back... If you want to go there, I can transform you.  No matter what you are, for this moment, you will become a ").
		BlueText().AddText("Dragon").
		BlackText().AddText("...").NewLine().
		OpenItem(0).BlueText().AddText("I want to become a dragon.").CloseItem()
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r Corba) Selection(selection int32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		character.UseItem(l)(c.CharacterId, item.MiniDracoTransformation)
		return script.WarpById(_map.WayToTempleOfTime, 0)(l, c)
	}
}
