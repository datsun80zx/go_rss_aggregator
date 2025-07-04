package commands

import (
	"context"
	"fmt"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/rss"
)

func HandlerAgg(s *internal.State, cmd Command) error {
	feedURL := "https://www.wagslane.dev/index.xml"

	feed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("problem encountered: %v", err)
	}

	fmt.Printf("%+v\n", feed)
	return nil
}
