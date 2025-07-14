package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/commands"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/config"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	// Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Print the initial config
	// fmt.Printf("Initial config: %+v\n", cfg)

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	dbQueries := database.New(db)

	programState := internal.State{
		Config:   &cfg,
		Database: dbQueries,
	}

	cmds := commands.Commands{
		Handlers: make(map[string]func(*internal.State, commands.Command) error),
	}

	err = cmds.Register("login", commands.HandlerLogin)
	if err != nil {
		log.Fatalf("Error registering login command: %v", err)
	}

	err = cmds.Register("register", commands.HandlerRegister)
	if err != nil {
		log.Fatalf("Error registering register command: %v", err)
	}

	err = cmds.Register("reset", commands.HandlerReset)
	if err != nil {
		log.Fatalf("Error registering reset command: %v", err)
	}

	err = cmds.Register("users", commands.HandlerUsers)
	if err != nil {
		log.Fatalf("Error registering users command: %v", err)
	}

	err = cmds.Register("agg", commands.HandlerAgg)
	if err != nil {
		log.Fatalf("Error registering agg command: %v", err)
	}

	err = cmds.Register("addfeed", commands.HandlerAddFeed)
	if err != nil {
		log.Fatalf("Error registering addfeed command: %v", err)
	}

	err = cmds.Register("feeds", commands.HandlerListFeeds)
	if err != nil {
		log.Fatalf("Error registering feeds command: %v", err)
	}

	err = cmds.Register("follow", commands.HandlerFollow)
	if err != nil {
		log.Fatalf("Error registering follow command: %v", err)
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
