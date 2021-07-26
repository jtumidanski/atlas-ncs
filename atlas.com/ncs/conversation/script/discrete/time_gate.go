package discrete

import (
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"github.com/sirupsen/logrus"
)

// TimeGate is located in Tera Forest   - Tera Forest Time Gate (240070000)
type TimeGate struct {
}

func (r TimeGate) NPCId() uint32 {
	return npc.TimeGate
}

func (r TimeGate) Initial(l logrus.FieldLogger, c script.Context) script.State {
	if !quest.IsCompleted(l)(c.CharacterId, 3718) {
		return r.NotActiveYet(l, c)
	}

	limit := 0
	quests := []uint32{3719, 3724, 3730, 3736, 3742, 3748}
	for _, q := range quests {
		if !quest.IsCompleted(l)(c.CharacterId, q) {
			break
		}
		limit++
	}

	if limit == 0 {
		return r.ProveYourValor(l, c)
	}

	warpStrings := []string{"Year 2021 - Average Town Entrance", "Year 2099 - Midnight Harbor Entrance", "Year 2215 - Bombed City Center Retail District", "Year 2216 - Ruined City Intersection", "Year 2230 - Dangerous Tower Lobby", "Year 2503 - Air Battleship Bow"}
	m := message.NewBuilder()
	for i := 0; i < limit; i++ {
		m = m.OpenItem(i).AddText(warpStrings[i]).CloseItem().NewLine()
	}
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r TimeGate) ProveYourValor(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Prove your valor against the ").
		BlueText().AddText("Guardian Nex").
		BlackText().AddText(" before unlocking next Neo City maps.")
	return script.SendOk(l, c, m.String())
}

func (r TimeGate) NotActiveYet(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("The time machine has not been activated yet.")
	return script.SendOk(l, c, m.String())
}

func (r TimeGate) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Warp(_map.Year2021AverageTownEntrance)
	case 1:
		return r.Warp(_map.Year2099MidnightHarborEntrance)
	case 2:
		return r.Warp(_map.Year2215BombedCityCenterRetailDistrict)
	case 3:
		return r.Warp(_map.Year2216RuinedCityIntersection)
	case 4:
		return r.Warp(_map.Year2230DangerousTowerLobby)
	case 5:
		return r.Warp(_map.Year2503AirBattleshipBow)
	}
	return r.CompleteYourMission
}

func (r TimeGate) Warp(mapId uint32) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		return script.WarpById(mapId, 1)(l, c)
	}
}

func (r TimeGate) CompleteYourMission(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Complete your mission first.")
	return script.SendOk(l, c, m.String())
}
