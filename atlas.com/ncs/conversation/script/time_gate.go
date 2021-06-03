package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// TimeGate is located in Tera Forest   - Tera Forest Time Gate (240070000)
type TimeGate struct {
}

func (r TimeGate) NPCId() uint32 {
	return npc.TimeGate
}

func (r TimeGate) Initial(l logrus.FieldLogger, c Context) State {
	if !character.QuestCompleted(l)(c.CharacterId, 3718) {
		return r.NotActiveYet(l, c)
	}

	limit := 0
	quests := []uint32{3719, 3724, 3730, 3736, 3742, 3748}
	for _, quest := range quests {
		if !character.QuestCompleted(l)(c.CharacterId, quest) {
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
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r TimeGate) ProveYourValor(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Prove your valor against the ").
		BlueText().AddText("Guardian Nex").
		BlackText().AddText(" before unlocking next Neo City maps.")
	return SendOk(l, c, m.String())
}

func (r TimeGate) NotActiveYet(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("The time machine has not been activated yet.")
	return SendOk(l, c, m.String())
}

func (r TimeGate) Selection(selection int32) StateProducer {
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

func (r TimeGate) Warp(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 1)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return Exit()(l, c)
	}
}

func (r TimeGate) CompleteYourMission(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Complete your mission first.")
	return SendOk(l, c, m.String())
}