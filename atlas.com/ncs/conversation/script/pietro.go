package script

import (
	"atlas-ncs/character"
	"atlas-ncs/item"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Pietro is located in Hidden Street - Receiving the Reward For the Event (109050000)
type Pietro struct {
}

func (r Pietro) NPCId() uint32 {
	return npc.Pietro
}

func (r Pietro) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Bam bam bam bam!! You have won the game from the ").NewLine().
		BlueText().AddText("EVENT").
		BlackText().AddText(". Congratulations on making it this far!")
	return SendNext(l, c, m.String(), r.Prize)
}

func (r Pietro) Prize(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You'll be awarded the ").
		BlueText().AddText("Scroll of Secrets").
		BlackText().AddText(" as the winning prize. On the scroll, it has secret information written in ancient characters.")
	return SendNext(l, c, m.String(), r.SomethingGood)
}

func (r Pietro) SomethingGood(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("The Scroll of Secrets can be deciphered by ").
		RedText().AddText("Chun Ji").
		BlackText().AddText(" or ").NewLine().
		RedText().AddText("Geanie").
		BlackText().AddText(" at Ludibrium. Bring it with you and something good's bound to happen.")
	return SendNext(l, c, m.String(), r.Validate)
}

func (r Pietro) Validate(l logrus.FieldLogger, c Context) State {
	if !character.CanHold(l)(c.CharacterId, item.ScrollOfSecrets) {
		return r.MakeRoom(l, c)
	}
	return r.Process(l, c)
}

func (r Pietro) Process(l logrus.FieldLogger, c Context) State {
	character.GainItem(l)(c.CharacterId, item.ScrollOfSecrets, 1)
	mapId := character.SavedLocation(l)(c.CharacterId, "EVENT")
	err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
	}
	return Exit()(l, c)
}

func (r Pietro) MakeRoom(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I think your Etc window is full. Please make room, then talk to me.")
	return SendOk(l, c, m.String())
}
