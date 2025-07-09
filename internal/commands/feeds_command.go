package commands

import (
	"context"
	"fmt"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
)

// HandlerListFeeds prints all feeds in the database along with their creators
func HandlerListFeeds(s *internal.State, cmd Command) error {
	// This command doesn't require any arguments
	// We're fetching all feeds regardless of the current user

	// Get all feeds from the database with user information
	feeds, err := s.Database.FetchFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't fetch feeds: %v", err)
	}

	// Check if we have any feeds
	if len(feeds) == 0 {
		fmt.Println("No feeds found in the database.")
		fmt.Println("Add a feed using: addfeed <name> <url>")
		return nil
	}

	// Print a header for clarity
	fmt.Printf("\n=== RSS Feeds (%d total) ===\n\n", len(feeds))

	// Print each feed with proper formatting
	for i, feed := range feeds {
		fmt.Printf("Feed #%d:\n", i+1)
		fmt.Printf("  Name: %s\n", feed.FeedName)
		fmt.Printf("  URL:  %s\n", feed.FeedUrl)
		fmt.Printf("  Added by: %s\n", feed.UserName)

		// Add a separator between feeds for readability
		if i < len(feeds)-1 {
			fmt.Println()
		}
	}

	fmt.Println("\n=========================")

	return nil
}
