package discrete

import (
	"atlas-ncs/character"
	"atlas-ncs/conversation/script"
	_map "atlas-ncs/map"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// KiriruEllinia is located in Victoria Road - Sky Ferry <To Ereve> (101000400)
type KiriruEllinia struct {
}

func (r KiriruEllinia) NPCId() uint32 {
	return npc.KiriruEllinia
}

func (r KiriruEllinia) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Eh... So... Um... Are you trying to leave Victoria to go to a different region? You can take this boat to ").
		BlueText().AddText("Ereve").
		BlackText().AddText(". There, you will see bright sunlight shinning on the leaves and feel a gentle breeze on your skin. It's where Shinsoo and Empress Cygnus are. Would you like to go to Ereve? It will take about ").
		BlueText().AddText("2 Minutes").
		BlackText().AddText(", and it will cost you ").
		BlueText().AddText("1000").
		BlackText().AddText(" mesos.").NewLine().
		OpenItem(0).BlueText().AddText(" Ereve (1000 mesos)").CloseItem()
	return script.SendListSelectionExit(l, c, m.String(), r.DestinationSelection, r.NotInterested)
}

func (r KiriruEllinia) NotInterested(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("If you're not interested, then oh well...")
	return script.SendOk(l, c, m.String())
}

func (r KiriruEllinia) DestinationSelection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.Validate
	}
	return nil
}

func (r KiriruEllinia) Validate(l logrus.FieldLogger, c script.Context) script.State {
	if !character.HasMeso(l)(c.CharacterId, 1000) {
		return r.NotEnoughMeso(l, c)
	}
	return r.Process(l, c)
}

func (r KiriruEllinia) NotEnoughMeso(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmm... Are you sure you have ").
		BlueText().AddText("1000").
		BlackText().AddText(" Mesos? Check your Inventory and make sure you have enough. You must pay the fee or I can't let you get on...")
	return script.SendOk(l, c, m.String())
}

func (r KiriruEllinia) Process(l logrus.FieldLogger, c script.Context) script.State {
	err := character.GainMeso(l)(c.CharacterId, -1000)
	if err != nil {
		l.WithError(err).Errorf("Unable to process payment by character %d.", c.CharacterId)
	}
	return script.Warp(_map.EmpressRoadToEreveFromEllinia)(l, c)
}
