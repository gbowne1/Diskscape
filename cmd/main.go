package main

import (
	"flag"
	"log"

	"github.com/gbowne1/Diskscape/core/config"
	"github.com/gbowne1/Diskscape/tui"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to the YAML configuration file")
	flag.Parse()

	if err := config.Load(*configPath); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if err := tui.Run(); err != nil {
		log.Fatalf("Failed to start TUI: %v", err)
	}
}
