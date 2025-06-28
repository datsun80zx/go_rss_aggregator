package commands

import (
	"context"
	"fmt"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
)

func HandlerReset(s *internal.State, cmd Command) error {
	err := s.Database.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("something went wrong with reseting database: %v", err)
	}
	fmt.Printf("database successfully reset!\n")
	return nil
}
