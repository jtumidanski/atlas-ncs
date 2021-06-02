package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"fmt"
	"github.com/sirupsen/logrus"
)

// HarpStringD is located in Hidden Street - Eliza's Garden (920020000)
type HarpStringD struct {
}

func (r HarpStringD) NPCId() uint32 {
	return npc.HarpStringD
}

func (r HarpStringD) Initial(l logrus.FieldLogger, c Context) State {
	_map.PlaySound(l)(c.WorldId, c.ChannelId, c.MapId, fmt.Sprintf("orbis/%s", HarpSounds[c.NPCId-2012027]))

	if !character.QuestStarted(l)(c.CharacterId, 3114) {
		return Exit()(l, c)
	}

	progress := -1 * character.QuestProgressInt(l)(c.CharacterId, 3114, 0)
	if progress <= -1 {
		return Exit()(l, c)
	}

	nextNote := HarpSong[progress]
	if 'D' != nextNote {
		character.SetQuestProgress(l)(c.CharacterId, 3114, 0, 0)
		character.ShowEffect(l)(c.CharacterId, "quest/party/wrong_kor")
		character.PlaySound(l)(c.CharacterId, "Party1/Failed")
		character.SendNotice(l)(c.CharacterId, "PINK_TEXT", "You've missed the note... Start over again.")
	} else {
		nextNote = HarpSong[progress+1]
		if nextNote == '|' {
			progress++
			if progress == 45 {
				character.SendNotice(l)(c.CharacterId, "PINK_TEXT", "Twinkle, twinkle, little star, how I wonder what you are.")
				character.SetQuestProgress(l)(c.CharacterId, 3114, 0, 42)
				character.ShowEffect(l)(c.CharacterId, "quest/party/clear")
				character.PlaySound(l)(c.CharacterId, "Party1/Clear")
			} else {
				if progress == 14 {
					character.SendNotice(l)(c.CharacterId, "PINK_TEXT", "Twinkle, twinkle, little star, how I wonder what you are.")
				} else if progress == 22 {
					character.SendNotice(l)(c.CharacterId, "PINK_TEXT", "Up above the world so high,")
				} else if progress == 30 {
					character.SendNotice(l)(c.CharacterId, "PINK_TEXT", "like a diamond in the sky.")
				}
			}
		}
		character.SetQuestProgress(l)(c.CharacterId, 3114, 0, uint32(-1*(progress+1)))
	}
	return Exit()(l, c)
}
