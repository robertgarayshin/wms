package main

import (
	"log"

	"github.com/robertgarayshin/wms/config"
	"github.com/robertgarayshin/wms/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("configuration error. %s", err)
	}

	app.Run(cfg)
}
