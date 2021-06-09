package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/event"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Arturo is located in Hidden Street - Abandoned Tower<Determine to adventure> (922011100)
type Arturo struct {
}

func (r Arturo) NPCId() uint32 {
	return npc.Arturo
}

func (r Arturo) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Congratulations on sealing the dimensional crack! For all of your hard work, I have a gift for you! Here take this prize.")
	return script.SendNext(l, c, m.String(), r.Validate)
}

func (r Arturo) Validate(l logrus.FieldLogger, c script.Context) script.State {
	ok := event.GiveEventReward(l)(c.CharacterId)
	if !ok {
		return r.NeedInventorySpace(l, c)
	}
	return script.WarpById(_map.EosTower101stFloor, 0)(l, c)
}

func (r Arturo) NeedInventorySpace(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("It seems you don't have a free slot in either your ").
		RedText().AddText("Equip").
		BlackText().AddText(", ").
		RedText().AddText("Use").
		BlackText().AddText(" or ").
		RedText().AddText("Etc").
		BlackText().AddText(" inventories. Please make some room and try again.")
	return script.SendOk(l, c, m.String())
}
