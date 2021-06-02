package script

import (
	"atlas-ncs/character"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/party"
	"fmt"
	"github.com/sirupsen/logrus"
)

// Lenario is located in Orbis - Guild Headquarters <Hall of Fame> (200000301)
type Lenario struct {
}

func (r Lenario) NPCId() uint32 {
	return npc.Lenario
}

func (r Lenario) Initial(l logrus.FieldLogger, c Context) State {
	if !character.IsGuildLeader(l)(c.CharacterId) {
		return r.NonGuildLeaderHello(l, c)
	}
	return r.Hello(l, c)
}

func (r Lenario) NonGuildLeaderHello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hello there! I'm ").
		BlueText().AddText("Lenario").
		BlackText().AddText(". Just guild masters can attempt to form guild unions.")
	return SendOk(l, c, m.String())
}

func (r Lenario) Hello(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Hello there! I'm ").
		BlueText().AddText("Lenario").
		BlackText().AddText(".").NewLine().
		OpenItem(0).BlueText().AddText("Can you please tell me what Guild Union is all about?").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("How do I make a Guild Union?").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("I want to make a Guild Union.").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("I want to add more guilds for the Guild Union.").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("I want to break up the Guild Union.").CloseItem()
	return SendListSelection(l, c, m.String(), r.Selection)
}

func (r Lenario) Selection(selection int32) StateProducer {
	switch selection {
	case 0:
		return r.Information
	case 1:
		return r.HowToCreate
	case 2:
		return r.Create
	case 3:
		return r.AddAnotherGuild
	case 4:
		return r.Disband
	}
	return nil
}

func (r Lenario) Information(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Guild Union is just as it says, a union of a number of guilds to form a super group. I am in charge of managing these Guild Unions.")
	return SendOk(l, c, m.String())
}

func (r Lenario) HowToCreate(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("To make a Guild Union, two and only ").
		BlueText().AddText("two Guild Masters need to be in a party").
		BlackText().AddText(" and ").
		BlueText().AddText("both must be present on this room").
		BlackText().AddText(" on the same channel. The leader of this party will be assigned as the Guild Union Master.").NewLine().NewLine().
		AddText("Initially, ").
		BlueText().AddText("only two guilds").
		BlackText().AddText(" can make part of the new Union, but over the time you can ").
		RedText().AddText("expand").
		BlackText().AddText(" the Union capacity by talking to me when the time comes and investing in an stipulated fee.")
	return SendOk(l, c, m.String())
}

func (r Lenario) Create(l logrus.FieldLogger, c Context) State {
	if !party.IsPartyLeader(l)(c.CharacterId) {
		return r.PartyLeaderMustTalk(l, c)
	}
	if !character.GuildHasAlliance(l)(c.CharacterId) {
		return r.AlreadyHasAlliance(l, c)
	}
	return r.ConfirmCreationFee(l, c)
}

func (r Lenario) PartyLeaderMustTalk(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("If you want to form a guild union, please tell your party leader to talk to me. He/She will be assigned as the Leader of the Guild Union.")
	return SendOk(l, c, m.String())
}

func (r Lenario) AlreadyHasAlliance(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("You can not create a Guild Union while your guild is already registered in another.")
	return SendOk(l, c, m.String())
}

func (r Lenario) ConfirmCreationFee(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Oh, are you interested in forming a Guild Union? The current fee for this operation is ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 2000000)).
		BlackText().AddText(".")
	return SendYesNo(l, c, m.String(), r.ValidateCreation, Exit())
}

func (r Lenario) AddAnotherGuild(l logrus.FieldLogger, c Context) State {
	if !character.AllianceLeader(l)(c.CharacterId) {
		return r.MustBeLeaderToExpand(l, c)
	}
	return r.ConfirmExpandCost(l, c)
}

func (r Lenario) MustBeLeaderToExpand(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You can not expand a Guild Union if you don't own one.")
	return SendOk(l, c, m.String())
}

func (r Lenario) ConfirmExpandCost(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Do you want to increase your Alliance by ").
		RedText().AddText("one guild").
		BlackText().AddText(" slot? The fee for this procedure is ").
		BlueText().AddText(fmt.Sprintf("%d mesos", 1000000)).
		BlackText().AddText(".")
	return SendYesNo(l, c, m.String(), r.ValidateExpansion, Exit())
}

func (r Lenario) Disband(l logrus.FieldLogger, c Context) State {
	if !character.AllianceLeader(l)(c.CharacterId) {
		return r.MustBeLeaderToDisband(l, c)
	}
	if !character.GuildHasAlliance(l)(c.CharacterId) {
		return r.MustHaveAllianceToDisband(l, c)
	}
	return r.PerformDisband(l, c)
}

func (r Lenario) MustBeLeaderToDisband(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You can not disband a Guild Union if you don't own one.")
	return SendOk(l, c, m.String())
}

func (r Lenario) PerformDisband(l logrus.FieldLogger, c Context) State {
	//TODO implement
	return r.DisbandSuccess(l, c)
}

func (r Lenario) ValidateCreation(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 2000000) {
		return r.NotEnoughMesos(l, c)
	}
	return r.GetUnionName(l, c)
}

func (r Lenario) ValidateExpansion(l logrus.FieldLogger, c Context) State {
	if !character.HasMeso(l)(c.CharacterId, 1000000) {
		return r.NotEnoughMesos(l, c)
	}
	if !character.AllianceAtCapacity(l)(c.CharacterId) {
		return r.AllianceAtCapacity(l, c)
	}
	return r.ProcessExpansion(l, c)
}

func (r Lenario) NotEnoughMesos(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You don't have enough mesos for this request.")
	return SendOk(l, c, m.String())
}

func (r Lenario) GetUnionName(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Now please enter the name of your new Guild Union. (max. 12 letters)")
	return SendGetText(l, c, m.String(), r.ConfirmUnionName)
}

func (r Lenario) ConfirmUnionName(text string) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		m := message.NewBuilder().AddText("Will ").AddText(text).AddText(" be the name of your Guild Union?")
		return SendYesNo(l, c, m.String(), r.ValidateUnionName(text), r.GetUnionName)
	}
}

func (r Lenario) ValidateUnionName(text string) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		if !character.ValidAllianceName(l)(text) {
			return r.InvalidName(l, c)
		}
		return r.PerformCreate(text)(l, c)
	}
}

func (r Lenario) InvalidName(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("This name is unavailable, please choose another one.")
	return SendNext(l, c, m.String(), r.GetUnionName)
}

func (r Lenario) PerformCreate(text string) StateProducer {
	return func(l logrus.FieldLogger, c Context) State {
		err := character.GainMeso(l)(c.CharacterId, 2000000)
		if err != nil {
			l.WithError(err).Errorf("Unable to process payment from character %d.", c.CharacterId)
		}
		err = character.CreateAlliance(l)(c.CharacterId, text)
		if err != nil {
			l.WithError(err).Errorf("Unable to create alliance.")
			return r.CreateAllianceError(l, c)
		}
		return r.CreationSuccess(l, c)
	}
}

func (r Lenario) CreationSuccess(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You have successfully formed a Guild Union.")
	return SendOk(l, c, m.String())
}

func (r Lenario) CreateAllianceError(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().
		AddText("Please check if you and the other one guild leader in your party are both here on this room right now, and make sure both guilds are currently unregistered on unions. No other guild leaders should be present with you 2 on this process.")
	return SendOk(l, c, m.String())
}

func (r Lenario) ProcessExpansion(l logrus.FieldLogger, c Context) State {
	err := character.GainMeso(l)(c.CharacterId, -1000000)
	if err != nil {
		l.WithError(err).Errorf("Unable to process payment from character %d.", c.CharacterId)
	}
	err = character.ExpandAlliance(l)(c.CharacterId)
	if err != nil {
		l.WithError(err).Errorf("Unable to expand alliance.")
		return r.ExpansionFailure(l, c)
	}
	return r.ExpansionSuccess(l, c)
}

func (r Lenario) ExpansionSuccess(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Your alliance can now accept one more guild.")
	return SendOk(l, c, m.String())
}

func (r Lenario) AllianceAtCapacity(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Your alliance already reached the maximum capacity for guilds.")
	return SendOk(l, c, m.String())
}

func (r Lenario) ExpansionFailure(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Unable to expand alliance.")
	return SendOk(l, c, m.String())
}

func (r Lenario) MustHaveAllianceToDisband(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("You can not expand a Guild Union if you don't own one.")
	return SendOk(l, c, m.String())
}

func (r Lenario) DisbandSuccess(l logrus.FieldLogger, c Context) State {
	m := message.NewBuilder().AddText("Your Guild Union has been disbanded.")
	return SendOk(l, c, m.String())
}
