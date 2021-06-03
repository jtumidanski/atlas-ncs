package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
	"math/rand"
)

// Karcasa is located in The Burning Sands - Tent of the Entertainers (260010600)
type Karcasa struct {
}

func (r Karcasa) NPCId() uint32 {
	return npc.Karcasa
}

func (r Karcasa) Initial(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("I don't know how you found out about this, but you came to the right place! For those that wandered around Nihal Desert and are getting homesick, I am offering a flight straight to Victoria Island, non-stop! Don't worry about the flying ship--it's only fallen once or twice! Don't you feel claustrophobic being in a long flight on that small ship?")
	return SendNext(l, c, m.String(), r.PleaseRemember)
}

func (r Karcasa) PleaseRemember(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Please remember two things. One, this line is actually for overseas shipping, so ").
		RedText().AddText("I cannot guarantee exactly which town you'll land").
		BlackText().AddText(". Two, since I am putting you in this special flight, it'll be a bit expensive. The service charge is ").
		BoldText().BlueText().AddText("10,000 mesos").
		NormalText().BlackText().AddText(". There's a flight that is about to take off. Are you interested in this direct flight?")
	return SendYesNo(l, c, m.String(), r.Validate, r.Scared)
}

func (r Karcasa) Scared(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Aye...are you scared of speed or heights? You can't trust my flying skills? Trust me, I've worked out all the kinks!")
	return SendOk(l, c, m.String())
}

func (r Karcasa) Validate(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 10000) {
		return r.ShortOnCash(l, c)
	}
	return r.OkGetReady(l, c)
}

func (r Karcasa) OkGetReady(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Okay, ready to takeoff~")
	return SendNext(l, c, m.String(), r.Process)
}

func (r Karcasa) ShortOnCash(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hey, are you short on cash? I told you you'll need ").
		BlueText().AddText("10,000").
		BlackText().AddText(" mesos to get on this.")
	return SendOk(l, c, m.String())
}

func (r Karcasa) Process(l logrus.FieldLogger, c Context) State {
	maps := []uint32{_map.Henesys, _map.Ellinia, _map.Perion, _map.KerningCity, _map.LithHarbor}
	mapId := maps[rand.Intn(len(maps))]
	err := character.GainMeso(l)(c.CharacterId, -int32(10000))
	if err != nil {
		l.WithError(err).Errorf("Unable to process payment for character %d.", c.CharacterId)
	}
	err = npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
	if err != nil {
		l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
	}
	return Exit()(l, c)
}
