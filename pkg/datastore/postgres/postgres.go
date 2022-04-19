package postgres

import (
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jkrus/Test_Seller/internal/config"
)

func postgresConnect(cfg *config.Config) (*gorm.DB, error) {
	dsn := "host=" + cfg.DB.Host + " user=" + cfg.DB.User + " password=" + cfg.DB.Pass + " dbname=" + cfg.DB.Name +
		" port=" + strconv.Itoa(cfg.DB.Port) + " sslmode=" + cfg.DB.SSLMode

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
