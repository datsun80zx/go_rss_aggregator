package main

import (
	"fmt"
	"log"
	"os"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/commands"
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

	programState := internal.State{
		Config: &cfg,
	}

	cmds := commands.Commands{
		Handlers: make(map[string]func(*internal.State, commands.Command) error),
	}

	err = cmds.Register("login", commands.HandlerLogin)
	if err != nil {
		log.Fatalf("Error registering login command: %v", err)
	}

	// parsing CLI arguments:
	if len(os.Args) < 2 {
		log.Fatalf("Not enough arguments provided")
	}

	commandName := os.Args[1]

	commandArgs := os.Args[2:]

	cmd := commands.Command{
		Name:      commandName,
		Arguments: commandArgs,
	}

	err = cmds.Run(&programState, cmd)
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	// // Set the current user to your name and update the config file
	// yourName := "datsun80zx" // Replace with your actual name
	// err = cfg.SetUser(yourName)
	// if err != nil {
	// 	log.Fatalf("Error setting user: %v", err)
	// }

	// // Read the config file again
	// updatedCfg, err := config.Read()
	// if err != nil {
	// 	log.Fatalf("Error reading updated config: %v", err)
	// }

	// // Print the updated config
	// fmt.Printf("Updated config: %+v\n", updatedCfg)
}
