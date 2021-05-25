package rest

import (
	"atlas-ncs/conversation"
	"atlas-ncs/npc"
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

func CreateRestService(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup) {
	go NewServer(l, ctx, wg, ProduceRoutes)
}

func ProduceRoutes(l logrus.FieldLogger) http.Handler {
	router := mux.NewRouter().PathPrefix("/ms/ncs").Subrouter().StrictSlash(true)
	router.Use(CommonHeader)

	router.HandleFunc("/script/{npcId}", conversation.GetConversation(l)).Methods(http.MethodGet)
	router.HandleFunc("/conversation/{characterId}", conversation.InConversation(l)).Methods(http.MethodGet)

	r := router.PathPrefix("/speak").Subrouter()
	r.HandleFunc("", npc.SendSpeech(l)).Methods(http.MethodPost)

	return router
}
