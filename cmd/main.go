package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gbowne1/Diskscape/core/config"
	"github.com/gbowne1/Diskscape/core/diskstats"
	"github.com/gbowne1/Diskscape/tui"
)

func main() {
	// Parse command-line arguments
	configPath := flag.String("config", "config.yaml", "Path to the YAML configuration file")
	flag.Parse()

	// Load the configuration
	if err := config.Load(*configPath); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	log.Println("Configuration loaded successfully.")

	// Retrieve the target directory from the loaded configuration
	targetDir := config.Get().TargetDir
	if targetDir == "" {
		log.Fatalf("Target directory not specified in the configuration file.")
	}

	// Get disk statistics for the target directory
	stats, err := diskstats.GetDiskStats(targetDir)
	if err != nil {
		log.Fatalf("Failed to retrieve disk statistics for directory '%s': %v", targetDir, err)
	}

	// Display the retrieved disk statistics in the logs
	fmt.Printf("Disk Stats for '%s':\n", targetDir)
	fmt.Printf("- Total Space: %d bytes\n", stats.TotalSpace)
	fmt.Printf("- Free Space: %d bytes\n", stats.FreeSpace)
	fmt.Printf("- Used Percentage: %.2f%%\n", stats.UsedPercentage)

	// Start the TUI (Terminal User Interface)
	if err := tui.Run(); err != nil {
		log.Fatalf("Failed to start TUI: %v", err)
	}
}
