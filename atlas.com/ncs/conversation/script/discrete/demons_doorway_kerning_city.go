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

// DemonsDoorwayKerningCity is located in Victoria Road - Kerning City Middle Forest III (103030200)
type DemonsDoorwayKerningCity struct {
}

func (r DemonsDoorwayKerningCity) NPCId() uint32 {
	return npc.DemonsDoorwayKerningCity
}

func (r DemonsDoorwayKerningCity) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !character.QuestStarted(l)(c.CharacterId, 28219) {
		m := message.NewBuilder().AddText("The entrance is blocked by a strange force.")
		return script.SendOk(l, c, m.String())
	}

	if !character.HasItem(l)(c.CharacterId, item.ValeforsEmblem) {
		m := message.NewBuilder().AddText("he entrance is blocked by a force that can only be lifted by those holding an emblem.")
		return script.SendOk(l, c, m.String())
	}

	m := message.NewBuilder().
		AddText("Would you like to move to ").
		BlueText().ShowMap(_map.ValeforStrollingPath).
		BlackText().AddText("?")
	return script.SendYesNo(l, c, m.String(), r.Process, script.Exit())
}

func (r DemonsDoorwayKerningCity) Process(l logrus.FieldLogger, c script.Context) script.State {
	if character.HasItem(l)(c.CharacterId, item.LargeModelOfACoin) {
		character.GainItem(l)(c.CharacterId, item.LargeModelOfACoin, -1)
	}
	if character.HasItem(l)(c.CharacterId, item.GoldenFeather) {
		character.GainItem(l)(c.CharacterId, item.GoldenFeather, -1)
	}
	return script.WarpById(_map.ValeforStrollingPath, 0)(l, c)
}