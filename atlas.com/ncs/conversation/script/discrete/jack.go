package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Jack is located in The Nautilus - Top Floor - Hallway (120000100)
type Jack struct {
}

func (r Jack) NPCId() uint32 {
	return npc.Jack
}

func (r Jack) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.HasItem(l)(c.CharacterId, item.DirtyTreasureMap) {
		return r.ScratchScratch(l, c)
	}
	return r.CanIKeepIt(l, c)
}

func (r Jack) ScratchScratch(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("(Scratch scratch...)")
	return script.SendOk(l, c, m.String())
}

func (r Jack) CanIKeepIt(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hey, nice ").
		BlueText().AddText("Treasure Map").
		BlackText().AddText(" you have there? ").
		RedText().AddText("Can I keep it").
		BlackText().AddText(" for the Nautilus crew, if you don't need it any longer?")
	return script.SendYesNo(l, c, m.String(), r.Remove, script.Exit())
}

func (r Jack) Remove(l logrus.FieldLogger, c script.Context) script.State {
	character.GainItem(l)(c.CharacterId, item.DirtyTreasureMap, -1)
	return script.Exit()(l, c)
}
