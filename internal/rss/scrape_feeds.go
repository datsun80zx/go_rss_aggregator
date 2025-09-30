package rss

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func ScrapeFeeds(ctx context.Context, s *internal.State) error {
	// Get next feed to fetch
	feed, err := s.Database.GetNextFeedToFetch(ctx)
	if err == sql.ErrNoRows {
		fmt.Printf("there are no feeds to fetch\n")
		return nil
	}

	if err != nil {
		return fmt.Errorf("there was a problem with connecting to the db: %v", err)
	}

	// Fetch the feed
	rssFeed, err := FetchFeed(ctx, feed.FeedsUrl)
	if err != nil {
		return fmt.Errorf("there was an issue fetching feed from db: %v", err)
	}

	// Mark it as fetched

	markedFeed := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        feed.FeedsID,
	}

	err = s.Database.MarkFeedFetched(ctx, markedFeed)
	if err != nil {
		return fmt.Errorf("there was an issue marking feed as marked: %v", err)
	}

	// print titles
	for _, item := range rssFeed.Channel.Item {

		itemPubTime := parseTime(item.PubDate)
		itemDescription := sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		}
		itemTitle := sql.NullString{
			String: item.Title,
			Valid:  item.Title != "",
		}

		newPost := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       itemTitle,
			Url:         item.Link,
			Description: itemDescription,
			PublishedAt: itemPubTime,
			FeedID:      markedFeed.ID,
		}

		_, err := s.Database.CreatePost(ctx, newPost)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" {
					continue
				}
			}

			return fmt.Errorf("failed to create post: %v", err)

		}

	}

	return nil
}
