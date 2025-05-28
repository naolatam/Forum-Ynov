package hostedservices

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func StartAllHostedServices() {
	// Start the hosted services
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()

	go startSessionCleanerHostedService(ctx)

	<-ctx.Done()
	log.Println("Shutting down all hosted services...")
}
