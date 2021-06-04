package script

import (
	"atlas-ncs/character"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"fmt"
	"github.com/sirupsen/logrus"
)

type Spinel struct {
}

func (r Spinel) NPCId() uint32 {
	return npc.Spinel
}

func (r Spinel) Initial(l logrus.FieldLogger, c Context) State {
	if c.MapId == _map.MushroomShrine || c.MapId == _map.TrendZoneMetropolis {
		return r.TravelingHello(l, c)
	}

	mapId, fee := r.GetDestinationAndFee(c)
	return r.TravelHello(mapId, fee)(l, c)
}

func (r Spinel) GetDestinationAndFee(c Context) (uint32, int) {
	mapId := _map.MushroomShrine
	fee := 3000
	if c.MapId == _map.BoatQuayTown {
		mapId = _map.TrendZoneMetropolis
		fee = 10000
	}
	return mapId, fee
}

func (r Spinel) TravelHello(mapId uint32, fee int) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("If you're tired of the monotonous daily life, how about getting out for a change? there's nothing quite like soaking up a new culture, learning something new by the minute! It's time for you to get out and travel. We, at the Maple Travel Agency recommend you going on a ").
			BlueText().AddText("World Tour").
			BlackText().AddText("! Are you worried about the travel expense? You shouldn't be! We, the ").
			BlueText().AddText("Maple Travel Agency").
			BlackText().AddText(", have carefully come up with a plan to let you travel for ONLY ").
			BlueText().AddText(fmt.Sprintf("%d mesos", fee)).
			BlackText().AddText("!")
		return SendNext(l, c, m.String(), r.CurrentlyOffer(mapId, fee))
	}
}

func (r Spinel) TravelingHello(l logrus.FieldLogger, c Context) State {
	mapId := character.SavedLocation(l)(c.CharacterId, "WORLDTOUR")
	m := message.NewBuilder().
		AddText("How's the traveling? Are you enjoying it?").NewLine().
		OpenItem(0).BlueText().AddText("Yes, I'm done with traveling. Can I go back to ").ShowMap(mapId).AddText("?").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("No, I'd like to continue exploring this place.").CloseItem()
	return SendListSelection(l, c, m.String(), r.TravelingSelection(mapId))
}

func (r Spinel) TravelingSelection(mapId uint32) ProcessSelection {
	return func(selection int32) StateProducer {
		switch selection {
		case 0:
			return r.AlrightHome(mapId)
		case 1:
			return r.LetMeKnow
		}
		return nil
	}
}

func (r Spinel) AlrightHome(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("Alright. I'll take you back to where you were before your visit. If you ever feel like traveling again down the road, please let me know!")
		return SendNext(l, c, m.String(), r.WarpBack(mapId))
	}
}

func (r Spinel) WarpBack(mapId uint32) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		character.ClearSavedLocation(l)(c.CharacterId, "WORLDTOUR")
		err := npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return Exit()(l, c)
	}
}

func (r Spinel) LetMeKnow(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("OK. If you ever change your mind, please let me know.")
	return SendOk(l, c, m.String())
}

func (r Spinel) CurrentlyOffer(mapId uint32, fee int) StateProducer {
	if mapId == _map.MushroomShrine {
		return r.CurrentlyOfferMushroomShrine(mapId, fee)
	} else {
		return r.CurrentlyOfferMalaysia(mapId, fee)
	}
}

func (r Spinel) CurrentlyOfferMushroomShrine(mapId uint32, fee int) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("We currently offer this place for you traveling pleasure: ").
			BlueText().AddText("Mushroom Shrine of Japan").
			BlackText().AddText(". I'll be there serving you as the travel guide. Rest assured, the number of destinations will be increase over time. Now, would you like to head over to the ").
			AddText("Mushroom Shrine?").NewLine().
			OpenItem(0).BlueText().AddText("Yes, take me to Mushroom Shrine (Japan)")
		return SendListSelection(l, c, m.String(), r.ConfirmMushroomShrine(mapId, fee))
	}
}

func (r Spinel) ConfirmMushroomShrine(mapId uint32, fee int) ProcessSelection {
	return func(selection int32) StateProducer {
		return func(l logrus.FieldLogger, c Context) State {
			m := message.NewBuilder().
				AddText("Would you like to travel to ").
				BlueText().AddText("Mushroom Shrine of Japan").
				BlackText().AddText("? ").AddText("If you desire to feel the essence of Japan, there's nothing like visiting the Shrine, a Japanese cultural melting pot. Mushroom Shrine is a mythical place that serves the incomparable Mushroom God from ancient times.")
			return SendYesNo(l, c, m.String(), r.Validate(fee, r.CheckoutMushroomShrine(mapId, fee)), Exit())
		}
	}
}

func (r Spinel) CurrentlyOfferMalaysia(mapId uint32, fee int) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("We currently offer this place for you traveling pleasure: ").
			BlueText().AddText("Trend Zone of Malaysia").
			BlackText().AddText(". ").
			RedText().ShowNPC(npc.Audrey).
			BlackText().AddText("'ll be there serving you as the travel guide. Rest assured, the number of destinations will be increase over time. Now, would you like to head over to the ").
			AddText("Metropolis?").NewLine().
			OpenItem(0).BlueText().AddText("Yes, take me to Metropolis (Malaysia)")
		return SendListSelection(l, c, m.String(), r.ConfirmMalaysia(mapId, fee))
	}
}

func (r Spinel) ConfirmMalaysia(mapId uint32, fee int) ProcessSelection {
	return func(selection int32) StateProducer {
		return func(l logrus.FieldLogger, c Context) State {
			m := message.NewBuilder().
				AddText("Would you like to travel to ").
				BlueText().AddText("Trend Zone of Malaysia").
				BlackText().AddText("? ").AddText("If you desire to feel the heat of the tropics on an upbeat environment, the residents of Malaysia are eager to welcome you. Also, the metropolis itself is the heart of the local economy, that place is known to always offer something to do or to visit around.")
			return SendYesNo(l, c, m.String(), r.Validate(fee, r.CheckoutMalaysia(mapId, fee)), Exit())
		}
	}
}

func (r Spinel) Validate(fee int, checkout StateProducer) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.HasMeso(l)(c.CharacterId, uint32(fee)) {
			return r.NotEnoughMeso(l, c)
		}
		return checkout(l, c)
	}
}

func (r Spinel) NotEnoughMeso(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You don't have enough mesos to take the travel.")
	return SendOk(l, c, m.String())
}

func (r Spinel) CheckoutMushroomShrine(mapId uint32, fee int) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("Check out the female shaman serving the Mushroom God, and I strongly recommend trying Takoyaki, Yakisoba, and other delicious food sold in the streets of Japan. Now, let's head over to #bMushroom Shrine#k, a mythical place if there ever was one.")
		return SendNext(l, c, m.String(), r.Process(mapId, fee))
	}
}

func (r Spinel) CheckoutMalaysia(mapId uint32, fee int) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().
			AddText("Once there, I strongly suggest you to schedule a visit to Kampung Village. Why? Surely you've come to know about the fantasy theme park Spooky World? No? It's simply put the greatest theme park around there, it's worth a visit! Now, let's head over to the ").
			BlueText().AddText("Trend Zone of Malaysia").
			BlackText().AddText(".")
		return SendNext(l, c, m.String(), r.Process(mapId, fee))
	}
}

func (r Spinel) Process(mapId uint32, fee int) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		character.SaveLocation(l)(c.CharacterId, "WORLDTOUR")
		err := character.GainMeso(l)(c.CharacterId, int32(-fee))
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment for character %d.", c.CharacterId)
		}
		err = npc.WarpById(l)(c.WorldId, c.ChannelId, c.CharacterId, mapId, 0)
		if err != nil {
			l.WithError(err).Errorf("Unable to warp character %d to %d as a result of a conversation with %d.", c.CharacterId, mapId, c.NPCId)
		}
		return Exit()(l, c)
	}
}
