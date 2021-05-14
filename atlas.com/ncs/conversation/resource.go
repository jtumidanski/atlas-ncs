package conversation

import (
	"atlas-ncs/conversation/script"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetConversation(l logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		npcId, err := strconv.Atoi(vars["npcId"])
		if err != nil {
			l.WithError(err).Errorf("Error parsing characterId as uint32")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = script.GetRegistry().GetScript(uint32(npcId))
		if err != nil {
			l.WithError(err).Debugf("Script for npc %d is not implemented.", npcId)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}

func InConversation(l logrus.FieldLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		characterId, err := strconv.Atoi(vars["characterId"])
		if err != nil {
			l.WithError(err).Errorf("Error parsing characterId as uint32")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = GetRegistry().GetPreviousContext(uint32(characterId))
		if err != nil {
			l.WithError(err).Debugf("Conversation with %d does not exist.", characterId)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}
