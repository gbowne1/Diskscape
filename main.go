package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gbowne1/Diskscape/config"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to the YAML configuration file for the application")
	flag.Parse()

	if config.Load == nil {
		log.Println("Error: 'Load' function is missing in the 'config' package")
		return
	}

	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		log.Printf("Configuration file not found at %s. Using default configuration.", *configPath)
		if err := config.Load("default-config.yaml"); err != nil {
			log.Fatalf("Failed to load default configuration: %v", err)
		}
		fmt.Printf("Disk Usage Analyzer initialized with configuration from %s\n", *configPath)
	} else if err != nil {
		log.Fatalf("Error checking configuration file: %v", err)
	} else {
		if err := config.Load(*configPath); err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}
	}

	fmt.Println("Disk Usage Analyzer initialized.")
}
