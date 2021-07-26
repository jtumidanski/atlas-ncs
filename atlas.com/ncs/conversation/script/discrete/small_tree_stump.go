package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// SmallTreeStump is located in Victoria Road - Top of the Tree That Grew (101010103)
type SmallTreeStump struct {
}

func (r SmallTreeStump) NPCId() uint32 {
	return npc.SmallTreeStump
}

func (r SmallTreeStump) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !quest.IsStarted(l)(c.CharacterId, 20716) {
		return r.NeverEndingFlow(l, c)
	}
	if character.HasItem(l)(c.CharacterId, item.ClearTreeSap) {
		return r.NeverEndingFlow(l, c)
	}
	if !character.CanHold(l)(c.CharacterId, item.ClearTreeSap) {
		return r.MakeRoom(l, c)
	}

	return r.GainTreeSap(l, c)
}

func (r SmallTreeStump) GainTreeSap(l logrus.FieldLogger, c script.Context) script.State {
	character.GainItem(l)(c.CharacterId, item.ClearTreeSap, 1)
	return r.ShowSuccess(l, c)
}

func (r SmallTreeStump) ShowSuccess(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("You bottled up some of the clear tree sap.  ").
		ShowItemImage2(item.ClearTreeSap)
	return script.SendOk(l, c, m.String())
}

func (r SmallTreeStump) NeverEndingFlow(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("A never ending flow of sap is coming from this small tree stump.")
	return script.SendOk(l, c, m.String())
}

func (r SmallTreeStump) MakeRoom(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Make sure you have a free spot in your ETC inventory.")
	return script.SendOk(l, c, m.String())
}
