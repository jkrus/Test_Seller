package main

import (
	"log"
	"os"

	"github.com/jkrus/Test_Seller/cmd/sale-announcements/app"
	"github.com/jkrus/Test_Seller/internal/config"
)

func main() {
	args := os.Args[1:]
	ctx := app.NewContext()
	wg := app.NewWaitGroup()
	cfg := config.NewConfig()

	err := app.Start(ctx, args, wg, cfg)
	if err != nil {
		log.Fatal(err)
	}
}
