package commands

import (
	"context"
	"fmt"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
)

func HandlerUsers(s *internal.State, cmd Command) error {

	userList, err := s.Database.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("something went wrong in database: %v", err)
	}

	if len(userList) == 0 {
		fmt.Printf("no registered users\n")
		return nil
	}

	currentUser := s.Config.CurrentUser

	for _, u := range userList {
		if u.Name == currentUser {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}
	return nil
}
