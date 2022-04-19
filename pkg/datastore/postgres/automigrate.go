package postgres

import (
	"log"

	"gorm.io/gorm"

	"github.com/jkrus/Test_Seller/pkg/models"
)

func autoMigrate(orm *gorm.DB) {
	log.Println("Auto-migration start...")
	// Try auto-migrate database tables with specified entities.
	if err := orm.AutoMigrate(models.Announcement{}); err != nil {
		log.Fatal("automigrate failed:", err)
	}

	// Try auto-migrate database tables with specified entities.

	log.Println("Auto-migration complete.")
}
