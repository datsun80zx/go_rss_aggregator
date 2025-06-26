package commands

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *internal.State, cmd Command) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("username required")
	}

	username := cmd.Arguments[0]

	_, err := s.Database.GetUser(context.Background(), username)
	if err == nil {
		fmt.Printf("User %s already exists\n", username)
		os.Exit(1)
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("something went wrong: %v", err)
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}
	createdUser, err := s.Database.CreateUser(context.Background(), newUser)
	if err != nil {
		return err
	}

	err = s.Config.SetUser(createdUser.Name)
	if err != nil {
		return err
	}
	fmt.Printf("User created successfully:\n")
	fmt.Printf("ID: %v\n", createdUser.ID)
	fmt.Printf("Name: %s\n", createdUser.Name)
	fmt.Printf("Created: %v\n", createdUser.CreatedAt)

	fmt.Printf("User set to %s\n", username)

	return nil
}
