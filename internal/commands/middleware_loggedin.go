package commands

import (
	"context"
	"fmt"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/database"
)

func MiddlewareLoggedIn(handler func(s *internal.State, cmd Command, user database.User) error) func(*internal.State, Command) error {
	return func(s *internal.State, cmd Command) error {
		// Get the current user from the config file
		currentUsername := s.Config.CurrentUser
		if currentUsername == "" {
			return fmt.Errorf("no user is currently logged in")
		}

		// Check if current user is in the database
		user, err := s.Database.GetUser(context.Background(), currentUsername)
		if err != nil {
			return fmt.Errorf("couldn't get current user: %v", err)
		}

		return handler(s, cmd, user)
	}
}
