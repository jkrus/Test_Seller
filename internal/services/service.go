package services

import (
	"context"
	"sync"

	"gorm.io/gorm"

	"github.com/jkrus/Test_Seller/internal/announcement"
	"github.com/jkrus/Test_Seller/internal/config"
	"github.com/jkrus/Test_Seller/internal/errors"
)

type (
	Services struct {
		Announcement announcement.Service
	}
)

func NewServices(ctx context.Context, wg *sync.WaitGroup, cfg *config.Config, orm *gorm.DB) (*Services, error) {
	// provide Announcement Service.
	announcementService := announcement.NewAnnouncementService(ctx, wg, cfg, orm)
	if err := announcementService.Start(); err != nil {
		return nil, errors.ErrStartAnnouncementService(err)
	}

	return &Services{
		Announcement: announcementService,
	}, nil
}
