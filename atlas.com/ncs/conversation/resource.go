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
			l.WithError(err).Errorf("Script for npc %d is not implemented.", npcId)
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
		npcId, err := strconv.Atoi(vars["npcId"])
		if err != nil {
			l.WithError(err).Errorf("Error parsing npcId as uint32")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		characterId, err := strconv.Atoi(vars["characterId"])
		if err != nil {
			l.WithError(err).Errorf("Error parsing characterId as uint32")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		c, err := GetRegistry().GetPreviousContext(uint32(characterId))
		if err != nil {
			l.WithError(err).Errorf("Conversation with %d does not exist.", characterId)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if c.ctx.NPCId != uint32(npcId) {
			l.WithError(err).Errorf("Conversation between npc %d and character %d does not exist.", npcId, characterId)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}
