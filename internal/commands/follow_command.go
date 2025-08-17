package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *internal.State, cmd Command, user database.User) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("usage: follow <url>")
	}

	// extract the feed url and get current user
	feedUrl := cmd.Arguments[0]
	// currentUsername := s.Config.CurrentUser
	// if currentUsername == "" {
	// 	return fmt.Errorf("no user is currently logged in")
	// }

	// user, err := s.Database.GetUser(context.Background(), currentUsername)
	// if err != nil {
	// 	return fmt.Errorf("couldn't get current user: %v", err)
	// }

	feed, err := s.Database.FetchFeed(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("couldn't find feed in database: %v", err)
	}

	newFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.FeedID,
	}

	createdFeedRecord, err := s.Database.CreateFeedFollow(context.Background(), newFollow)
	if err != nil {
		return fmt.Errorf("problem creating feed record: %v", err)
	}

	fmt.Printf("Feed follow successful!\n")
	fmt.Printf("User: %v is now following feed: %v\n", createdFeedRecord.UserName, createdFeedRecord.FeedName)

	return nil

}
