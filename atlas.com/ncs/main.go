package main

import (
	"atlas-ncs/conversation"
	"atlas-ncs/kafka"
	"atlas-ncs/logger"
	"atlas-ncs/npc"
	"atlas-ncs/rest"
	"atlas-ncs/tracing"
	"context"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const serviceName = "atlas-ncs"
const consumerGroupId = "NPC Conversation Service"

func main() {
	l := logger.CreateLogger(serviceName)
	l.Infoln("Starting main service.")

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	tc, err := tracing.InitTracer(l)(serviceName)
	if err != nil {
		l.WithError(err).Fatal("Unable to initialize tracer.")
	}
	defer func(tc io.Closer) {
		err := tc.Close()
		if err != nil {
			l.WithError(err).Errorf("Unable to close tracer.")
		}
	}(tc)

	kafka.CreateConsumers(l, ctx, wg,
		conversation.StartConsumer(consumerGroupId),
		conversation.ContinueConsumer(consumerGroupId),
		conversation.SetReturnTextConsumer(consumerGroupId))

	rest.CreateService(l, ctx, wg, "/ms/ncs", conversation.InitResource, npc.InitResource)

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.Infof("Initiating shutdown with signal %s.", sig)
	cancel()
	wg.Wait()
	l.Infoln("Service shutdown.")
}
