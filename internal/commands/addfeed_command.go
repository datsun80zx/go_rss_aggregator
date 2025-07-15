package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *internal.State, cmd Command) error {
	// Check that we have the required arguments
	if len(cmd.Arguments) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	// Extract the name and URL from arguments
	feedName := cmd.Arguments[0]
	feedURL := cmd.Arguments[1]

	// Get the current user from the database
	currentUsername := s.Config.CurrentUser
	if currentUsername == "" {
		return fmt.Errorf("no user is currently logged in")
	}

	user, err := s.Database.GetUser(context.Background(), currentUsername)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %v", err)
	}

	// Create the feed parameters
	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	}

	// Create the feed in the database
	feed, err := s.Database.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("couldn't create feed: %v", err)
	}

	// Print the feed details
	fmt.Printf("Feed created successfully:\n")
	fmt.Printf("ID: %v\n", feed.ID)
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("URL: %s\n", feed.Url)
	fmt.Printf("User ID: %v\n", feed.UserID)
	fmt.Printf("Created: %v\n\n", feed.CreatedAt)

	newFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	createdFeedRecord, err := s.Database.CreateFeedFollow(context.Background(), newFollow)
	if err != nil {
		return fmt.Errorf("problem creating feed record: %v", err)
	}

	fmt.Printf("Feed follow successful!\n")
	fmt.Printf("User: %v is now following feed: %v\n", createdFeedRecord.UserName, createdFeedRecord.FeedName)

	return nil
}
