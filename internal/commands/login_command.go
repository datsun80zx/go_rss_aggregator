package commands

import (
	"errors"
	"fmt"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
)

func HandlerLogin(s *internal.State, cmd Command) error {
	if len(cmd.Arguments) < 1 {
		return errors.New("username required")
	}

	// get's the username from the arguments.
	username := cmd.Arguments[0]

	err := s.Config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User set to %s\n", username)

	return nil
}
