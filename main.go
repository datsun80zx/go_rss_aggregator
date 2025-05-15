package main

import (
	"fmt"
	"log"

	"github.com/datsun80zx/go_rss_aggregator.git/internal/config"
)

func main() {
	// Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Print the initial config
	fmt.Printf("Initial config: %+v\n", cfg)

	// Set the current user to your name and update the config file
	yourName := "datsun80zx" // Replace with your actual name
	err = cfg.SetUser(yourName)
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}

	// Read the config file again
	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading updated config: %v", err)
	}

	// Print the updated config
	fmt.Printf("Updated config: %+v\n", updatedCfg)
}
