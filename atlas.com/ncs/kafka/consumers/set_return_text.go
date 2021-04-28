package consumers

import "github.com/sirupsen/logrus"

type setReturnTextCommand struct {
	CharacterId uint32 `json:"characterId"`
	Text        string `json:"text"`
}

func SetReturnTextCommandCreator() EmptyEventCreator {
	return func() interface{} {
		return &setReturnTextCommand{}
	}
}

func HandleSetReturnTextCommand() EventProcessor {
	return func(l logrus.FieldLogger, e interface{}) {
		if event, ok := e.(*setReturnTextCommand); ok {
			l.Debugf("Handling SetReturnText command for character %d, with value %s.", event.CharacterId, event.Text)
		} else {
			l.Errorf("Unable to cast event provided to handler")
		}
	}
}