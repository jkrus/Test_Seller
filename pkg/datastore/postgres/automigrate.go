package postgres

import (
	"log"

	"gorm.io/gorm"
)

func autoMigrate(orm *gorm.DB) {
	log.Println("Auto-migration start...")

	// Try auto-migrate database tables with specified entities.

	log.Println("Auto-migration complete.")
}
