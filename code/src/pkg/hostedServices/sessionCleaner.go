package hostedservices

import (
	"context"
	"log"
	"time"

	"Forum-back/internal/config"
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
	if err != nil {
		panic(err)
	}

	row, err := db.QueryContext(ctx, "DELETE FROM sessions WHERE expireAt < NOW()")
	if err != nil {
		log.Printf("[HostedService] Error cleaning sessions: %v", err)
		return
	}
	row.Close()
	db.Close()
	log.Println("[HostedService] Session cleaner service completed successfully")
}
