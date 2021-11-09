package npc

import (
	"atlas-ncs/json"
	"atlas-ncs/rest"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
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

func InitResource(router *mux.Router, l logrus.FieldLogger) {
	r := router.PathPrefix("/speak").Subrouter()
	r.HandleFunc("", registerSendSpeech(l)).Methods(http.MethodPost)
}

func registerSendSpeech(l logrus.FieldLogger) http.HandlerFunc {
	return rest.RetrieveSpan("send_speech", SendSpeech(l))
}

func SendSpeech(l logrus.FieldLogger) rest.SpanHandler {
	return func(span opentracing.Span) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			li := &SpeechInputDataContainer{}
			err := json.FromJSON(li, r.Body)
			if err != nil {
				l.WithError(err).Errorf("Deserializing input.")
				rw.WriteHeader(http.StatusBadRequest)
				return
			}
			attr := li.Data.Attributes
			SendSimple(l, span)(attr.CharacterId, attr.NPCId)(attr.Message)
			rw.WriteHeader(http.StatusNoContent)
		}
	}
}
