package discrete

import (
	"atlas-ncs/conversation/script"
	"atlas-ncs/npc"
	"atlas-ncs/npc/message"
	"atlas-ncs/quest"
	"fmt"
	"github.com/sirupsen/logrus"
)

// WizTheLibrarian is located in Ludibrium - Helios Tower <Library> (222020000)
type WizTheLibrarian struct {
}

func (r WizTheLibrarian) NPCId() uint32 {
	return npc.WizTheLibrarian
}

func (r WizTheLibrarian) Initial(l logrus.FieldLogger, c script.Context) script.State {
	questIds := []uint32{3615, 3616, 3617, 3618, 3630, 3633, 3639, 3920}
	questItems := []uint32{4031235, 4031236, 4031237, 4031238, 4031270, 4031280, 4031298, 4031591}

	cl := message.NewBuilder()
	counter := 0
	for i, questId := range questIds {
		if quest.IsCompleted(l)(c.CharacterId, questId) {
			counter++
			cl = cl.NewLine().ShowItemImage2(questItems[i]).AddText(" ").BlueText().ShowItemName1(questItems[i])
		}
	}
	if counter == 0 {
		return r.NoStoryBook(l, c)
	}
	return r.Progress(counter, cl.String())(l, c)
}

func (r WizTheLibrarian) NoStoryBook(l logrus.FieldLogger, c script.Context) script.State {
	m := message.NewBuilder().
		BlueText().ShowCharacterName().AddText(" ").
		BlackText().AddText("has not returned a single storybook yet.")
	return script.SendOk(l, c, m.String())
}

func (r WizTheLibrarian) Progress(counter int, books string) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("Let's see.. ").
			BlueText().ShowCharacterName().AddText(" ").
			BlackText().AddText(" has returned a total of ").
			BlueText().AddText(fmt.Sprintf("%d", counter)).
			BlackText().AddText(" books. The list of returned books is as follows:").
			AddText(books)
		return script.SendNext(l, c, m.String(), r.SettlingDown(counter, books))
	}
}

func (r WizTheLibrarian) SettlingDown(counter int, books string) script.StateProducer {
	return func(l logrus.FieldLogger, c script.Context) script.State {
		m := message.NewBuilder().
			AddText("The library is settling down now thanks chiefly to you, ").
			BlueText().ShowCharacterName().
			BlackText().AddText("'s immense help. If the story gets mixed up once again, then I'll be counting on you to fix it once more.")
		return script.SendNextPrevious(l, c, m.String(), script.Exit(), r.Progress(counter, books))
	}
}
