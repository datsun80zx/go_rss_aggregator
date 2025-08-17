package commands

import (
	"context"
	"fmt"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/database"
)

func HandlerUnfollow(s *internal.State, cmd Command, user database.User) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("usage: follow <url>")
	}

	unfollowParams := database.UnfollowFeedParams{
		UserID: user.ID,
		Url:    cmd.Arguments[0],
	}

	err := s.Database.UnfollowFeed(context.Background(), unfollowParams)
	if err != nil {
		return fmt.Errorf("there was an issue with deleting the feed: %v", err)
	}
	fmt.Printf("successfully deleted: %v\n", unfollowParams.Url)
	return nil
}
