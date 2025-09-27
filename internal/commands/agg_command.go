package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/datsun80zx/go_rss_aggregator.git/internal"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/rss"
)

func HandlerAgg(s *internal.State, cmd Command) error {
	if len(cmd.Arguments) < 1 {
		fmt.Printf("\ntime between fetch required\n\n")
		return nil
	}

	duration, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("there was an error with parsing time: %v", err)
	}

	fmt.Printf("Collecting feeds every %v\n\n", duration)

	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		err := rss.ScrapeFeeds(context.Background(), s)
		if err != nil {
			fmt.Printf("\nthere was an error with scraping feeds: %v\n\n", err)
		}

	}

}
