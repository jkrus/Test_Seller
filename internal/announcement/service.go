package announcement

import (
	"context"
	"log"
	"sync"

	"gorm.io/gorm"

	store "github.com/jkrus/Test_Seller/internal/announcement/repo"
	"github.com/jkrus/Test_Seller/internal/config"
	"github.com/jkrus/Test_Seller/pkg/service"
)

type (
	Service interface {
		service.Service
		Add([]byte) (int64, error)
		GetById(id int64, opt []string) ([]byte, error)
		GetList(page, limit int, sortBy, sort string) ([]byte, error)
		AddImage(id int, path string) ([]byte, error)
	}

	announcementService struct {
		ctx    context.Context
		mainWG *sync.WaitGroup
		cfg    *config.Config

		repo store.Storage
	}
)

func NewAnnouncementService(ctx context.Context, wg *sync.WaitGroup, cfg *config.Config, orm *gorm.DB) Service {

	return &announcementService{ctx: ctx, mainWG: wg, cfg: cfg, repo: store.NewAnnouncementStorage(orm)}
}

func (as *announcementService) Add(bytes []byte) (int64, error) {
	id, err := as.repo.Create(bytes)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (as *announcementService) AddImage(id int, path string) ([]byte, error) {
	data, err := as.repo.Update(id, path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *announcementService) GetById(id int64, opt []string) ([]byte, error) {
	data, err := as.repo.Get(id, opt)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (as *announcementService) GetList(page, limit int, sortBy, sort string) ([]byte, error) {
	list, err := as.repo.GetList(page, limit, sortBy, sort)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (as *announcementService) Start() error {
	log.Println("Start Announcement service...")

	as.createHandlerContext()

	log.Println("Announcement service start success.")

	return nil
}

func (as *announcementService) Stop() error {
	log.Println("Stop Announcement Service...")

	log.Println("Info Announcement stopped.")

	return nil
}

func (as *announcementService) createHandlerContext() {
	as.mainWG.Add(1)
	go func() {
		for {
			<-as.ctx.Done()
			_ = as.Stop()
			as.mainWG.Done()
			return
		}
	}()

}
