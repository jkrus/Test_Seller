package app

import (
	"context"
	"flag"
	"log"
	"net/http"
	"sync"

	"github.com/pkg/errors"

	"github.com/jkrus/Test_Seller/internal/config"
	app_err "github.com/jkrus/Test_Seller/internal/errors"
	app "github.com/jkrus/Test_Seller/internal/handlers/http"
	"github.com/jkrus/Test_Seller/internal/services"
	"github.com/jkrus/Test_Seller/pkg/datastore/postgres"
	"github.com/jkrus/Test_Seller/pkg/server"
)

type (
	startCmd struct {
		fs   *flag.FlagSet
		name string
	}
)

func newStartCmd() Runner {
	sc := &startCmd{
		fs: flag.NewFlagSet("start", flag.ContinueOnError),
	}

	sc.fs.StringVar(&sc.name, "Start", "start", "use for start app")

	return sc
}

func (s startCmd) Init(args []string) error {
	return s.fs.Parse(args)
}

func (s startCmd) Run(ctx context.Context, wg *sync.WaitGroup, cfg *config.Config) error {
	err := cfg.Load()
	if err != nil {
		return app_err.ErrLoadConfig(err)
	}

	orm, err := postgres.Start(ctx, wg, cfg)
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return app_err.ErrOpenDatabase(err)
		}
	}

	newServices, err := services.NewServices(ctx, wg, cfg, orm)
	if err != nil {
		return err
	}

	handlers := app.NewHandlers(newServices)

	newHTTP := server.NewHTTP(ctx, wg, cfg, handlers)

	handlers.Register()

	if err = newHTTP.Start(); err != nil {
		if !errors.Is(http.ErrServerClosed, err) {
			return app_err.ErrStartHTTPServer(err)
		}
	}

	<-ctx.Done()
	wg.Wait()
	log.Println("Application shutdown complete.")

	return nil
}

func (s startCmd) Name() string {
	return s.name
}
