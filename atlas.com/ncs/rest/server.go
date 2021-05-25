package rest

import (
	"atlas-ncs/conversation"
	"atlas-ncs/npc"
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"sync"
	"time"
)

type ConfigFunc func(config *Config)

type Config struct {
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
	addr         string
}

func NewServer(cl *logrus.Logger, ctx context.Context, wg *sync.WaitGroup, configurators ...ConfigFunc) {
	l := cl.WithFields(logrus.Fields{"originator": "HTTPServer"})
	w := cl.Writer()
	defer func() {
		err := w.Close()
		if err != nil {
			l.WithError(err).Errorf("Closing log writer.")
		}
	}()

	config := &Config{
		readTimeout:  time.Duration(5) * time.Second,
		writeTimeout: time.Duration(10) * time.Second,
		idleTimeout:  time.Duration(120) * time.Second,
		addr:         ":8080",
	}

	for _, configurator := range configurators {
		configurator(config)
	}

	router := mux.NewRouter().PathPrefix("/ms/ncs").Subrouter().StrictSlash(true)
	router.Use(commonHeader)

	router.HandleFunc("/script/{npcId}", conversation.GetConversation(l)).Methods(http.MethodGet)
	router.HandleFunc("/conversation/{characterId}", conversation.InConversation(l)).Methods(http.MethodGet)

	r := router.PathPrefix("/speak").Subrouter()
	r.HandleFunc("", npc.SendSpeech(l)).Methods(http.MethodPost)

	hs := http.Server{
		Addr:         config.addr,
		Handler:      router,
		ErrorLog:     log.New(w, "", 0),
		ReadTimeout:  config.readTimeout,
		WriteTimeout: config.writeTimeout,
		IdleTimeout:  config.idleTimeout,
	}

	l.Infoln("Starting server on port 8080")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		wg.Add(1)
		defer wg.Done()
		err := hs.ListenAndServe()
		if err != http.ErrServerClosed {
			l.WithError(err).Errorf("Error while serving.")
			return
		}
	}()

	<-ctx.Done()
	l.Infof("Shutting down server on port 8080")
	err := hs.Close()
	if err != nil {
		l.WithError(err).Errorf("Error shutting down HTTP service.")
	}
}

func commonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
