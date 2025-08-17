package commands

import (
	"context"
	"fmt"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/database"
)

func HandlerFollowing(s *internal.State, cmd Command, user database.User) error {
	// // First get the username of the current user.
	// currentUsername := s.Config.CurrentUser
	// if currentUsername == "" {
	// 	return fmt.Errorf("no user currently logged in")
	// }

	// // Next I need to get the user id of that user from the users database so I can use that as the parameter for my query.
	// currentUser, err := s.Database.GetUser(context.Background(), currentUsername)
	// if err != nil {
	// 	return fmt.Errorf("there was an issue with finding current user: %v", err)
	// }

	followList, err := s.Database.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error with retrieving feeds followed: %v", err)
	}

	if len(followList) == 0 {
		fmt.Println("You are not following any feeds.")
		fmt.Println("follow a feed using: follow <url>")
		return nil
	}

	fmt.Printf("\n=== RSS Feeds Followed (%d total) ===\n\n", len(followList))

	for i, feed := range followList {
		fmt.Printf("Feed #%d:\n", i+1)
		fmt.Printf("  Name: %s\n", feed.FeedName)
		fmt.Printf("  URL:  %s\n\n", feed.FeedUrl)
	}

	fmt.Println("\n=========================")

	return nil

}
