//testing
package main

import (
	"github.com/pocket-id/pocket-id/backend/internal/service"
	"log"
	"github.com/pocket-id/pocket-id/backend/testing/db"
	"github.com/pocket-id/pocket-id/backend/testing/appconfig"
)

func main() {
	log.Println("program starting")
	dbService, err := db.NewDBService()
	if err != nil {
		log.Fatalf("Unable to create mock db service: %s", err)
	}
	appConfig, err := appconfig.NewAppConfigService()
	if err != nil {
		log.Fatalf("Unable to create mock db service: %s", err)
	}

	emailService, err := service.NewEmailService(appConfig, dbService)
	if err != nil {
		log.Fatalf("Unable to create email service: %s", err)
	}

	errx := emailService.SendTestEmail("alexlehm@tilde.green")
	if errx != nil {
		log.Fatalf("Unable to create mock db service: %s", errx)
	}
}
