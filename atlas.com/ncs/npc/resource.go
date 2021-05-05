package npc

import (
	"atlas-ncs/rest/attributes"
	"github.com/sirupsen/logrus"
	"net/http"
)

type SpeechInputDataContainer struct {
	Data SpeechData `json:"data"`
}

type SpeechData struct {
	Id         string           `json:"id"`
	Type       string           `json:"type"`
	Attributes SpeechAttributes `json:"attributes"`
}

type SpeechAttributes struct {
	CharacterId uint32 `json:"characterId"`
	NPCId       uint32 `json:"npcId"`
	Message     string `json:"message"`
}

func SendSpeech(l logrus.FieldLogger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		li := &SpeechInputDataContainer{}
		err := attributes.FromJSON(li, r.Body)
		if err != nil {
			l.WithError(err).Errorf("Deserializing input.")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		attr := li.Data.Attributes
		err = Processor(l).Conversation(attr.CharacterId, attr.NPCId).SendSimple(attr.Message)
		if err != nil {
			l.WithError(err).Errorf("Error sending simple message to %d on behalf of %d.", attr.CharacterId, attr.NPCId)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusNoContent)
	}
}
