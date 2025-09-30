package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/database"
)

func HandlerBrowse(s *internal.State, cmd Command, user database.User) error {

	var numPosts int32

	if len(cmd.Arguments) < 1 {
		numPosts = 2
	} else {
		num, err := strconv.Atoi(cmd.Arguments[0])
		if err != nil {
			return fmt.Errorf("invalid integer argument: %v", err)
		}
		numPosts = int32(num)
	}

	postParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  numPosts,
	}

	// request posts from database
	posts, err := s.Database.GetPostsForUser(context.Background(), postParams)
	if err != nil {
		return fmt.Errorf("couldn't fetch posts: %v", err)
	}

	// Check if we have any posts
	if len(posts) == 0 {
		fmt.Println("No posts found in the database.")
		return nil
	}

	// Print a header for clarity
	fmt.Printf("\n=== RSS Posts (%d total) ===\n\n", len(posts))

	// Print each post with proper formatting
	for i, post := range posts {
		fmt.Printf("Post #%d:\n", i+1)
		fmt.Printf("  Name: %v\n", post.Title.String)
		fmt.Printf("  URL:  %v\n", post.Url)
		fmt.Printf("  Published on: %v\n", post.PublishedAt.Time)
		fmt.Printf("  Description: %v\n", post.Description.String)

		// Add a separator between feeds for readability
		if i < len(posts)-1 {
			fmt.Println()
		}
	}

	fmt.Println("\n=========================")

	return nil

}
