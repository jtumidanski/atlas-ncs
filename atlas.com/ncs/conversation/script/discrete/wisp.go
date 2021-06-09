package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"github.com/sirupsen/logrus"
)

// Wisp is located in Ludibrium - Eos Tower Entrance (220000400)
type Wisp struct {
}

func (r Wisp) NPCId() uint32 {
	return npc.Wisp
}

func (r Wisp) Initial(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hello there, I'm ").
		BlueText().AddText("Mar the Fairy").
		BlackText().AddText(" of Victoria Island's main disciple. Mar the Fairy summoned me here to see if the pets are being taken care of here in Ludibrium. What can I do for you?").NewLine().
		OpenItem(0).BlueText().AddText("Who are you?").CloseItem().NewLine().
		OpenItem(1).BlueText().AddText("Tell me more about Pets.").CloseItem().NewLine().
		OpenItem(2).BlueText().AddText("How do I raise Pets?").CloseItem().NewLine().
		OpenItem(3).BlueText().AddText("Do Pets die too?").CloseItem().NewLine().
		OpenItem(4).BlueText().AddText("What are the commands for brown and black kitty?").CloseItem().NewLine().
		OpenItem(5).BlueText().AddText("What are the commands for brown puppy?").CloseItem().NewLine().
		OpenItem(6).BlueText().AddText("What are the commands for pink and white bunny?").CloseItem().NewLine().
		OpenItem(7).BlueText().AddText("What are the commands for Mini Cargo?").CloseItem().NewLine().
		OpenItem(8).BlueText().AddText("What are the commands for Husky?").CloseItem().NewLine().
		OpenItem(9).BlueText().AddText("What are the commands for Black Pig?").CloseItem().NewLine().
		OpenItem(10).BlueText().AddText("What are the commands for Panda").CloseItem().NewLine().
		OpenItem(11).BlueText().AddText("What are the commands for Dino Boy & Girl?").CloseItem().NewLine().
		OpenItem(12).BlueText().AddText("What are the commands for Rudolph?").CloseItem().NewLine().
		OpenItem(13).BlueText().AddText("What are the commands for Monkey?").CloseItem().NewLine().
		OpenItem(14).BlueText().AddText("What are the commands for Robot?").CloseItem().NewLine().
		OpenItem(15).BlueText().AddText("What are the commands for Elephant?").CloseItem().NewLine().
		OpenItem(16).BlueText().AddText("What are the commands for Golden Pig?").CloseItem().NewLine().
		OpenItem(17).BlueText().AddText("What are the commands for Penguin?").CloseItem().NewLine().
		OpenItem(18).BlueText().AddText("What are the commands for Mini Yeti?").CloseItem().NewLine().
		OpenItem(19).BlueText().AddText("What are the commands for Jr. Balrog?").CloseItem().NewLine().
		OpenItem(20).BlueText().AddText("What are the commands for Baby Dragon?").CloseItem().NewLine().
		OpenItem(21).BlueText().AddText("What are the commands for Green/Red/Blue Dragon?").CloseItem().NewLine().
		OpenItem(22).BlueText().AddText("What are the commands for Black Dragon?").CloseItem().NewLine().
		OpenItem(23).BlueText().AddText("What are the commands for Snowman?").CloseItem().NewLine().
		OpenItem(24).BlueText().AddText("What are the commands for Sun Wu Kong?").CloseItem().NewLine().
		OpenItem(25).BlueText().AddText("What are the commands for Jr. Reaper?").CloseItem().NewLine().
		OpenItem(26).BlueText().AddText("What are the commands for Crystal Rudolph?").CloseItem().NewLine().
		OpenItem(27).BlueText().AddText("What are the commands for Kino?").CloseItem().NewLine().
		OpenItem(28).BlueText().AddText("What are the commands for White Duck?").CloseItem().NewLine().
		OpenItem(29).BlueText().AddText("What are the commands for Pink Bean?").CloseItem().NewLine().
		OpenItem(30).BlueText().AddText("What are the commands for Porcupine?").CloseItem()
	return script.SendListSelection(l, c, m.String(), r.Selection)
}

func (r Wisp) Selection(selection int32) script.StateProducer {
	switch selection {
	case 0:
		return r.WhoAreYou
	case 1:
		return r.MoreAboutPets
	case 2:
		return r.HowToRaise
	case 3:
		return r.DoPetsDie
	case 4:
		return r.Kitty
	case 5:
		return r.Puppy
	case 6:
		return r.Bunny
	case 7:
		return r.MiniCargo
	case 8:
		return r.Husky
	case 9:
		return r.BlackPig
	case 10:
		return r.Panda
	case 11:
		return r.Dino
	case 12:
		return r.Rudolph
	case 13:
		return r.Monkey
	case 14:
		return r.Robot
	case 15:
		return r.Elephant
	case 16:
		return r.GoldenPig
	case 17:
		return r.Penguin
	case 18:
		return r.MiniYeti
	case 19:
		return r.JrBalrog
	case 20:
		return r.BabyDragon
	case 21:
		return r.GreenRedBlueDragon
	case 22:
		return r.BlackDragon
	case 23:
		return r.Snowman
	case 24:
		return r.SunWuKong
	case 25:
		return r.JrReaper
	case 26:
		return r.CrystalRudolph
	case 27:
		return r.Kino
	case 28:
		return r.WhiteDuck
	case 29:
		return r.PinkBean
	case 30:
		return r.Porcupine
	}
	return nil
}

func (r Wisp) WhoAreYou(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("I'm Wisp, continuing on with the studies that my Master Mar the Fairy assigned me. There seems to be a lot of pets even here in Ludibrium. I need to get back to my studies, so if you'll excuse me...")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) MoreAboutPets(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Hmmmm,you must have a lot of questions regarding the pets. Long ago, a person by the name ").
		BlueText().AddText("Cloy").
		BlackText().AddText(", sprayed Water of Life on it, and cast spell on it to create a magical animal. I know it sounds unbelievable, but it's a doll that became an actual living thing. They understand and follow people very well.")
	return script.SendNext(l, c, m.String(), r.But)
}

func (r Wisp) HowToRaise(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Depending on the command you give, pets can love it, hate, and display other kinds of reactions to it. If you give the pet a command and it follows you well, your closeness goes up. Double click on the pet and you can check the closeness, level, fullness and etc...")
	return script.SendNext(l, c, m.String(), r.Talk)
}

func (r Wisp) DoPetsDie(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Dying... well, they aren't technically ALIVE per se, so I don't know if dying is the right term to use. They are dolls with my magical power and the power of Water of Life to become a live object. Of course while it's alive, it's just like a live animal...")
	return script.SendNext(l, c, m.String(), r.AfterSomeTime)
}

func (r Wisp) Kitty(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rBrown Kitty and Black Kitty#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no, badgirl, badboy#k (level 1 ~ 30)").NewLine().
		AddText("#bstupid, ihateyou, dummy#k (level 1 ~ 30)").NewLine().
		AddText("#biloveyou#k (level 1~30)").NewLine().
		AddText("#bpoop#k (level 1 ~ 30)").NewLine().
		AddText("#btalk, say, chat#k (level 10 ~ 30)").NewLine().
		AddText("#bcutie#k (level 10 ~ 30)").NewLine().
		AddText("#bup, stand, rise#k (level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) Puppy(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rBrown Puppy#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no, badgirl, badboy#k (level 1 ~ 30)").NewLine().
		AddText("#bstupid, ihateyou, baddog, dummy#k (level 1 ~ 30)").NewLine().
		AddText("#biloveyou#k (level 1~30)").NewLine().
		AddText("#bpee#k (level 10 ~ 30)").NewLine().
		AddText("#btalk, say, chat, bark#k (level 10 ~ 30)").NewLine().
		AddText("#bdown#k (level 10 ~ 30)").NewLine().
		AddText("#bup, stand, rise#k (level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) Bunny(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rPink Bunny and White Bunny#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no, badgirl, badboy#k (level 1 ~ 30)").NewLine().
		AddText("#bup, stand#k (level 1 ~ 30)").NewLine().
		AddText("#biloveyou#k (level 1~30)").NewLine().
		AddText("#bpoop#k (level 1 ~ 30)").NewLine().
		AddText("#btalk, say, chat#k (level 10 ~ 30)").NewLine().
		AddText("#bhug#k (level 10 ~ 30)").NewLine().
		AddText("#bsleep, sleepy, gotobed#k (level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) MiniCargo(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rMini Cargo#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no, badgirl, badboy#k (level 1 ~ 30)").NewLine().
		AddText("#bup, stand#k (level 1 ~ 30)").NewLine().
		AddText("#biloveyou#k (level 1~30)").NewLine().
		AddText("#bpee#k (level 1 ~ 30)").NewLine().
		AddText("#btalk, say, chat#k (level 10 ~ 30)").NewLine().
		AddText("#bthelook, charisma#k (level 10 ~ 30)").NewLine().
		AddText("#bgoodboy, good#k (level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) Husky(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rHusky#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no, badgirl, badboy#k (level 1 ~ 30)").NewLine().
		AddText("#bstupid, ihateyou, baddog, dummy#k (level 1 ~ 30)").NewLine().
		AddText("#biloveyou#k (level 1 ~ 30)").NewLine().
		AddText("#bpee#k (level 1 ~ 30)").NewLine().
		AddText("#btalk, say, chat, bark#k (level 10 ~ 30)").NewLine().
		AddText("#bdown#k (level 10 ~ 30)").NewLine().
		AddText("#bup, stand, rise#k (level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) BlackPig(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rBlack Pig#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no, badgirl, badboy#k (level 1 ~ 30)").NewLine().
		AddText("#bpoop#k (level 1 ~ 30)").NewLine().
		AddText("#biloveyou#k (level 1~30)").NewLine().
		AddText("#bpoop#k (level 1 ~ 30)").NewLine().
		AddText("#bhand, up, stand#k (level 1 ~ 30)").NewLine().
		AddText("#btalk, say, chat, hug#k (level 10 ~ 30)").NewLine().
		AddText("#bsmile#k (level 10 ~ 30)").NewLine().
		AddText("#blaugh, smile#k (level 10 ~ 30)").NewLine().
		AddText("#bcharisma, sleep, sleepy, gotobed#k(level 20~30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) Panda(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rPanda#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no, badgirl, badboy#k (level 1 ~ 30)").NewLine().
		AddText("#biloveyou#k (level 1 ~ 30)").NewLine().
		AddText("#bpee#k(level 1 ~ 30)").NewLine().
		AddText("#bup, stand, hug#k (level 1 ~ 30)").NewLine().
		AddText("#btalk, chat#k (level 10 ~ 30)").NewLine().
		AddText("#bplay#k (level 20 ~ 30)").NewLine().
		AddText("#bmeh, bleh#k (level 10 ~ 30)").NewLine().
		AddText("#bsleep, sleepy, gotobed#k (level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) Dino(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rDino Boy and Dino Girl#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no,, stupid, ihateyou, badboy, badgirl#k (evel 1 ~ 30)").NewLine().
		AddText("#biloveyou, dummy#k (level 1 ~ 30)").NewLine().
		AddText("#bpoop#k (level 1 ~ 30)").NewLine().
		AddText("#btalk, chat(level 10 ~ 30)").NewLine().
		AddText("#bsmile, laugh#k (level 1 ~ 30)").NewLine().
		AddText("#bcutie#k (level 10 ~ 30)").NewLine().
		AddText("#bsleep, nap, sleepy#k (level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) Rudolph(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rRudolph#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k(level 1 ~30) ").NewLine().
		AddText("#bbad, no, badgirl, badboy#k(level 1~30)").NewLine().
		AddText("#bup, stand#k(level 1 ~ 30) ").NewLine().
		AddText("#bstupid, ihateyou, dummy#k(level 1 ~ 30) ").NewLine().
		AddText("#bmerryxmas, merrychristmas#k(level 11 ~ 30)").NewLine().
		AddText("#biloveyou#k(level 1 ~ 30)").NewLine().
		AddText("#bpoop#k(level 1 ~ 30)").NewLine().
		AddText("#btalk, say, chat#k(level 11 ~ 30)").NewLine().
		AddText("#blonely, alone, down, rednose#k(level 11~30),").NewLine().
		AddText("#bcutie#k(level 11 ~ 30)").NewLine().
		AddText("#bmush, go#k(level 21 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) Monkey(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rMonkey#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit, rest#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no, badboy, badgirl#k (level 1 ~ 30)").NewLine().
		AddText("#bup, stand#k(level 1 ~ 30)").NewLine().
		AddText("#biloveyou, pee#k (level 1 ~ 30)").NewLine().
		AddText("#btalk, say, chat#k (level 11 ~ 30)").NewLine().
		AddText("#bplay, melong#k (level 11 ~ 30)").NewLine().
		AddText("#bsleep, sleepy, gotobed#k (level 21 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) Robot(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rRobot#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit, stand, rise#k (level 1 ~ 30)").NewLine().
		AddText("#battack, bad, no, badboy#k (level 1 ~ 30)").NewLine().
		AddText("#bstupid, ihateyou, dummy#k (level 1 ~ 30)").NewLine().
		AddText("#biloveyou, good#k (level 1 ~ 30)").NewLine().
		AddText("#bspeak, disguise#k (level 11 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) Elephant(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rElephant#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit, rest#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no, badboy, badgirl#k (level 1 ~ 30)").NewLine().
		AddText("#bup, stand, rise#k(level 1 ~ 30)").NewLine().
		AddText("#biloveyou, pee#k (level 1 ~ 30)").NewLine().
		AddText("#btalk, say, chat, play#k (level 11 ~ 30)").NewLine().
		AddText("#bsleep, sleepy, gotobed#k (level 21 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) GoldenPig(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rGolden Pig#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no, badboy, badgirl#k (level 1 ~ 30)").NewLine().
		AddText("#bpoop, iloveyou#k (level 1 ~ 30)").NewLine().
		AddText("#btalk, say, chat#k (level 11 ~ 30)").NewLine().
		AddText("#bloveme, hugme#k (level 11 ~ 30)").NewLine().
		AddText("#bsleep, sleepy, gotobed#k (level 21 ~ 30)").NewLine().
		AddText("#bimpressed, outofhere#k (level 21 ~ 30)").NewLine().
		AddText("#broll, showmethemoney#k (level 21 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) Penguin(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rPenguin#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no, badboy, badgirl#k (level 1 ~ 30)").NewLine().
		AddText("#bpoop#k (level 1 ~ 30)").NewLine().
		AddText("#bup, stand, rise#k (level 1 ~ 30)").NewLine().
		AddText("#biloveyou#k (level 1 ~ 30)").NewLine().
		AddText("#btalk, chat, say#k (level 10 ~ 30)").NewLine().
		AddText("#bhug, hugme#k (level 10 ~ 30)").NewLine().
		AddText("#bwing, hand#k (level 10 ~ 30)").NewLine().
		AddText("#bsleep#k (level 20 ~ 30)").NewLine().
		AddText("#bkiss, smooch, muah#k (level 20 ~ 30)").NewLine().
		AddText("#bfly#k (level 20 ~ 30)").NewLine().
		AddText("#bcute, adorable#k (level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) MiniYeti(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rMini Yeti#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no, badboy, badgirl#k (level 1 ~ 30)").NewLine().
		AddText("#bpoop#k (level 1 ~ 30)").NewLine().
		AddText("#bdance, boogie, shakeit#k (level 1 ~ 30)").NewLine().
		AddText("#bcute, cutie, pretty, adorable#k (level 1 ~ 30)").NewLine().
		AddText("#biloveyou, likeyou, mylove#k (level 1 ~ 30)").NewLine().
		AddText("#btalk, chat, say#k (level 10 ~ 30)").NewLine().
		AddText("#bsleep, nap#k (level 10 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) JrBalrog(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rJr. Balrog#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bliedown#k (level 1 ~ 30)").NewLine().
		AddText("#bno|bad|badgirl|badboy#k (level 1 ~ 30)").NewLine().
		AddText("#biloveyou|mylove|likeyou#k (level 1 ~ 30)").NewLine().
		AddText("#bcute|cutie|pretty|adorable#k (level 1 ~ 30)").NewLine().
		AddText("#bpoop#k (level 1 ~ 30)").NewLine().
		AddText("#bsmirk|crooked|laugh#k (level 1 ~ 30)").NewLine().
		AddText("#bmelong#k (level 11 ~ 30)").NewLine().
		AddText("#bgood|thelook|charisma#k (level 11 ~ 30)").NewLine().
		AddText("#bspeak|talk|chat|say#k (level 11 ~ 30)").NewLine().
		AddText("#bsleep|nap|sleepy#k (level 11 ~ 30)").NewLine().
		AddText("#bgas#k (level 21 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) BabyDragon(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rBaby Dragon#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bno|bad|badgirl|badboy#k (level 1 ~ 30)").NewLine().
		AddText("#biloveyou|loveyou#k (level 1 ~ 30)").NewLine().
		AddText("#bpoop#k (level 1 ~ 30)").NewLine().
		AddText("#bstupid|ihateyou|dummy#k (level 1 ~ 30)").NewLine().
		AddText("#bcutie#k (level 11 ~ 30)").NewLine().
		AddText("#btalk|chat|say#k (level 11 ~ 30)").NewLine().
		AddText("#bsleep|sleepy|gotobed#k (level 11 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) GreenRedBlueDragon(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rGreen/Red/Blue Dragon#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 15 ~ 30)").NewLine().
		AddText("#bno|bad|badgirl|badboy#k (level 15 ~ 30)").NewLine().
		AddText("#biloveyou|loveyou#k (level 15 ~ 30)").NewLine().
		AddText("#bpoop#k (level 15 ~ 30)").NewLine().
		AddText("#bstupid|ihateyou|dummy#k (level 15 ~ 30)").NewLine().
		AddText("#btalk|chat|say#k (level 15 ~ 30)").NewLine().
		AddText("#bsleep|sleepy|gotobed#k (level 15 ~ 30)").NewLine().
		AddText("#bchange#k (level 21 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) BlackDragon(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rBlack Dragon#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 15 ~ 30)").NewLine().
		AddText("#bno|bad|badgirl|badboy#k (level 15 ~ 30)").NewLine().
		AddText("#biloveyou|loveyou#k (level 15 ~ 30)").NewLine().
		AddText("#bpoop#k (level 15 ~ 30)").NewLine().
		AddText("#bstupid|ihateyou|dummy#k (level 15 ~ 30)").NewLine().
		AddText("#btalk|chat|say#k (level 15 ~ 30)").NewLine().
		AddText("#bsleep|sleepy|gotobed#k (level 15 ~ 30)").NewLine().
		AddText("#bcutie, change#k (level 21 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) Snowman(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rSnowman#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bstupid, ihateyou, dummy#k (level 1 ~ 30)").NewLine().
		AddText("#bloveyou, mylove, ilikeyou#k (level 1 ~ 30)").NewLine().
		AddText("#bmerrychristmas#k (level 1 ~ 30)").NewLine().
		AddText("#bcutie, adorable, cute, pretty#k (level 1 ~ 30)").NewLine().
		AddText("#bbad, no, badgirl, badboy#k (level 1 ~ 30)").NewLine().
		AddText("#btalk, chat, say/sleep, sleepy, gotobed#k (level 10 ~ 30)").NewLine().
		AddText("#bchang#k (level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) SunWuKong(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rSun Wu Kong#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k(level 1 ~ 30)").NewLine().
		AddText("#bno,bad,badgirl,badboy#k(level 1 ~ 30) ").NewLine().
		AddText("#bpoope#k(level 1 ~ 30) ").NewLine().
		AddText("#bcutie,adorable,cute,pretty#k(level 1 ~ 30) ").NewLine().
		AddText("#biloveyou,loveyou,luvyou,ilikeyou,mylove#k(level 1 ~ 30) ").NewLine().
		AddText("#btalk,chat,say/sleep,sleepy,gotobed#k(level 10 ~ 30) ").NewLine().
		AddText("#btransform#k(level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) JrReaper(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rJr. Reaper#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bno|bad|badgirl|badboy#k (level 1 ~ 30)").NewLine().
		AddText("#bplaydead, poop#k (level 1 ~ 30)").NewLine().
		AddText("#btalk|chat|say#k (level 1 ~ 30)").NewLine().
		AddText("#biloveyou, hug#k (level 1 ~ 30)").NewLine().
		AddText("#bsmellmyfeet, rockout, boo#k (level 1 ~ 30)").NewLine().
		AddText("#btrickortreat#k (level 1 ~ 30)").NewLine().
		AddText("#bmonstermash#k (level 1 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) CrystalRudolph(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rCrystal Rudolph#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bno|badgirl|badboy#k (level 1 ~ 30)").NewLine().
		AddText("#bbleh|joke#k(level 1~30)").NewLine().
		AddText("#bdisguise|transform#k(level 1 ~ 30) ").NewLine().
		AddText("#bawesome|feelgood|lalala#k(level 1 ~ 30) ").NewLine().
		AddText("#bloveyou|heybabe#k(level 1 ~ 30) ").NewLine().
		AddText("#btalk|say|chat#k(level 10 ~ 30) ").NewLine().
		AddText("#bsleep|sleepy|nap|gotobed#k(level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) Kino(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rKino#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bbad|no|badgirl|badboy#k (level 1 ~ 30)").NewLine().
		AddText("#bpoop#k (level 1 ~ 30)").NewLine().
		AddText("#bsleep|nap|sleepy|gotobed#k(level 1 ~ 30) ").NewLine().
		AddText("#btalk|say|chat#k(level 10 ~ 30) ").NewLine().
		AddText("#biloveyou|mylove|likeyou#k(level 10 ~ 30) ").NewLine().
		AddText("#bmeh|bleh#k(level 10 ~ 30) ").NewLine().
		AddText("#bdisguise|change|transform#k(level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) WhiteDuck(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rWhite Duck#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k(level 1 ~ 30) ").NewLine().
		AddText("#bbad|no|badgirl|badboy#k(level 1 ~ 30) ").NewLine().
		AddText("#bup|stand#k(level 1 ~ 30) ").NewLine().
		AddText("#bpoop#k(level 1 ~ 30) ").NewLine().
		AddText("#btalk|chat|say#k(level 1 ~ 30) ").NewLine().
		AddText("#bhug#k(level 1 ~ 30) ").NewLine().
		AddText("#bloveyou#k(level 1 ~ 30) ").NewLine().
		AddText("#bcutie#k(level 1 ~ 30) ").NewLine().
		AddText("#bsleep#k(level 1 ~ 30) ").NewLine().
		AddText("#bsmarty(level 10 ~ 30) ").NewLine().
		AddText("#bdance#k (level 20 ~ 30) ").NewLine().
		AddText("#bswan#k(level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) PinkBean(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rPink Bean#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k(level 1 ~ 30) ").NewLine().
		AddText("#bbad|no|badgirl|badboy|poop#k(level 1 ~ 30) ").NewLine().
		AddText("#blazy|dummy|ihateyoutalk|chat|say|mumbleiloveyou|hugme|loveyou|#k(level 1 ~ 30) ").NewLine().
		AddText("#bshake|music|charmbleh|joke|boo#k(level 20 ~ 30) ").NewLine().
		AddText("#bgotobed|sleep|sleepypoke|stinky|dummy|ihateyou#k(level 20 ~ 30)").NewLine().
		AddText("#bkongkong#k(level 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) Porcupine(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("These are the commands for #rPorcupine#k. The level mentioned next to the command shows the pet level required for it to respond.").NewLine().
		AddText("#bsit#k (level 1 ~ 30)").NewLine().
		AddText("#bno|bad|badgirl|badboy#k (level 1 ~ 30)").NewLine().
		AddText("#bhugcushion|sleep|knit|poop#k (level 1 ~ 30)").NewLine().
		AddText("#bcomb|beach#k (level 10 ~ 30)").NewLine().
		AddText("#btreeninja|dart#k (level 20 ~ 30)")
	return script.SendOk(l, c, m.String())
}

func (r Wisp) But(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("But Water of Life only comes out little at the very bottom of the World Tree, so those babies can't be alive forever... I know, it's very unfortunate... but even if it becomes a doll again they can be brought back to life so be good to it while you're with it.")
	return script.SendNextPrevious(l, c, m.String(), r.SpecialCommands, r.MoreAboutPets)
}

func (r Wisp) SpecialCommands(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().AddText("Oh yeah, they'll react when you give them special commands. You can scold them, love them.. it all depends on how you take care of them. They are afraid to leave their masters so be nice to them, show them love. They can get sad and lonely fast..")
	return script.SendPrevious(l, c, m.String(), r.But)
}

func (r Wisp) Talk(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Talk to the pet, pay attention to it and its closeness level will go up and eventually his overall level will go up too. As the closeness rises, the pet's overall level will rise soon after. As the overall level rises, one day the pet may even talk like a person a little bit, so try hard raising it. Of course it won't be easy doing so...")
	return script.SendNextPrevious(l, c, m.String(), r.HowToRaise, r.Hungry)
}

func (r Wisp) Hungry(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("It may be a live doll but they also have life so they can feel the hunger too. ").
		BlueText().AddText("Fullness").
		BlackText().AddText(" shows the level of hunger the pet's in. 100 is the max, and the lower it gets, it means that the pet is getting hungrier. After a while, it won't even follow your command and be on the offensive, so watch out over that.")
	return script.SendNextPrevious(l, c, m.String(), r.NotNormalFood, r.Talk)
}

func (r Wisp) NotNormalFood(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("That's right! Pets can't eat the normal human food. Instead a teddy bear in Ludibrium called ").
		BlueText().AddText("Patricia").
		BlackText().AddText(" sells ").
		BlueText().AddText("Pet Food").
		BlackText().AddText(" so if you need food for your pet, find ").
		BlueText().AddText("Patricia").
		BlackText().AddText(" It'll be a good idea to buy the food in advance and feed the pet before it gets really hungry.")
	return script.SendNextPrevious(l, c, m.String(), r.GoesHome, r.Hungry)
}

func (r Wisp) GoesHome(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Oh, and if you don't feed the pet for a long period of time, it goes back home by itself. You can take it out of its home and feed it but it's not really good for the pet's health, so try feeding him on a regular basis so it doesn't go down to that level, alright? I think this will do.")
	return script.SendPrevious(l, c, m.String(), r.NotNormalFood)
}

func (r Wisp) AfterSomeTime(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("After some time... that's correct, they stop moving. They just turn back to being a doll, after the effect of magic dies down and Water of Life dries out. But that doesn't mean it's stopped forever, because once you pour Water of Life over, it's going to be back alive.")
	return script.SendNextPrevious(l, c, m.String(), r.Sad, r.GoesHome)
}

func (r Wisp) Sad(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		AddText("Even if it someday moves again, it's sad to see them stop altogether. Please be nice to them while they are alive and moving. Feed them well, too. Isn't it nice to know that there's something alive that follows and listens to only you?")
	return script.SendPrevious(l, c, m.String(), r.AfterSomeTime)
}
