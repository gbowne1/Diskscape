package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yourusername/diskusage-analyzer/config"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to configuration file")
	flag.Parse()

	if err := config.Load(*configPath); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	fmt.Println("Disk Usage Analyzer initialized.")
}
