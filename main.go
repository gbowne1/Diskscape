package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gbowne1/Diskscape/config"
    "github.com/gbowne1/Diskscape/diskstats"
    "github.com/gbowne1/Diskscape/report"
    "github.com/gbowne1/Diskscape/utils"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to configuration file")
	flag.Parse()

	if err := config.Load(*configPath); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	fmt.Println("Disk Usage Analyzer initialized.")
}
