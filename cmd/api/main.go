package main

import (
	"github.com/qtheflash/bookmark-management/internal/api"
	"github.com/qtheflash/bookmark-management/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	app := api.NewEngine(cfg)
	err = app.Start()
	if err != nil {
		panic(err)
	}
}
