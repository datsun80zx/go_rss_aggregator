package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
)

func HandlerLogin(s *internal.State, cmd Command) error {
	if len(cmd.Arguments) < 1 {
		return errors.New("username required")
	}

	// get's the username from the arguments.
	username := cmd.Arguments[0]

	_, err := s.Database.GetUser(context.Background(), username)
	if err == sql.ErrNoRows {
		fmt.Printf("User %s doesn't exist\n", username)
		os.Exit(1)
	}

	if err != nil {
		return fmt.Errorf("something went wrong with database: %v", err)
	}

	err = s.Config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User set to %s\n", username)

	return nil
}
