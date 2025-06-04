package hostedservices

import (
	"context"
	"log"
	"time"

	"Forum-back/internal/config"
	"Forum-back/pkg/services"
)

func startSessionCleanerHostedService(ctx context.Context) {
	log.Println("[HostedService] Session cleaner service started")
	ticker := time.NewTicker(3 * time.Hour) // Service will run every 3 hours
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			log.Println("[HostedService] Running session cleaner service...")
			cleanSession(ctx)

		case <-ctx.Done():
			log.Println("[HostedService] Session cleaner service stopped")
			ticker.Stop()
			return
		}
	}

}

func cleanSession(ctx context.Context) {

	db, err := config.OpenDBConnection()
	sessionService := services.NewSessionService(db)
	if err != nil {
		panic(err)
	}
	sessionService.DeleteExpiredSessions(time.Now())

	log.Println("[HostedService] Session cleaner service completed successfully")
}
