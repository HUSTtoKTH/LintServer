package main

import (
	"log"

	"github.com/HUSTtoKTH/lintserver/config"
	"github.com/HUSTtoKTH/lintserver/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
