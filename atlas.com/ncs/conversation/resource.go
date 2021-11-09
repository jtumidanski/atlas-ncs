package conversation

import (
	registry2 "atlas-ncs/conversation/script/registry"
	"atlas-ncs/rest"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

const (
	getConversation = "get_conversation"
	inConversation  = "in_conversation"
)

func InitResource(router *mux.Router, l logrus.FieldLogger) {
	router.HandleFunc("/script/{npcId}", registerGetConversation(l)).Methods(http.MethodGet)
	router.HandleFunc("/conversation/{characterId}", registerInConversation(l)).Methods(http.MethodGet)
}

type npcIdHandler func(npcId uint32) http.HandlerFunc

func parseNpcId(l logrus.FieldLogger, next npcIdHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		monsterId, err := strconv.Atoi(vars["npcId"])
		if err != nil {
			l.WithError(err).Errorf("Error parsing npcId as uint32")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next(uint32(monsterId))(w, r)
	}
}

type characterIdHandler func(characterId uint32) http.HandlerFunc

func parseCharacterId(l logrus.FieldLogger, next characterIdHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		monsterId, err := strconv.Atoi(vars["characterId"])
		if err != nil {
			l.WithError(err).Errorf("Error parsing characterId as uint32")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		next(uint32(monsterId))(w, r)
	}
}

func registerGetConversation(l logrus.FieldLogger) http.HandlerFunc {
	return rest.RetrieveSpan(getConversation, func(span opentracing.Span) http.HandlerFunc {
		return parseNpcId(l, func(npcId uint32) http.HandlerFunc {
			return handleGetConversation(l)(span)(npcId)
		})
	})
}

func registerInConversation(l logrus.FieldLogger) http.HandlerFunc {
	return rest.RetrieveSpan(inConversation, func(span opentracing.Span) http.HandlerFunc {
		return rest.RetrieveSpan(inConversation, func(span opentracing.Span) http.HandlerFunc {
			return parseCharacterId(l, func(characterId uint32) http.HandlerFunc {
				return handleInConversation(l)(span)(characterId)
			})
		})
	})
}

func handleGetConversation(l logrus.FieldLogger) func(span opentracing.Span) func(npcId uint32) http.HandlerFunc {
	return func(span opentracing.Span) func(npcId uint32) http.HandlerFunc {
		return func(npcId uint32) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				_, err := registry2.GetRegistry().GetScript(npcId)
				if err != nil {
					l.WithError(err).Debugf("Script for npc %d is not implemented.", npcId)
					w.WriteHeader(http.StatusNotFound)
					return
				}
				w.WriteHeader(http.StatusOK)
				return
			}
		}
	}
}

func handleInConversation(l logrus.FieldLogger) func(span opentracing.Span) func(characterId uint32) http.HandlerFunc {
	return func(span opentracing.Span) func(characterId uint32) http.HandlerFunc {
		return func(characterId uint32) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				_, err := GetRegistry().GetPreviousContext(characterId)
				if err != nil {
					l.WithError(err).Debugf("Conversation with %d does not exist.", characterId)
					w.WriteHeader(http.StatusNotFound)
					return
				}
				w.WriteHeader(http.StatusOK)
				return
			}
		}
	}
}
